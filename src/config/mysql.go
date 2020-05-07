package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var Mysql *sqlx.DB

const dataSourceName = "root:.Fangke123@tcp(39.108.145.221:3306)/wechat?charset=utf8"

func SqlInit() {

	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("数据库连接异常", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Second)
	Mysql = db

}
