package main

import "github.com/maxlcoder/gin-web/cmd"

//func main() {
//	r := router.InitRouter()
//	s := &http.Server{
//		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
//		Handler: r,
//		ReadTimeout: setting.ReadTimeout,
//		WriteTimeout: setting.WriteTimeout,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	s.ListenAndServe()
//}

func main() {
	cmd.Execute()
}
