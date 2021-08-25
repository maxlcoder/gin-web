package main

import (
	"fmt"
	"github.com/maxlcoder/gin-web/pkg/setting"
	"github.com/maxlcoder/gin-web/router"
	"net/http"
)

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: r,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
