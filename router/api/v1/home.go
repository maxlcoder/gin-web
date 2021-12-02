package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context)  {


	c.JSON(http.StatusOK, gin.H{
		"code": 123,
		"message": "ffg",
	})
}