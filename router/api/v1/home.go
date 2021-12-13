package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/maxlcoder/gin-web/pkg/logging"
	"net/http"
)

func Home(c *gin.Context)  {

	logging.Debug("test logging debug log")
	log.Debug("test debug log")
	log.Info("test info log")
	c.JSON(http.StatusOK, gin.H{
		"code": 123,
		"message": "ffg",
	})
}