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
