package main

import (
	"log"
	"net/http"
	"wechat/src/config"
	"wechat/src/web/handle"
)

func main() {
	http.HandleFunc("/", handle.ResponseHandle)
	config.SqlInit()
	//config.RedisInit()
	log.Fatal(http.ListenAndServe(":8090", nil))

}
