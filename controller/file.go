package controller

import (
	resp2 "github.com/NuLink-network/nulink-node/resp"
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
)

func UploadFile(c *gin.Context) {
	req := &entity.UploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp2.ParameterErr(c)
		return
	}

	if code := logic.UploadFile(req.AccountID, req.PolicyID, req.Files); code != resp2.CodeSuccess {
		resp2.Error(c, code)
		return
	}
	resp2.SuccessNil(c)
}

func CreatePolicyAndUploadFile(c *gin.Context) {
	req := &entity.CreatePolicyAndUploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp2.ParameterErr(c)
		return
	}

	if code := logic.CreatePolicyAndUploadFile(req.AccountID, req.PolicyID, req.PolicyLabel, req.EncryptedPK, req.Files); code != resp2.CodeSuccess {
		resp2.Error(c, code)
		return
	}
	resp2.SuccessNil(c)
}

func GetFileList(c *gin.Context) {
	req := &entity.GetFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp2.ParameterErr(c)
		return
	}

	response, err := logic.GetFileList(req.AccountID, req.FileName, req.Paginate.Page, req.Paginate.PageSize)
	if err != nil {
		resp2.InternalServerError(c)
		return
	}
	resp2.Success(c, response)
}

func GetOthersFileList(c *gin.Context) {
	req := &entity.GetOthersFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp2.ParameterErr(c)
		return
	}

	response, err := logic.GetOthersFileList(req.AccountID, req.FileName, req.Paginate.Page, req.Paginate.PageSize)
	if err != nil {
		resp2.InternalServerError(c)
		return
	}
	resp2.Success(c, response)
}

func DeleteFile(c *gin.Context) {
	req := &entity.DeleteFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp2.ParameterErr(c)
		return
	}

	if err := logic.DeleteFile(req.AccountID, req.FileIDs); err != nil {
		resp2.InternalServerError(c)
		return
	}
	resp2.SuccessNil(c)
}
