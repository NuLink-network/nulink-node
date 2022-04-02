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

func GetFileList(accountID uint64, address string) ([]*entity.GetFileListResponse, error) {
	file := &dao.File{
		AccountID: accountID,
		Address:   address,
	}
	files, err := file.Find()
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.GetFileListResponse, 0, 10)
	for _, f := range files {
		resp = append(resp, &entity.GetFileListResponse{
			AccountID: f.AccountID,
			Address:   f.Address,
			CreatedAt: f.CreatedAt,
		})
	}
	return resp, nil
}

func GetOthersFileList(accountID uint64) ([]*entity.GetOthersFileListResponse, error) {
	file := &dao.File{}
	files, err := file.FindNotAccountID(accountID)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.GetOthersFileListResponse, 0, len(files))
	for _, f := range files {
		resp = append(resp, &entity.GetOthersFileListResponse{
			AccountID: f.AccountID,
			Address:   f.Address,
			CreatedAt: f.CreatedAt,
		})
	}
	return resp, nil
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
