package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

type Policy struct {
	ID          uint64         `gorm:"primarykey"`
	AccountID   uint64         `gorm:"column:account_id" json:"account_id" sql:"bigint(20)"`
	Label       string         `gorm:"column:label" json:"label" sql:"varchar()"`               // todo length?
	EncryptedPK string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar()"` // todo length?
	VerifyPK    string         `gorm:"verify_pk:" json:"verify_pk" sql:"varchar()"`             // todo length?
	Signature   string         `gorm:"column:signature" json:"signature" sql:"varchar()"`       // todo length?
	IsPublish   bool           `gorm:"column:is_publish" json:"is_publish" sql:"tinyint(1)"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
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
