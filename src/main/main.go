package main

import (
	"chatroom/src/common"
	"chatroom/src/config"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", common.HandleResponse)
	config.Init()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
