package models

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Book     Book `gorm:"foreignkey:BookID"`
	BookID   uint
	Stand    Stand `gorm:"foreignkey:StandID"`
	StandID  uint
	Quantity int32 `gorm:"column:quantity"`
	StockMin int32 `gorm:"column:stock_min"`
}
