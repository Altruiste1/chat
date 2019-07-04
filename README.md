# chat
beego + WebSocket聊天室，数据库使用的是MySQL，使用websocket保持长连接，后期会加上token，及session
#

1.环境准备

   (1)安装beego及bee工具,这个可以网上搜下教程
      可以参考 https://beego.me/

  （2) 安装mysql，并且创建数据库，安装mysql的教程可以网上找下，笔者装的是windows版的。#
       配置时需要记住：用户名，密码，hostname，port

  （3）安装mysql驱动
      go get github.com/go-sql-driver/mysql, 或者可以下载源码放在GOPATH的src下

  （4）安装websocket包
       go get  github.com/gorilla/websocket
 
2.  获取代码:  go get github.com/Altruiste1/chat

3. 配置chat/conf/app.conf文件，这个文件下是一些配置信息，包括端口号，数据库的一些参数设定，
   包括session，xsrf等等设置开启
  
   笔者在app.conf里关于mysql的配置如下，若您有所改动，请对对应参数做修改，database为数据库名，需要您手动先创建一个数据库，名字为mytest,或者其他

   driverName= mysql
   user=root
   password=123456
   hostName=localhost
   port=3306
   database=mytest
  
 4. 进入到github.com/Altruiste1/chat目录，执行bee run,运行成功后

    打开浏览器，输入  http://127.0.0.1:8080/
 
    下面是效果图
![https://github.com/Altruiste1/image/blob/master/login.png]
![https://github.com/Altruiste1/image/blob/master/chat.png]  
