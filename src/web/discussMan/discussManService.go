package discussMan

import (
	"time"
	"wechat/src/common/enum"
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
	//e := exception.Error{}
	util.HandleParamsToStruct(param, &discussMan)
	discussMan.Status = enum.NormalMan
	discussMan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	discussMan.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscussMan(discussMan)

	return []byte(result), exception.Success
}
func EditDiscussMan(param []byte) (interface{}, exception.Error) {
	discussMan := DiscussMan{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &discussMan)
	dbDiscussMan := GetDiscussManById(discussMan.ManId)
	if dbDiscussMan.ManId == 0 {
		e = exception.DiscussManNotExist
		return nil, e
	}
	discussMan.ManId = dbDiscussMan.ManId
	discussMan.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateDiscussMan(discussMan)
	return []byte(result), exception.Success
}
func DeleteDiscussMan(param []byte) (interface{}, exception.Error) {
	discussMan := DiscussMan{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &discussMan)
	dbDiscussMan := GetDiscussManById(discussMan.ManId)
	if dbDiscussMan.ManId == 0 {
		e = exception.DiscussManNotExist
		return nil, e
	}
	discussMan.ManId = dbDiscussMan.ManId
	discussMan.Status = enum.DeleteMan
	discussMan.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	UpdateDiscussMan(discussMan)

	return []byte(result), exception.Success
}
