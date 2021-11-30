package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/pkg/response"
	"github.com/maxlcoder/gin-web/services"
	"net/http"
)

func CreateUser(c *gin.Context)  {
	user := services.User{
		Name: "zhangsan",
	}

	result := user.CreateUser()

	g := response.Gin{C: c}
	g.Success("success", result)
}

func Me(c *gin.Context) {
	user := services.User{
		Id: 1,
	}
	result, err := user.GetUser()
	g := response.Gin{C: c}
	if err != nil {
		g.Fail(http.StatusNotFound, 404, "当前用户不存在", err.Error())
		return
	}

	g.Success("success", result)
}
