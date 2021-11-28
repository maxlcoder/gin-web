package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/models"
	"github.com/maxlcoder/gin-web/models/user_model"
	"github.com/maxlcoder/gin-web/pkg/setting"
	"github.com/maxlcoder/gin-web/router"
)

func init() {
	// migration
	models.DB.AutoMigrate(&user_model.User{})
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 全局中间件
	r.Use()
	// 加载路由
	router.LoadApiRouter(r)
	// 启动端口
	r.Run(setting.ServerSetting.HTTTPort)
}
