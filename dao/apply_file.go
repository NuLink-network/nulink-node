package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

const (
	StatusAll      = 0
	StatusApplying = 1
	StatusApprove  = 2
	StatusReject   = 3
)

type ApplyFile struct {
	ID                  uint64         `gorm:"primarykey"`
	FileID              uint64         `gorm:"column:file_id" json:"file_id" sql:"bigint(20)"`
	FileName            uint64         `gorm:"column:file_name" json:"file_name" sql:"varchar()"`
	Proposer            string         `gorm:"column:proposer" json:"proposer" sql:"varchar()"`
	ProposerAccountID   string         `gorm:"column:proposer_account_id" json:"proposer_account_id" sql:"varchar(36)"`
	Proprietor          string         `gorm:"column:proprietor" json:"proprietor" sql:"bigint(20)"`
	ProprietorAccountID string         `gorm:"column:proprietor_account_id" json:"proprietor_account_id" sql:"varchar(36)"`
	Status              uint8          `gorm:"column:approve_status" json:"approve_status" sql:"tinyint(4)" comment:"1: applying, 2: approve 3: reject"`
	StartAt             time.Time      `gorm:"column:start_at" json:"start_at,omitempty" sql:"datetime"`
	FinishAt            time.Time      `gorm:"column:finish_at" json:"finish_at" sql:"datetime"`
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewAppleFile() *File {
	return &File{}
}

func (a *ApplyFile) TableName() string {
	return "apple_file"
}

func (a *ApplyFile) Create() (id uint64, err error) {
	err = db.GetDB().Create(a).Error
	return a.ID, err
}

func (a *ApplyFile) BatchCreate(as []*ApplyFile) error {
	return db.GetDB().Create(as).Error
}

func (a *ApplyFile) Find(page, pageSize int) (files []*ApplyFile, err error) {
	err = db.GetDB().Where(a).Scopes(Paginate(page, pageSize)).Find(&files).Error
	return files, err
}

func (a *ApplyFile) BatchDelete(ids []uint64) error {
	return db.GetDB().Where(a).Delete(a, ids).Error
}

func (a *ApplyFile) Updates(new *ApplyFile) error {
	return db.GetDB().Where(a).Updates(new).Error
}
