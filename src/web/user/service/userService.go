package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wechat/src/web/user/dao"
	"wechat/src/web/user/model"
)

func SelectUserList(param []byte) []byte {

	form := model.UserForm{}
	_ = json.Unmarshal(param, &form)
	log.Println(form)
	list := dao.SelectUserList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons
}
func RegisterUser(request *http.Request) {

}
