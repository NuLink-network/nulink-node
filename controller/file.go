package controller

import (
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
)

func UploadFile(c *gin.Context) {
	req := &entity.UploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.UploadFile(req.AccountID, req.PolicyID, req.Files); code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessNil(c)
}

func CreatePolicyAndUploadFile(c *gin.Context) {
	req := &entity.CreatePolicyAndUploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.CreatePolicyAndUploadFile(req.AccountID, req.PolicyLabel, req.PolicyLabel, req.EncryptedPK, req.Files); code != resp.CodeSuccess {
		resp.Error(c, code)
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

	list, code := logic.GetFileList(req.AccountID, req.FileName, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, len(list))
}

func GetOthersFileList(c *gin.Context) {
	req := &entity.GetOthersFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	list, code := logic.GetOthersFileList(req.AccountID, req.FileName, req.Category, req.Format, req.Desc, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, len(list))
}

func DeleteFile(c *gin.Context) {
	req := &entity.DeleteFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	if code := logic.DeleteFile(req.AccountID, req.FileIDs); code != resp.CodeSuccess {
		resp.InternalServerError(c)
		return
	}
	resp.SuccessNil(c)
}
