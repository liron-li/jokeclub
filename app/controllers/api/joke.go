package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"jokeclub/app/models"
	"jokeclub/pkg/e"
	"jokeclub/pkg/util"
	"net/http"
)

type paginateRequest struct {
	Page     string `valid:"Required; MaxSize(50)"`
	PageSize string `valid:"Required; MaxSize(50)"`
}

/**
 * @api {get} /api/jokes 获取段子列表数据
 * @apiName getJokes
 * @apiGroup jokes
 *
 * @apiParam {Int} [page] 页码
 * @apiParam {Int} [pageSize] 每页条数
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 *
 * @apiSampleRequest http://localhost:8000/api/jokes
 */
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
