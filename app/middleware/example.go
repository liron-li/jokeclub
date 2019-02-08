package middleware

import (
	"github.com/gin-gonic/gin"
)

func Example() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
