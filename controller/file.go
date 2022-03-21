package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/controller/resp"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
)

func UploadFile(c *gin.Context) {
	req := &entity.UploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}
	if len(req.Addresses) < 1 {
		resp.ParameterErr(c)
		return
	}

	if err := logic.UploadFile(req.AccountID, req.Addresses); err != nil {
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}
