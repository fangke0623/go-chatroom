package main

import (
	"log"
	"net/http"
	"wechat/src/common"
	"wechat/src/config"
)

func main() {
	http.HandleFunc("/", common.HandleResponse)
	config.Init()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
