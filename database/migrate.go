package database

import (
	"github.com/dusanbre/goPgsqlDocker/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
