package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	// 用户语法错误
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	// 没有权限
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
)