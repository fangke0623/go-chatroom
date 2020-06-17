package main

import (
	"log"
	"net/http"
	"wechat/src/config"
	"wechat/src/web/handle"
)

func init() {
	//config.LogInit()
	config.SqlInit()
	//config.RedisInit()

}
func main() {

	http.HandleFunc("/", handle.ResponseHandle)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Println(err)
	}
}
