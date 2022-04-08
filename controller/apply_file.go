package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/controller/resp"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
)

func ApplyFile(c *gin.Context) {
	req := &entity.ApplyFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	if err := logic.ApplyFile(req.FileIDs, req.ProposerID, req.Signature, req.StartAt, req.FinishAt); err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}

func ApplyFileList(c *gin.Context) {
	req := &entity.ApplyFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	response, err := logic.ApplyFileList(req.FileID, req.Status, req.ProposerAccountID, req.ProprietorAccountID, req.Paginate.Page, req.Paginate.PageSize)
	if err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}

func RevokeApply(c *gin.Context) {
	req := &entity.RevokeApplyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	if err := logic.RevokeApply(req.ProposerAccountID, req.ApplyIDs); err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}

func ApproveApply(c *gin.Context) {
	req := &entity.ApproveApplyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	if err := logic.ApproveApply(req.ApplyID); err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}

func RejectApply(c *gin.Context) {
	req := &entity.RejectApplyRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		// todo log
		resp.ParameterErr(c)
		return
	}

	if err := logic.RejectApply(req.ApplyID); err != nil {
		// todo log
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}
