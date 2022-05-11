package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"github.com/NuLink-network/nulink-node/utils"
	"gorm.io/gorm"
	"time"
)

const (
	ApplyStatusAll      = 0
	ApplyStatusApplying = 1
	ApplyStatusApproved = 2
	ApplyStatusRejected = 3
)

type ApplyFile struct {
	ID     uint64 `gorm:"primarykey"`
	FileID string `gorm:"column:file_id" json:"file_id" sql:"char(36)"`
	//FileName    string         `gorm:"column:file_name" json:"file_name" sql:"varchar()"`
	Proposer    string         `gorm:"column:proposer" json:"proposer" sql:"varchar(32)" comment:"申请者"`
	ProposerID  string         `gorm:"column:proposer_id" json:"proposer_id" sql:"char(36)"`
	FileOwner   string         `gorm:"column:file_owner" json:"file_owner" sql:"varchar(32)" comment:"文件拥有者"`
	FileOwnerID string         `gorm:"column:file_owner_id" json:"file_owner_id" sql:"char(36)"`
	Status      uint8          `gorm:"column:status;default:1" json:"status" sql:"tinyint(4)" comment:"1: applying, 2: approved 3: rejected"`
	StartAt     time.Time      `gorm:"column:start_at" json:"start_at" sql:"datetime"`
	FinishAt    time.Time      `gorm:"column:finish_at" json:"finish_at" sql:"datetime"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at" sql:"datetime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at" sql:"datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at" sql:"datetime"`
}

func NewAppleFile() *ApplyFile {
	return &ApplyFile{}
}

func (a *ApplyFile) TableName() string {
	return "apply_file"
}

func (a *ApplyFile) Create() (id uint64, err error) {
	err = db.GetDB().Create(a).Error
	return a.ID, err
}

func (a *ApplyFile) BatchCreate(as []*ApplyFile) error {
	return db.GetDB().Create(as).Error
}

func (a *ApplyFile) Get() (af *ApplyFile, err error) {
	err = db.GetDB().Where(a).First(&af).Error
	return af, err
}

func (a *ApplyFile) Find(pager Pager) (afs []*ApplyFile, err error) {
	tx := db.GetDB().Where(a)
	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&afs).Error
	return afs, err
}

func (a *ApplyFile) FindAny(ext *QueryExtra, pager Pager) (afs []*ApplyFile, err error) {
	tx := db.GetDB().Where(a)
	if ext != nil {
		if ext.Conditions != nil {
			for k, v := range ext.Conditions {
				tx = tx.Where(k, v)
			}
		}
		if !utils.IsEmpty(ext.OrderStr) {
			tx.Order(ext.OrderStr)
		}
	}

	if pager != nil {
		tx = tx.Scopes(pager)
	}
	err = tx.Find(&afs).Error
	return afs, err
}

func (a *ApplyFile) Delete() (rows int64, err error) {
	ret := db.GetDB().Where(a).Delete(a)
	return ret.RowsAffected, ret.Error
}

func (a *ApplyFile) BatchDelete(ids []uint64) error {
	return db.GetDB().Where(a).Delete(a, ids).Error
}

func (a *ApplyFile) DeleteByFileIDs(fids []string) error {
	return db.GetDB().Where(a).Where("file_id in ?", fids).Delete(a).Error
}

func (a *ApplyFile) Updates(new *ApplyFile) error {
	return db.GetDB().Where(a).Updates(new).Error
}
