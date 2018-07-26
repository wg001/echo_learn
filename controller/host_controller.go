package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func HostPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index",nil)
}
