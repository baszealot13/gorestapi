package services

import (
	"fmt"
	"gorestapi/src/config"
	"gorestapi/src/entities"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]entities.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser models. ... Insert New data
func CreateUser(user *entities.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *entities.User, id string) (err error) {
	if err = config.DB.Where("user_id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *entities.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *entities.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
