package models

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	FullName string `gorm:"column:full_name"`
}
