package controller

import (
	"QuickShare/utility"
	. "QuickShare/utility"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	utility.CleanTheFiles()
	ResponseOK(c, "OK", nil)
}
