package api

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"jokeclub/app/models"
	"jokeclub/pkg/cache"
	"jokeclub/pkg/e"
	"jokeclub/pkg/util"
	"net/http"
)

/**
 * @api {get} /api/jokes 获取段子列表数据
 * @apiGroup jokes
 *
 * @apiParam {string} token token
 * @apiParam {int} [page] 页码
 * @apiParam {int} [pageSize] 每页条数
 * @apiParam {int} [type] 1:推荐 2：最新 3：图片 4：视频
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 * @apiSuccess {int} data.id  id
 * @apiSuccess {int} data.user_id  用户id
 * @apiSuccess {object} data.user  用户信息
 * @apiSuccess {string} data.user.name  用户名
 * @apiSuccess {string} data.user.nickname  昵称
 * @apiSuccess {string} data.user.email  邮箱
 * @apiSuccess {string} data.user.avatar  头像
 * @apiSuccess {string} data.user.slogan  签名

 * @apiSuccess {string} data.user.created_at  创建时间
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

	order := "id desc"
	maps := make(map[string]interface{})

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	jokeType := c.DefaultQuery("type", "1")

	rules := govalidator.MapData{
		"page":     []string{"numeric", "numeric_between:1,9999999"},
		"pageSize": []string{"numeric", "numeric_between:1,100"},
		"jokeType": []string{"digits:1"},
	}

	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		RequiredDefault: false,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	res := v.Validate()

	// 如果参数验证失败
	if len(res) > 0 {
		util.ReturnInvalidParamsJson(c, res)
		return
	}

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

	data := models.JokePaginate(c, page, pageSize, maps, order)

	util.ReturnSuccessJson(c, data)
}

/**
 * @api {post} /api/user/up 支持
 * @apiGroup jokes
 *
 * @apiParam {string} token token
 * @apiParam {int} joke_id 段子id
 * @apiParam {int} cancel 是否取消 0：否 1:是
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/joke/up
 */
func JokeUp(c *gin.Context) {

	rules := govalidator.MapData{
		"joke_id": []string{"numeric", "required"},
		"cancel":  []string{"numeric"},
	}

	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		RequiredDefault: false,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	res := v.Validate()

	// 如果参数验证失败
	if len(res) > 0 {
		util.ReturnInvalidParamsJson(c, res)
		return
	}

	jokeId := c.PostForm("joke_id")
	cancel := c.DefaultPostForm("cancel", "0")

	idInt, err := com.StrTo(jokeId).Int()
	joke := models.GetJoke(idInt)

	if joke.ID <= 0 || err != nil {
		util.ReturnErrorJson(c, e.Error)
		c.Abort()
		return
	}

	user := cache.UserProfile(c)

	if cancel != "0" {
		joke.Up(user.ID, true)
	} else {
		joke.Up(user.ID, false)
	}

	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {post} /api/jokes/down 反对
 * @apiGroup jokes
 *
 * @apiParam {string} token token
 * @apiParam {int} joke_id 段子id
 * @apiParam {int} cancel 是否取消 0：否 1:是
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/jokes/down
 */
func JokeDown(c *gin.Context) {
	rules := govalidator.MapData{
		"joke_id": []string{"numeric", "required"},
		"cancel":  []string{"numeric"},
	}

	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		RequiredDefault: false,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	res := v.Validate()

	// 如果参数验证失败
	if len(res) > 0 {
		util.ReturnInvalidParamsJson(c, res)
		return
	}

	jokeId := c.PostForm("joke_id")
	cancel := c.DefaultPostForm("cancel", "0")

	idInt, err := com.StrTo(jokeId).Int()
	joke := models.GetJoke(idInt)

	if joke.ID <= 0 || err != nil {
		util.ReturnErrorJson(c, e.Error)
		c.Abort()
		return
	}

	user := cache.UserProfile(c)

	if cancel != "0" {
		joke.Down(user.ID, true)
	} else {
		joke.Down(user.ID, false)
	}

	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {post} /api/jokes/favorite 收藏
 * @apiGroup jokes
 *
 * @apiParam {string} token token
 * @apiParam {int} joke_id 段子id
 * @apiParam {int} cancel 是否取消 0：否 1:是
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/jokes/favorite
 */
func JokeFavorite(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {get} /api/jokes/comments 获取评论
 * @apiGroup jokes
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/jokes/comments
 */
func Comments(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {post} /api/jokes/comments 发起评论
 * @apiGroup jokes
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/jokes/comments
 */
func PostComments(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}
