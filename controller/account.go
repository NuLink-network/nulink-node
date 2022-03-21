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

	response, err := logic.CreateAccount(req.Name, req.EthereumAddr, req.EncryptedPK, req.VerifyPK, req.Signature)
	if err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}

func GetAccount(c *gin.Context) {
	req := &entity.GetAccountRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	response, err := logic.GetAccount(req.AccountID)
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
