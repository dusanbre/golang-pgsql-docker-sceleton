package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique_index;not null"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}
