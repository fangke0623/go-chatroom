package user

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"wechat/src/common/exception"
	"wechat/src/common/util"
)

type Service interface {
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

	user := GetUserById(form.Id)
	jsons, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons, e
}
func (form Form) RegisterUser(param []byte) (interface{}, exception.Error) {

	user := User{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserByUsername(user.UserName)
	if dbUser.UserName != "" {
		e = exception.UserNameIsExist
		return result, e
	}
	if strings.Compare(user.Password, user.RePassword) == 0 {
		user.Id = util.GenerateUUID()
		user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		user.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
		SaveUser(user)
		e.ErrorMsg = "注册成功"
	} else {
		e = exception.PassWordIsInconsistent
		return result, e
	}
	return result, e
}
func (form Form) Login(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	user := GetUserByUsername(form.Username)
	if user.Password != "" {
		if strings.Compare(user.Password, form.Password) == 0 {
			return user, e
		} else {
			e = exception.PassWordIsWrong
		}
	} else {
		e = exception.UserNotExist
	}
	return user, e

}
func (form Form) EditUser(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	user := User{}
	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserById(user.Id)
	if dbUser.UserName == "" {
		e = exception.UserNotExist
		return "", e
	}

	user.Id = dbUser.Id
	user.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateUserById(user)
	e.ErrorMsg = "修改成功"

	return GetUserById(user.Id), e
}
