package discuss

import (
	"log"
	"wechat/src/config"
)

func SaveDiscuss(discuss Discuss) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into d_discuss values (:discuss_id,:discuss_title,:user_id,:visible_type,:status,:create_date,:update_date,:update_id)", &discuss)
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
func SelectDiscussList(form Form) []Discuss {
	var discuss []Discuss
	mysql := config.Mysql
	queryString := "select * from d_discuss where status = 1"
	if form.UserId != "" && form.DiscussTitle != "" {
		if form.UserId != "" {
			queryString += "and user_id = \"" + form.UserId + "\""
		}
		if form.DiscussTitle != "" {
			queryString += "and discuss_title like \"" + form.DiscussTitle + "%\""
		}
	}
	err := mysql.Select(&discuss, queryString)
	if err != nil {
		log.Println(err)
	}
	return discuss
}
