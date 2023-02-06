package errorx

import "google.golang.org/grpc/status"

var (
	ErrDataBase   = status.Error(10002, "database error")
	ErrNoThisItem = status.Error(10003, "no this item")
	ErrOutOfTime  = status.Error(100004, "out of time")
	ErrMsgEncoder = status.Error(16003, "message encoding error")
	ErrMsgDecoder = status.Error(16004, "message decoding error")
	ErrMsgQ       = status.Error(16005, "message queue request failed")
)
