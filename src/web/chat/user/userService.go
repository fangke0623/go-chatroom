package user

import (
	"strings"
	"time"
	"wechat/src/common/enum"
	"wechat/src/common/exception"
	"wechat/src/common/util"
	"wechat/src/config"
)

type ServiceUser interface {
	FindUserList(param []byte) (interface{}, exception.Error)

	DetailUser(param []byte) (interface{}, exception.Error)

	RegisterUser(param []byte) (interface{}, exception.Error)

	Login(param []byte) (interface{}, exception.Error)

	EditUser(param []byte) (interface{}, exception.Error)
}

func (form Form) FindUserList(param []byte) (interface{}, exception.Error) {

	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)

	list := SelectUserList(form)
	return list, e
}
func (form Form) DetailUser(param []byte) (interface{}, exception.Error) {

	e := exception.Error{}

	util.HandleParamsToStruct(param, &form)
	if form.Id == "" {
		return exception.ThrowException(exception.ParamMiss)
	}
	user := GetUserById(form.Id)

	return user, e
}
func (form Form) RegisterUser(param []byte) (interface{}, exception.Error) {
	user := User{}
	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserByUsername(user.UserName)
	if dbUser.UserName != "" {
		return exception.ThrowException(exception.UserNameIsExist)
	}
	if strings.Compare(user.Password, user.RePassword) == 0 {
		user.Id = util.GenerateUUID()
		user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		user.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
		SaveUser(user)
		return exception.ThrowException(exception.Success)
	} else {
		return exception.ThrowException(exception.PassWordIsInconsistent)
	}
}
func (form Form) Login(param []byte) (interface{}, exception.Error) {
	e := exception.Success
	util.HandleParamsToStruct(param, &form)
	user := GetUserByUsername(form.Username)
	if user.Password != "" {
		if strings.Compare(user.Password, form.Password) == 0 {
			return user, e
		} else {
			return exception.ThrowException(exception.PassWordIsWrong)
		}
	} else {
		return exception.ThrowException(exception.UserNotExist)
	}

}
func (form Form) EditUser(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	user := User{}
	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserById(user.Id)
	if dbUser.UserName == "" {
		e = exception.UserNotExist
		return exception.ThrowException(exception.UserNotExist)
	}

	user.Id = dbUser.Id
	user.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateUserById(user)
	e.ErrorMsg = "修改成功"
	config.DoDel(enum.UserCache + user.Id)
	return GetUserById(user.Id), e
}
