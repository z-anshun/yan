package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"` //系统内部查错误
}
type ErrorResponse struct {
	HttpSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSc: 400, Error: Err{Error: "Request body is failed",
		ErrorCode: "001"}} //消息体无法解析
	ErrorNotAuthUser = ErrorResponse{HttpSc: 401, Error: Err{Error: "User authentication failed.",
		ErrorCode: "002"}} //验证不通过
	ErrorDBError = ErrorResponse{HttpSc: 500, Error: Err{Error: "DB ops failed.",
		ErrorCode: "003"}}
	ErrorInternalFaults = ErrorResponse{HttpSc: 500, Error: Err{"Internal service error.",
		"004"}}
	ErrorNoOne= ErrorResponse{HttpSc: 500, Error: Err{"Select user error.",
		"005"}}//未找到
)
