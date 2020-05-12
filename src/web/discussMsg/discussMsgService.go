package discussMsg

import (
	"encoding/json"
	"fmt"
	"time"
	"wechat/src/common/util"
)

func FindDiscussMsgList(param []byte) []byte {
	form := Form{}

	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussMsgList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons

}
func AddDiscussMsg(param []byte) []byte {
	discussMsg := DiscussMsg{}
	result := ""

	util.HandleParamsToStruct(param, &discussMsg)

	discussMsg.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	discussMsg.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscussMsg(discussMsg)
	result = "注册成功"
	return []byte(result)
}
