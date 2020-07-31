package mtproto

import (
	cryptoRand "crypto/rand"
	"fmt"
	"github.com/ansel1/merry"
	"math/rand"
	"net"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/proxy"
	"golang.org/x/sync/semaphore"
)

//go:generate go run schema/tl_generate.go 116 < schema/tl-schema-116.tl > tl_schema.go
//go:generate gofmt -w tl_schema.go

const RoutinesCount = 4

var log = new(Logger)

type AppConfig struct {
	AppID          int32
	AppHash        string
	AppVersion     string
	DeviceModel    string
	SystemVersion  string
	SystemLangCode string
	LangPack       string
	LangCode       string
}

type MTProto struct {
	useIPv6    bool
	session    *Session
	appCfg     *AppConfig
	connDialer proxy.Dialer
	conn       net.Conn

	// Two queues here.
	// First (external) has limited size and contains external requests.
	// Second (internal) is unlimited. Special goroutine transfers messages
	// from external to internal queue while len(interbal) < cap(external).
	// This allows throttling (as same as single limited queue).
	// And this allow to safely (without locks) return any failed (due to
	// network probles for example) messages back to internal queue and retry later.
	extSendQueue chan *packetToSend //external
	sendQueue    chan *packetToSend //internal

	routinesStop chan struct{}
	routinesWG   sync.WaitGroup

	mutex            *sync.Mutex
	connectSemaphore *semaphore.Weighted
	reconnSemaphore  *semaphore.Weighted

	encryptionReady bool
	lastSeqNo       int32
	msgsByID        map[int64]*packetToSend
	seqNo           int32
	msgId           int64
	handleEvent     func(TL)

	dcOptions []*DcOption
}

type packetToSend struct {
	msgID   int64
	seqNo   int32
	msg     TL
	resp    chan TL
	needAck bool
}

func newPacket(msg TL, resp chan TL) *packetToSend {
	return &packetToSend{msg: msg, resp: resp}
}

func NewAppConfig(appID int32, appHash string) *AppConfig {
	if appID == 0 {
		appID = int32(48841)
	}
	if appHash == "" {
		appHash = "3151c01673d412c18c055f089128be50"
	}
	config := &AppConfig{
		AppID:          appID,
		AppHash:        appHash,
		AppVersion:     "0.0.1",
		DeviceModel:    "Unknown",
		SystemVersion:  runtime.GOOS + "/" + runtime.GOARCH,
		SystemLangCode: "en",
		LangPack:       "",
		LangCode:       "en",
	}
	return config
}

func NewMTProto(cfg *AppConfig) *MTProto {
	if cfg == nil {
		cfg = NewAppConfig(0, "")
	}

	m := &MTProto{
		useIPv6:    false,
		session:    nil,
		connDialer: &net.Dialer{},
		appCfg:     cfg,

		extSendQueue: make(chan *packetToSend, 64),
		sendQueue:    make(chan *packetToSend, 1024),
		routinesStop: make(chan struct{}, RoutinesCount),

		msgsByID: make(map[int64]*packetToSend),
		mutex:    &sync.Mutex{},

		connectSemaphore: semaphore.NewWeighted(1),
		reconnSemaphore:  semaphore.NewWeighted(1),
	}

	go m.debugRoutine()
	return m
}

func (m *MTProto) SetSession(s *Session) {
	m.session = s
}

func (m *MTProto) UseIPv6(b bool) {
	m.useIPv6 = b
}

func (m *MTProto) SetDialer(d proxy.Dialer) {
	m.connDialer = d
}

func (m *MTProto) InitSessAndConnect() error {
	if err := m.InitSession(true); err != nil {
		return merry.Wrap(err)
	}
	if err := m.Connect(); err != nil {
		return merry.Wrap(err)
	}
	return nil
}

func (m *MTProto) InitSession(isReady bool) error {
	if m.session != nil {
		m.encryptionReady = isReady
	} else {
		m.encryptionReady = false
		// Create New Session
		if !m.useIPv6 {
			m.session = &Session{Addr: "149.154.167.50:443"}
		} else {
			m.session = &Session{Addr: "[2001:67c:4e8:f002::a]:443"}
		}
	}

	rand.Seed(time.Now().UnixNano())
	m.session.sessionId = rand.Int63()
	return nil
}

func (m *MTProto) CopySession() *Session {
	sess := *m.session
	return &sess
}

func (m *MTProto) DCAddr(dcID int32) (string, bool) {
	for _, o := range m.dcOptions {
		if o.ID == dcID && o.Ipv6 == m.useIPv6 && !o.Cdn {
			if o.Ipv6 {
				return fmt.Sprintf("[%s]:%d", o.IpAddress, o.Port), true
			}
			return fmt.Sprintf("%s:%d", o.IpAddress, o.Port), true
		}
	}
	return "", false
}

func (m *MTProto) SetEventsHandler(handler func(TL)) {
	m.handleEvent = handler
}

func (m *MTProto) initConfig() error {

	// getting connection configs
	log.Debug("connecting: getting config...")
	x := m.SendSyncRetry(InvokeWithLayer{
		Layer,
		InitConnection{
			Flags:          0,
			ApiID:          m.appCfg.AppID,
			DeviceModel:    m.appCfg.DeviceModel,
			SystemVersion:  m.appCfg.SystemVersion,
			AppVersion:     m.appCfg.AppVersion,
			SystemLangCode: m.appCfg.SystemLangCode,
			LangPack:       m.appCfg.LangPack,
			LangCode:       m.appCfg.LangCode,
			Proxy:          nil, //flagged
			Query:          HelpGetConfig{},
		},
	}, 5, 10, 10)
	if cfg, ok := x.(Config); ok {
		m.session.DcID = cfg.ThisDc
		for _, v := range cfg.DcOptions {
			v := v.(DcOption)
			m.dcOptions = append(m.dcOptions, &v)
		}
	} else {
		return WrongRespError(x)
	}
	return nil
}

func (m *MTProto) initConnection() error {
	log.Info("connecting to DC %d (%s)...", m.session.DcID, m.session.Addr)
	var err error
	m.conn, err = m.connDialer.Dial("tcp", m.session.Addr)
	if err != nil {
		return merry.Wrap(err)
	}
	_, err = m.conn.Write([]byte{0xef})
	if err != nil {
		return merry.Wrap(err)
	}

	// getting new authKey if need
	if !m.encryptionReady {
		if err = m.makeAuthKey(); err != nil {
			return merry.Wrap(err)
		}
		m.encryptionReady = true
	}
	return nil
}

func (m *MTProto) Connect() error {
	if !m.connectSemaphore.TryAcquire(1) {
		log.Info("connection already in progress, aborting")
		return nil
	}
	defer m.connectSemaphore.Release(1)

	m.routinesWG.Add(4)

	retry := 0
	for {
		err := m.initConnection()
		if err == nil {
			break
		}
		log.Warn("failed to connect, Retrying (%d)", retry)
		if retry > 10 {
			log.Error(err, "failed to connect, Abort.")
			break
		}
		retry += 1
		time.Sleep(time.Second)
	}

	// starting goroutines
	log.Debug("connecting: starting routines...")
	go m.sendRoutine()
	go m.readRoutine()
	go m.queueTransferRoutine() // starting messages transfer from external to internal queue
	go m.pingRoutine()          // starting keepalive pinging

	err := m.initConfig()
	if err != nil {
		log.Error(err, "failed to get config")
	}

	log.Info("connected to DC %d (%s)...", m.session.DcID, m.session.Addr)
	return nil
}

func (m *MTProto) reconnectLogged() {
	log.Info("reconnecting...")
	if !m.reconnSemaphore.TryAcquire(1) {
		log.Info("reconnection already in progress, aborting")
		return
	}
	defer func() { m.reconnSemaphore.Release(1) }()

	for {
		err := m.reconnect(0)
		if err == nil {
			return
		}
		log.Error(err, "failed to reconnect")
		log.Info("retrying in 5 seconds")
		time.Sleep(5 * time.Second)
		// and trying to reconnect again
	}
}

func (m *MTProto) Reconnect() error {
	return m.reconnect(0)
}

func (m *MTProto) reconnect(newDcID int32) error {
	log.Info("reconnecting: DC %d -> %d", m.session.DcID, newDcID)

	// stopping routines
	log.Debug("stopping routines...")
	for i := 0; i < RoutinesCount; i++ {
		m.routinesStop <- struct{}{}
	}

	// closing connection, readRoutine will then fail to read() and will handle stop signal
	if m.conn != nil {
		if err := m.conn.Close(); err != nil && !IsClosedConnErr(err) {
			return merry.Wrap(err)
		}
	}

	// waiting for all routines to stop
	log.Debug("waiting for routines...")
	m.routinesWG.Wait()
	log.Debug("done stopping routines...")

	// removing unused stop signals (if any)
	for empty := false; !empty; {
		select {
		case <-m.routinesStop:
		default:
			empty = true
		}
	}

	// saving IDs of messages from msgsByID[],
	// some of them may not have been sent, so we'll resend them after reconnection
	pendingIDs := make([]int64, 0, len(m.msgsByID))
	m.mutex.Lock()
	for id := range m.msgsByID {
		pendingIDs = append(pendingIDs, id)
	}
	m.mutex.Unlock()
	log.Debug("found %d pending packet(s)", len(pendingIDs))

	if newDcID != 0 {
		// renewing connection
		if newDcID != m.session.DcID {
			m.encryptionReady = false //TODO: export auth here (if authed)

			//https://github.com/sochix/TLSharp/blob/0940d3d982e9c22adac96b6c81a435403802899a/TLSharp.Core/TelegramClient.cs#L84
		}
		newDcAddr, ok := m.DCAddr(newDcID)
		if !ok {
			return merry.Errorf("wrong DC number: %d", newDcID)
		}
		m.session.DcID = newDcID
		m.session.Addr = newDcAddr
	}

	if err := m.Connect(); err != nil {
		return merry.Wrap(err)
	}

	// Checking pending messages.
	// 1) some of them may have been answered, so they will not be in msgsByID[]
	// 2) some of them may have been received by TG, but response has not reached us yet
	// 3) some of them may be actually lost and must be resend
	// Here we resend messages both from (2) and (3). But since msgID and seqNo
	// are preserved, TG will ignore doubles from (2). And (3) will finally reach TG.
	if len(pendingIDs) > 0 {
		var packets []*packetToSend
		m.mutex.Lock()
		for _, id := range pendingIDs {
			packet, ok := m.msgsByID[id]
			if ok {
				packets = append(packets, packet)
			}
		}
		m.pushPendingPacketsUnlocked(packets)
		m.mutex.Unlock()
	}

	log.Info("reconnected to DC %d (%s)", m.session.DcID, m.session.Addr)
	return nil
}

func (m *MTProto) NewConnection(dcID int32) (*MTProto, error) {
	session := m.CopySession()
	log.Info("making new connection to DC %d (current: %d)", dcID, session.DcID)
	isOnSameDC := session.DcID == dcID
	encrIsReady := isOnSameDC
	session.DcID = dcID
	var ok bool
	if session.Addr, ok = m.DCAddr(dcID); !ok {
		return nil, merry.Errorf("unable find address for DC #%d", dcID)
	}

	newMT := NewMTProto(m.appCfg)
	newMT.SetDialer(m.connDialer)
	newMT.SetSession(session)

	if err := newMT.InitSession(encrIsReady); err != nil {
		return nil, merry.Wrap(err)
	}
	if err := newMT.Connect(); err != nil {
		return nil, merry.Wrap(err)
	}

	if !isOnSameDC {
		res := m.SendSync(AuthExportAuthorization{DcID: dcID})
		exported, ok := res.(AuthExportedAuthorization)
		if !ok {
			return nil, merry.New(UnexpectedTL("auth export", res))
		}
		res = newMT.SendSync(AuthImportAuthorization{ID: exported.ID, Bytes: exported.Bytes})
		if _, ok := res.(AuthAuthorization); !ok {
			return nil, merry.New(UnexpectedTL("auth import", res))
		}
	}
	return newMT, nil
}

func (m *MTProto) Send(msg TLReq) chan TL {
	resp := make(chan TL, 1)
	m.extSendQueue <- newPacket(msg, resp)
	return resp
}

func (m *MTProto) SendSync(msg TLReq) TL {
	resp := make(chan TL, 1)
	m.extSendQueue <- newPacket(msg, resp)
	return <-resp
}

func (m *MTProto) SendSyncRetry(
	msg TLReq, failRetryInterval time.Duration,
	floodNumShortRetries int, floodMaxWait time.Duration,
) TL {
	retryNum := -1
	for {
		retryNum += 1
		res := m.SendSync(msg)

		if IsError(res, "RPC_CALL_FAIL") {
			log.Warn("got RPC error, retrying in %s", failRetryInterval)
			time.Sleep(failRetryInterval)
			continue
		}

		if floodWait, ok := IsFloodError(res); ok {
			if retryNum < floodNumShortRetries {
				floodWait = time.Second
			}
			if floodWait > floodMaxWait {
				return res
			}
			log.Warn("got flood-wait, retrying in %s, retry #%d of %d short",
				floodWait, retryNum, floodNumShortRetries)
			time.Sleep(floodWait)
			continue
		}

		return res
	}
}

type AuthDataProvider interface {
	PhoneNumber() (string, error)
	Code() (string, error)
	Token() (string, error)
	Password() (string, error)
}

type ScanfAuthDataProvider struct{}

func (ap ScanfAuthDataProvider) PhoneNumber() (string, error) {
	var phonenumber string
	fmt.Print("Enter phone number: ")
	_, _ = fmt.Scanf("%s", &phonenumber)
	return phonenumber, nil
}

func (ap ScanfAuthDataProvider) Code() (string, error) {
	var code string
	fmt.Print("Enter code: ")
	_, _ = fmt.Scanf("%s", &code)
	return code, nil
}

func (ap ScanfAuthDataProvider) Token() (string, error) {
	var code string
	fmt.Print("Enter token: ")
	_, _ = fmt.Scanf("%s", &code)
	return code, nil
}

func (ap ScanfAuthDataProvider) Password() (string, error) {
	var passwd string
	fmt.Print("Enter password: ")
	_, _ = fmt.Scanf("%s", &passwd)
	return passwd, nil
}

func (m *MTProto) Auth(authData AuthDataProvider) error {
	phonenumber, err := authData.PhoneNumber()
	if err != nil {
		return merry.Wrap(err)
	}

	var authSentCode AuthSentCode
	flag := true
	for flag {
		x := m.SendSync(AuthSendCode{
			PhoneNumber: phonenumber,
			ApiID:       m.appCfg.AppID,
			ApiHash:     m.appCfg.AppHash,
			Settings:    CodeSettings{Flags: 1, CurrentNumber: true},
		})
		switch x.(type) {
		case AuthSendCode:
			authSentCode = x.(AuthSentCode)
			flag = false
		case RpcError:
			x := x.(RpcError)
			if x.ErrorCode != ErrSeeOther {
				return WrongRespError(x)
			}
			var newDc int32
			n, _ := fmt.Sscanf(x.ErrorMessage, "PHONE_MIGRATE_%d", &newDc)
			if n != 1 {
				n, _ := fmt.Sscanf(x.ErrorMessage, "NETWORK_MIGRATE_%d", &newDc)
				if n != 1 {
					return merry.Errorf("RPC error_string: %s", x.ErrorMessage)
				}
			}

			if err := m.reconnect(newDc); err != nil {
				return merry.Wrap(err)
			}
		default:
			return WrongRespError(x)
		}
	}

	code, err := authData.Code()
	if err != nil {
		return merry.Wrap(err)
	}

	//if authSentCode.Phone_registered
	x := m.SendSync(AuthSignIn{phonenumber, authSentCode.PhoneCodeHash, code})
	if IsError(x, "SESSION_PASSWORD_NEEDED") {
		x = m.SendSync(AccountGetPassword{})
		accPasswd, ok := x.(AccountPassword)
		if !ok {
			return WrongRespError(x)
		}

		passwd, err := authData.Password()
		if err != nil {
			return merry.Wrap(err)
		}

		algo, ok := accPasswd.CurrentAlgo.(PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow)
		if !ok {
			return merry.Errorf("unknown password algo %T, application update is maybe needed to log in",
				accPasswd.CurrentAlgo)
		}
		passwdSRP, err := calcInputCheckPasswordSRP(algo, accPasswd, passwd, cryptoRand.Read, log.Debug)
		if err != nil {
			return merry.Wrap(err)
		}
		x = m.SendSync(AuthCheckPassword{passwdSRP.(InputCheckPasswordSRP)})
		if _, ok := x.(RpcError); ok {
			return WrongRespError(x)
		}
	}
	auth, ok := x.(AuthAuthorization)
	if !ok {
		return merry.Errorf("RPC: %#v", x)
	}
	userSelf := auth.User.(User)
	fmt.Printf("Signed in: id %d name <%s %s>\n", userSelf.ID, userSelf.FirstName, userSelf.LastName)
	return nil
}

func (m *MTProto) AuthBot(authData AuthDataProvider, tokenList ...string) error {
	var token string
	var err error
	if len(tokenList) == 0 {
		token, err = authData.Token()
		if err != nil {
			return merry.Wrap(err)
		}
	} else {
		token = tokenList[0]
	}

	var auth AuthAuthorization
	flag := true
	for flag {
		x := m.SendSync(AuthImportBotAuthorization{
			ApiID:        m.appCfg.AppID,
			ApiHash:      m.appCfg.AppHash,
			BotAuthToken: token,
		})
		switch x.(type) {
		case AuthAuthorization:
			auth = x.(AuthAuthorization)
			flag = false
		case RpcError:
			x := x.(RpcError)
			if x.ErrorCode != ErrSeeOther {
				return WrongRespError(x)
			}
			var newDc int32
			n, _ := fmt.Sscanf(x.ErrorMessage, "USER_MIGRATE_%d", &newDc)
			if n != 1 {
				n, _ := fmt.Sscanf(x.ErrorMessage, "NETWORK_MIGRATE_%d", &newDc)
				if n != 1 {
					return merry.Errorf("RPC error_string: %s", x.ErrorMessage)
				}
			}

			if err := m.reconnect(newDc); err != nil {
				return merry.Wrap(err)
			}
		default:
			return WrongRespError(x)
		}
	}

	userSelf := auth.User.(User)
	fmt.Printf("Signed in: id %d name <%s %s>\n", userSelf.ID, userSelf.FirstName, userSelf.LastName)
	return nil
}

func (m *MTProto) popPendingPacketsUnlocked() []*packetToSend {
	packets := make([]*packetToSend, 0, len(m.msgsByID))
	msgs := make([]TL, 0)
	for id, packet := range m.msgsByID {
		delete(m.msgsByID, id)
		packets = append(packets, packet)
		msgs = append(msgs, packet.msg)
	}
	log.Debug("popped %d pending packet(s): %#v", len(packets), msgs)
	return packets
}
func (m *MTProto) popPendingPackets() []*packetToSend {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.popPendingPacketsUnlocked()
}
func (m *MTProto) pushPendingPacketsUnlocked(packets []*packetToSend) {
	for _, packet := range packets {
		m.sendQueue <- packet
	}
	log.Debug("pushed %d pending packet(s)", len(packets))
}
func (m *MTProto) pushPendingPackets(packets []*packetToSend) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.pushPendingPacketsUnlocked(packets)
}
func (m *MTProto) resendPendingPackets() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	packets := m.popPendingPacketsUnlocked()
	m.pushPendingPacketsUnlocked(packets)
}

func (m *MTProto) GetContacts() error {
	x := m.SendSync(ContactsGetContacts{0})
	list, ok := x.(ContactsContacts)
	if !ok {
		return merry.Errorf("RPC: %#v", x)
	}

	contacts := make(map[int32]User)
	for _, v := range list.Users {
		if v, ok := v.(User); ok {
			contacts[v.ID] = v
		}
	}
	fmt.Printf(
		"\033[33m\033[1m%10s    %10s    %-30s    %-20s\033[0m\n",
		"id", "mutual", "name", "username",
	)
	for _, v := range list.Contacts {
		v := v.(Contact)
		fmt.Printf(
			"%10d    %10t    %-30s    %-20s\n",
			v.UserID,
			toBool(v.Mutual),
			fmt.Sprintf("%s %s", contacts[v.UserID].FirstName, contacts[v.UserID].LastName),
			contacts[v.UserID].Username,
		)
	}

	return nil
}

func (m *MTProto) pingRoutine() {
	defer func() {
		log.Debug("pingRoutine done")
		m.routinesWG.Done()
	}()
	for {
		select {
		case <-m.routinesStop:
			return
		case <-time.After(60 * time.Second):
			m.extSendQueue <- newPacket(Ping{0xCADACADA}, nil)
		}
	}
}

func (m *MTProto) sendRoutine() {
	defer func() {
		log.Debug("sendRoutine done")
		m.routinesWG.Done()
	}()
	for {
		select {
		case <-m.routinesStop:
			return
		case x := <-m.sendQueue:
			err := m.send(x)
			if IsClosedConnErr(err) {
				continue //closed connection, should receive stop signal now
			}
			if err != nil {
				log.Error(err, "sending failed")
				go m.reconnectLogged()
				return
			}
		}
	}
}

func (m *MTProto) readRoutine() {
	defer func() {
		log.Debug("readRoutine done")
		m.routinesWG.Done()
	}()
	for {
		select {
		case <-m.routinesStop:
			return
		default:
		}

		data, err := m.read()
		if IsClosedConnErr(err) {
			continue //closed connection, should receive stop signal now
		}
		if err != nil {
			log.Error(err, "reading failed")
			go m.reconnectLogged()
			return
		}
		m.process(m.msgId, m.seqNo, data, true)
	}
}

func (m *MTProto) queueTransferRoutine() {
	defer func() {
		log.Debug("queueTransferRoutine done")
		m.routinesWG.Done()
	}()
	for {
		if len(m.sendQueue) < cap(m.extSendQueue) {
			select {
			case <-m.routinesStop:
				return
			case msg := <-m.extSendQueue:
				m.sendQueue <- msg
			}
		} else {
			select {
			case <-m.routinesStop:
				return
			default:
				time.Sleep(10000 * time.Microsecond)
			}
		}
	}
}

// Periodically checks messages in "msgsByID" and warns if they stay there too long
func (m *MTProto) debugRoutine() {
	for {
		m.mutex.Lock()
		count := 0
		for id := range m.msgsByID {
			delta := time.Now().Unix() - (id >> 32)
			if delta > 5 {
				log.Warn("msgsByID: #%d: is here for %ds", id, delta)
			}
			count++
		}
		m.mutex.Unlock()
		log.Debug("msgsByID: %d total", count)
		time.Sleep(5 * time.Second)
	}
}

func (m *MTProto) clearPacketData(msgID int64) {
	m.mutex.Lock()
	packet, ok := m.msgsByID[msgID]
	if ok {
		if packet.resp != nil {
			close(packet.resp)
		}
		delete(m.msgsByID, msgID)
	}
	m.mutex.Unlock()
}
func (m *MTProto) respAndClearPacketData(msgID int64, response TL) {
	m.mutex.Lock()
	packet, ok := m.msgsByID[msgID]
	if ok {
		if packet.resp == nil {
			log.Warn("second response to message #%d %#v", msgID, packet.msg)
		} else {
			packet.resp <- response
			close(packet.resp)
			packet.resp = nil
		}
		delete(m.msgsByID, msgID)
	}
	m.mutex.Unlock()
}

func (m *MTProto) process(msgId int64, seqNo int32, dataTL TL, mayPassToHandler bool) {
	switch data := dataTL.(type) {
	case MsgContainer:
		for _, v := range data.Items {
			m.process(v.MsgID, v.SeqNo, v.Data, true)
		}

	case BadServerSalt:
		m.session.ServerSalt = data.NewServerSalt
		m.resendPendingPackets()

	case BadMsgNotification:
		m.respAndClearPacketData(data.BadMsgID, data)

	case MsgsStateInfo:
		m.respAndClearPacketData(data.ReqMsgID, data)

	case NewSessionCreated:
		m.session.ServerSalt = data.ServerSalt

	case Ping:
		m.sendQueue <- newPacket(Pong{msgId, data.PingID}, nil)

	case Pong:
		// (ignore)

	case MsgsAck:
		m.mutex.Lock()
		for _, id := range data.MsgIds {
			packet, ok := m.msgsByID[id]
			if ok {
				packet.needAck = false
				// if request is not waiting for response, removing it
				if m.msgsByID[id].resp == nil { //TODO: packet.resp
					delete(m.msgsByID, id)
				}
			}
		}
		m.mutex.Unlock()

	case RpcResult:
		m.process(msgId, 0, data.obj, false)
		m.respAndClearPacketData(data.reqMsgID, data.obj)

	default:
		if mayPassToHandler && m.handleEvent != nil {
			go m.handleEvent(dataTL)
		}
	}

	// should acknowledge odd ids
	if (seqNo & 1) == 1 {
		m.sendQueue <- newPacket(MsgsAck{[]int64{msgId}}, nil)
	}
}
