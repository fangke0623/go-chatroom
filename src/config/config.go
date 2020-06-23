package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"wechat/src/common/fileHandle"
	"wechat/src/common/util"
)

type Conf struct {
	Db    DbConf              `json:"db"`
	Redis RedisConf           `json:"redis"`
	Log   LogConf             `json:"log"`
	File  fileHandle.FileConf `json:"file"`
}

func (c Conf) GetConf() Conf {
	jsonFile, err := ioutil.ReadFile("/usr/src/go/config/conf.json")

	if util.IsNil(jsonFile) {
		jsonFile, _ = ioutil.ReadFile("conf.json")
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(jsonFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
