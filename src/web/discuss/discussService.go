package discuss

import (
	"encoding/json"
	"fmt"
	"time"
	"wechat/src/common/util"
)

func AddDiscuss(param []byte) []byte {

	discuss := Discuss{}
	result := ""

	util.HandleParamsToStruct(param, &discuss)

	discuss.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	discuss.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscuss(discuss)
	result = "注册成功"

	return []byte(result)
}
func UpdateDiscuss(param []byte) []byte {

	discuss := Discuss{}
	result := ""

	util.HandleParamsToStruct(param, &discuss)

	discuss.CreateDate = time.Now().String()
	SaveDiscuss(discuss)
	result = "修改成功"

	return []byte(result)
}
func DeleteDiscuss(param []byte) []byte {

	discuss := Discuss{}
	result := ""

	util.HandleParamsToStruct(param, &discuss)

	discuss.CreateDate = time.Now().String()
	SaveDiscuss(discuss)
	result = "删除成功"

	return []byte(result)
}
func FindDiscussList(param []byte) []byte {

	form := Form{}

	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons
}
