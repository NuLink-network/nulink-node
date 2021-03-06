package controller

import (
	"strings"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/NuLink-network/nulink-node/utils"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	req := &entity.UploadFileRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	fileNames := make([]string, 0, len(req.Files))
	for _, f := range req.Files {
		fileNames = append(fileNames, f.Name)
	}
	names, code := logic.DuplicateFilename(req.AccountID, fileNames)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	if len(names) > 0 {
		ret := map[string][]string{
			"filenames": names,
		}
		resp.Resp(c, resp.CodeDuplicateFilename, ret)
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

	fileNames := make([]string, 0, len(req.Files))
	for _, f := range req.Files {
		fileNames = append(fileNames, f.Name)
	}
	names, code := logic.DuplicateFilename(req.AccountID, fileNames)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	if len(names) > 0 {
		ret := map[string][]string{
			"filenames": names,
		}
		resp.Resp(c, resp.CodeDuplicateFilename, ret)
		return
	}

	if code := logic.CreatePolicyAndUploadFile(req.AccountID, req.PolicyLabelID, req.PolicyLabel, req.EncryptedPK, req.Files); code != resp.CodeSuccess {
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

	list, count, code := logic.GetFileList(req.AccountID, req.FileName, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, count)
}

func GetOthersFileList(c *gin.Context) {
	req := &entity.GetOthersFileListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}
	category := strings.ToLower(req.Category)
	if !utils.IsEmpty(category) {
		if _, ok := utils.FileCategory[category]; !ok {
			resp.ParameterErr(c)
			return
		}
	}
	format := strings.ToLower(req.Format)
	if !utils.IsEmpty(format) {
		if _, ok := utils.FileFormat[format]; !ok {
			resp.ParameterErr(c)
			return
		}
	}

	list, count, code := logic.GetOthersFileList(req.AccountID, req.Include, req.FileName, category, format, req.Desc, req.Paginate.Page, req.Paginate.PageSize)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, count)
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

func FileDetail(c *gin.Context) {
	req := &entity.FileDetailRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	ret, code := logic.FileDetail(req.FileID, req.ConsumerID)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.Success(c, ret)
}
