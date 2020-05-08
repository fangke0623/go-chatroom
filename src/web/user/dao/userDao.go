package dao

import (
	"log"
	"wechat/src/config"
	"wechat/src/web/user/model"
)

func SelectUserList(form model.UserForm) []model.User {
	var userList []model.User
	mysql := config.Mysql
	queryString := "select * from f_user"
	if form.Id != "" {
		queryString += " where id = " + form.Id
	}
	_ = mysql.Select(&userList, queryString)
	return userList
}

func SaveUser(user model.User) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into f_user values (:id,:email,:user_name,:password,:create_time,:nickname,:mobile)", &user)
	if err != nil {
		log.Println(err)
	}
	if result != nil {
		log.Println(result)
	}
	tx.Commit()
}
