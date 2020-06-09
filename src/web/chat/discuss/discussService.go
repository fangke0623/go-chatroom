package discuss

import (
	"strconv"
	"time"
	"wechat/src/common/enum"
	"wechat/src/common/exception"
	"wechat/src/common/util"
	"wechat/src/web/chat/discussMan"
)

func AddDiscuss(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	form := Form{}
	result := ""

	util.HandleParamsToStruct(param, &form)
	discuss := formToStruct(form)
	discuss.Status = enum.Normal
	discuss.VisibleType = form.VisibleType
	discuss.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	discuss.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	discussId := SaveDiscuss(discuss)
	discuss.DiscussId = discussId
	saveDiscussManForCreateMan(discuss)
	e.ErrorMsg = "创建成功"

	return []byte(result), e
}
func saveDiscussManForCreateMan(discuss Discuss) {
	man := discussMan.DiscussMan{}
	man.UserId = discuss.UserId
	man.DiscussId = discuss.DiscussId
	man.ManType = "1"
	man.Status = enum.Normal
	man.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	man.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	man.Remind = "1"
	discussMan.SaveDiscussMan(man)
}
func formToStruct(form Form) Discuss {
	discuss := Discuss{}
	discuss.DiscussId, _ = strconv.ParseInt(form.DiscussId, 10, 64)
	discuss.DiscussTitle = form.DiscussTitle
	discuss.UserId = form.UserId
	discuss.VisibleType = form.VisibleType
	return discuss
}
func EditDiscuss(param []byte) (interface{}, exception.Error) {
	e := exception.Error{}
	form := Form{}

	util.HandleParamsToStruct(param, &form)
	discuss := formToStruct(form)
	dbDiscuss := GetDiscussById(discuss.DiscussId)
	if dbDiscuss.DiscussTitle == "" {
		e = exception.DiscussNotExist
		return "", e
	}
	discuss.DiscussId = dbDiscuss.DiscussId
	discuss.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateDiscuss(discuss)
	e.ErrorMsg = "修改成功"

	return "", e
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
func GetDiscussDetail(param []byte) (interface{}, exception.Error) {

	form := Form{}
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	discussId, _ := strconv.ParseInt(form.DiscussId, 10, 64)
	discuss := GetDiscussById(discussId)

	return discuss, e
}
func FindDiscussList(param []byte) (interface{}, exception.Error) {

	form := Form{}
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussList(form)

	return list, e
}
