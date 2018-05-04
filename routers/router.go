package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"echo_learn/controller"
)

func GetRouters(echo *echo.Echo)  {
	echo.POST("/wg001/:p",wg001)
	echo.GET("/wg002",controller.Wg002)
}

func wg001(c echo.Context) error {
	log.Info("jskajfdjasldfjsa")
	c.Logger().Infof("get token is:xx")
	header := c.Request().Header
	token := header.Get("token")
	c.Logger().Debug("get token is:", token)
	return nil
}
