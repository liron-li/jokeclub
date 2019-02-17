package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	PasswordError: "密码错误",
	InvalidParams: "参数错误",
	AccountExist:  "账号已经存在",
	NicknameExist: "昵称已经存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
