package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/pkg/e"
	"net/http"
)

func GetTags(c *gin.Context)  {
	code := e.SUCCESS
	message := e.GetMsg(code)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
		"data": "data",
	})
}

func CreateTag(c *gin.Context)  {
	name := c.PostForm("name")
	fmt.Println(name)
}

func UpdateTag(c *gin.Context)  {

}

func DeleteTag(c *gin.Context)  {

}