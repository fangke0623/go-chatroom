package config

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

var RedisConn redis.Conn

func RedisInit() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("redis connect error", err)
	} else {
		log.Println("redis connect success")
	}
	RedisConn = c
}

func DoSet(key string, value interface{}) {

	reply, err := RedisConn.Do("SET", key, value)
	if err != nil {
		log.Println("redis set key value error", err)
	}
	if reply != nil {
		log.Println("set key reply", reply)
	}
}
func DoGet(key string) []byte {

	reply, err := redis.Bytes(RedisConn.Do("GET", key))
	if err != nil {
		log.Println("redis get key value error", err)
	}
	return reply
}
func DoDel(key string) interface{} {

	reply, err := redis.Bool(RedisConn.Do("Del", key))
	if err != nil {
		log.Println("redis del key value error", err)
	}
	return reply
}
func DoExpire(key string, time int64) {

	reply, err := RedisConn.Do("EXPIRE", key, time)
	if err != nil {
		log.Println("do expire error", err)
	}
	if reply != nil {
		log.Println("do expire reply", reply)
	}
}

func RedisClose() {

	RedisConn.Close()
}
