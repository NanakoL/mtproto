package mtproto

import (
	"encoding/json"
	"fmt"
	"os"
)

type Session struct {
	DcID        int32  `json:"dc_id"`
	AuthKey     []byte `json:"auth_key"`
	AuthKeyHash []byte `json:"auth_key_hash"`
	ServerSalt  int64  `json:"server_salt"`
	Addr        string `json:"addr"`
	sessionId   int64
}

func (s *Session) Save(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return WrapError(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(s); err != nil {
		return WrapError(err)
	}
	return nil
}

func LoadSession(path string) (*Session, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("no session data")
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var sess Session
	if err := json.NewDecoder(f).Decode(&sess); err != nil {
		return nil, err
	}
	return &sess, nil
}
