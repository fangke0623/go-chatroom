package discussMsg

import (
	"time"
	"wechat/src/common/enum"
	"wechat/src/common/exception"
	"wechat/src/common/util"
	"wechat/src/web/chat/discussMan"
	"wechat/src/web/chat/user"
)

func FindDiscussMsgList(param []byte) (interface{}, exception.Error) {
	form := Form{}
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussMsgList(form)
	for key, msg := range list {
		man := discussMan.GetDiscussManById(msg.ManId)
		u := user.GetUserById(man.UserId)
		list[key].UserId = u.Id
		list[key].Nickname = u.Nickname
	}

	return list, e

}
func AddDiscussMsg(param []byte) (interface{}, exception.Error) {
	form := Form{}
	result := ""
	e := exception.Error{}
	util.HandleParamsToStruct(param, &form)
	discussManForm := discussMan.Form{}
	discussManForm.UserId = form.UserId
	discussManForm.DiscussId = form.DiscussId
	man := discussMan.FindDiscussMan(discussManForm)
	discussMsg := DiscussMsg{}
	discussMsg.DiscussId = man.DiscussId
	discussMsg.ManId = man.ManId
	discussMsg.MsgContent = form.MsgContent
	discussMsg.Status = enum.Normal
	discussMsg.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	discussMsg.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscussMsg(discussMsg)
	e.ErrorMsg = "发送成功"
	return []byte(result), e
}
