package errorx

const (
	defaultErrCode  = 1001
	RPCErrCode      = 1002
	NotFoundCode    = 1003
	ParamErrCode    = 1004
	InternalErrCode = 1005
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}
func NewDefaultCodeError(code int, msg string) error {
	return &CodeError{
		Code: defaultErrCode,
		Msg:  msg,
	}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

var _ error = (*CodeError)(nil)
