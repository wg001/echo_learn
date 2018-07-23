package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"echo_learn/controller"
	"echo_learn/dto"
	"strings"
	"echo_learn/utils"
)

func GetRouters(echo *echo.Echo)  {
	echo.POST("/wg001/:p",wg001)
	echo.GET("/wg002",controller.Wg002)
	echo.POST("/reg",controller.Register)
}

func wg001(c echo.Context) error {
	log.Info(c.Request().Header)
	contentType:=c.Request().Header.Get("Content-Type")
	registerorm := new(dto.Register)
	c.Bind(registerorm)
	if contentType=="" || !strings.HasPrefix(contentType,"application/json"){
		return utils.GenResponse(c,dto.ERROR_CODE,dto.STATUS_REQUEST_NOT_ALLOWED, nil)
	}
	log.Info("-------phone:",registerorm.Phone)
	log.Info("content-type---:",contentType)
	log.Info("jskajfdjasldfjsa")
	c.Logger().Infof("get token is:xx")
	header := c.Request().Header
	token := header.Get("token")
	c.Logger().Debug("get token is:", token)
	return utils.GenResponse(c,dto.RIGHT_CODE,dto.STATUS_RIGHT,nil)
}
