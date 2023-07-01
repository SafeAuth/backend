package model

import (
	"gorm.io/gorm"
)

type Programs struct {
	gorm.Model
	Owner       string `gorm:"not null"`
	ProgramName string `gorm:"not null"`
	ProgramKey  string `gorm:"not null"`
	EncKey      string `gorm:"not null"`
	DL          string `gorm:"not null"`
	Version     string `gorm:"not null"`
	KillSwitch  bool   `gorm:"not null"`
	HwidLock    bool   `gorm:"not null"`
}
