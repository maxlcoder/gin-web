package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/models"
	"net/http"
)

func Home(c *gin.Context)  {

	//var user = models.User{Name: "hello"}

	user := new(models.User)
	user.Name = "fdsf"
	name := user.Bye()



	c.JSON(http.StatusOK, gin.H{
		"code": 123,
		"message": name + "----" + user.Name,
	})
}