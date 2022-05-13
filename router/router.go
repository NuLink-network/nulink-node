package router

import (
	"github.com/NuLink-network/nulink-node/controller"
	"github.com/gin-gonic/gin"
)

const (
	AccountGroup = "/account"
	LabelGroup   = "/label"
	PolicyGroup  = "/policy"
	FileGroup    = "/file"
	ApplyGroup   = "/apply"
)

func Register(engine *gin.Engine) {
	engine.GET("/ping", controller.Ping)

	account := engine.Group(AccountGroup)
	account.POST("/create", controller.CreateAccount)
	account.POST("/get", controller.GetAccount)
	account.POST("/isexist", controller.AccountIsExist)

	label := engine.Group(LabelGroup)
	label.POST("/list", controller.PolicyLabelList)

	policy := engine.Group(PolicyGroup)
	policy.POST("/revoke", controller.RevokePolicy)
	policy.POST("/list", controller.PolicyList)
	policy.POST("/file-detail-list", controller.FileDetailList)

	file := engine.Group(FileGroup)
	file.POST("/upload", controller.UploadFile)
	file.POST("/create-policy-and-upload", controller.CreatePolicyAndUploadFile)
	file.POST("/list", controller.GetFileList)
	file.POST("/others-list", controller.GetOthersFileList)
	file.POST("/delete", controller.DeleteFile)
	file.POST("/detail", controller.FileDetail)

	apply := engine.Group(ApplyGroup)
	apply.POST("/file", controller.ApplyFile)
	apply.POST("/list", controller.ApplyFileList)
	apply.POST("/revoke", controller.RevokeApply)
	apply.POST("/approve", controller.ApproveApply)
	apply.POST("/reject", controller.RejectApply)
}
