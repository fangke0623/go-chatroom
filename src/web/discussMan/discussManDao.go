package discussMan

import (
	"log"
	"wechat/src/config"
)

func SelectDiscussManList(form Form) []DiscussMan {
	var list []DiscussMan
	mysql := config.Mysql
	queryString := "select * from d_discuss_man where status != 3"
	if form.DiscussId != 0 {
		if form.DiscussId != 0 {
			queryString += "and discuss_id = " + string(form.DiscussId)
		}
	}
	err := mysql.Select(&list, queryString)
	if err != nil {
		log.Println(err)
	}
	return list
}
func SaveDiscussMan(discussMan DiscussMan) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into d_discuss_man values (:man_id,:discuss_id,:member_id,:man_type,:remind,:status,:create_date,:update_date)", &discussMan)
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
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("update d_discuss_man set man_type = :man_type,remind = :remind,status = :status,update_date=:update_date where man_id=:man_id", &discussMan)
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
