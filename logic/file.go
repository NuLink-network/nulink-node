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
	files, err := file.Find(page, pageSize)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.GetFileListResponse, 0, 10)
	for _, f := range files {
		ret = append(ret, &entity.GetFileListResponse{
			AccountID: f.OwnerID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			CreatedAt: f.CreatedAt,
		})
	}
	return ret, nil
}

func GetOthersFileList(accountID string, fileName string, page, pageSize int) ([]*entity.GetOthersFileListResponse, error) {
	file := &dao.File{
		Name: fileName,
	}
	files, err := file.FindNotAccountID(accountID, page, pageSize)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.GetOthersFileListResponse, 0, len(files))
	for _, f := range files {
		ret = append(ret, &entity.GetOthersFileListResponse{
			AccountID: f.OwnerID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			CreatedAt: f.CreatedAt,
		})
	}
	return ret, nil
}

func DeleteFile(accountID string, fileIDs []string) error {
	// todo signature verification

	fs := make([]*dao.File, 0, len(fileIDs))
	for _, fid := range fileIDs {
		fs = append(fs, &dao.File{
			OwnerID: accountID,
			FileID:  fid,
		})
	}
	return dao.NewFile().BatchDelete(fs)
}
