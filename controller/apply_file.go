package controller

import (
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/NuLink-network/nulink-node/utils"
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
)

func ApplyFile(c *gin.Context) {
	req := &entity.ApplyFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.ApplyFile(req.FileIDs, req.ProposerID, req.StartAt, req.EndAt); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}

func ApplyFileList(c *gin.Context) {
	req := &entity.ApplyFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}
	if utils.IsEmpty(req.FileOwnerID) && utils.IsEmpty(req.ProposerID) {
		resp.ParameterErr(c)
		return
	}

	list, count, code := logic.ApplyFileList(req.FileID, req.Status, req.ProposerID, req.FileOwnerID, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, count)
}

func ApplyDetail(c *gin.Context) {
	req := &entity.ApplyDetailRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	ret, code := logic.ApplyDetail(req.ApplyID)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, ret)
}

func RevokeApply(c *gin.Context) {
	req := &entity.RevokeApplyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.RevokeApply(req.ProposerID, req.ApplyIDs); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}

func ApproveApply(c *gin.Context) {
	req := &entity.ApproveApplyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		log.Logger().Error(err.Error())
		return
	}
	if len(req.Remark) > 128 {
		resp.ParameterErr(c)
		return
	}

	if code := logic.ApproveApply(req.AccountID, req.ApplyID, req.Remark, req.Policy); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}

func RejectApply(c *gin.Context) {
	req := &entity.RejectApplyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}
	if len(req.Remark) > 128 {
		resp.ParameterErr(c)
		return
	}

	if code := logic.RejectApply(req.AccountID, req.ApplyID, req.Remark); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}
