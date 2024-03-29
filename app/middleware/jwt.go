package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"jokeclub/pkg/e"
	"jokeclub/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.Success
		token := util.GetToken(c)

		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.InvalidJwtToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.JwtTokenTimeout
			}
		}

		if code != e.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
