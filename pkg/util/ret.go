package util

import (
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/e"
	"net/http"
)

func RetJson(code int, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	}
}

func ReturnInvalidParamsJson(c *gin.Context, errMsg map[string][]string) {
	c.JSON(http.StatusOK, RetJson(e.InvalidParams, errMsg))
}

func ReturnSuccessJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, RetJson(e.Success, data))
}

func ReturnErrorJson(c *gin.Context, code int) {
	c.JSON(http.StatusOK, RetJson(code, nil))
}

func ReturnJson(c *gin.Context, code int, data interface{})  {
	c.JSON(http.StatusOK, RetJson(code, data))
}
