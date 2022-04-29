package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

const (
	PolicyStatusAll         = 0
	PolicyStatusUnpublished = 1
	PolicyStatusPublished   = 2
)

type Policy struct {
	ID               uint64         `gorm:"primarykey"`
	Hrac             string         `gorm:"column:hrac" json:"hrac" sql:"varchar(256)"`
	Label            string         `gorm:"column:label" json:"label" sql:"varchar(32)"`
	PolicyID         string         `gorm:"column:policy_id" json:"policy_id" sql:"char(36)"`
	Creator          string         `gorm:"column:creator" json:"creator" sql:"varchar(32)"`
	CreatorID        string         `gorm:"column:creator_id" json:"creator_id" sql:"char(36)"`
	Consumer         string         `gorm:"column:consumer" json:"consumer" sql:"varchar(32)"`
	ConsumerID       string         `gorm:"column:consumer_id" json:"consumer_id" sql:"char(36)"`
	EncryptedPK      string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar(256)"`           // todo length?
	EncryptedAddress string         `gorm:"column:encrypted_address" json:"encrypted_address" sql:"varchar(256)"` // todo length?
	Status           uint8          `gorm:"column:status;default:1" json:"status" sql:"tinyint(4)" comment:"1: unpublished, 2: published"`
	Gas              string         `gorm:"column:gas" json:"gas" sql:"varchar(32)"`
	TxHash           string         `gorm:"column:tx_hash" json:"tx_hash" sql:"char(66)"`
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

func (p *Policy) Get() (policy *Policy, err error) {
	err = db.GetDB().Where(p).First(&policy).Error
	return policy, err
}

func (p *Policy) Find(pager func(*gorm.DB) *gorm.DB) (ps []*Policy, err error) {
	tx := db.GetDB().Where(p)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&ps).Error
	return ps, err
}

func (p *Policy) FindPolicyIDs() (policyIDs []string, err error) {
	err = db.GetDB().Model(p).Where(p).Pluck("policy_id", &policyIDs).Error
	return policyIDs, err
}

func (p *Policy) Updates(new *Policy) error {
	return db.GetDB().Where(p).Updates(new).Error
}

func (p *Policy) Delete() (rows int64, err error) {
	ret := db.GetDB().Where(p).Delete(p)
	return ret.RowsAffected, ret.Error
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

func CreatePolicyAndFiles(policy *Policy, files []*File) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(policy).Error; err != nil {
			return err
		}
		if err := tx.Create(files).Error; err != nil {
			return err
		}
		return nil
	})
}

func Tx(models ...interface{}) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, m := range models {
			if err := tx.Create(m).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
