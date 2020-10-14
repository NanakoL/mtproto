package mtproto

import (
	"crypto/aes"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

const (
	tgPublicKeyN  = "24403446649145068056824081744112065346446136066297307473868293895086332508101251964919587745984311372853053253457835208829824428441874946556659953519213382748319518214765985662663680818277989736779506318868003755216402538945900388706898101286548187286716959100102939636333452457308619454821845196109544157601096359148241435922125602449263164512290854366930013825808102403072317738266383237191313714482187326643144603633877219028262697593882410403273959074350849923041765639673335775605842311578109726403165298875058941765362622936097839775380070572921007586266115476975819175319995527916042178582540628652481530373407"
	tgPublicKeyE  = 65537
	tgPublicKeyFP = -4344800451088585951 // 14101943622620965665 as int64
)

var telegramPublicKey rsa.PublicKey

func init() {
	telegramPublicKey.N, _ = new(big.Int).SetString(tgPublicKeyN, 10)
	telegramPublicKey.E = tgPublicKeyE
}

func sha1Sum(data []byte) []byte {
	r := sha1.Sum(data)
	return r[:]
}

func sha256Sum(buffers ...[]byte) []byte {
	s := sha256.New()
	for _, buf := range buffers {
		s.Write(buf)
	}
	return s.Sum(nil)
}

func xor(dst, src []byte) {
	for i := range dst {
		dst[i] = dst[i] ^ src[i]
	}
}

func bigIntPaddedBytes(num *big.Int, size int) []byte {
	buf := make([]byte, size)
	numBuf := num.Bytes()
	copy(buf[size-len(numBuf):], numBuf)
	return buf
}

func doRSAEncrypt(em []byte) []byte {
	z := make([]byte, 255)
	copy(z, em)

	c := new(big.Int)
	c.Exp(new(big.Int).SetBytes(z), big.NewInt(int64(telegramPublicKey.E)), telegramPublicKey.N)

	res := make([]byte, 256)
	copy(res, c.Bytes())

	return res
}

func splitPQ(pq *big.Int) (p1, p2 *big.Int) {
	value0 := big.NewInt(0)
	value1 := big.NewInt(1)
	value15 := big.NewInt(15)
	value17 := big.NewInt(17)
	rndmax := big.NewInt(0).SetBit(big.NewInt(0), 64, 1)

	what := big.NewInt(0).Set(pq)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := big.NewInt(0)
	i := 0
	for !(g.Cmp(value1) == 1 && g.Cmp(what) == -1) {
		q := big.NewInt(0).Rand(rnd, rndmax)
		q = q.And(q, value15)
		q = q.Add(q, value17)
		q = q.Mod(q, what)

		x := big.NewInt(0).Rand(rnd, rndmax)
		whatnext := big.NewInt(0).Sub(what, value1)
		x = x.Mod(x, whatnext)
		x = x.Add(x, value1)

		y := big.NewInt(0).Set(x)
		lim := 1 << (uint(i) + 18)
		j := 1
		flag := true

		for j < lim && flag {
			a := big.NewInt(0).Set(x)
			b := big.NewInt(0).Set(x)
			c := big.NewInt(0).Set(q)

			for b.Cmp(value0) == 1 {
				b2 := big.NewInt(0)
				if b2.And(b, value1).Cmp(value0) == 1 {
					c.Add(c, a)
					if c.Cmp(what) >= 0 {
						c.Sub(c, what)
					}
				}
				a.Add(a, a)
				if a.Cmp(what) >= 0 {
					a.Sub(a, what)
				}
				b.Rsh(b, 1)
			}
			x.Set(c)

			z := big.NewInt(0)
			if x.Cmp(y) == -1 {
				z.Add(what, x)
				z.Sub(z, y)
			} else {
				z.Sub(x, y)
			}
			g.GCD(nil, nil, z, what)

			if (j & (j - 1)) == 0 {
				y.Set(x)
			}
			j = j + 1

			if g.Cmp(value1) != 0 {
				flag = false
			}
		}
		i = i + 1
	}

	p1 = big.NewInt(0).Set(g)
	p2 = big.NewInt(0).Div(what, g)

	if p1.Cmp(p2) == 1 {
		p1, p2 = p2, p1
	}

	return
}

func makeGAB(g int32, gA, dhPrime *big.Int) (b, gB, gAb *big.Int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndMax := big.NewInt(0).SetBit(big.NewInt(0), 2048, 1)
	b = big.NewInt(0).Rand(rnd, rndMax)
	gB = big.NewInt(0).Exp(big.NewInt(int64(g)), b, dhPrime)
	gAb = big.NewInt(0).Exp(gA, b, dhPrime)

	return
}

func generateAES(msgKey, authKey []byte, decode bool) ([]byte, []byte) {
	var x int
	if decode {
		x = 8
	} else {
		x = 0
	}
	aesKey := make([]byte, 0, 32)
	aesIv := make([]byte, 0, 32)
	tA := make([]byte, 0, 48)
	tB := make([]byte, 0, 48)
	tC := make([]byte, 0, 48)
	tD := make([]byte, 0, 48)

	tA = append(tA, msgKey...)
	tA = append(tA, authKey[x:x+32]...)

	tB = append(tB, authKey[32+x:32+x+16]...)
	tB = append(tB, msgKey...)
	tB = append(tB, authKey[48+x:48+x+16]...)

	tC = append(tC, authKey[64+x:64+x+32]...)
	tC = append(tC, msgKey...)

	tD = append(tD, msgKey...)
	tD = append(tD, authKey[96+x:96+x+32]...)

	sha1A := sha1.Sum(tA)
	sha1B := sha1.Sum(tB)
	sha1C := sha1.Sum(tC)
	sha1D := sha1.Sum(tD)

	aesKey = append(aesKey, sha1A[0:8]...)
	aesKey = append(aesKey, sha1B[8:8+12]...)
	aesKey = append(aesKey, sha1C[4:4+12]...)

	aesIv = append(aesIv, sha1A[8:8+12]...)
	aesIv = append(aesIv, sha1B[0:8]...)
	aesIv = append(aesIv, sha1C[16:16+4]...)
	aesIv = append(aesIv, sha1D[0:8]...)

	return aesKey, aesIv
}

func doAES256IGEencrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("AES256IGE: data too small to encrypt: %d < %d", len(data), aes.BlockSize)
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("AES256IGE: data not divisible by block size: %d %% %d != 0", len(data), aes.BlockSize)
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(x, data[i:i+aes.BlockSize])
		block.Encrypt(t, x)
		xor(t, y)
		x, y = t, data[i:i+aes.BlockSize]
		copy(encrypted[i:], t)
		i += aes.BlockSize
	}

	return encrypted, nil
}

func doAES256IGEdecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("AES256IGE: data too small to decrypt: %d < %d", len(data), aes.BlockSize)
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("AES256IGE: data not divisible by block size: %d %% %d != 0", len(data), aes.BlockSize)
	}

	t := make([]byte, aes.BlockSize)
	x := make([]byte, aes.BlockSize)
	y := make([]byte, aes.BlockSize)
	copy(x, iv[:aes.BlockSize])
	copy(y, iv[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	i := 0
	for i < len(data) {
		xor(y, data[i:i+aes.BlockSize])
		block.Decrypt(t, y)
		xor(t, x)
		y, x = t, data[i:i+aes.BlockSize]
		copy(decrypted[i:], t)
		i += aes.BlockSize
	}

	return decrypted, nil

}

func calcInputCheckPasswordSRP(
	algo PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow,
	accPassword AccountPassword,
	password string,
	randFunc func([]byte) (int, error),
	logDebug func(string, ...interface{}),
) (TL, error) {
	logDebug(" --- SRP calculation start --- ")
	logDebug("algo.Salt1 (client): %#v", algo.Salt1)
	logDebug("algo.Salt2 (server): %#v", algo.Salt2)
	logDebug("algo.G:              %#v", algo.G)
	logDebug("algo.P:              %#v", algo.P)
	logDebug("accPassword.SrpB:    %#v", accPassword.SrpB)
	logDebug("accPassword.SrpID:   %#v", accPassword.SrpID)
	logDebug("password:            %#v", password)
	defer logDebug(" --- SRP calculation end --- ")

	if len(password) == 0 {
		return nil, fmt.Errorf("password is empty")
	}
	// TODO: check g and p
	// https://github.com/tdlib/td/blob/d9a18a064fa99130dc9214fb6131ba59e5660892/td/telegram/PasswordManager.cpp#L79
	// DhHandshake::check_config(g, p, ...

	clientSalt := algo.Salt1
	serverSalt := algo.Salt2
	gVal := algo.G
	pBuf := algo.P
	BBuf := accPassword.SrpB
	srpID := accPassword.SrpID

	if len(BBuf) != 256 {
		return nil, fmt.Errorf("wrong SrpB size, expected 256 bytes, got %d", len(BBuf))
	}

	gNum := new(big.Int).SetInt64(int64(gVal))
	gBuf := bigIntPaddedBytes(gNum, 256)
	pNum := new(big.Int).SetBytes(pBuf)
	BNum := new(big.Int).SetBytes(BBuf)

	if BNum.Cmp(pNum) != -1 {
		return nil, fmt.Errorf("expected SrpB < P, got: SrpB = %s, P = %s", BNum, pNum)
	}

	// calc_password_hash
	buf := sha256Sum(clientSalt, []byte(password), clientSalt)
	buf = sha256Sum(serverSalt, buf, serverSalt)
	hash := pbkdf2.Key(buf, clientSalt, 100000, 64, sha512.New)
	xBuf := sha256Sum(serverSalt, hash, serverSalt)
	xNum := new(big.Int).SetBytes(xBuf)
	// /calc_password_hash

	aBuf := make([]byte, 2048/8)
	_, err := randFunc(aBuf)
	if err != nil {
		return nil, err
	}
	aNum := new(big.Int).SetBytes(aBuf)
	logDebug("a: %s %#v", aNum, aBuf)

	ANum := new(big.Int).Exp(gNum, aNum, pNum)
	ABuf := bigIntPaddedBytes(ANum, 256)

	uBuf := sha256Sum(ABuf, BBuf)
	uNum := new(big.Int).SetBytes(uBuf)
	kBuf := sha256Sum(pBuf, gBuf)
	kNum := new(big.Int).SetBytes(kBuf)

	vNum := new(big.Int).Exp(gNum, xNum, pNum)
	kvNum := new(big.Int).Mul(kNum, vNum)
	kvNum.Mod(kvNum, pNum)
	tNum := new(big.Int).Sub(BNum, kvNum)
	if tNum.Sign() == -1 {
		logDebug("LZ!!! %s", tNum)
		tNum.Add(tNum, pNum)
	}
	expNum := new(big.Int).Mul(uNum, xNum)
	expNum.Add(expNum, aNum)

	SNum := new(big.Int).Exp(tNum, expNum, pNum)
	SBuf := bigIntPaddedBytes(SNum, 256)
	KBuf := sha256Sum(SBuf)

	h1 := sha256Sum(pBuf)
	h2 := sha256Sum(gBuf)
	xor(h1, h2)
	clientSaltHash := sha256Sum(clientSalt)
	serverSaltHash := sha256Sum(serverSalt)
	MBuf := sha256Sum(h1, clientSaltHash, serverSaltHash, ABuf, BBuf, KBuf)
	logDebug("srpID: %#v", srpID)
	logDebug("ABuf:  %#v", ABuf)
	logDebug("MBuf:  %#v", MBuf)

	return InputCheckPasswordSRP{SrpID: srpID, A: ABuf, M1: MBuf}, nil
}
