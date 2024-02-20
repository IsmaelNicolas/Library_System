package utils

import (
	"github.com/IsmaelNicolas/Library_System/models"
	"github.com/jinzhu/gorm"
)

// CreateBook crea un nuevo libro en la base de datos
func CreateBook(db *gorm.DB, title string, year int32, authorID uint) (*models.Book, error) {
	book := &models.Book{Title: title, Year: year, AuthorID: authorID}
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// ReadBookByID lee un libro de la base de datos por su ID
func ReadBookByID(db *gorm.DB, id uint64) (*models.Book, error) {
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// UpdateBookByID actualiza un libro en la base de datos por su ID
func UpdateBookByID(db *gorm.DB, id uint64, title string, year int32, authorID uint) (*models.Book, error) {
	var book models.Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	book.Title = title
	book.Year = year
	book.AuthorID = authorID
	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
