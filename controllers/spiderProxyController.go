package controllers

import (
	"github.com/astaxie/beego"
	"github.com/PuerkitoBio/goquery"
	"spiter/models"
	"time"
	"strconv"
	"fmt"
	"net/http"
	"github.com/astaxie/beego/httplib"
	"net/url"
)
//http://www.xicidaili.com
const (
	ANONYMOUS = "http://www.xicidaili.com/nn/"			//匿名
	ORDINARY = "http://www.xicidaili.com/nt/"			//普通
	HTTPSPROXY = "http://www.xicidaili.com/wn/"			//Https代理
	HTTPPROXY = "http://www.xicidaili.com/wt/"			//Http代理

)

type ProxyController struct {
	beego.Controller
}
//代理匿名代理IP爬取
func (c *ProxyController) GetAysProxyIp() {
	for i:=1;i<=3265 ;i++  {
		url := fmt.Sprintf("%s%d", ANONYMOUS, i)
		transport := getTransport()
		client := http.Client{Transport:transport}
		req,_ := http.NewRequest("GET",url,nil)
		req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36 Maxthon/5.2.3.1000")
		resp,_ := client.Do(req)
		go ProxyIp(resp)
	}
}

//代普通代理IP爬取
func (c *ProxyController) GetOryProxyIp() {
	for i:=1;i<=682 ;i++  {
		url := fmt.Sprintf("%s%d", ORDINARY, i)
		get := httplib.Get(url)
		get.Header("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36 Maxthon/5.2.3.1000")
		transport := getTransport()
		get.SetTransport(transport)
		response, _ := get.Response()
		go ProxyIp(response)
	}
}

//代理HTTPS代理IP爬取
func (c *ProxyController) GetHttpsProxyIp() {
	for i:=1;i<=1339 ;i++  {
		url := fmt.Sprintf("%s%d", HTTPSPROXY, i)
		transport := getTransport()
		client := http.Client{Transport: transport}
		req,_ := http.NewRequest("GET",url,nil)
		req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36 Maxthon/5.2.3.1000")
		resp,_ := client.Do(req)
		go ProxyIp(resp)
	}
}

//代理HTTP代理IP爬取
func (c *ProxyController) GetHttpProxyIp() {
	for i:=1;i<=1835 ;i++  {
		requsetUrl := fmt.Sprintf("%s%d",HTTPPROXY , i)
		transport := getTransport()
		client := http.Client{Transport:transport}
		req,_ := http.NewRequest("GET",requsetUrl,nil)
		req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36 Maxthon/5.2.3.1000")
		resp,_ := client.Do(req)
		go ProxyIp(resp)
	}
}

//通过代理IP爬取
func getTransport()( http.RoundTripper){
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://118.190.95.43:9001")
	}
	transport := &http.Transport{Proxy: proxy}

	return transport
}

// http://www.xicidaili.com/wn/代理爬取
func ProxyIp(response *http.Response){
	doc,_ := goquery.NewDocumentFromReader(response.Body)
	doc.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
		proxy := models.TbSpiderProxyIp{}
		selection.Children().Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 1:
				proxy.Ip = selection.Text()
			case 2:
				port, _ :=strconv.ParseInt(selection.Text(), 10, 64)
				proxy.Port = port
			case 3:
				proxy.Address = selection.Text()
			case 4:
				proxy.Anonymous = selection.Text()
			case 5:
				ty := selection.Text()
				if ty == "HTTPS"{
					proxy.Https = 1
				}else {
					proxy.Https = 0
				}
			case 9:
				proxy.Check_date = selection.Text()
			default:
			}
		})
		proxy.Status = 1
		proxy.Createdate = time.Now().Format("2006-01-02 15:04:05")

		models.InsertTbSpiderProxy(&proxy)

	})
}

//检测代理IP是否可用

func (c *ProxyController) CheckProxyIp(){
	var boo bool
	proxys, _ := models.SelectTbSpiderProxys()
	for _,p := range proxys {
		id :=p.Id
		ip := p.Ip
		port := p.Port
		url := fmt.Sprintf("%s%d", ip, port)
		get := httplib.Get(url)
		string, _ := get.String()
		//判断是否有效
		 boo = string == ip
		if boo {
			//修改当前状态
			p.Anonymous = "高匿"
			p.Https = 1
			models.UpdataTbSpiderProxy(&p)
		}else {
			models.DeleteTbSpiderProxy(id)
		}

	}
}

