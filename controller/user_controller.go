package controller

import (
	"github.com/labstack/echo"
	"echo_learn/domains"
	"echo_learn/dto"
)

func Wg002(c echo.Context) error {
	user:=domains.GetUserInfo(1,dto.GetDbStruct().GetDBObj())
	return c.JSON(200,dto.RightResponse{dto.RIGHT_CODE,"success",user})
}
