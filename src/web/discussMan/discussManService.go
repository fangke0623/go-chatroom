package discussMan

import (
	"encoding/json"
	"fmt"
	"time"
	"wechat/src/common/util"
)

func FindDiscussManList(param []byte) []byte {
	form := Form{}

	util.HandleParamsToStruct(param, &form)
	list := SelectDiscussManList(form)
	jsons, err := json.Marshal(list)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsons

}

func AddDiscussMan(param []byte) []byte {
	discussMan := DiscussMan{}
	result := ""

	util.HandleParamsToStruct(param, &discussMan)

	discussMan.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	discussMan.UpdateDate = time.Now().Format("2006-01-02 15:04:05")
	SaveDiscussMan(discussMan)
	result = "注册成功"

	return []byte(result)
}
