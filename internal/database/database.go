package database

import (
	"github.com/ekideno/gosnip/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(domain.URL{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
