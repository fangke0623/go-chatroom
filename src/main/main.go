package main

import (
	"log"
	"net/http"
	"wechat/src/config"
	"wechat/src/web/handle"
)

func main() {
	config.LogInit()
	http.HandleFunc("/", handle.ResponseHandle)
	config.SqlInit()
	//config.RedisInit()
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Println(err)
	}
}
