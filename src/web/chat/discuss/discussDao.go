package discuss

import (
	"log"
	"strconv"
	"wechat/src/config"
)

func SaveDiscuss(discuss Discuss) int64 {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into d_discuss values (:discuss_id,:discuss_title,:user_id,:visible_type,:status,:create_time,:modify_time,:modify_id)", &discuss)
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
	discussId, _ := result.LastInsertId()
	return discussId
}
func SelectDiscussList(form Form) []Discuss {
	var discuss []Discuss
	mysql := config.Mysql
	queryString := "select * from d_discuss where status = 1"
	if form.UserId != "" {
		queryString += " and user_id = \"" + form.UserId + "\""
	}
	if form.DiscussTitle != "" {
		queryString += " and discuss_title like \"" + form.DiscussTitle + "%\""
	}
	err := mysql.Select(&discuss, queryString)
	if err != nil {
		log.Println(err)
	}
	return discuss
}
func GetDiscussById(discussId int64) Discuss {
	var discuss Discuss
	mysql := config.Mysql
	queryString := "select * from d_discuss where status = 1" + " and discuss_id = \"" + strconv.FormatInt(discussId, 10) + "\" limit 1"
	err := mysql.Get(&discuss, queryString)
	if err != nil {
		log.Println(err)
	}
	return discuss
}
func UpdateDiscuss(discuss Discuss) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	queryString := "update d_discuss set "
	if discuss.DiscussTitle != "" {
		queryString += "discuss_title=:discuss_title,"
	}
	if discuss.VisibleType != "" {
		queryString += "visible_type=:visible_type,"
	}
	if discuss.Status != "" {
		queryString += "status=:status,"
	}
	queryString += "modify_time=:modify_time where discuss_id=:discuss_id"
	result, err := tx.NamedExec(queryString, &discuss)
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
