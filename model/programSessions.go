package models

import (
	"gorm.io/gorm"
)

type ProgramSessions struct {
	gorm.Model
	Program  string `gorm:"not null"`
	Session  string `gorm:"not null"`
	Expires  int64  `gorm:"not null"`
	IV       string `gorm:"not null"`
	LoggedIn bool   `gorm:"not null"`
	IP       string `gorm:"not null"`
}
