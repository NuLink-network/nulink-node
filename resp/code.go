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
	CodeApplyIrrevocable  = 3001
	CodeApplyRejected     = 3002
	CodeApplyApproved     = 3003
	CodePolicyUnpublished = 3004
)

const (
	CodePolicyNotExist = 4001
	CodeApplyNotExist  = 4002
)

const (
	CodePolicyIsExist = 4101
)

const (
	MsgPolicyIsExist     = "policy already exists"
	MsgPolicyNotExist    = "policy does not exist"
	MsgApplyNotExist     = "apply does not exist"
	MsgApplyIrrevocable  = "apply is irrevocable"
	MsgApplyApproved     = "apply approved"
	MsgApplyRejected     = "apply rejected"
	MsgPolicyUnpublished = "policy unpublished"
)

var code2msg = map[int]string{
	CodeSuccess:           MsgSuccess,
	CodeApplyIrrevocable:  MsgApplyIrrevocable,
	CodeApplyApproved:     MsgApplyApproved,
	CodeApplyRejected:     MsgApplyRejected,
	CodePolicyUnpublished: MsgPolicyUnpublished,
	CodeParameterErr:      MsgParameterErr,
	CodePolicyNotExist:    MsgPolicyNotExist,
	CodeApplyNotExist:     MsgApplyNotExist,
	CodePolicyIsExist:     MsgPolicyIsExist,
}
