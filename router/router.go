package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/router/api"
	v1 "github.com/maxlcoder/gin-web/router/api/v1"
	"github.com/maxlcoder/gin-web/router/api/v1/user"
)

func LoadApiRouter(r *gin.Engine)  {
	r.POST("login", api.Login)

	r.GET("/", v1.Home)

	apiV1 := r.Group("api/v1")
	apiV1.Use()
	{
		apiV1.POST("/users", user.CreateUser)
	}

}