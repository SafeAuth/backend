package model

import (
	"gorm.io/gorm"
)

type ProgramFiles struct {
	gorm.Model
	Program      string `gorm:"not null"`
	FileId       string `gorm:"not null"`
	FileName     string `gorm:"not null"`
	FileLocation string `gorm:"not null"`
	EncKey       string `gorm:"not null"`
}
