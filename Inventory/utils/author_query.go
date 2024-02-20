package utils

import (
	"github.com/IsmaelNicolas/Library_System/models"
	"github.com/jinzhu/gorm"
)

// CreateAuthor crea un nuevo autor en la base de datos
func CreateAuthor(db *gorm.DB, fullName string) (*models.Author, error) {
	author := &models.Author{FullName: fullName}
	if err := db.Create(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func ReadAuthors(db *gorm.DB) ([]*models.Author, error) {
	var authors []*models.Author
	if err := db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
