package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID             uint64         `gorm:"primarykey"`
	FileID         string         `gorm:"column:file_id" json:"file_id" sql:"char(36)"`
	Name           string         `gorm:"column:name" json:"name" sql:"varchar(512)"`
	Address        string         `gorm:"column:address" json:"addr" sql:"varchar(512)"`
	Owner          string         `gorm:"column:owner" json:"owner" sql:"varchar(512)"`
	OwnerAccountID string         `gorm:"column:owner_account_id" json:"owner_account_id" sql:"char(36)"`
	Thumbnail      string         `gorm:"column:thumbnail" json:"thumbnail" sql:"varchar(512)"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at,omitempty" sql:"datetime"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at,omitempty" sql:"datetime"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at,index" json:"deleted_at,omitempty" sql:"datetime"`
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

func (f *File) Find(page, pageSize int) (files []File, err error) {
	err = db.GetDB().Where(f).Scopes(Paginate(page, pageSize)).Find(&files).Error
	return files, err
}

func (f *File) FindNotAccountID(accountID string, page, pageSize int) (files []*File, err error) {
	err = db.GetDB().Where(f).Where("owner_account_id != ?", accountID).Scopes(Paginate(page, pageSize)).Find(&files).Error
	return files, err
}

func (f *File) Delete() error {
	return db.GetDB().Delete(f).Error
}
