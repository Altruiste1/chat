package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TextModel struct{

}

// table:user_info
type UserInfo struct{
	Name string `orm:"pk;column(name)"`
	Password string `orm:column(password)`
}

func init(){
	driverName,dataSource:=fetch_config()
	//orm.RegisterDataBase("sys", driverName, "root:123456@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	orm.RegisterDataBase("default", driverName, dataSource)
	orm.RegisterModel(new(UserInfo))
//	orm.RunSyncdb("default",true,true)
}

func fetch_config()(string,string){
	driverName:=beego.AppConfig.String("driverName")
	user:=beego.AppConfig.String("user")
	password:=beego.AppConfig.String("password")
	host:=beego.AppConfig.String("hostName")
	port:=beego.AppConfig.String("port")
	database:=beego.AppConfig.String("database")
	dataSource:=user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8"
	return driverName,dataSource
}

func Insert(info *UserInfo) error {
	o := orm.NewOrm()
	_, err := o.Insert(info)
	if err != nil {
		return  err
	}

	return nil
}

func Read(u *UserInfo)error{

	o:=orm.NewOrm()
	if err := o.Read(u);err!=nil{
		return err
	}

	return nil
}

func IsUserExist(info *UserInfo)bool{
	o := orm.NewOrm()
	exist := o.QueryTable("user_info").Filter("name", info.Name).Filter("password",info.Password).Exist()
	return exist
}


