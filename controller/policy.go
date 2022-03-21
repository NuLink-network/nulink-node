package controller

import (
	"github.com/NuLink-network/nulink-node/controller/resp"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/gin-gonic/gin"
)

func CreatePolicy(c *gin.Context) {
	req := &entity.CreatePolicyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	err := logic.CreatePolicy(req.AccountID, req.Label, req.EncryptedPK, req.VerifyPK, req.Signature)
	if err != nil {
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}
