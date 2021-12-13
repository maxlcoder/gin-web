package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/pkg/setting"
	"github.com/maxlcoder/gin-web/router"
	"io"
	"os"
)

func main() {
	// 设置模式
	gin.SetMode(os.Getenv("GIN_MODE"))
	// 日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 全局中间件
	r.Use()
	// 加载路由
	router.LoadApiRouter(r)
	// 启动端口
	r.Run(setting.ServerSetting.HTTTPort)
}
