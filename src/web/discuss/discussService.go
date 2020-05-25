package discuss

import (
	"time"
	"wechat/src/common/enum"
	"wechat/src/common/exception"
	"wechat/src/common/util"
)

func AddDiscuss(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	discuss := Discuss{}
	result := ""

	util.HandleParamsToStruct(param, &discuss)

	discuss.Status = enum.Normal
	discuss.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	discuss.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscuss(discuss)
	e.ErrorMsg = "创建成功"

	return []byte(result), e
}
func EditDiscuss(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	discuss := Discuss{}
	result := ""

	util.HandleParamsToStruct(param, &discuss)
	dbDiscuss := GetDiscussById(discuss.DiscussId)
	if dbDiscuss.DiscussTitle == "" {
		e = exception.DiscussNotExist
	}
	discuss.DiscussId = dbDiscuss.DiscussId
	discuss.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateDiscuss(discuss)
	e.ErrorMsg = "修改成功"

	return []byte(result), e
}
func DeleteDiscuss(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	discuss := Discuss{}
	result := ""

	util.HandleParamsToStruct(param, &discuss)
	dbDiscuss := GetDiscussById(discuss.DiscussId)
	if dbDiscuss.DiscussTitle == "" {
		e = exception.DiscussNotExist
	}
	discuss.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	discuss.Status = enum.Delete
	UpdateDiscuss(discuss)
	e.ErrorMsg = "删除成功"

	return []byte(result), e
}
func FindDiscussList(param []byte) (interface{}, exception.Error) {

	form := Form{}
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussList(form)

	return list, e
}
