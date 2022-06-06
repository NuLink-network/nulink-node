package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"github.com/NuLink-network/nulink-node/utils"
	"gorm.io/gorm"
	"time"
)

type File struct {
	ID            uint64         `gorm:"primarykey"`
	FileID        string         `gorm:"column:file_id" json:"file_id" sql:"char(36)"`
	MD5           string         `gorm:"column:md5" json:"md5" sql:"varchar(32)"`
	Name          string         `gorm:"column:name" json:"name" sql:"varchar(32)"`
	Suffix        string         `gorm:"column:suffix" json:"suffix" sql:"varchar(16)"`
	Category      string         `gorm:"column:category" json:"category" sql:"varchar(32)"`
	Address       string         `gorm:"column:address" json:"addr" sql:"varchar(512)"`
	Thumbnail     string         `gorm:"column:thumbnail" json:"thumbnail" sql:"varchar(512)"`
	Owner         string         `gorm:"column:owner" json:"owner" sql:"varchar(512)"`
	OwnerID       string         `gorm:"column:owner_id" json:"owner_id" sql:"char(36)"`
	OwnerAddress  string         `gorm:"column:owner_address" json:"owner_address" sql:"char(42)"`
	PolicyLabelID string         `gorm:"column:policy_label_id" json:"policy_label_id" sql:"char(36)"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at,omitempty" sql:"datetime"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at,omitempty" sql:"datetime"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at,omitempty" sql:"datetime"`
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

func (f *File) Get() (file *File, err error) {
	err = db.GetDB().Where(f).First(&file).Error
	return file, err
}

func (f *File) Find(pager Pager) (files []*File, err error) {
	name := f.Name
	f.Name = ""

	tx := db.GetDB().Where("name like ?", name).Where(f)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&files).Error
	return files, err
}

//func (f *File) FindAny(query interface{}, args ...interface{}) (files []File, err error) {
//	err = db.GetDB().Where(query, args).Find(&files).Error
//	return files, err
//}

func (f *File) FindByFileIDs(fileIDs []string, pager Pager) (files []*File, err error) {
	tx := db.GetDB().Where("file_id in ?", fileIDs)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&files).Error
	return files, err
}

func (f *File) FindAny(ext *QueryExtra, pager Pager) (files []*File, err error) {
	tx := db.GetDB().Where(f)
	if ext != nil {
		if ext.Conditions != nil {
			for k, v := range ext.Conditions {
				tx = tx.Where(k, v)
			}
		}
		if !utils.IsEmpty(ext.OrderStr) {
			tx = tx.Order(ext.OrderStr)
		}
	}

	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&files).Error
	return files, err
}

func (f *File) Delete() error {
	return db.GetDB().Delete(f).Error
}

func (f *File) DeleteByFilesIDs(fileIDs []string) error {
	return db.GetDB().Where(f).Where("file_id in ?", fileIDs).Delete(f).Error
}

func (f *File) Updates(new *File) error {
	return db.GetDB().Where(f).Updates(new).Error
}
