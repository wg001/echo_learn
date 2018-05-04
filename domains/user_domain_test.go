package domains_test

import (
	"testing"

	"echo_learn/domains"
	"fmt"
	"echo_learn/utils"
	"github.com/labstack/gommon/log"
)

const TIME_COMPARE = "2018-05-01 12:10:02"

func TestGetUserInfo(t *testing.T) {
	user,userError:=domains.GetUserInfo(1)
	if userError!=nil{
		panic(userError.Error())
	}
	if user.UpdateTime.After(user.CreateTime){
		fmt.Println("ok")
	}
	compareTime:=utils.FormatStrTimeToLocal(TIME_COMPARE)
	log.Infof("get comparetime is %v\n",compareTime)
	if user.UpdateTime.After(compareTime){
		fmt.Println("ok1111")
	}
	if user.CreateTime.After(compareTime){
		fmt.Println("ok2222")
	}
	//fmt.Println(utils.FormatTime(user.UpdateTime))
}
