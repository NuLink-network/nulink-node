package dao

import (
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/resource/db"
)

type Account struct {
	ID           uint64         `gorm:"primarykey"`
	Account      string         `gorm:"column:account_id" json:"account_id" sql:"char(36)"`
	Name         string         `gorm:"column:name" json:"name" sql:"varchar(32)"` // // todo length?
	EthereumAddr string         `gorm:"column:ethereum_addr" json:"ethereum_addr" sql:"char(42)"`
	EncryptedPK  string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar(256)"` // todo length?
	VerifyPK     string         `gorm:"verify_pk:" json:"verify_pk" sql:"varchar(256)"`             // todo length?
	Status       int8           `gorm:"column:status;default:1" json:"status" sql:"tinyint(4)" comment:"1: default, "`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewAccount(name, account, ethereumAddr, encryptedPK, verifyPK string) *Account {
	return &Account{
		Name:         name,
		Account:      account,
		EthereumAddr: ethereumAddr,
		EncryptedPK:  encryptedPK,
		VerifyPK:     verifyPK,
	}
}

func (a *Account) TableName() string {
	return "account"
}

//func (a *AccountID) BeforeFind(tx *gorm.DB) (err error) {
//	tx.Where("deleted_at not null")
//	return
//}

func (a *Account) Create() (id uint64, err error) {
	err = db.GetDB().Create(a).Error
	return a.ID, err
}

func (a *Account) Get() (account *Account, err error) {
	err = db.GetDB().Where(a).First(&account).Error
	return account, err
}

func (a *Account) Count() (n int64, err error) {
	err = db.GetDB().Where(a).Count(&n).Error
	return n, err
}

func (a *Account) IsExist() (isExist bool, err error) {
	acc := Account{}
	if err = db.GetDB().Where(a).Limit(1).First(&acc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
