package main

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized   = Err{"api not recognized,bad request", "001"} //请求方式错误
	ErrorRequestBodyParseFailed = Err{"request body is not correct", "002"}    //请求内容错误
	ErrorInternalFaults=Err{"internal service error","003"}
)
