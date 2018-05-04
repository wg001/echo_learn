package controller

import (
	"github.com/labstack/echo"
	"echo_learn/domains"
	"echo_learn/dto"
	"net/http"
)

func Wg002(c echo.Context) error {
	user,getUserError:=domains.GetUserInfo(1)
	if getUserError!=nil{
		//return dto.GenResponse(c,dto.ERROR_CODE,getUserError.Error(),nil)
		return c.JSON(http.StatusOK,dto.Response{dto.ERROR_CODE,getUserError.Error(),nil})
	}
	return c.JSON(http.StatusOK,dto.Response{dto.RIGHT_CODE,"success",user})
	//return dto.GenResponse(c,dto.ERROR_CODE,getUserError.Error(),user)
}
