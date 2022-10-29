package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, message string, data interface{}) {
	if data == nil {
		c.JSON(code, gin.H{
			"message": message,
		})
	} else {
		c.JSON(code, gin.H{
			"message": message,
			"data":    data,
		})
	}
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	Response(c, http.StatusOK, "OK", nil)
}
