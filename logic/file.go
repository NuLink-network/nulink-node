package logic

import (
	"github.com/NuLink-network/nulink-node/controller/resp"
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
)

func UploadFile(accountID, fileOwner, policyID string, files []entity.File) (code int) {
	p := &dao.Policy{PolicyID: policyID}
	isExist, err := p.IsExist()
	if err != nil {
		// todo log
		return resp.CodeInternalServerError
	}
	if !isExist {
		return resp.CodePolicyNotExist
	}

	file := &dao.File{}
	fs := make([]*dao.File, 0, len(files))
	for _, f := range files {
		fs = append(fs, &dao.File{
			Name:           f.Name,
			Address:        f.Address,
			Owner:          fileOwner,
			OwnerAccountID: accountID,
		})
	}
	if err := file.BatchCreate(fs); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func GetFileList(accountID string, fileName string, page, pageSize int) ([]*entity.GetFileListResponse, error) {
	file := &dao.File{
		OwnerAccountID: accountID,
		Name:           fileName,
	}
	files, err := file.Find(page, pageSize)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.GetFileListResponse, 0, 10)
	for _, f := range files {
		ret = append(ret, &entity.GetFileListResponse{
			AccountID: f.OwnerAccountID,
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
			AccountID: f.OwnerAccountID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			CreatedAt: f.CreatedAt,
		})
	}
	return ret, nil
}

func DeleteFile(accountID, fileID uint64, signature string) error {
	file := &dao.File{
		ID:        fileID,
		AccountID: accountID,
		Signature: signature,
	}
	// todo signature verification
	return file.Delete()
}
