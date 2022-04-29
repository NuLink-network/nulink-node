package controller

import (
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/gin-gonic/gin"
)

func RevokePolicy(c *gin.Context) {
	req := &entity.RevokePolicyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.RevokePolicy(req.AccountID, req.PolicyID); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}

func PolicyList(c *gin.Context) {
	req := &entity.PolicyListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	response, code := logic.PolicyList(req.PolicyID, req.CreatorID, req.ConsumerID, req.Status, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, response)

}

func FileDetailList(c *gin.Context) {
	req := &entity.FileDetailListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Logger().WithField("error", err).Error("params error")
		resp.ParameterErr(c)
		return
	}

	response, code := logic.FileDetailList(req.CreatorID, req.ConsumerID, req.Status, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, response)

}
