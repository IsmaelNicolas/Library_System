package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Importa el driver de PostgreSQL para GORM
)

func Connect() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := "user=" + user + " dbname=" + dbname + " password=" + password + " host=" + host + " port=" + port

	db, err := gorm.Open(dsn)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
