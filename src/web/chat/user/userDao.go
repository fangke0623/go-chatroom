package user

import (
	"encoding/json"
	"log"
	"wechat/src/common/enum"
	"wechat/src/common/query"
	"wechat/src/common/util"
	"wechat/src/config"
)

func SelectUserList(form Form) []User {
	var userList []User
	mysql := config.Mysql
	queryString := "select * from f_user"
	if form.Id != "" {
		queryString += " where id = " + form.Id
	}
	queryString += " order by modify_time"
	if form.PageSize != "" && form.CurrentPage != "" {
		queryString += " limit " + query.GetStartRow(form.Page) + " , " + form.PageSize
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
	result, err := tx.NamedExec("insert into f_user values (:id,:email,:user_name,:password,:create_time,:modify_time,:nickname,:mobile)", &user)
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
func GetUserById(id string) User {
	var user = User{}

	jsons := config.DoGet(enum.UserCache + id)

	if util.IsNil(jsons) {
		mysql := config.Mysql
		queryString := "select * from f_user where id = \"" + id + "\" limit 1"
		err := mysql.Get(&user, queryString)
		if err != nil {
			log.Println(err)
		}
		config.DoSet(enum.UserCache+user.Id, jsons)
		config.DoExpire(enum.UserCache+user.Id, 7*24*3600)
	} else {
		err := json.Unmarshal(jsons, &user)
		if err != nil {
			log.Println(err)
		}
	}
	return user
}
func UpdateUserById(user User) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	queryString := "update f_user set "
	if user.Password != "" {
		queryString += "password=:password,"
	}
	if user.Nickname != "" {
		queryString += "nickname=:nickname,"
	}
	if user.Mobile != "" {
		queryString += "mobile=:mobile,"
	}
	if user.Email != "" {
		queryString += "email=:email,"
	}
	queryString += "modify_time=:modify_time where id=:id"
	result, err := tx.NamedExec(queryString, &user)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if result != nil {
		log.Println(result)
	}

	if err != nil {
		log.Fatal(err)
	}
}
