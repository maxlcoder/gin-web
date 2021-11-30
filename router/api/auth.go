package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/models"
	"github.com/maxlcoder/gin-web/pkg/e"
	"github.com/maxlcoder/gin-web/pkg/logging"
	"github.com/maxlcoder/gin-web/pkg/util"
	"net/http"
)

type Auth struct {
	Username string `form:"username" binding:"required,max=10"`
	Password string `form:"password" binding:"required,max=10"`
}

func Login(c *gin.Context)  {
	var auth Auth
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if err := c.ShouldBind(&auth); err != nil {
		logging.Info(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"message": err.Error(),
			"data": data,
		})
		return
	}

	isExist := models.CheckAuth(auth.Username, auth.Password)
	if !isExist {
		code = e.ERROR_AUHT
	} else {
		token, err := util.GenerateToken(auth.Username, auth.Password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": e.GetMsg(code),
		"data": data,
	})
}