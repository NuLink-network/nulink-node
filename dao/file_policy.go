package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

type FilePolicy struct {
	ID        uint64         `gorm:"primarykey"`
	FileID    string         `gorm:"column:file_id" json:"file_id" sql:"char(36)"`
	PolicyID  string         `gorm:"column:policy_id" json:"policy_id" sql:"char(36)"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at,omitempty" sql:"datetime"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at,omitempty" sql:"datetime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty" sql:"datetime"`
}

func NewFilePolicy() *FilePolicy {
	return &FilePolicy{}
}

func (f *FilePolicy) TableName() string {
	return "file_policy"
}

func (f *FilePolicy) Create() (id uint64, err error) {
	err = db.GetDB().Create(f).Error
	return f.ID, err
}

func (f *FilePolicy) BatchCreate(fps []*FilePolicy) error {
	return db.GetDB().Create(fps).Error
}

func (f *FilePolicy) Get() (fp *FilePolicy, err error) {
	err = db.GetDB().Where(f).First(&fp).Error
	return fp, err
}

func (f *FilePolicy) Find(pager func(*gorm.DB) *gorm.DB) (fps []*FilePolicy, err error) {
	tx := db.GetDB().Where(f)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&fps).Error
	return fps, err
}

func (f *FilePolicy) FindFileIDsByPolicyIDs(policyIDs []string, pager func(*gorm.DB) *gorm.DB) (fileIDs []string, err error) {
	tx := db.GetDB().Model(f).Where("policy_id in ?", policyIDs)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Distinct().Pluck("file_id", &fileIDs).Error
	return fileIDs, err
}

func (f *FilePolicy) Delete() (rows int64, err error) {
	ret := db.GetDB().Where(f).Delete(f)
	return ret.RowsAffected, ret.Error
}
