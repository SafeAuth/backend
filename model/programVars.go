package model

import (
	"gorm.io/gorm"
)

type ProgramVars struct {
	gorm.Model
	Program string `gorm:"not null"`
	Name    string `gorm:"not null"`
	Value   string `gorm:"not null"`
	EncKey  string `gorm:"not null"`
}
