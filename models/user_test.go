package models

import (
	"testing"
)

var info = &UserInfo{
	Name:"root",
	Password:"123",
}

func TestList(t *testing.T){
	t.Run("Add", testInsert)
	t.Run("Get", testReadInfo)
}

func testInsert(t *testing.T) {
	// t.SkipNow 会跳过该用例
	t.SkipNow()
	if err:=Insert(info);err!=nil{
		t.Errorf("insertfailed:%v\n",err)
	}
}

func testReadInfo(t *testing.T){
	if err:=Read(info);err!=nil{
		t.Errorf("read failed:%v\n",err)
	}
}

func TestRead(t *testing.T){
	if err:=Read(info);err!=nil{
		t.Errorf("read failed:%v\n",err)
	}
}