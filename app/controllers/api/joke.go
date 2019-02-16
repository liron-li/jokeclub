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
 * @apiParam {int} [page] 页码
 * @apiParam {int} [pageSize] 每页条数
 * @apiParam {int} [type] 1:推荐 2：最新 3：图片 4：视频
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 * @apiSuccess {int} data.id  id
 * @apiSuccess {int} data.user_id  用户id
 * @apiSuccess {string} data.content  内容
 * @apiSuccess {string} data.image  图片资源路径
 * @apiSuccess {string} data.video  视频资源路径
 * @apiSuccess {int} data.up_num  支持数
 * @apiSuccess {int} data.down_num  反对数
 * @apiSuccess {int} data.favorite_num  收藏数
 * @apiSuccess {int} data.comment_num  投票数
 * @apiSuccess {int} data.type  类型 0: 文本段子 1：图片 2：视频
 * @apiSuccess {string} data.created_at  创建时间
 * @apiSuccess {string} data.updated_at  更新时间
 * @apiSuccess {string} data.deleted_at  删除时间
 *
 * @apiSampleRequest http://localhost:8000/api/jokes
 */
func Jokes(c *gin.Context) {

	var data interface{}
	order := "id desc"
	maps := make(map[string]interface{})

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	jokeType := c.DefaultQuery("type", "1")

	valid := validation.Validation{}
	a := paginateRequest{Page: page, PageSize: pageSize}
	ok, _ := valid.Valid(&a)

	code := e.InvalidParams

	switch jokeType {
	case "1": // 推荐
		order = "up_num desc"
	case "2": // 最新
		order = "id desc"
	case "3": // 图片
		maps["type"] = models.PicType
	case "4": // 视频
		maps["type"] = models.VideoTYpe
	}

	if ok {
		data = models.JokePaginate(c, page, pageSize, maps, order)
		code = e.Success
	}

	c.JSON(http.StatusOK, util.RetJson(code, data))
}
