package controller

import (
	. "QuickShare/db"
	"QuickShare/db/model"
	"QuickShare/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	name := user.Username
	pwd := user.Password
	u := model.User{}
	err := DB.Where("name=?", name).First(&u).Error
	if err != nil {
		utility.Response(c, 404, "No user", nil)
		return
	}
	if utility.PasswordVerify(pwd, u.Password) {
		jwtData := utility.JwtData{
			ID: strconv.Itoa(int(u.ID)),
		}
		c.SetCookie("token", utility.GenerateStandardJwt(&jwtData), 3600, "/", "localhost", false, true)
		utility.ResponseOK(c, "OK", nil)
		return
	}
}
