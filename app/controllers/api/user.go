package api

import (
	"net/http"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"jokeclub/app/models"
	"jokeclub/pkg/e"
	"jokeclub/pkg/logging"
	"jokeclub/pkg/util"
)

/**
 * @api {get} /api/user/profile 获取用户详细信息
 * @apiName userProfile
 * @apiGroup user
 *
 * @apiParam {string} token token

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

	util.ReturnSuccessJson(c, models.GetUserProfile(claims.UserId))
}

/**
 * @api {post} /api/login 登录
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

	rules := govalidator.MapData{
		"username": []string{"required", "max:32", "alpha_num", "min:4"},
		"password": []string{"max:32", "min:6", "alpha_num"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:账号不能为空",
			"max:账号最多32个字符",
			"alpha_num:账号只能是数字和字母",
			"min:账号至少有4个字符",
		},
		"password": []string{"max:密码最多32个字符", "min:密码至少有6个字符", "alpha_num:密码只能是数字和字母"},
	}

	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		Messages:        messages,  // custom message map (Optional)
		RequiredDefault: false,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	res := v.Validate()

	// 如果参数验证失败
	if len(res) > 0 {
		util.ReturnInvalidParamsJson(c, res)
		return
	}

	var code int
	data := make(map[string]interface{})

	userAuth, isExist := models.CheckAuth(username, password)
	if isExist {
		token, err := util.GenerateToken(userAuth.UserId, username)
		if err != nil {
			code = e.Error
		} else {
			data["token"] = token
			code = e.Success
		}

	} else {
		code = e.PasswordError
	}

	util.ReturnJson(c, code, data)
}

/**
 * @api {post} /api/register 注册
 * @apiGroup user
 *
 * @apiParam {int} type 类型 0：账号 1：手机号 2：邮箱 3：微信
 * @apiParam {string} identify 授权标识
 * @apiParam {string} [password] 密码
 * @apiParam {string} nickname 昵称
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/register
 */
func Register(c *gin.Context) {

	typeValue := c.PostForm("type")
	identify := c.PostForm("identify")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	typeInt, err := com.StrTo(typeValue).Int()

	if err != nil {
		util.ReturnErrorJson(c, e.Error)
		return
	}

	rules := govalidator.MapData{
		"type": []string{"required", "digits:1"},
		"identify": func(typeValue int) []string {
			var r []string
			switch typeValue {
			case 1: // 手机
				r = []string{"required", "max:32", "digits_between:6,11", "min:4"}
			case 2: // 邮箱
				r = []string{"required", "max:32", "email", "min:4"}
			default:
				r = []string{"required", "max:32", "alpha_num", "min:4"}
			}
			return r
		}(typeInt),
		"password": []string{"max:32", "min:6", "alpha_num"},
		"nickname": []string{"required", "max:18", "min:4"},
	}

	messages := govalidator.MapData{
		"identify": []string{
			"required:账号不能为空",
			"max:账号最多32个字符",
			"alpha_num:账号只能是数字和字母",
			"min:账号至少有4个字符",
			"email:邮箱格式错误",
			"digits_between:手机号格式不正确",
		},
		"password": []string{"max:密码最多32个字符", "min:密码至少有6个字符", "alpha_num:密码只能是数字和字母"},
		"nickname": []string{"required:昵称不能为空", "max:账号最多32个字符", "min:昵称至少4个字符"},
	}

	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		Messages:        messages,  // custom message map (Optional)
		RequiredDefault: false,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	res := v.Validate()

	// 如果参数验证失败
	if len(res) > 0 {
		util.ReturnInvalidParamsJson(c, res)
		return
	}

	// 如果账号已经存在
	if models.CheckUserAuthExist(identify, typeInt) {
		util.ReturnErrorJson(c, e.AccountExist)
		return
	}
	// 昵称存在
	if models.CheckUserExist(models.User{Nickname: nickname}) {
		util.ReturnErrorJson(c, e.NicknameExist)
		return
	}

	errors := models.DoRegister(identify, typeInt, password, nickname)

	if errors != nil {
		logging.Error(errors)
		util.ReturnErrorJson(c, e.Error)
		return
	}

	util.ReturnSuccessJson(c, nil)
}

/**
 * @api {get} /api/user/my-message 私信
 * @apiGroup user
 *
 * @apiParam {string} token token
 *
 * @apiSuccess {int} code  状态码 0：成功，其他表示错误
 * @apiSuccess {string} msg  消息
 * @apiSuccess {array} data  数据体
 *
 * @apiSampleRequest http://localhost:8000/api/user/my-message
 */
func MyMessages(c *gin.Context) {

	order := "id desc"
	maps := make(map[string]interface{})

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	rules := govalidator.MapData{
		"page":     []string{"numeric", "numeric_between:1,9999999"},
		"pageSize": []string{"numeric", "numeric_between:1,100"},
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

	data := models.MessageSessionPaginate(c, page, pageSize, maps, order)

	c.JSON(http.StatusOK, util.RetJson(e.Success, data))
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
