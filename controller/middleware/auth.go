package middleware

import (
	"QuickShare/db/model"
	"QuickShare/utility"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	jwtData, err := utility.ParseToken(token)
	log.Println(jwtData, err)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
	id, err := strconv.Atoi(jwtData.ID)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	user, err := model.GetUserByID(id)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	if !user.Admin {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()
}
