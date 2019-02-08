package util

import (
	"github.com/gin-gonic/gin"
	"goweb/pkg/e"
)

func RetJson(code int, data interface{}) gin.H {
	return gin.H{
		"code": e.Success,
		"msg":  e.GetMsg(e.Success),
		"data": data,
	}
}
