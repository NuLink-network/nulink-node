package logic

import (
	"errors"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/NuLink-network/nulink-node/utils"
)

func CreateAccount(name, account, ethereumAddr, encryptedPK, verifyPK string) (code int) {
	acc := &dao.Account{
		Name:         name,
		AccountID:    account,
		EthereumAddr: ethereumAddr,
		EncryptedPK:  encryptedPK,
		VerifyPK:     verifyPK,
	}
	_, err := acc.Create()
	if err != nil {
		if utils.IsDuplicateError(err.Error()) {
			return resp.CodeAccountIsExist
		}
		log.Logger().WithField("account", utils.JSON(acc)).WithField("error", err).Error("create account failed")
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
		Avatar:       a.Avatar,
		UserSite:     a.UserSite,
		Twitter:      a.Twitter,
		Instagram:    a.Instagram,
		Facebook:     a.Facebook,
		Profile:      a.Profile,
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

func UpdateAccount(accountID string, update *entity.UpdateAccountRequest) (code int) {
	acc := &dao.Account{AccountID: accountID}
	newAccount := &dao.Account{
		Name:      update.Name,
		Avatar:    update.Avatar,
		UserSite:  update.UserSite,
		Twitter:   update.Twitter,
		Instagram: update.Instagram,
		Facebook:  update.Facebook,
		Profile:   update.Profile,
	}
	if err := acc.Updates(newAccount); err != nil {
		log.Logger().WithField("account_id", accountID).WithField("account", newAccount).WithField("error", err).Error("update account failed")
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}
