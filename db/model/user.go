package model

import (
	. "QuickShare/db"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) GetByID(id uint) error {
	return DB.First(user, id).Error
}
