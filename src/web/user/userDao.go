package user

import (
	"log"
	"wechat/src/config"
)

func SelectUserList(form Form) []User {
	var userList []User
	mysql := config.Mysql
	queryString := "select * from f_user"
	if form.Id != "" {
		queryString += " where id = " + form.Id
	}
	err := mysql.Select(&userList, queryString)
	if err != nil {
		log.Println(err)
	}
	return userList
}
func GetUserByUsername(username string) User {
	var user = User{}
	mysql := config.Mysql
	queryString := "select * from f_user where user_name = \"" + username + "\""
	err := mysql.Get(&user, queryString)
	if err != nil {
		log.Println(err)
	}
	return user
}

func SaveUser(user User) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into f_user values (:id,:email,:user_name,:password,:create_time,:nickname,:mobile)", &user)
	if err != nil {
		log.Println(err)
	}
	err = tx.Commit()
	if result != nil {
		log.Println(result)
	}

	if err != nil {
		log.Println(err)
	}
}
