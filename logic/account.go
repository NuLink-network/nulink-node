package logic

import (
	"errors"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func CreateAccount(name, account, ethereumAddr, encryptedPK, verifyPK string) (code int) {
	acc := dao.NewAccount(name, account, ethereumAddr, encryptedPK, verifyPK)
	_, err := acc.Create()
	if err != nil {
		log.Logger().WithField("account", acc).WithField("error", err).Error("create account failed")
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func GetAccount(accountID string) (*entity.GetAccountResponse, int) {
	acc := &dao.Account{AccountID: accountID}
	a, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeAccountNotExist
		}
		log.Logger().WithField("account", a).WithField("error", err).Error("get account failed")
		return nil, resp.CodeInternalServerError
	}
	return &entity.GetAccountResponse{
		Name:         a.Name,
		AccountID:    a.AccountID,
		EthereumAddr: a.EthereumAddr,
		EncryptedPK:  a.EncryptedPK,
		VerifyPK:     a.VerifyPK,
		Status:       a.Status,
		CreatedAt:    a.CreatedAt.Unix(),
	}, resp.CodeSuccess
}

func AccountIsExist(name, account, ethereumAddr, encryptedPK, verifyPK string) (*entity.AccountIsExistResponse, int) {
	acc := &dao.Account{
		Name:         name,
		AccountID:    account,
		EthereumAddr: ethereumAddr,
		EncryptedPK:  encryptedPK,
		VerifyPK:     verifyPK,
	}
	isExist, err := acc.IsExist()
	if err != nil {
		log.Logger().WithField("account", acc).WithField("error", err).Error("query account failed")
		return nil, resp.CodeInternalServerError
	}
	return &entity.AccountIsExistResponse{IsExist: isExist}, resp.CodeSuccess
}
