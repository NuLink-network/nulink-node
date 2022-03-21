package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

const (
	StatusApprove = 1
	StatusReject  = 2
)

type AppleFile struct {
	ID           uint64         `gorm:"primarykey"`
	FileID       uint64         `gorm:"column:file_id" json:"file_id" sql:"bigint(20)"`
	ProposerID   uint64         `gorm:"column:proposer_id" json:"proposer_id" sql:"bigint(20)"`
	ProprietorID uint64         `gorm:"column:proprietor_id" json:"proprietor_id" sql:"bigint(20)"`
	Signature    string         `gorm:"column:signature" json:"signature" sql:"varchar()"` // vSK Signature todo length?
	StartAt      time.Time      `gorm:"column:start_at" json:"start_at,omitempty" sql:"datetime"`
	Status       int8           `gorm:"column:approve_status" json:"approve_status" sql:"bigint(20)" comment:"0: default, 1: approve 2: reject"`
	FinishAt     time.Time      `gorm:"column:finish_at" json:"finish_at" sql:"datetime"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewAppleFile() *File {
	return &File{}
}

func (a *AppleFile) TableName() string {
	return "apple_file"
}

func (a *AppleFile) Create() (id uint64, err error) {
	err = db.GetDB().Create(a).Error
	return a.ID, err
}

func (a *AppleFile) BatchCreate(as []*AppleFile) error {
	return db.GetDB().Create(as).Error
}

func (a *AppleFile) Find() (files []*AppleFile, err error) {
	err = db.GetDB().Where(a).Find(&files).Error
	return files, err
}

func (a *AppleFile) BatchDelete(ids []uint64) error {
	return db.GetDB().Delete(a, ids).Error
}

func (a *AppleFile) Updates(new *AppleFile) error {
	return db.GetDB().Where(a).Updates(new).Error
}
