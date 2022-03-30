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

func GetAccount(accountID uint64) (*entity.GetAccountResponse, error) {
	account := &dao.Account{ID: accountID}
	a, err := account.Get()
	if err != nil {
		return nil, err
	}
	return &entity.GetAccountResponse{
		ID:           a.ID,
		Name:         a.Name,
		EthereumAddr: a.EthereumAddr,
		EncryptedPK:  a.EncryptedPK,
		VerifyPK:     a.VerifyPK,
		CreatedAt:    a.CreatedAt,
	}, nil
}

func AccountIsExist(accountID uint64, name, ethereumAddr, encryptedPK, verifyPK string) (*entity.AccountIsExistResponse, error) {
	account := &dao.Account{
		ID:           accountID,
		Name:         name,
		EthereumAddr: ethereumAddr,
		EncryptedPK:  encryptedPK,
		VerifyPK:     verifyPK,
	}
	n, err := account.Count()
	if err != nil {
		return nil, err
	}
	return &entity.AccountIsExistResponse{IsExist: n > 0}, nil
}
