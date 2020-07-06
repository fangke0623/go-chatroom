package main

import (
	"log"
	"net"
	"net/http"
	"wechat/src/config"
	"wechat/src/web/handle"
)

var c = config.Conf{}

func init() {
	conf := c.GetConf()
	config.LogInit(conf.Log)
	config.SqlInit(conf.Db)
	config.RedisInit(conf.Redis)

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8090")
	go handle.Tcp(listener)
	//timer.TimerFunc()
	http.HandleFunc("/", handle.ResponseHandle)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("error", err.Error())
	}

}
