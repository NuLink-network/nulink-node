package controller

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func CreateAccount(c *gin.Context) {
	req := &entity.CreateAccountRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	if err := logic.CreateAccount(req.Name, req.Account, req.EthereumAddr, req.EncryptedPK, req.VerifyPK); err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}

func GetAccount(c *gin.Context) {
	req := &entity.GetAccountRequest{}
	//c.ShouldBindQuery()
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}
	if len(strings.TrimSpace(req.AccountID)) == 0 {
		resp.ParameterErr(c)
		return
	}

	response, err := logic.GetAccount(req.AccountID)
	if err != nil {
		log.Logger().WithField("account_id", req.AccountID).Error("get account failed, error: ", err)
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}

func AccountIsExist(c *gin.Context) {
	req := &entity.AccountIsExistRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}
	if len(req.Account) == 0 {
		resp.ParameterErr(c)
		return
	}

	response, err := logic.AccountIsExist(req.Name, req.Account, req.EthereumAddr, req.EncryptedPK, req.VerifyPK)
	if err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}
