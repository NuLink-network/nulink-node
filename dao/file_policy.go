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
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at,index" json:"deleted_at,omitempty" sql:"datetime"`
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
