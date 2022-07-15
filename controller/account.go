package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/NuLink-network/nulink-node/utils"
)

func CreateAccount(c *gin.Context) {
	req := &entity.CreateAccountRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.CreateAccount(req.Name, req.AccountID, req.EthereumAddr, req.EncryptedPK, req.VerifyPK); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}

func GetAccount(c *gin.Context) {
	req := &entity.GetAccountRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	response, code := logic.GetAccount(req.AccountID)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, response)
}

func AccountIsExist(c *gin.Context) {
	req := &entity.AccountIsExistRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	response, code := logic.AccountIsExist(req.Name, req.AccountID, req.EthereumAddr, req.EncryptedPK, req.VerifyPK)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, response)
}

func UpdateAccount(c *gin.Context) {
	req := &entity.UpdateAccountRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if utils.IsEmpty(req.Avatar) && utils.IsEmpty(req.Name) &&
		utils.IsEmpty(req.UserSite) && utils.IsEmpty(req.Twitter) &&
		utils.IsEmpty(req.Instagram) && utils.IsEmpty(req.Facebook) && utils.IsEmpty(req.Profile) {
		resp.SuccessNil(c)
		return
	}

	if code := logic.UpdateAccount(req.AccountID, req); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}
