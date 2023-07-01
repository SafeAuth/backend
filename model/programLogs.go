package model

import (
	"gorm.io/gorm"
)

type ProgramLogs struct {
	gorm.Model
	Program  string `gorm:"not null"`
	Username string `gorm:"not null"`
	Message  string `gorm:"not null"`
	Time     string `gorm:"not null"`
	IP       string `gorm:"not null"`
}
