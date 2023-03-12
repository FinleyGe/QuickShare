package router

import (
	. "QuickShare/controller"
	. "QuickShare/controller/middleware"
)

func init() {
	Router.GET("/info/:hash", GetFileInfo)
	Router.GET("/get/:hash", Download)

	Router.POST("/upload", Admin, Upload)
	Router.GET("/all_info", Admin, GetAllInfo)

	Router.POST("/login", Login)
	Router.POST("/register", Register)
}
