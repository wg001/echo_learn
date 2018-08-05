package domains_test

import (
	"testing"

	"echo_learn/domains"
	"fmt"
	"echo_learn/utils"
	"github.com/labstack/gommon/log"
	"echo_learn/dto"
	"time"
	"github.com/jinzhu/gorm"
)

const TIME_COMPARE = "2018-05-01 12:10:02"

func TestGetUserInfo(t *testing.T) {
	userDomainObj := domains.NewUserDomain()
	user, userError := userDomainObj.GetUserInfo(1)
	if userError != nil {
		panic(userError.Error())
	}
	if user.UpdateTime.After(user.CreateTime) {
		fmt.Println("ok")
	}
	compareTime := utils.FormatStrTimeToLocal(TIME_COMPARE)
	log.Infof("get comparetime is %v\n", compareTime)
	if user.UpdateTime.After(compareTime) {
		fmt.Println("ok1111")
	}
	if user.CreateTime.After(compareTime) {
		fmt.Println("ok2222")
	}
	//fmt.Println(utils.FormatTime(user.UpdateTime))
}

func TestSaveUser(t *testing.T) {
	userRegister := dto.Register{}
	userRegister.Phone = "1582553872"
	userDomainObj := domains.NewUserDomain()
	res := userDomainObj.SaveUser(userRegister)
	fmt.Println(res)
}

func TestGetUserByCondition(t *testing.T) {
	condition := make(map[string]interface{})
	updatedata :=make(map[string]interface{})
	//updatedata["phone"] = "15225538470"
	updatedata["username"]= "王刚"
	updatedata["is_delete"] = 1
	updatedata["level"] = gorm.Expr("level+?",1)
	condition["phone"] = "13739289281"
	condition["id"] = 2
	userDomainObj := domains.NewUserDomain()
	for i := 0; i < 10; i++ {
		go func(index int) {
			err := userDomainObj.UpdateUser(updatedata, condition)
			if err != nil {
				panic(err)
			}
			fmt.Println("---", index)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("ok")
}
