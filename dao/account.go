package dao

import (
	"github.com/NuLink-network/nulink-node/utils"
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/resource/db"
)

type Account struct {
	ID           uint64         `gorm:"primarykey"`
	AccountID    string         `gorm:"column:account_id" json:"account_id" sql:"char(36)"`
	Name         string         `gorm:"column:name" json:"name" sql:"varchar(32)"`
	Avatar       string         `gorm:"column:avatar" json:"avatar" sql:"varchar(1024)"`
	UserSite     string         `gorm:"column:user_site" json:"user_site" sql:"varchar(1024)"`
	Twitter      string         `gorm:"column:twitter" json:"twitter" sql:"varchar(1024)"`
	Instagram    string         `gorm:"column:instagram" json:"instagram" sql:"varchar(1024)"`
	Facebook     string         `gorm:"column:facebook" json:"facebook" sql:"varchar(1024)"`
	Profile      string         `gorm:"column:profile" json:"profile" sql:"varchar(1024)"`
	EthereumAddr string         `gorm:"column:ethereum_addr" json:"ethereum_addr" sql:"char(42)"`
	EncryptedPK  string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar(256)"`
	VerifyPK     string         `gorm:"column:verify_pk" json:"verify_pk" sql:"varchar(256)"`
	Status       int8           `gorm:"column:status;default:1" json:"status" sql:"tinyint(4)" comment:"1: default, "`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

type AccountPart struct {
	Name         string
	Avatar       string
	EthereumAddr string
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) TableName() string {
	return "account"
}

func (a *Account) Create() (id uint64, err error) {
	err = db.GetDB().Create(a).Error
	return a.ID, err
}

func (a *Account) Get() (account *Account, err error) {
	err = db.GetDB().Where(a).First(&account).Error
	return account, err
}
func (a *Account) FindAny(ext *QueryExtra, pager Pager) (accounts []*Account, count int64, err error) {
	tx := db.GetDB().Where(a)
	if ext != nil {
		if ext.Conditions != nil {
			for k, v := range ext.Conditions {
				tx = tx.Where(k, v)
			}
		}
		if !utils.IsEmpty(ext.OrderStr) {
			tx = tx.Order(ext.OrderStr)
		}
	}

	if pager != nil {
		if err := tx.Model(a).Count(&count).Error; err != nil {
			return nil, count, err
		}
		if count == 0 {
			return nil, 0, nil
		}
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&accounts).Error
	return accounts, count, err
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

func (a *Account) Updates(new *Account) error {
	return db.GetDB().Where(a).Updates(new).Error
}

func (a *Account) FindAccountByAccountIDs(accountIDs []string) (accounts map[string]*AccountPart, err error) {
	query := &QueryExtra{
		Conditions: map[string]interface{}{
			"account_id in ?": accountIDs,
		},
	}
	as, _, err := a.FindAny(query, nil)
	if err != nil {
		return nil, err
	}
	accounts = make(map[string]*AccountPart, 0)
	for _, a := range as {
		accounts[a.AccountID] = &AccountPart{
			Name:         a.Name,
			Avatar:       a.Avatar,
			EthereumAddr: a.EthereumAddr,
		}
	}
	return accounts, nil
}
