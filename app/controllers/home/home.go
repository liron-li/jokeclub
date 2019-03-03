package home

import (
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/util"
)

func Index(c *gin.Context) {
	//c.HTML(http.StatusOK, "home.html", gin.H{"data": "world"})
	util.ReturnSuccessJson(c, []string{})
}
