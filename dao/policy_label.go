package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

type PolicyLabel struct {
	ID            uint64         `gorm:"primarykey"`
	PolicyLabelID string         `gorm:"column:policy_label_id" json:"policy_label_id" sql:"char(36)"`
	Label         string         `gorm:"column:label" json:"label" sql:"varchar(128)"`
	Creator       string         `gorm:"column:creator" json:"creator" sql:"varchar(32)"`
	CreatorID     string         `gorm:"column:creator_id" json:"creator_id" sql:"char(36)"`
	EncryptedPK   string         `gorm:"column:encrypted_pk" json:"encrypted_pk" sql:"varchar(256)"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewPolicyLabel() *PolicyLabel {
	return &PolicyLabel{}
}

func (p *PolicyLabel) TableName() string {
	return "policy_label"
}

func (p *PolicyLabel) Create() (id uint64, err error) {
	err = db.GetDB().Create(p).Error
	return p.ID, err
}

func (p *PolicyLabel) Get() (pl *PolicyLabel, err error) {
	err = db.GetDB().Where(p).First(&pl).Error
	return pl, err
}

func (p *PolicyLabel) Find(pager Pager) (ps []*PolicyLabel, err error) {
	tx := db.GetDB().Where(p)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&ps).Error
	return ps, err
}

func (p *PolicyLabel) FindPolicyIDs() (policyIDs []string, err error) {
	err = db.GetDB().Model(p).Where(p).Pluck("policy_id", &policyIDs).Error
	return policyIDs, err
}

func (p *PolicyLabel) Updates(new *PolicyLabel) error {
	return db.GetDB().Where(p).Updates(new).Error
}

func (p *PolicyLabel) Delete() (rows int64, err error) {
	ret := db.GetDB().Where(p).Delete(p)
	return ret.RowsAffected, ret.Error
}

func (p *PolicyLabel) IsExist() (isExist bool, err error) {
	pl := PolicyLabel{}
	if err = db.GetDB().Where(p).First(&pl).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func CreatePolicyAndFiles(policy *PolicyLabel, files []*File) error {
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
