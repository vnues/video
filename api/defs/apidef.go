package defs

//requests
//credential 证书；凭据；国书
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}