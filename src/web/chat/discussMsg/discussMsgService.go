package discussMsg

import (
	"encoding/json"
	"log"
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
	e := exception.Error{}
	err := json.Unmarshal(param, &form)
	if err != nil {
		log.Println(err.Error())
	}
	//util.HandleParamsToStruct(param, &form)
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
	return getMsgDetail(discussMsg), e
}
func getMsgDetail(msg DiscussMsg) DiscussMsg {
	man := discussMan.GetDiscussManById(msg.ManId)
	u := user.GetUserById(man.UserId)
	msg.UserId = u.Id
	msg.Nickname = u.Nickname
	return msg
}
