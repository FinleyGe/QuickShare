package router

import (
	. "QuickShare/controller"
)

func init() {
	Router.GET("/test", Test)
}
