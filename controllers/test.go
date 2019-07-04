package controllers

import (
	"github.com/Altruiste1/chat/pkg/impl"
	"time"
)
type TestController struct {
	MainController
}

type UserSession struct{
	username string
	conn *impl.Connection
}

func (this *TestController)Get(){
	this.SetSession("hl",UserSession{
		"hl",
		nil,
	})

	
}

func formatTimeStamp(in interface{}, layout string) (out string) {
	timeStr := in.(int)
	month, err := time.Parse(layout, string(timeStr))
	if err != nil {
		return time.Now().Format(layout)
	}
	return month.Format(layout)
}