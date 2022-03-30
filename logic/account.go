package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
)

func CreateAccount(name, account, ethereumAddr, encryptedPK, verifyPK string) error {
	acc := dao.NewAccount(name, account, ethereumAddr, encryptedPK, verifyPK)
	_, err := acc.Create()
	if err != nil {
		return err
	}
	return nil
}

func GetAccount(account string) (*entity.GetAccountResponse, error) {
	acc := &dao.Account{Account: account}
	a, err := acc.Get()
	if err != nil {
		return nil, err
	}
	return &entity.GetAccountResponse{
		Name:         a.Name,
		Account:      a.Account,
		EthereumAddr: a.EthereumAddr,
		EncryptedPK:  a.EncryptedPK,
		VerifyPK:     a.VerifyPK,
		Status:       a.Status,
		CreatedAt:    a.CreatedAt,
	}, nil
}

func AccountIsExist(name, account, ethereumAddr, encryptedPK, verifyPK string) (*entity.AccountIsExistResponse, error) {
	acc := &dao.Account{
		Name:         name,
		Account:      account,
		EthereumAddr: ethereumAddr,
		EncryptedPK:  encryptedPK,
		VerifyPK:     verifyPK,
	}
	isExist, err := acc.IsExist()
	if err != nil {
		return nil, err
	}
	return &entity.AccountIsExistResponse{IsExist: isExist}, nil
}
