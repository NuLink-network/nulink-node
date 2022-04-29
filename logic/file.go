package logic

import (
	"errors"
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resp"
	"gorm.io/gorm"
)

func UploadFile(accountID, policyID string, files []entity.File) (code int) {
	p := &dao.Policy{PolicyID: policyID}
	policy, err := p.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodePolicyNotExist
		}
		return resp.CodeInternalServerError
	}
	if policy.Status != dao.PolicyStatusPublished {
		return resp.CodePolicyUnpublished
	}
	if policy.CreatorID != accountID {
		return resp.CodePolicyNotYours
	}

	fs := make([]*dao.File, 0, len(files))
	fps := make([]*dao.FilePolicy, 0, len(files))
	for _, f := range files {
		fs = append(fs, &dao.File{
			FileID:  f.ID,
			Name:    f.Name,
			Address: f.Address,
			//Owner:          fileOwner,
			OwnerID: accountID,
		})
		fps = append(fps, &dao.FilePolicy{
			FileID:   f.ID,
			PolicyID: policyID,
		})
	}
	if err := dao.NewFile().BatchCreate(fs); err != nil {
		return resp.CodeInternalServerError
	}
	if err := dao.NewFilePolicy().BatchCreate(fps); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func CreatePolicyAndUploadFile(accountID, policyID, policyLabel, encryptedPK string, files []entity.File) (code int) {
	p := &dao.Policy{PolicyID: policyID}
	isExist, err := p.IsExist()
	if err != nil {
		// todo log
		return resp.CodeInternalServerError
	}
	if isExist {
		return resp.CodePolicyIsExist
	}

	policy := &dao.Policy{
		PolicyID:    policyID,
		Label:       policyLabel,
		CreatorID:   accountID,
		EncryptedPK: encryptedPK,
	}

	fs := make([]*dao.File, 0, len(files))
	fps := make([]*dao.FilePolicy, 0, len(files))
	for _, f := range files {
		fs = append(fs, &dao.File{
			FileID:  f.ID,
			Name:    f.Name,
			Address: f.Address,
			//Owner:          fileOwner, // todo
			OwnerID: accountID,
		})
		fps = append(fps, &dao.FilePolicy{
			FileID:   f.ID,
			PolicyID: policyID,
		})
	}
	if err = dao.CreatePolicyAndFiles(policy, fs); err != nil {
		return resp.CodeInternalServerError
	}

	if err := dao.NewFilePolicy().BatchCreate(fps); err != nil {
		return resp.CodeInternalServerError
	}

	return resp.CodeSuccess
}

func GetFileList(accountID string, fileName string, page, pageSize int) ([]*entity.GetFileListResponse, error) {
	file := &dao.File{
		OwnerID: accountID,
		Name:    fileName,
	}
	files, err := file.Find(Paginate(page, pageSize))
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.GetFileListResponse, 0, 10)
	for _, f := range files {
		ret = append(ret, &entity.GetFileListResponse{
			FileID:    f.FileID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			Owner:     f.Owner,
			OwnerID:   f.OwnerID,
			CreatedAt: f.CreatedAt.Unix(),
		})
	}
	return ret, nil
}

func GetOthersFileList(accountID string, fileName string, page, pageSize int) ([]*entity.GetFileListResponse, error) {
	file := &dao.File{
		Name: fileName,
	}
	files, err := file.FindNotAccountID(accountID, page, pageSize)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.GetFileListResponse, 0, len(files))
	for _, f := range files {
		ret = append(ret, &entity.GetFileListResponse{
			FileID:    f.FileID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			Owner:     f.Owner,
			OwnerID:   f.OwnerID,
			CreatedAt: f.CreatedAt.Unix(),
		})
	}
	return ret, nil
}

func DeleteFile(accountID string, fileIDs []string) (code int) {
	// todo signature verification
	file := &dao.File{
		OwnerID: accountID,
	}

	// 删除文件和策略的关系
	// 删除文件的使用申请
	// 删除文件

	//files, err := file.FindByFileIDs(fileIDs, nil)
	//if err != nil {
	//	return resp.CodeInternalServerError
	//}

	if err := file.DeleteByFilesIDs(fileIDs); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}
