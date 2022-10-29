package router

import (
	. "QuickShare/controller/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(Cors())
}
