package database

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/SafeAuth/backend/jwt"
	"github.com/SafeAuth/backend/model"
	argonpass "github.com/dwin/goArgonPass"
	"github.com/google/uuid"
)

// "math"
// "os"

// "github.com/SafeAuth/backend/model"
// "gorm.io/gorm"

func VerUser(token string) model.ValidateUser {
	db := DB
	result := map[string]any{}
	db.Model(&model.Users{}).Select("banned, admin, username, id, admin_token, token").Where("token = ? ", token, token).Find(&result)
	if len(result) == 0 {
		return model.ValidateUser{Admin: false, ValidUser: false, Banned: false, Username: "", Uid: -1}
	}
	admin, _ := strconv.ParseBool(fmt.Sprintf("%v", result["admin"]))
	banned := result["banned"] != nil
	username := fmt.Sprintf("%v", result["username"])
	jwt := fmt.Sprintf("%v", result["token"])
	adminToken := fmt.Sprintf("%v", result["admin_token"])
	uid, _ := strconv.ParseInt(fmt.Sprintf("%v", result["id"]), 10, 64)
	return model.ValidateUser{Admin: admin, ValidUser: true, Banned: banned, Username: username, Uid: int(uid), JWT: jwt, ApiKey: adminToken}
}

func Login(username string, password string) (string, error) {
	db := DB
	result := map[string]any{}
	db.Model(&model.Users{}).Select("password, id, token, banned").Where("username iLIKE ? OR email iLIKE ?", username, username).Find(&result)

	hashedPassword := fmt.Sprintf("%v", result["password"])
	err := argonpass.Verify(password, hashedPassword)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	banned := result["banned"] != nil
	if banned {
		return "", errors.New("You are banned")
	}

	return fmt.Sprintf("%v", result["token"]), nil
}

func Register(username string, email string, password string, ip string) (string, error) {
	db := DB
	result := map[string]any{}
	db.Model(&model.Users{}).Select("id").Where("username iLIKE ? OR email iLIKE ?", username, email).Find(&result)
	if len(result) != 0 {
		return "", errors.New("Username or email already taken")
	}

	hashedPassword, err := argonpass.Hash(password, nil)
	if err != nil {
		return "", errors.New("Error hashing password")
	}

	token, err := jwt.GenerateUserToken("system", username)
	if err != nil {
		return "", errors.New("Error generating token")
	}

	adminToken := uuid.New().String()
	db.Create(&model.Users{Username: username, Email: email, Password: hashedPassword, Token: token, AdminToken: adminToken, Premium: "false", Admin: "false", Verified: "false", Ip: ip, LastIp: ip, LastLogin: "never", Registered: "never"})
	return token, nil
}
