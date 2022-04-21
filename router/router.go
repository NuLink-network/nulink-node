package router

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/controller"
)

const (
	AccountGroup = "/account"
	PolicyGroup  = "/policy"
	FileGroup    = "/file"
)

func Register(engine *gin.Engine) {
	account := engine.Group(AccountGroup)
	account.POST("/create", controller.CreateAccount)
	account.GET("/get", controller.GetAccount)
	account.GET("/isexist", controller.AccountIsExist)

	policy := engine.Group(PolicyGroup)
	policy.POST("/revoke", controller.RevokePolicy)
	policy.GET("/list", controller.PolicyList)

	file := engine.Group(FileGroup)
	file.POST("/upload", controller.UploadFile)
	file.POST("/create-policy-and-upload", controller.CreatePolicyAndUploadFile)
	file.GET("/list", controller.GetFileList)
	file.GET("/others-list", controller.GetOthersFileList)
}
