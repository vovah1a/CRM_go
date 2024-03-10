package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" gorm:"text; not null; unique"`
	Password  string `json:"password" gorm:"text; not null"`
}
