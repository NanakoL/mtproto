package mtproto

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"
)

func (m *MTProto) justSend(msg TLReq) error {
	return WrapError(m.send(newPacket(msg, nil)))
}

func (m *MTProto) send(packet *packetToSend) error {
	if packet.msgID == 0 {
		packet.msgID = GenerateMessageId()
	}
	log.Message(false, packet.msg, packet.msgID)
	obj := packet.msg.encode()

	x := NewEncodeBuf(256)

	// padding for tcpsize
	x.Int(0)

	if m.encryptionReady {
		packet.needAck = true
		switch packet.msg.(type) {
		case Ping, MsgsAck:
			packet.needAck = false
		}
		z := NewEncodeBuf(256)
		z.Long(m.session.ServerSalt)
		z.Long(m.session.sessionId)
		z.Long(packet.msgID)
		if packet.seqNo == 0 {
			if packet.needAck {
				packet.seqNo = m.lastSeqNo | 1
			} else {
				packet.seqNo = m.lastSeqNo
			}
			m.lastSeqNo += 2
		}
		z.Int(packet.seqNo)
		z.Int(int32(len(obj)))
		z.Bytes(obj)

		msgKey := sha1Sum(z.buf)[4:20]
		aesKey, aesIV := generateAES(msgKey, m.session.AuthKey, false)

		y := make([]byte, len(z.buf)+((16-(len(obj)%16))&15))
		copy(y, z.buf)
		encryptedData, err := doAES256IGEencrypt(y, aesKey, aesIV)
		if err != nil {
			return WrapError(err)
		}

		x.Bytes(m.session.AuthKeyHash)
		x.Bytes(msgKey)
		x.Bytes(encryptedData)

		if packet.resp != nil || packet.needAck {
			m.mutex.Lock()
			m.msgsByID[packet.msgID] = packet
			m.mutex.Unlock()
		}
	} else {
		x.Long(0)
		x.Long(packet.msgID)
		x.Int(int32(len(obj)))
		x.Bytes(obj)
	}

	// minus padding
	size := len(x.buf)/4 - 1

	if size < 127 {
		x.buf[3] = byte(size)
		x.buf = x.buf[3:]
	} else {
		binary.LittleEndian.PutUint32(x.buf, uint32(size<<8|127))
	}
	_ = m.conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	if _, err := m.conn.Write(x.buf); err != nil {
		return WrapError(err)
	}

	return nil
}

func (m *MTProto) read() (TL, error) {
	var err error
	var n int
	var size int
	var data TL

	err = m.conn.SetReadDeadline(time.Now().Add(300 * time.Second))
	if err != nil {
		return nil, WrapError(err)
	}
	b := make([]byte, 1)
	n, err = m.conn.Read(b)
	if err != nil {
		return nil, WrapError(err)
	}

	if b[0] < 127 {
		size = int(b[0]) << 2
	} else {
		b := make([]byte, 3)
		n, err = m.conn.Read(b)
		if err != nil {
			return nil, WrapError(err)
		}
		size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
	}

	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = m.conn.Read(buf[size-left:])
		if err != nil {
			return nil, WrapError(err)
		}
		left -= n
	}

	if size == 4 {
		return nil, fmt.Errorf("server response error: %d", int32(binary.LittleEndian.Uint32(buf)))
	}

	dbuf := NewDecodeBuf(buf)

	authKeyHash := dbuf.Bytes(8)
	if binary.LittleEndian.Uint64(authKeyHash) == 0 {
		m.msgId = dbuf.Long()
		messageLen := dbuf.Int()
		if int(messageLen) != dbuf.size-20 {
			return nil, fmt.Errorf("message len: %d (need %d)", messageLen, dbuf.size-20)
		}
		m.seqNo = 0

		data = dbuf.Object()
		if dbuf.err != nil {
			return nil, WrapError(dbuf.err)
		}
	} else {
		msgKey := dbuf.Bytes(16)
		encryptedData := dbuf.Bytes(dbuf.size - 24)
		aesKey, aesIV := generateAES(msgKey, m.session.AuthKey, true)
		x, err := doAES256IGEdecrypt(encryptedData, aesKey, aesIV)
		if err != nil {
			return nil, WrapError(err)
		}
		dbuf = NewDecodeBuf(x)
		_ = dbuf.Long() // salt
		_ = dbuf.Long() // session_id
		m.msgId = dbuf.Long()
		m.seqNo = dbuf.Int()
		messageLen := dbuf.Int()
		if int(messageLen) > dbuf.size-32 {
			return nil, fmt.Errorf("message len: %d (need <= %d)", messageLen, dbuf.size-32)
		}
		// DEBUG vvv
		if dbuf.err != nil {
			panic(dbuf.err)
		}
		if int(messageLen)+32 > len(dbuf.buf) {
			panic(fmt.Sprintf("AHTUNG!!! %d(%d) + 32 > %d", messageLen, int(messageLen), len(dbuf.buf)))
		}
		hash := sha1.Sum(dbuf.buf[0 : 32+messageLen])
		if !bytes.Equal(hash[4:20], msgKey) {
			return nil, fmt.Errorf("wrong msg_key")
		}
		// if !bytes.Equal(sha1.Sum(dbuf.buf[0 : 32+messageLen])[4:20], msgKey) {
		// 	return nil, fmt.Errorf("Wrong msg_key")
		// }
		// DEBUG ^^^

		data = m.decodeMessage(dbuf, nil)
		if dbuf.err != nil {
			return nil, WrapError(dbuf.err)
		}
	}
	mod := m.msgId & 3
	if mod != 1 && mod != 3 {
		return nil, fmt.Errorf("wrong bits of message_id: %d", mod)
	}

	log.Message(true, data, 0)
	return data, nil
}

func (m *MTProto) makeAuthKey() error {
	var x []byte
	var err error
	var data interface{}

	// (send) req_pq
	nonceFirst := GenerateNonce(16)
	err = m.justSend(ReqPqMulti{nonceFirst})
	if err != nil {
		return WrapError(err)
	}

	// (parse) resPQ
	data, err = m.read()
	if err != nil {
		return WrapError(err)
	}
	res, ok := data.(ResPQ)
	if !ok {
		return fmt.Errorf("handshake: Need resPQ, got %#v", data)
	}
	if !bytes.Equal(nonceFirst, res.Nonce) {
		return fmt.Errorf("handshake: Wrong nonce")
	}
	found := false
	for _, b := range res.ServerPublicKeyFingerprints {
		if b == tgPublicKeyFP {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("handshake: No fingerprint")
	}

	// (encoding) p_q_inner_data
	p, q := splitPQ(str2big(res.Pq))
	nonceSecond := GenerateNonce(32)
	nonceServer := res.ServerNonce
	innerData1 := (PQInnerDataDc{res.Pq, big2str(p), big2str(q), nonceFirst, nonceServer, nonceSecond, m.session.DcID}).encode()

	x = make([]byte, 255)
	copy(x[0:], sha1Sum(innerData1))
	copy(x[20:], innerData1)
	encryptedData1 := doRSAEncrypt(x)

	// (send) req_DH_params
	err = m.justSend(ReqDHParams{nonceFirst, nonceServer, big2str(p), big2str(q), tgPublicKeyFP, string(encryptedData1)})
	if err != nil {
		return WrapError(err)
	}

	// (parse) server_DH_params_{ok, fail}
	data, err = m.read()
	if err != nil {
		return WrapError(err)
	}
	dh, ok := data.(ServerDHParamsOk)
	if !ok {
		return fmt.Errorf("handshake: Need server_DH_params_ok, got %#v", data)
	}
	if !bytes.Equal(nonceFirst, dh.Nonce) {
		return fmt.Errorf("handshake: Wrong nonce")
	}
	if !bytes.Equal(nonceServer, dh.ServerNonce) {
		return fmt.Errorf("handshake: Wrong server_nonce")
	}
	t1 := make([]byte, 48)
	copy(t1[0:], nonceSecond)
	copy(t1[32:], nonceServer)
	hash1 := sha1Sum(t1)

	t2 := make([]byte, 48)
	copy(t2[0:], nonceServer)
	copy(t2[16:], nonceSecond)
	hash2 := sha1Sum(t2)

	t3 := make([]byte, 64)
	copy(t3[0:], nonceSecond)
	copy(t3[32:], nonceSecond)
	hash3 := sha1Sum(t3)

	tmpAESKey := make([]byte, 32)
	tmpAESIV := make([]byte, 32)

	copy(tmpAESKey[0:], hash1)
	copy(tmpAESKey[20:], hash2[0:12])

	copy(tmpAESIV[0:], hash2[12:20])
	copy(tmpAESIV[8:], hash3)
	copy(tmpAESIV[28:], nonceSecond[0:4])

	// (parse-thru) server_DH_inner_data
	decodedData, err := doAES256IGEdecrypt([]byte(dh.EncryptedAnswer), tmpAESKey, tmpAESIV)
	if err != nil {
		return WrapError(err)
	}
	innerbuf := NewDecodeBuf(decodedData[20:])
	data = innerbuf.Object()
	if innerbuf.err != nil {
		return WrapError(innerbuf.err)
	}
	dhi, ok := data.(ServerDHInnerData)
	if !ok {
		return fmt.Errorf("handshake: Need server_DH_inner_data")
	}
	if !bytes.Equal(nonceFirst, dhi.Nonce) {
		return fmt.Errorf("handshake: Wrong nonce")
	}
	if !bytes.Equal(nonceServer, dhi.ServerNonce) {
		return fmt.Errorf("handshake: Wrong server_nonce")
	}

	_, gB, gAb := makeGAB(dhi.G, str2big(dhi.GA), str2big(dhi.DhPrime))
	m.session.AuthKey = gAb.Bytes()
	if m.session.AuthKey[0] == 0 { //TODO: what?
		m.session.AuthKey = m.session.AuthKey[1:]
	}
	m.session.AuthKeyHash = sha1Sum(m.session.AuthKey)[12:20]
	t4 := make([]byte, 32+1+8)
	copy(t4[0:], nonceSecond)
	t4[32] = 1
	copy(t4[33:], sha1Sum(m.session.AuthKey)[0:8])
	nonceHash1 := sha1Sum(t4)[4:20]
	saltBuf := make([]byte, 8)
	copy(saltBuf, nonceSecond[:8])
	xor(saltBuf, nonceServer[:8])
	m.session.ServerSalt = int64(binary.LittleEndian.Uint64(saltBuf))

	// (encoding) client_DH_inner_data
	innerData2 := (ClientDHInnerData{nonceFirst, nonceServer, 0, big2str(gB)}).encode()
	x = make([]byte, 20+len(innerData2)+(16-((20+len(innerData2))%16))&15)
	copy(x[0:], sha1Sum(innerData2))
	copy(x[20:], innerData2)
	encryptedData2, err := doAES256IGEencrypt(x, tmpAESKey, tmpAESIV)

	// (send) set_client_DH_params
	err = m.justSend(SetClientDHParams{nonceFirst, nonceServer, string(encryptedData2)})
	if err != nil {
		return WrapError(err)
	}

	// (parse) dh_gen_{ok, retry, fail}
	data, err = m.read()
	if err != nil {
		return WrapError(err)
	}
	dhg, ok := data.(DhGenOk)
	if !ok {
		return fmt.Errorf("handshake: Need dh_gen_ok, got %#v", data)
	}
	if !bytes.Equal(nonceFirst, dhg.Nonce) {
		return fmt.Errorf("handshake: Wrong nonce")
	}
	if !bytes.Equal(nonceServer, dhg.ServerNonce) {
		return fmt.Errorf("handshake: Wrong server_nonce")
	}
	if !bytes.Equal(nonceHash1, dhg.NewNonceHash1) {
		return fmt.Errorf("handshake: Wrong new_nonce_hash1")
	}
	return nil
}
