package user

import (
	"log"
	"wechat/src/config"
)

func SelectUserList(form UserForm) []User {
	var userList []User
	mysql := config.Mysql
	queryString := "select * from f_user"
	if form.Id != "" {
		queryString += " where id = " + form.Id
	}
	_ = mysql.Select(&userList, queryString)
	return userList
}
func GetUserByUsername(username string) User {
	var user = User{}
	mysql := config.Mysql
	queryString := "select * from f_user where username = " + username
	_ = mysql.Select(&user, queryString)
	return user
}

func SaveUser(user User) {
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
