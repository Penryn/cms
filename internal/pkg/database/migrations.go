package database

import (
	"cms/internal/model"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB)error {
	return db.AutoMigrate(
		model.User{},
		model.Contact{},
	)
}