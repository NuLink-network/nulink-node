package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

type Policy struct {
	ID               uint64         `gorm:"primarykey"`
	Hrac             uint64         `gorm:"column:hrac" json:"hrac" sql:"varchar()"`
	Label            string         `gorm:"column:label" json:"label" sql:"varchar()"` // todo length?
	PolicyID         string         `gorm:"column:policy_id" json:"policy_id" sql:"char(36)"`
	Publisher        string         `gorm:"column:publisher" json:"publisher" sql:"varchar()"` // todo length?
	PublisherID      string         `gorm:"column:publisher_id" json:"publisher_id" sql:"char(36)"`
	Consumer         string         `gorm:"column:consumer" json:"consumer" sql:"varchar()"` // todo length?
	ConsumerID       string         `gorm:"column:consumer_id" json:"consumer_id" sql:"char(36)"`
	EncryptedPK      string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar()"`           // todo length?
	EncryptedAddress string         `gorm:"column:encrypted_address" json:"encrypted_address" sql:"varchar()"` // todo length?
	VerifyPK         string         `gorm:"verify_pk:" json:"verify_pk" sql:"varchar()"`                       // todo length?
	Status           int8           `gorm:"column:status" json:"status" sql:"tinyint(4)" comment:"1: default, "`
	Gas              string         `gorm:"column:gas" json:"gas" sql:"varchar()"`
	TxHash           string         `gorm:"column:tx_hash" json:"tx_hash" sql:"char(66)" comment:"1: default, "`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewPolicy() *Policy {
	return &Policy{}
}

func (p *Policy) TableName() string {
	return "policy"
}

func (p *Policy) Create() (id uint64, err error) {
	err = db.GetDB().Create(p).Error
	return p.ID, err
}

func (p *Policy) Find() (ps []Policy, err error) {
	err = db.GetDB().Where(p).Find(&ps).Error
	return ps, err
}

func (p *Policy) Updates(new *Policy) error {
	return db.GetDB().Where(p).Updates(new).Error
}

func (p *Policy) IsExist() (isExist bool, err error) {
	pl := Policy{}
	if err = db.GetDB().Where(p).First(&pl).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
