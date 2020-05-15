package main

import (
	"log"
	"net/http"
	"wechat/src/config"
	"wechat/src/web/handle"
)

func main() {
	//ticker := time.NewTicker(time.Second * 10)
	//go func() {
	//	for _ = range ticker.C {
	//		log.Println("ticked at %v", time.Now())
	//	}
	//}()
	http.HandleFunc("/", handle.ResponseHandle)
	config.SqlInit()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
