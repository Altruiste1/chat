package controllers

import (
	"fmt"
	"github.com/Altruiste1/chat/models"
	"net/http"
)

const(
	UserNameLenth = 2
	PasswordLenth = 3
)

type RegisterController struct{
	MainController
	request  RequestBody
	token  string
}

type RequestBody struct{
	UserName string `json:"user_name"`
	Password  string  `json:"password"`
}

func (c *RegisterController)Register(){
	info := c.fetchRegisterMsg()
	// 这里遗留一个问题，用户名被占用等错误的请求及数据库错误等统一归类到了http.StatusBadRequest
	if err := models.Insert(info);err!=nil{
		//http.Error(c.Ctx.ResponseWriter,fmt.Sprintf("%v",err),http.StatusBadRequest)
		c.Ctx.WriteString(fmt.Sprintf("%v",err))
		return
	}
	
	c.Ctx.WriteString("注册成功")
}

func (c *RegisterController)Login(){
	//if err := c.SetUserSession(); err!=nil{
	//	http.Error(c.Ctx.ResponseWriter,fmt.Sprintf("%v",err),http.StatusBadRequest)
	//	return
	//}

	info:=c.fetchRegisterMsg()
	if exist:=models.IsUserExist(info);!exist{
		http.Error(c.Ctx.ResponseWriter,fmt.Sprintf("Bad Request:User not exist",),http.StatusBadRequest)
		return
	}

	//c.SetSession("hl",UserSession{
	//	"hl",
	//	nil,
	//})

	c.Data["name"] = c.request.UserName
	c.TplName = "web.html"
}

func (c *RegisterController)SetUserSession()error{
	if v:=c.GetSession(c.request.UserName);v!=nil{
		return fmt.Errorf("错误，该用户已登陆，请确认是否为本人登陆,%s",v)
	}else{
		c.SetSession(c.request.UserName,
			UserSession{
			c.request.UserName,
			nil,
		})
	}

	return nil
}

// 获得注册的信息
func (c *RegisterController)fetchRegisterMsg()*models.UserInfo{
	info := &models.UserInfo{c.request.UserName,
		c.request.Password,
	}
	return info
}

func (c *RegisterController)Prepare(){
    if err:=c.Parse();err!=nil{
    	fmt.Println(err)
		http.Error(c.Ctx.ResponseWriter,fmt.Sprintf("%v",err),500)
        return
	}
    
    if err:=c.InvalidCheck();err!=nil{
		http.Error(c.Ctx.ResponseWriter,fmt.Sprintf("%v",err),400)
    	return
	}

}

func (c *RegisterController)Parse()error{
	//fmt.Printf("%s",c.Ctx.Input.RequestBody)
	//if err := json.Unmarshal(c.Ctx.Input.RequestBody,&c.request); err != nil {
	//	return fmt.Errorf(fmt.Sprintf("unmarshal failed",err))
	//}
	c.request.UserName = c.GetString("user_name")
	c.request.Password = c.GetString("password")
    return nil
}

// 对传入的数据进行检测
func (c *RegisterController)InvalidCheck()error{
	if len(c.request.Password)<PasswordLenth||len(c.request.UserName)<UserNameLenth{
		return fmt.Errorf(fmt.Sprintf("用户名长度or用户密码过短：用户名:%d" +
			"最短密码长度为;%d", UserNameLenth,PasswordLenth))
	}

	return nil
}

