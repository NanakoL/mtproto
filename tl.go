package mtproto

type TL interface {
	encode() []byte
}

type TLReq interface {
	TL
	decodeResponse(*DecodeBuf) TL
}

type MsgContainer struct {
	Items []MTMessage
}

func (e MsgContainer) encode() []byte { return nil }

type MTMessage struct {
	MsgID int64
	SeqNo int32
	Size  int32
	Data  TL
}

type RpcResult struct {
	reqMsgID int64
	obj      TL
}

func (e RpcResult) encode() []byte { return nil }

type VectorInt []int32

func (e VectorInt) encode() []byte { return nil }

type VectorLong []int64

func (e VectorLong) encode() []byte { return nil }

type VectorObject []TL

func (e VectorObject) encode() []byte { return nil }
