package infrastructure

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
}
