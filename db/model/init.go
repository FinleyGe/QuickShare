package model

import (
	. "QuickShare/db"
)

func init() {
	DB.AutoMigrate(&User{}, &File{})
}
