package jwt

import (
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/maxlcoder/gin-web/pkg/e"
	"github.com/maxlcoder/gin-web/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//var data interface{}
		authorizationHeader := request.AuthorizationHeaderExtractor
		token, err := authorizationHeader.ExtractToken(c.Request)
		code = e.SUCCESS
		if err != nil {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"message": e.GetMsg(code),
				"data": new(struct{}),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
