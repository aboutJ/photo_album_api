package dao

import (
	"photo_album/config"
	"photo_album/model"
)

func AddUser(user *model.User) error {
	config.DB.Create(user)
	return config.DB.Error
}

func DeleteUser(id uint64) {
	config.DB.Delete(&model.User{}, id)
}

func FindUserByEmail(email string) ([]model.User, error) {
	var users []model.User
	result := config.DB.Find(&users, "email=?", email)
	return users, result.Error
}
