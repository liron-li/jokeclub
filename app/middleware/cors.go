package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:6666"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return false
		},
		MaxAge: 12 * time.Hour,
	})
}
