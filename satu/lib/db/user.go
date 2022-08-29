package db

import (
	"Satu/config"
	"Satu/middlewares"
	"Satu/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers() (interface{}, error) {
	var user []models.User

	if err := config.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(id string) (interface{}, error) {
	var user []models.User

	if err := config.DB.First(&user, "user_id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id string, user *models.User) (interface{}, error) {
	var existingUser models.User
	if err := config.DB.First(&existingUser, "user_id=?", id).Error; err != nil {
		return nil, err
	}
	existingUser.UserName = user.UserName
	existingUser.Password = user.Password
	if updateErr := config.DB.Save(&existingUser).Where("user_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingUser, nil
}

func DeleteUser(id string) (interface{}, error) {
	var user models.User
	if err := config.DB.First(&user, "user_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&user).Where("user_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return user, nil
}

func RegisterUser(user *models.User) (interface{}, error) {
	var users models.User
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	fmt.Println(string(bytes))
	users.Name = user.Name
	users.UserName = user.UserName
	users.Password = string(bytes)
	users.CreatedAt = time.Now()

	if err := config.DB.Create(&users).Error; err != nil {
		return nil, err
	}

	var err error
	user.Token, err = middlewares.CreateToken(string(user.Id))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func LoginUser(user *models.User) (interface{}, error) {
	var users models.User

	if err := config.DB.First(&users, "name=?", user.Name).Error; err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password))

	if err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(string(user.Id))
	if err != nil {
		return nil, err
	}

	return user, nil
}
