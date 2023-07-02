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
	// Get the database connection
	db := DB
	// Create a map to hold the result
	result := map[string]any{}
	// Query the database for the user, selecting the fields we need
	db.Model(&model.Users{}).Select("banned, admin, username, id, admin_token, token").Where("token = ? ", token, token).Find(&result)
	// If the result is empty, return an empty ValidateUser
	if len(result) == 0 {
		return model.ValidateUser{Admin: false, ValidUser: false, Banned: false, Username: "", Uid: -1}
	}
	// Convert the admin field to a boolean
	admin, _ := strconv.ParseBool(fmt.Sprintf("%v", result["admin"]))
	// If the banned field is not nil, the user is banned
	banned := result["banned"] != nil
	// Cast the username field to a string
	username := fmt.Sprintf("%v", result["username"])
	// Cast the id field to an integer
	uid, _ := strconv.ParseInt(fmt.Sprintf("%v", result["id"]), 10, 64)
	// Cast the token field to a string
	jwt := fmt.Sprintf("%v", result["token"])
	// Cast the admin_token field to a string
	adminToken := fmt.Sprintf("%v", result["admin_token"])
	// Return a ValidateUser object
	return model.ValidateUser{Admin: admin, ValidUser: true, Banned: banned, Username: username, Uid: int(uid), JWT: jwt, ApiKey: adminToken}
}

func Login(username string, password string, ip string) (string, error) {
	// Create a new database connection
	db := DB

	// Create a new variable to store the result of the query
	result := map[string]any{}

	// Query the database for the password, id, token, and banned status
	db.Model(&model.Users{}).Select("password, id, token, banned").Where("username iLIKE ? OR email iLIKE ?", username, username).Find(&result)

	// Get the password from the result
	hashedPassword := fmt.Sprintf("%v", result["password"])

	// Verify the password
	err := argonpass.Verify(password, hashedPassword)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	// Check if the user is banned
	banned := result["banned"] != nil
	if banned {
		return "", errors.New("You are banned")
	}

	// Return the token
	return fmt.Sprintf("%v", result["token"]), nil
}

func Register(username string, email string, password string, ip string) (string, error) {
	// Search for existing username or email
	db := DB                                                                                                        // Get database connection
	result := map[string]any{}                                                                                      // Create empty map for results
	db.Model(&model.Users{}).Select("id").Where("username iLIKE ? OR email iLIKE ?", username, email).Find(&result) // Query database
	// Check if results were returned
	if len(result) != 0 {
		return "", errors.New("Username or email already taken")
	}

	// hash the password using argon2id
	hashedPassword, err := argonpass.Hash(password, nil)
	if err != nil {
		return "", errors.New("Error hashing password")
	}

	// Generate token for user

	token, err := jwt.GenerateUserToken("system", username)
	if err != nil {
		return "", errors.New("Error generating token")
	}
	// Generate a token for the admin
	adminToken := uuid.New().String()
	// Create the user in the database
	db.Create(&model.Users{Username: username, Email: email, Password: hashedPassword, Token: token, AdminToken: adminToken, Premium: "false", Admin: "false", Verified: "false", Ip: ip, LastIp: ip, LastLogin: "never", Registered: "never"})
	// Return the token
	return fmt.Sprintf("%v", result["token"]), nil
}

func Reset(email string) (string, error) {
	db := DB
	result := map[string]any{}

	db.Model(&model.Users{}).Select("id").Where("email iLIKE ?", email).Find(&result)
	if len(result) == 0 {
		return "", errors.New("No user with that email")
	}

	token := uuid.New().String()
	db.Exec("INSERT INTO resets (token, email, expires) VALUES (?, ?, NOW() + INTERVAL '15 minutes')", token, email)
	return token, nil

}

func ResetPassword(token string, password string) error {
	db := DB
	result := map[string]any{}

	db.Model(&model.Resets{}).Select("id").Where("token = ? AND expires > NOW()", token).Find(&result)
	if len(result) == 0 {
		return errors.New("Invalid token")
	}

	hashedPassword, err := argonpass.Hash(password, nil)
	if err != nil {
		return errors.New("Error hashing password")
	}

	db.Exec("UPDATE users SET password = ? WHERE email = ?", hashedPassword, result["email"])
	return nil
}

func CreateProgram(programName string, encKey string, username string) error {
	db := DB
	result := map[string]any{}

	db.Model(&model.Programs{}).Select("id").Where("program_name = ?", programName).Find(&result)
	if len(result) != 0 {
		return errors.New("Program name already taken")
	}

	db.Create(&model.Programs{ProgramName: programName, EncKey: encKey, Username: username})
	return nil
}
