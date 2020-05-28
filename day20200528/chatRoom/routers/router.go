package routers

import (
	"github.com/astaxie/beego"
	"github.com/zhangatle/learngo/day20200528/chatRoom/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/chatRoom", &controllers.ServerController{})
	beego.Router("/chatRoom/WS", &controllers.ServerController{}, "get:WS")
}
