package dao

import (
	"wechat/src/config"
	"wechat/src/web/user/model"
)

func SelectUserList(form model.UserForm) []model.User {
	var userList []model.User
	mysql := config.Mysql
	queryString := "select * from f_user"
	if form.Id[0] != "" {
		queryString += " where id = " + form.Id[0]
	}
	_ = mysql.Select(&userList, queryString)
	return userList
}
