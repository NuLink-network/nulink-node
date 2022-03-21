package dao

import (
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/resource/db"
)

type Account struct {
	ID           uint64         `gorm:"primarykey"`
	Name         string         `gorm:"column:name" json:"name" sql:"char(18)"` // // todo length?
	EthereumAddr string         `gorm:"column:ethereum_addr" json:"ethereum_addr" sql:"char(42)"`
	EncryptedPK  string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar()"` // todo length?
	VerifyPK     string         `gorm:"verify_pk:" json:"verify_pk" sql:"varchar()"`             // todo length?
	Signature    string         `gorm:"column:signature" json:"signature" sql:"varchar()"`       // todo length?
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewAccount(name, ethereumAddr, encryptedPK, verifyPK, signature string) *Account {
	return &Account{
		Name:         name,
		EthereumAddr: ethereumAddr,
		EncryptedPK:  encryptedPK,
		VerifyPK:     verifyPK,
		Signature:    signature,
	}
}

func (a *Account) TableName() string {
	return "account"
}

//func (a *Account) BeforeFind(tx *gorm.DB) (err error) {
//	tx.Where("deleted_at not null")
//	return
//}

func (a *Account) Create() (id uint64, err error) {
	err = db.GetDB().Create(a).Error
	return a.ID, err
}

func (a *Account) Get() (account *Account, err error) {
	err = db.GetDB().Where(a).Limit(1).Find(account).Error
	return account, err
}

func (a *Account) Count() (n int64, err error) {
	err = db.GetDB().Where(a).Count(&n).Error
	return n, err
}