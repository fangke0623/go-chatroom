package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var Mysql *sqlx.DB

const d = "root:.Fangke123@tcp(39.108.145.221:3306)/chatroom?charset=utf8"

type DbConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func SqlInit(conf DbConf) {
	dataSourceName := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ")/" + conf.Dbname + "?charset=utf8"
	db, err := sqlx.Open("mysql", dataSourceName)
	if db == nil {
		return
	}
	if err != nil {
		log.Println("mysql connect error", err.Error())
	} else {
		log.Println("mysql connect success")
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Second)
	Mysql = db

}
