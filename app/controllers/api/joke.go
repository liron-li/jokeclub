package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/e"
	"jokeclub/pkg/util"
	"github.com/astaxie/beego/validation"
	"jokeclub/app/models"
)

type paginateRequest struct {
	Page     string `valid:"Required; MaxSize(50)"`
	PageSize string `valid:"Required; MaxSize(50)"`
}

func Jokes(c *gin.Context) {

	var data interface{}
	maps := make(map[string]interface{})

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	valid := validation.Validation{}
	a := paginateRequest{Page: page, PageSize: pageSize}
	ok, _ := valid.Valid(&a)

	code := e.InvalidParams

	if ok {
		data = models.JokePaginate(c, page, pageSize, maps)
		code = e.Success
	}

	c.JSON(http.StatusOK, util.RetJson(code, data))
}
