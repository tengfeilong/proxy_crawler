package routers

import (
	"spiter/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/getMovieInfo", &controllers.MovieInfoController{},"*:GetMovieInfo")
	beego.Router("/getAysProxyIp", &controllers.ProxyController{},"*:GetAysProxyIp")
	beego.Router("/getOryProxyIp", &controllers.ProxyController{},"*:GetOryProxyIp")
	beego.Router("/getHttpsProxyIp", &controllers.ProxyController{},"*:GetHttpsProxyIp")
	beego.Router("/getHttpProxyIp", &controllers.ProxyController{},"*:GetHttpProxyIp")
	beego.Router("/checkProxyIp", &controllers.ProxyController{},"get:CheckProxyIp")

	beego.Router("/getKuaiAysProxyIp", &controllers.KuaiProxyCortroller{},"get:GetKuaiAysProxyIp")
	beego.Router("/getKuaiOryProxyIp", &controllers.KuaiProxyCortroller{},"get:GetKuaiOryProxyIp")
}
