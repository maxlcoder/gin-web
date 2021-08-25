package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/middleware/jwt"
	"github.com/maxlcoder/gin-web/pkg/setting"
	"github.com/maxlcoder/gin-web/router/api"
	v1 "github.com/maxlcoder/gin-web/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.POST("login", api.Login)

	apiV1 := r.Group("api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)

		apiV1.POST("/tags", v1.CreateTag)

		apiV1.PUT("/tags/:id", v1.UpdateTag)

		apiV1.DELETE("/tags", v1.DeleteTag)
	}

	return r
}