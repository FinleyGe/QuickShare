package model

import (
	. "QuickShare/db"
	"log"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

func AddUser(User User) error {
	return DB.Create(&User).Error
}

func GetUserByID(id int) (User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func IsAdminExist() bool {
	var user User
	DB.Where("admin = ?", true).First(&user)
	return user.ID != 0
}

func IsUserExist(name string) bool {
	var user User
	DB.Where("user_name = ?", name).First(&user)
	if (user != User{}) {
		log.Println("user already exists")
		return true
	}
	return false
}
