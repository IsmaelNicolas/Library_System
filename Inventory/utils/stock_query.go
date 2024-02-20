package utils

import (
	"github.com/IsmaelNicolas/Library_System/models"
	"github.com/jinzhu/gorm"
)

// CreateStock crea un nuevo registro de stock en la base de datos

func CreateStock(db *gorm.DB, bookID uint64, standID uint64, quantity int32, stockMin int32) (*models.Stock, error) {
	stock := &models.Stock{BookID: uint(bookID), StandID: uint(standID), Quantity: quantity, StockMin: stockMin}
	if err := db.Create(stock).Error; err != nil {
		return nil, err
	}
	return stock, nil
}

func ReadStocks(db *gorm.DB) ([]*models.Stock, error) {
	var stocks []*models.Stock
	if err := db.Find(&stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}
