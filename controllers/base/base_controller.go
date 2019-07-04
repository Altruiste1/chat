package base

import (
	"github.com/astaxie/beego/session"
)

var GlobalSessions *session.Manager

//NewManager 函数的参数的函数如下所示:引擎名字，可以是 memory、file、mysql 或 redis。
//managerconfig: 一个 JSON 字符串,传入 Manager 的配置信息
//cookieName: 客户端存储 cookie 的名字。
//enableSetCookie,omitempty: 是否开启 SetCookie,omitempty 这个设置
//gclifetime: 触发 GC 的时间。
//maxLifetime: 服务器端存储的数据的过期时间
//secure: 是否开启 HTTPS，在 cookie 设置的时候有 cookie.Secure 设置。
//sessionIDHashFunc: sessionID 生产的函数，默认是 sha1 算法。
//sessionIDHashKey: hash 算法中的 key。
//cookieLifeTime: 客户端存储的 cookie 的时间，默认值是 0，即浏览器生命周期。
//providerConfig: 配置信息，根据不同的引擎设置不同的配置信息，
func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}