package discussMsg

import (
	"log"
	"wechat/src/config"
)

func SelectDiscussMsgList(form Form) []DiscussMsg {
	var list []DiscussMsg
	mysql := config.Mysql
	queryString := "select * from d_discuss_msg where status = 1"
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
func SaveDiscussMsg(msg DiscussMsg) {
	mysql := config.Mysql
	tx := mysql.MustBegin()
	result, err := tx.NamedExec("insert into d_discuss_msg values (:msg_id,:discuss_id,:man_id,:msg_content,:create_date,:update_date,:msg_type,:status)", &msg)
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
