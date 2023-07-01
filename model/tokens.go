package model

import (
	"gorm.io/gorm"
)

type Tokens struct {
	gorm.Model
	Token   string `gorm:"unique;not null"`
	Used    string `gorm:"not null"`
	UsedBy  string `gorm:"not null"`
	SubTime string `gorm:"not null"`
}
