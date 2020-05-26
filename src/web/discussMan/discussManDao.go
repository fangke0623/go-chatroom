package discussMan

import (
	"log"
	"wechat/src/config"
)

func SelectDiscussManList(form Form) []DiscussMan {
	var list []DiscussMan
	mysql := config.Mysql
	queryString := "select * from d_discuss_man where status != 3 "
	if form.DiscussId != "" {
		queryString += "and discuss_id=" + form.DiscussId

	}
	err := mysql.Select(&list, queryString)
	if err != nil {
		log.Println(err)
	}
	return list
}

func GetDiscussManById(manId int) DiscussMan {
	var discussMan DiscussMan
	mysql := config.Mysql
	queryString := "select * from d_discuss_man where man_id=" + string(manId) + "and status != 3 limit 1"

	err := mysql.Get(&discussMan, queryString)
	if err != nil {
		log.Println(err)
	}
	return discussMan
}

func findDiscussMan(form Form) DiscussMan {
	var discussMan DiscussMan
	mysql := config.Mysql
	queryString := "select * from d_discuss_man where status != 3"
	if form.DiscussId != "" {
		queryString += " and discuss_id = " + form.DiscussId
	}
	if form.UserId != "" {
		queryString += " and user_id = " + form.UserId
	}
	queryString += "limit 1"
	err := mysql.Get(&discussMan, queryString)
	if err != nil {
		log.Println(err)
	}
	return discussMan
}
func SaveDiscussMan(discussMan DiscussMan) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into d_discuss_man values (:man_id,:discuss_id,:user_id,:man_type,:remind,:status,:create_time,:modify_time)", &discussMan)
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

func UpdateDiscussMan(discussMan DiscussMan) {
	mysql := config.Mysql
	queryString := "update d_discuss_man set "
	if discussMan.ManType != "" {
		queryString += "man_type = :man_type,"
	}
	if discussMan.Remind != "" {
		queryString += "remind = :remind,"
	}
	if discussMan.Status != "" {
		queryString += "status = :status,"
	}
	queryString += "update_date=:update_date where man_id=:man_id"
	tx := mysql.MustBegin()
	result, err := tx.NamedExec(queryString, &discussMan)
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
