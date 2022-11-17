package router

import (
	. "QuickShare/controller"
	. "QuickShare/controller/middleware"
)

func init() {
	Router.GET("/info/:hash", GetFileInfo)
	Router.GET("/get/:hash", Download)
	Router.POST("/upload", Auth, Upload)
	Router.GET("/all_info", Auth, GetAllInfo)
}
