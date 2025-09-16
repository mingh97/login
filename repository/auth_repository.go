package repository

import (
	"loginRegister/global"
	"loginRegister/model"
)

func CreateUser(user *model.User) error {
	return global.DB.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := global.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
