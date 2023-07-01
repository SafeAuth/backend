package database

import (
	"fmt"
	"os"

	"github.com/SafeAuth/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_SSLMODE"),
		os.Getenv("DATABASE_TIMEZONE"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the DB")
	}
	fmt.Println("Connected to the DB")

	DB.AutoMigrate(&model.ProgramFiles{}, &model.ProgramLogs{}, &model.Programs{}, &model.ProgramSessions{}, &model.ProgramTokens{}, &model.ProgramUsers{}, &model.ProgramVars{}, &model.Resets{}, &model.Users{}, &model.Tokens{}, &model.Users{})

}
