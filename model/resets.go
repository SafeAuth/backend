package model

import (
	"gorm.io/gorm"
)

type Resets struct {
	gorm.Model

	Email     string `gorm:"unique;not null"`
	Code      string `gorm:"unique;not null"`
	ExpiresAt int64  `gorm:"not null"`
}
