package domains_test

import (
	"testing"
	//_ ""
	"echo_learn/dto"
	"echo_learn/domains"
	"fmt"
	"echo_learn/utils"
	"github.com/labstack/gommon/log"
)

const TIME_COMPARE = "2018-05-01 12:10:02"

func TestGetUserInfo(t *testing.T) {
	db := dto.GetDbStruct().GetDBObj()
	user:=domains.GetUserInfo(1, db)
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
