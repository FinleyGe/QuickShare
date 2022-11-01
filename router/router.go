package router

import (
	. "QuickShare/controller"
)

func init() {
	Router.GET("/test", Test)
	Router.GET("/get/:hash", Download)
	Router.POST("/upload", Upload)
}
