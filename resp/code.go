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
	CodePolicyNotExist  = 4001
	CodePolicyNotYours  = 4002
	CodeApplyNotExist   = 4003
	CodeAccountNotExist = 4004
	CodeFileNotExist    = 4005
)

const (
	CodeApplyIrrevocable  = 4101
	CodeApplyRejected     = 4102
	CodeApplyApproved     = 4103
	CodePolicyIsExist     = 4104
	CodePolicyUnpublished = 4105
)

const (
	MsgPolicyIsExist     = "policy already exists"
	MsgPolicyNotExist    = "policy does not exist"
	MsgPolicyNotYours    = "please choose your own published policy"
	MsgApplyNotExist     = "apply does not exist"
	MsgAccountNotExist   = "account does not exist"
	MsgFileNotExist      = "file does not exist"
	MsgApplyIrrevocable  = "apply is irrevocable"
	MsgApplyApproved     = "apply approved"
	MsgApplyRejected     = "apply rejected"
	MsgPolicyUnpublished = "policy unpublished"
)

var code2msg = map[int]string{
	CodeSuccess:             MsgSuccess,
	CodeParameterErr:        MsgParameterErr,
	CodeInternalServerError: MsgInternalServerError,
	CodePolicyNotExist:      MsgPolicyNotExist,
	CodePolicyNotYours:      MsgPolicyNotYours,
	CodeApplyNotExist:       MsgApplyNotExist,
	CodeAccountNotExist:     MsgAccountNotExist,
	CodeFileNotExist:        MsgFileNotExist,
	CodeApplyIrrevocable:    MsgApplyIrrevocable,
	CodeApplyApproved:       MsgApplyApproved,
	CodeApplyRejected:       MsgApplyRejected,
	CodePolicyIsExist:       MsgPolicyIsExist,
	CodePolicyUnpublished:   MsgPolicyUnpublished,
}
