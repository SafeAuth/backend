package models

import (
	"gorm.io/gorm"
)

type ProgramUsers struct {
	gorm.Model
	Program    string `gorm:"not null"`
	Username   string `gorm:"not null"`
	Email      string `gorm:"not null"`
	Password   string `gorm:"not null"`
	ExpiresAt  int64  `gorm:"not null"`
	Paused     bool   `gorm:"not null"`
	HWID       string `gorm:"not null"`
	Rank       string `gorm:"not null"`
	Var        string `gorm:"not null"`
	Banned     bool   `gorm:"not null"`
	BanReason  string `gorm:"not null"`
	BanExpires int64  `gorm:"not null"`
	LastLogin  int64  `gorm:"not null"`
	LastIP     string `gorm:"not null"`
	LastHWID   string `gorm:"not null"`
}
