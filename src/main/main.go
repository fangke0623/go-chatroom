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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
