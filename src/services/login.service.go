package services

import (
	"errors"
	"gorestapi/src/config"
	"gorestapi/src/entities"

	_ "github.com/go-sql-driver/mysql"
)

// LogIn ... LogIn Service
func LogIn(user *entities.User, username string, password string) (err error) {
	if err = config.DB.Where("username = ?", username).Find(user).Error; err != nil {
		return err
	}

	if user.Password != password {
		return errors.New("Unauthorize")
	}

	return nil
}
