package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"jokeclub/app/models"
	"jokeclub/pkg/e"
	"jokeclub/pkg/logging"
	"jokeclub/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

/**
 * @api {get} /api/user/profile 获取用户详细信息
 * @apiName userProfile
 * @apiGroup user
 *
 * @apiParam {string} token 页码

 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/user/profile
 */
func Profile(c *gin.Context) {
	token := c.Query("token")
	claims, err := util.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(401)
	}

	c.JSON(http.StatusOK, util.RetJson(e.Success, claims))
}

/**
 * @api {get} /api/login 登录
 * @apiName userLogin
 * @apiGroup user
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/login
 */
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.InvalidParams

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.Error
			} else {
				data["token"] = token
				code = e.Success
			}

		} else {
			code = e.PasswordError
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, util.RetJson(code, data))
}

/**
 * @api {get} /api/register 注册
 * @apiGroup user
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/register
 */
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {get} /api/user/my-message 私信
 * @apiGroup user
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/user/my-message
 */
func MyMessage(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {get} /api/user/my-up-jokes 我赞过的
 * @apiGroup user
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/user/my-up-jokes
 */
func MyUpedJokes(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {get} /api/user/my-favorite 我的收藏
 * @apiGroup user
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/user/my-favorite
 */
func MyFavorite(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}

/**
 * @api {get} /api/user/my-feedback 意见反馈
 * @apiGroup user
 *
 * @apiParam {string} username 用户名称
 * @apiParam {string} password 密码
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/user/my-feedback
 */
func Feedback(c *gin.Context) {
	c.JSON(http.StatusOK, util.RetJson(e.Success, ""))
}
