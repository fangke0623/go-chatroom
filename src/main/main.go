package main

import (
	"log"
	"net/http"
	"wechat/src/config"
)

func main() {
	http.HandleFunc("/", HandleResponse)
	config.SqlInit()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
