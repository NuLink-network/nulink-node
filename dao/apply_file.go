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

type ApplyFile struct {
	ID           uint64         `gorm:"primarykey"`
	FileID       uint64         `gorm:"column:file_id" json:"file_id" sql:"bigint(20)"`
	FileName     uint64         `gorm:"column:file_name" json:"file_name" sql:"varchar()"`
	Proposer     string         `gorm:"column:proposer" json:"proposer" sql:"varchar()"`
	ProposerID   uint64         `gorm:"column:proposer_id" json:"proposer_id" sql:"bigint(20)"`
	Proprietor   string         `gorm:"column:proprietor" json:"proprietor" sql:"bigint(20)"`
	ProprietorID uint64         `gorm:"column:proprietor_id" json:"proprietor_id" sql:"bigint(20)"`
	Status       int8           `gorm:"column:approve_status" json:"approve_status" sql:"tinyint(4)" comment:"1: default, 2: approve 3: reject"`
	StartAt      time.Time      `gorm:"column:start_at" json:"start_at,omitempty" sql:"datetime"`
	FinishAt     time.Time      `gorm:"column:finish_at" json:"finish_at" sql:"datetime"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
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

func (a *ApplyFile) Find() (files []*ApplyFile, err error) {
	err = db.GetDB().Where(a).Find(&files).Error
	return files, err
}

func (a *ApplyFile) BatchDelete(ids []uint64) error {
	return db.GetDB().Delete(a, ids).Error
}

func (a *ApplyFile) Updates(new *ApplyFile) error {
	return db.GetDB().Where(a).Updates(new).Error
}
