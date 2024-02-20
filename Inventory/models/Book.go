package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title    string `gorm:"column:title"`
	Year     int32  `gorm:"column:year"`
	Author   Author `gorm:"foreignkey:AuthorID"`
	AuthorID uint
}
