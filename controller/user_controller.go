package controller

import (
	"github.com/labstack/echo"
	"echo_learn/domains"
	"echo_learn/dto"
	"echo_learn/utils"
	"github.com/sirupsen/logrus"
)

func Wg002(c echo.Context) error {
	userObj:=domains.NewUserDomain()
	user, getUserError := userObj.GetUserInfo(1)
	if getUserError != nil {
		return utils.GenResponse(c, dto.RIGHT_CODE, getUserError.Error(), nil)
	}
	return utils.GenResponse(c, dto.RIGHT_CODE, dto.STATUS_RIGHT, user)
}

//用户注册
func Register(c echo.Context) error {
	register := dto.Register{}
	c.Bind(&register)
	if phoneCheck, err := utils.PhoneCheck(register.Phone); !phoneCheck || err != nil {
		logrus.Errorf("-----register:%v",register)
		logrus.Errorf("phone check wrong:%v%v",phoneCheck,err)
		return utils.GenResponse(c, dto.ERROR_CODE, "手机号不合法", nil)
	}
	userDomain:=domains.NewUserDomain()
	userDomain.SaveUser(register)

	return utils.GenResponse(c, dto.RIGHT_CODE, dto.STATUS_RIGHT, register)
}

func Login(c echo.Context) error{
	return nil
}

//func UpdateUser(c echo.Echo) error{
//
//}