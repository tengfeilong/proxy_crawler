package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	)

var(
	db orm.Ormer
)

type TbSpiderProxyIp struct {
	Id int64
	Ip string
	Port int64
	Address string
	Anonymous string
	Operator string
	Check_date string
	Createdate string
	Status int64
	Https int64
}


func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:Lrw360+@tcp(192.168.1.121:3306)/lrw360-map?charset=utf8", 30)
	orm.RegisterModel(new(TbSpiderProxyIp),new(MovieInfo))
	db = orm.NewOrm()
}

func InsertTbSpiderProxy(proxy *TbSpiderProxyIp)(int64,error){
	insert, error := db.InsertOrUpdate(proxy)
	return insert,error
}

func DeleteTbSpiderProxy(id int64)(int64,error)  {
	info := TbSpiderProxyIp{Id: id}
	delete, error := db.Delete(&info)
	return delete, error
}
func UpdataTbSpiderProxy(proxy *TbSpiderProxyIp)(int64,error){
	update, error := db.Update(proxy)
	return update,error
}

func SelectTbSpiderProxy(id int64)(TbSpiderProxyIp,error){
	proxyInfo := TbSpiderProxyIp{}
	error := db.Raw("select * from tb_spider_proxy_ip where id=?", id).QueryRow(&proxyInfo)
	return proxyInfo,error
}

func SelectTbSpiderProxys()([]TbSpiderProxyIp,error){
	proxyInfo := []TbSpiderProxyIp{}
	_, error:= db.Raw("select * from tb_spider_proxy_ip").QueryRows(&proxyInfo)
	return proxyInfo,error
}
