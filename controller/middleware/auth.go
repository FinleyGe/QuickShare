package middleware

import (
	"QuickShare/db/model"
	"QuickShare/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
	t, err := utility.ParseToken(token)
	id, _ := strconv.Atoi(t)
	user := model.User{}
	err = user.GetByID(uint(id))
	if (err != nil || user == model.User{}) {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
