package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     uint8  `json:"role"`
	Avatar   string `json:"avatar"`
	Gender   uint8  `json:"gender"`
	Status   uint8  `json:"status"`
}

func (user User) TableName() string {
	return "user"
}
