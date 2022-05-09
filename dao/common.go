package dao

import (
	"github.com/NuLink-network/nulink-node/resource/db"
	"gorm.io/gorm"
)

type QueryExtra struct {
	Conditions  map[string]interface{}
	OrderStr    string
	DistinctStr []string
}

type Pager func(*gorm.DB) *gorm.DB

func Paginate(page, pageSize int) Pager {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Tx(models ...interface{}) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, m := range models {
			if err := tx.Create(m).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
