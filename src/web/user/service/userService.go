package service

import (
	"encoding/json"
	"fmt"
	"time"
	"wechat/src/common/util"
	"wechat/src/web/user/dao"
	"wechat/src/web/user/model"
)

func SelectUserList(param []byte) []byte {

	form := model.UserForm{}
	util.HandleParamsToStruct(param, &form)

	list := dao.SelectUserList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons
}
func RegisterUser(param []byte) bool {

	user := model.User{}

	util.HandleParamsToStruct(param, &user)
	user.Id = util.GenerateUUID()
	user.CreateTime = time.Now().String()
	dao.SaveUser(user)

	return true

}
