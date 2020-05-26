package discussMsg

import (
	"time"
	"wechat/src/common/enum"
	"wechat/src/common/exception"
	"wechat/src/common/util"
)

func FindDiscussMsgList(param []byte) (interface{}, exception.Error) {
	form := Form{}
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussMsgList(form)

	return list, e

}
func AddDiscussMsg(param []byte) (interface{}, exception.Error) {
	discussMsg := DiscussMsg{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &discussMsg)
	discussMsg.Status = enum.Normal
	discussMsg.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	discussMsg.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscussMsg(discussMsg)
	e.ErrorMsg = "发送成功"
	return []byte(result), e
}
