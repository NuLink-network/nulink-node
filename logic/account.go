package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
)

func CreateAccount(name, ethereumAddr, encryptedPK, verifyPK, signature string) (*entity.CreateAccountResponse, error) {
	account := dao.NewAccount(name, ethereumAddr, encryptedPK, verifyPK, signature)
	accountID, err := account.Create()
	if err != nil {
		return nil, err
	}
	return &entity.CreateAccountResponse{AccountID: accountID}, nil
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
		Signature:    a.Signature,
		CreatedAt:    a.CreatedAt,
	}, nil
}
