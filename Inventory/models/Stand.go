package models

import (
	"github.com/jinzhu/gorm"
)

type Stand struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Location string `gorm:"column:location"`
}
