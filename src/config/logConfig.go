package config

import (
	"log"
	"os"
)

type LogConf struct {
	FileName string `json:"fileName"`
	IsOpen   bool   `json:"isOpen"`
}

func LogInit(conf LogConf) {

	if conf.IsOpen {
		// 获取日志文件句柄
		// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
		logFile, err := os.OpenFile(conf.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err.Error())
		}
		// 设置存储位置
		log.SetOutput(logFile)
	}
}
