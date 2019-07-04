package routers

import (
	"github.com/Altruiste1/chat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Get")
	beego.Router("/msg", &controllers.MainController{},"get:Put")
	beego.Router("/register", &controllers.RegisterController{},"post:Register")
	beego.Router("/login", &controllers.RegisterController{},"post:Login")
    beego.Router("/ws/join",&controllers.WsController{},"*:Ws")
	beego.Router("/test",&controllers.TestController{},"*:Get")
    
}
