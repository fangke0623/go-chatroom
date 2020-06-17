package user

import "wechat/src/common/exception"

func Handle(path string, param []byte) (interface{}, exception.Error) {
	var result interface{}
	var userForm Form
	var e exception.Error
	switch path {
	case "/user/list":
		result, e = userForm.FindUserList(param)
		break
	case "/user/detail":
		result, e = userForm.DetailUser(param)
		break
	case "/user/register":
		result, e = userForm.RegisterUser(param)
		break
	case "/user/login":
		result, e = userForm.Login(param)
		//cookie.SetCookie(writer, result)
		break
	case "/user/edit":
		result, e = userForm.EditUser(param)
		break
	}
	return result, e
}
