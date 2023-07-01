package models

import (
	"gorm.io/gorm"
)

type ProgramTokens struct {
	gorm.Model
	Program string `gorm:"not null"`
	token   string `gorm:"not null"`
	Days    string `gorm:"not null"`
	Rank    string `gorm:"not null"`
	Used    string `gorm:"not null"`
	UsedBy  string `gorm:"not null"`
	UsedAt  string `gorm:"not null"`
}
