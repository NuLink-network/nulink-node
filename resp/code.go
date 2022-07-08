package resp

const (
	CodeSuccess             = 2000
	CodeUnauthorized        = 3000
	CodeParameterErr        = 4000
	CodeInternalServerError = 5000
)

const (
	MsgSuccess             = "Success"
	MsgUnauthorized        = "Unauthorized Operation"
	MsgInternalServerError = "Internal Server Error"
	MsgParameterErr        = "Invalid Parameter"
)

const (
	CodePolicyLabelNotExist = 4001
	CodePolicyLabelNotYours = 4002
	CodePolicyNotExist      = 4003
	CodePolicyNotYours      = 4004
	CodePolicyIsExist       = 4005
	CodeApplyNotExist       = 4006
	CodeAccountNotExist     = 4007
	CodeFileNotExist        = 4008
	CodeDuplicateFilename   = 4009
	CodeAccountIsExist      = 4010
	CodePolicyLabelIsExist  = 4011
)

const (
	CodeApplyIrrevocable  = 4101
	CodeApplyRejected     = 4102
	CodeApplyApproved     = 4103
	CodeApplyUnApproved   = 4104
	CodePolicyUnpublished = 4106
	CodeFileApplied       = 4107
)

const (
	MsgPolicyLabelNotExist = "policy label does not exist"
	MsgPolicyLabelNotYours = "please choose your own policy label"
	MsgPolicyIsExist       = "policy already exists"
	MsgPolicyNotExist      = "policy does not exist"
	MsgPolicyNotYours      = "please choose your own published policy"
	MsgApplyNotExist       = "apply does not exist"
	MsgAccountNotExist     = "account does not exist"
	MsgFileNotExist        = "file does not exist"
	MsgDuplicateFilename   = "duplicate filename"
	MsgAccountIsExist      = "account already exists"
	MsgFileApplied         = "you have already applied for file and cannot apply again"
	MsgApplyIrrevocable    = "apply is irrevocable"
	MsgApplyApproved       = "apply approved"
	MsgApplyRejected       = "apply rejected"
	MsgApplyUnapproved     = "apply unapproved"
	MsgPolicyUnpublished   = "policy unpublished"
	MsgPolicyLabelIsExist  = "policy label already exists"
)

var code2msg = map[int]string{
	CodeSuccess:             MsgSuccess,
	CodeUnauthorized:        MsgUnauthorized,
	CodeParameterErr:        MsgParameterErr,
	CodeInternalServerError: MsgInternalServerError,
	CodePolicyLabelNotExist: MsgPolicyLabelNotExist,
	CodePolicyLabelNotYours: MsgPolicyLabelNotYours,
	CodePolicyNotExist:      MsgPolicyNotExist,
	CodePolicyNotYours:      MsgPolicyNotYours,
	CodeApplyNotExist:       MsgApplyNotExist,
	CodeAccountNotExist:     MsgAccountNotExist,
	CodeFileNotExist:        MsgFileNotExist,
	CodeDuplicateFilename:   MsgDuplicateFilename,
	CodeAccountIsExist:      MsgAccountIsExist,
	CodeFileApplied:         MsgFileApplied,
	CodeApplyIrrevocable:    MsgApplyIrrevocable,
	CodeApplyApproved:       MsgApplyApproved,
	CodeApplyUnApproved:     MsgApplyUnapproved,
	CodeApplyRejected:       MsgApplyRejected,
	CodePolicyIsExist:       MsgPolicyIsExist,
	CodePolicyUnpublished:   MsgPolicyUnpublished,
	CodePolicyLabelIsExist:  MsgPolicyLabelIsExist,
}
