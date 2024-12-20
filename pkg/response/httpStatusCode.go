package response

const (
	ErrorCodeSuccess      = 20001 // success
	ErrorCodeParamInvalid = 20003 // Email is invalid
	ErrorInvalidToken     = 30001 // token is invalid
)

// message
var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorInvalidToken:     "token is invalid",
}
