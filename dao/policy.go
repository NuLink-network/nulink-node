package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"github.com/NuLink-network/nulink-node/utils"
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
	PolicyLabelID    string         `gorm:"column:policy_label_id" json:"policy_label_id" sql:"char(36)"`
	Creator          string         `gorm:"column:creator" json:"creator" sql:"varchar(32)"`
	CreatorID        string         `gorm:"column:creator_id" json:"creator_id" sql:"char(36)"`
	Consumer         string         `gorm:"column:consumer" json:"consumer" sql:"varchar(32)"`
	ConsumerID       string         `gorm:"column:consumer_id" json:"consumer_id" sql:"char(36)"`
	EncryptedPK      string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar(256)"`
	EncryptedAddress string         `gorm:"column:encrypted_address" json:"encrypted_address" sql:"varchar(256)"`
	Gas              string         `gorm:"column:gas" json:"gas" sql:"varchar(32)"`
	TxHash           string         `gorm:"column:tx_hash" json:"tx_hash" sql:"char(128)"`
	StartAt          time.Time      `gorm:"column:start_at" json:"start_at" sql:"datetime"`
	EndAt            time.Time      `gorm:"column:end_at" json:"end_at" sql:"datetime"`
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

func (p *Policy) Find(ext *QueryExtra, pager Pager) (ps []*Policy, err error) {
	tx := db.GetDB().Where(p)
	if ext != nil && ext.Conditions != nil {
		for k, v := range ext.Conditions {
			tx = tx.Where(k, v)
		}
	}
	if !utils.IsEmpty(ext.OrderStr) {
		tx.Order(ext.OrderStr)
	}
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
