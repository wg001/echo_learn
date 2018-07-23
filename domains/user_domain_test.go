package domains_test

import (
	"testing"

	"echo_learn/domains"
	"fmt"
	"echo_learn/utils"
	"github.com/labstack/gommon/log"
	"echo_learn/dto"
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

func TestSaveUser(t *testing.T) {
	userRegister:=dto.Register{}
	userRegister.Phone="1582553872"
	res:=domains.SaveUser(userRegister)
	fmt.Println(res)
}

func TestGetUserByCondition(t *testing.T) {
	conditionUserOrm :=dto.UserORM{Phone:"13739289281"}
	res,err:=domains.GetUserByCondition(conditionUserOrm)
	if err!=nil{
		panic(err)
	}
	fmt.Println(res)
}
