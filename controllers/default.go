package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName="login.html"
}

func (c *MainController)Put(){
	//c.token = c.formToken()
	//c.SetSession("token",c.token)
	//c.Data["name"] = c.request.UserName
	//c.Data["token"] = c.token
	//c.Redirect("/multichat",302)
	////fmt.Println(c.token)
	////c.TplName = "web.html"
}

func (c * MainController)Prepare(){

}