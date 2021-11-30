package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
	Err interface{} `json:"err,omitempty"`
}

func (g *Gin) Success(message string, data interface{})  {
	g.C.JSON(http.StatusOK, Response{
		Code: 200,
		Message: message,
		Data: data,
	})
	return
}

func (g *Gin) Fail(status int, code int, message string, err interface{})  {
	g.C.JSON(status, Response{
		Code: code,
		Message: message,
		Err: err,
	})
	return
}
