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

func GetFileList(c *gin.Context) {
	req := &entity.GetFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	response, err := logic.GetFileList(req.AccountID, req.Address)
	if err != nil {
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}

func GetOthersFileList(c *gin.Context) {
	req := &entity.GetOthersFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	response, err := logic.GetOthersFileList(req.AccountID)
	if err != nil {
		resp.InternalServerError(c)
		return
	}
	resp.Success(c, response)
}

func DeleteFile(c *gin.Context) {
	req := &entity.DeleteFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if err := logic.DeleteFile(req.FileID, req.AccountID, req.Signature); err != nil {
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}
