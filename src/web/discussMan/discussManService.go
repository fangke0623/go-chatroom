package discussMan

import (
	"time"
	"wechat/src/common/exception"
	"wechat/src/common/util"
)

func FindDiscussManList(param []byte) (interface{}, exception.Error) {
	form := Form{}
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussManList(form)

	return list, e

}

func AddDiscussMan(param []byte) (interface{}, exception.Error) {
	discussMan := DiscussMan{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &discussMan)

	discussMan.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	discussMan.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscussMan(discussMan)
	result = "注册成功"

	return []byte(result), e
}
func EditDiscussMan(param []byte) (interface{}, exception.Error) {
	discussMan := DiscussMan{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &discussMan)

	discussMan.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
	UpdateDiscussMan(discussMan)
	result = "修改成功"

	return []byte(result), e
}
