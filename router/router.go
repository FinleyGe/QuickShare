package router

import (
	. "QuickShare/controller"
	. "QuickShare/controller/middleware"
)

func init() {
	Router.POST("/login", Login)
	Router.POST("/register", Register)

	Router.GET("/info/:hash", GetFileInfo)
	Router.GET("/get/:hash", Download)
	Router.DELETE("/delete/:hash", Admin, DeleteFile)
	Router.POST("/upload", Admin, Upload)

	Router.GET("/all_info", Admin, GetAllInfo)
	Router.GET("/all_info_type", Admin, GetAllInfoByType)
	Router.GET("/search", Admin, SearchFile)

	//! will be removed in the future
	// Router.GET("/all_info", Admin, GetAllInfo)
}
