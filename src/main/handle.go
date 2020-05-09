package main

import (
	"encoding/json"
	"log"
	"net/http"
	"wechat/src/web/user"
)

func HandleResponse(writer http.ResponseWriter, request *http.Request) {

	param, err := handleParams(request)
	if err != nil {
		log.Print("form转换json异常：", err)
	}
	switch request.URL.Path {
	case "/user/list":
		writer.Write(user.FindUserList(param))
	case "/user/register":
		writer.Write(user.RegisterUser(param))
	case "/user/login":
		writer.Write(user.Login(param))
	}
}
func handleParams(request *http.Request) ([]byte, error) {
	var param []byte
	var err error
	_ = request.ParseForm()
	if request.Method == "POST" {
		param, err = json.Marshal(request.PostForm)
	}
	if request.Method == "GET" {
		param, err = json.Marshal(request.Form)
	}
	return param, err
}
