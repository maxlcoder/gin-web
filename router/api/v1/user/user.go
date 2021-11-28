package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/services/user_service"
	"net/http"
)

func CreateUser(c *gin.Context)  {
	user := user_service.User{
		Name: "xxxx",
	}
	result := user.CreateUser()
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
		"data": result,
	})
}
