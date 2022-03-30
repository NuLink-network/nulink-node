package controller

import (
	"github.com/NuLink-network/nulink-node/controller/resp"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}
	if len(req.Account) == 0 {
		resp.ParameterErr(c)
		return
	}

	response, err := logic.GetAccount(req.Account)
	if err != nil {
		// todo log
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

	response, err := logic.AccountIsExist(req.AccountID, req.Name, req.EthereumAddr, req.EncryptedPK, req.VerifyPK)
	if err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}
