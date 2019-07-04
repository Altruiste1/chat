package controllers

import (
	"fmt"
	"net/http"
	
	"github.com/Altruiste1/chat/pkg/impl"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

const (
	MaxOnline=10
)
var (
	OnlineCount=0
	Client =make([]oneClient,0,MaxOnline)
	connMap = make(map[*impl.Connection]string)
)

type oneClient struct{
	conn *impl.Connection
	name string
	count int
}

var(
	upgrader = websocket.Upgrader{
		CheckOrigin:func(r *http.Request)bool{
			return true
		},
	}
)

type WsController struct{
	beego.Controller
	wsClient oneClient
}

func (c *WsController)Ws(){
	c.wsHandler(c.Ctx.ResponseWriter,c.Ctx.Request)
}

func (c *WsController)Prepare(){
	//c.wsClient.name=c.Input().Get("name")
	c.wsClient.name=c.GetString("name")
	if c.GetSession("token") != c.GetString("token"){
		fmt.Println("token is：",c.GetString("token"))
		fmt.Println("token is：",c.GetSession("token"))
		fmt.Println("token is invalid")
		return
	}

	fmt.Println("token is：",c.GetSession("token"))

}

func(c *WsController) wsHandler(w http.ResponseWriter, r *http.Request){

	var(
		wsConn *websocket.Conn
		conn *impl.Connection
		err error
		data []byte
	)

	if wsConn,err = upgrader.Upgrade(w,r,nil);err!=nil{
	    http.Error(w,fmt.Sprintf("%v",err),http.StatusInternalServerError)
		return
	}

	if conn,err = impl.InitConnection(wsConn);err!=nil{
		goto ERR
	}

	// 当用户成功连接时，为其session（UserSession）加上conn
	//if err = c.ChangeUserSession(c.wsClient.name,conn);err!=nil{
	//	goto ERR
	//}

	if err = c.editClient(conn);err!=nil{
		goto ERR
	}

	for{
		if data,err = readMsg(conn);err!=nil{
			goto ERR
		}

		if err = sendMsg(data,conn); err != nil{
			goto ERR
		}
	}

ERR:
	if err!=nil{
		conn.Close()
		http.Error(w,fmt.Sprintf("%v",err),http.StatusInternalServerError)
	}

	conn.Close()
}

func (c *WsController)editClient(conn *impl.Connection)error{
	if  conn==nil{
		return fmt.Errorf("invalid conn")
	}

	Client=append(Client,oneClient{conn,c.wsClient.name,OnlineCount})
	connMap[conn]=c.wsClient.name
	OnlineCount++
	return nil
}

func readMsg(conn *impl.Connection)([]byte,error){
	var(
		data []byte
		err error
	)

	if data,err = conn.ReadMessage();err!=nil{
		return nil,err
	}

	if _,ok:=connMap[conn];!ok{
		return nil,fmt.Errorf("查询信息发送方失败")
	}

	data=[]byte(connMap[conn]+":"+string(data))
	fmt.Printf("get:%s",data)
	return data ,nil
	
}

func sendMsg(data []byte,conn *impl.Connection)error{
	var(
		err  error
	)

	for _,client:=range Client{
		// 将收到的消息发给客户端
		if err = client.conn.WriteMessage(data);err!=nil{
			return err
		}
		
	}

	return nil
}

func (c *WsController)ChangeUserSession(name string,conn *impl.Connection)error{
	v:=c.GetSession(name)
	userSession :=v.(UserSession)
	if v==nil{
		return fmt.Errorf(fmt.Sprintf("session %s not exist",name))
	}else if userSession.conn!=nil{
		userSession.conn.Close()
	}

	userSession.conn = conn
	return nil
}