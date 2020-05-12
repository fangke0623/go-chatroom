package user

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"wechat/src/common/util"
)

func FindUserList(param []byte) []byte {

	form := Form{}
	util.HandleParamsToStruct(param, &form)

	list := SelectUserList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons
}
func RegisterUser(param []byte) []byte {

	user := User{}
	result := ""

	util.HandleParamsToStruct(param, &user)
	dbUser := GetUserByUsername(user.UserName)
	if dbUser.UserName != "" {
		result = "用户名已存在"
		return []byte(result)
	}
	if strings.Compare(user.Password, user.RePassword) == 1 {
		user.Id = util.GenerateUUID()
		user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		SaveUser(user)
		result = "注册成功"
	} else {
		result = "两次密码输入不一致"
		return []byte(result)
	}
	return []byte(result)
}
func Login(param []byte) []byte {
	var result []byte
	form := Form{}
	util.HandleParamsToStruct(param, &form)
	user := GetUserByUsername(form.Username)
	if user.Password != "" {
		if strings.Compare(user.Password, form.Password) == 1 {
			jsons, err := json.Marshal(user)
			if err != nil {
				fmt.Println("error:", err)
			}
			result = jsons
		} else {
			result = []byte("密码不正确，请重新输入")

		}
	} else {
		result = []byte("该用户不存在")
	}
	return result

}
