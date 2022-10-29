package controller

import (
	. "QuickShare/utility"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	ResponseOK(c, "OK", nil)
}
