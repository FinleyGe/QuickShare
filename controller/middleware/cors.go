package middleware

import (
	. "QuickShare/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	if Config.Env == "dev" || Config.Env == "test" {
		corsConfig.AllowAllOrigins = true
	} else if Config.Env == "prod" {
		// TODO: add AllowHeaders
		/* cors.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization") */
		corsConfig.AllowOrigins = Config.Server.AllowOrigins
	} else {
		panic("Invalid Env")
	}
	return cors.New(corsConfig)
}
