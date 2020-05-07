package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var Sql *sqlx.DB

func Init() {

	db, err := sqlx.Open("mysql", "root:.Fangke123@tcp(39.108.145.221:3306)/chatroom?charset=utf8")
	if err != nil {
		log.Println("数据库连接异常", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Second)
	Sql = db

}
