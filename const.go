package mtproto

const (
	crcVector       = 0x1cb5c415
	crcRpcResult    = 0xf35c6d01
	crcMsgContainer = 0x73f1f8dc
	crcGzipPacked   = 0x3072cfa1
)

const (
	ErrSeeOther     = int32(303)
	ErrBadRequest   = int32(400)
	ErrUnauthorized = int32(401)
	ErrForbidden    = int32(403)
	ErrNotFound     = int32(404)
	ErrFlood        = int32(420)
	ErrInternal     = int32(500)
)
