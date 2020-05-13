package exception

var (
	//user
	PassWordIsWrong        = Error{10, "密码错误"}
	PassWordIsInconsistent = Error{11, "两次输入密码不一致"}
	UserNotExist           = Error{12, "用户不存在"}
	UserNameIsExist        = Error{13, "用户名已存在"}

	//discuss

	//discussMan

	//discussMsg

)
