package dao

import (
	"wechat/src/config"
	"wechat/src/web/user/model"
)

func SelectUserList(form model.UserForm) []model.User {
	var userList []model.User
	sql := config.Sql
	queryString := "select * from f_user"
	if form.Id != "" {
		queryString += " where id = " + form.Id
	}
	_ = sql.Select(&userList, queryString)
	return userList
}
