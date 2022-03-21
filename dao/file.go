package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID        uint64         `gorm:"primarykey"`
	AccountID uint64         `gorm:"column:account_id" json:"account_id" sql:"bigint(20)"`
	Address   string         `gorm:"column:address" json:"addr" sql:"varchar(512)"`
	Signature string         `gorm:"column:signature" json:"signature" sql:"varchar()"` // todo length?
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at,omitempty" sql:"datetime"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at,omitempty" sql:"datetime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at,index" json:"deleted_at,omitempty" sql:"datetime"`
}

func NewFile() *File {
	return &File{}
}

func (f *File) TableName() string {
	return "file"
}

func (f *File) Create() (id uint64, err error) {
	err = db.GetDB().Create(f).Error
	return f.ID, err
}

func (f *File) BatchCreate(fs []*File) error {
	return db.GetDB().Create(fs).Error
}

func (f *File) Find() (files []File, err error) {
	err = db.GetDB().Where(f).Find(&files).Error
	return files, err
}

func (f *File) FindNotAccountID(accountID uint64) (files []*File, err error) {
	err = db.GetDB().Where(f).Where("account_id != ?", accountID).Find(&files).Error
	return files, err
}

func (f *File) Delete() error {
	return db.GetDB().Delete(f).Error
}
