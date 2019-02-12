package util

import (
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/e"
)

func RetJson(code int, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	}
}
