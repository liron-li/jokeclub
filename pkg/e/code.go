package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	PasswordError: "密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
