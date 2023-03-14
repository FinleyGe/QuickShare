package controller

import (
	"QuickShare/config"
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
	u.Username = name
	err := DB.First(&u).Error
	if err != nil {
		utility.Response(c, 404, "No user", nil)
		return
	}

	if utility.PasswordVerify(pwd, u.Password) {
		token, err := utility.GenerateToken(strconv.Itoa(int(u.ID)))
		if err != nil {
			utility.Response(c, 500, "Token generate error", nil)
			return
		}

		// bug: should be different in dev and prod
		if config.Config.Env == "dev" {
			c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		} else if config.Config.Env == "prod" || config.Config.Env == "test" {
			c.SetCookie("token", token, 3600, "/api", "qs.f1nley.xyz", false, true)
		}

		utility.ResponseOK(c, "OK", nil)
		return
	}
	utility.Response(c, 400, "Wrong password", nil)

}

func Register(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	name := user.Username
	pwd := user.Password
	u := model.User{}
	u.Admin = false
	u.Username = name
	if model.IsUserExist(name) {
		utility.Response(c, 400, "User already exists", nil)
		return
	}
	var err error
	user.Password, err = utility.PasswordHash(pwd)
	if err != nil {
		utility.Response(c, 500, "Password hash error", nil)
		return
	}
	err = model.AddUser(user)
	if err != nil {
		utility.Response(c, 400, "User already exists", nil)
		return
	}
	utility.ResponseOK(c, "OK", nil)
}
