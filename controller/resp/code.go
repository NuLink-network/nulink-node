package resp

const (
	CodeSuccess             = 2000
	CodeParameterErr        = 4000
	CodeInternalServerError = 5000
)

const (
	MsgSuccess             = "Success"
	MsgInternalServerError = "Internal Server Error"
	MsgParameterErr        = "Invalid Parameter"
)

const (
	CodePolicyIsExist  = 4001
	CodePolicyNotExist = 4002
)

const (
	MsgPolicyIsExist  = "policy already exists"
	MsgPolicyNotExist = "policy does not exist"
)

var code2msg = map[int]string{
	CodeSuccess:        MsgSuccess,
	CodeParameterErr:   MsgParameterErr,
	CodePolicyIsExist:  MsgPolicyIsExist,
	CodePolicyNotExist: MsgPolicyNotExist,
}
