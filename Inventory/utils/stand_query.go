package utils

import (
	"github.com/IsmaelNicolas/Library_System/models"
	"github.com/jinzhu/gorm"
)

// CreateStand crea un nuevo stand en la base de datos
func CreateStand(db *gorm.DB, name, location string) (*models.Stand, error) {
	stand := &models.Stand{Name: name, Location: location}
	if err := db.Create(stand).Error; err != nil {
		return nil, err
	}
	return stand, nil
}

// ReadStands lee todos los stands de la base de datos
func ReadStands(db *gorm.DB) ([]*models.Stand, error) {
	var stands []*models.Stand
	if err := db.Find(&stands).Error; err != nil {
		return nil, err
	}
	return stands, nil
}
