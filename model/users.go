package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	AdminToken string `gorm:"not null"`
	Premium    string `gorm:"not null"`
	Admin      string `gorm:"not null"`
	Verified   string `gorm:"not null"`
	Ip         string `gorm:"not null"`
	LastLogin  string `gorm:"not null"`
	Registered string `gorm:"not null"`
	LastIp     string `gorm:"not null"`
}
