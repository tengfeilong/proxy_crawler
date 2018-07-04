package models

import (
	"github.com/astaxie/goredis"
	"github.com/astaxie/beego"
)

var(
	redis goredis.Client
)

const (
	URL_QUEUE  =  "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

func ConnectRedis(){
	address := beego.AppConfig.String("redis_addres")
	redis.Addr = address
}

func SetString(key string,value string)(bool){
	error := redis.Set(key, []byte(value))
	if error != nil{
		return false
	}
	return true
}

func GetString(key string){
	redis.Get(key)
}

func Lpush(value string)(bool){
	error := redis.Lpush(URL_QUEUE, []byte(value))
	if error != nil{
		return false
	}
	return true
}

func Rpop()(string){
	rpop, error := redis.Rpop(URL_QUEUE)
	if error != nil{
		panic(error)
	}
	return string(rpop)
}

func GetQueueLength() int{
	length,err := redis.Llen(URL_QUEUE)
	if err != nil{
		return 0
	}

	return length
}

func AddToSet(url string){
	redis.Sadd(URL_VISIT_SET, []byte(url))
}

func IsVisit(url string) bool{
	bIsVisit, err := redis.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil{
		return false
	}

	return bIsVisit
}