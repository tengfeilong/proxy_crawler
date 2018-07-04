package controllers

import (
	"fmt"
	"net/http"
	"github.com/astaxie/beego/httplib"
	"time"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"spiter/models"
	"github.com/astaxie/beego"
	"log"
)

const (
	KUAIANONYMOUS = "https://www.kuaidaili.com/free/inha/"			//匿名
	KUAIORDINARY = "https://www.kuaidaili.com/free/intr/"			//普通
)

type KuaiProxyCortroller struct {
	beego.Controller
}

//代理匿名代理IP爬取
func (c *KuaiProxyCortroller) GetKuaiAysProxyIp() {
	for i:=1;i<=2380 ;i++  {
		url := fmt.Sprintf("%s%d", KUAIANONYMOUS, i)
		//transport := getTransport()
		client := http.Client{}
		req,_ := http.NewRequest("GET",url,nil)
		req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36 Maxthon/5.2.3.1000")
		resp,_ := client.Do(req)
		go KuaiProxyIp(resp)
	}
}

//代普通代理IP爬取
func (c *KuaiProxyCortroller) GetKuaiOryProxyIp() {
	for i:=1302;i<=2380 ;i++  {
		url := fmt.Sprintf("%s%d", KUAIORDINARY, i)
		log.Println(url)
		get := httplib.Get(url)
		get.Header("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36 Maxthon/5.2.3.1000")
		/*transport := getTransport()
		get.SetTransport(transport)*/
		response, _ := get.Response()
		KuaiProxyIp(response)
	}
}

// http://www.xicidaili.com/wn/代理爬取
func KuaiProxyIp(response *http.Response){
	doc,_ := goquery.NewDocumentFromReader(response.Body)
	doc.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
		proxy := models.TbSpiderProxyIp{}
		selection.Children().Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				proxy.Ip = selection.Text()
			case 1:
				port, _ :=strconv.ParseInt(selection.Text(), 10, 64)
				proxy.Port = port
			case 2:
				proxy.Anonymous = selection.Text()
			case 4:
				proxy.Address = selection.Text()

			case 3:
				ty := selection.Text()
				if ty == "HTTPS"{
					proxy.Https = 1
				}else {
					proxy.Https = 0
				}
			case 6:
				proxy.Check_date = selection.Text()
			default:
			}
		})
		proxy.Status = 1
		proxy.Createdate = time.Now().Format("2006-01-02 15:04:05")

		models.InsertTbSpiderProxy(&proxy)

	})
}
