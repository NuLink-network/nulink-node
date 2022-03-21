package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
)

func UploadFile(accountID uint64, addresses []string) error {
	f := &dao.File{}
	fs := make([]*dao.File, 0, len(addresses))
	for _, addr := range addresses {
		fs = append(fs, &dao.File{
			AccountID: accountID,
			Address:   addr,
		})
	}
	return f.BatchCreate(fs)
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