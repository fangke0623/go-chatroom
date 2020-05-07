package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wechat/src/web/user/dao"
	"wechat/src/web/user/model"
)

func SelectUserList(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var param = &request.PostForm

	form := model.UserForm{}
	form.Id = param.Get("id")
	form.Name = param.Get("name")
	log.Println(form)
	list := dao.SelectUserList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	writer.Write(jsons)
}
