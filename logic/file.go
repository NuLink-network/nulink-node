package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
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
