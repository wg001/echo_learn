package utils

import (
	"echo_learn/dto"
	//"go/src/net/http"
	"github.com/labstack/echo"
)

func GenResponse(c echo.Context, status int16, message string, data interface{}) error {
	if data == nil {
		data = struct{}{}
	}
	return c.JSON(dto.HTTP_OK, dto.Response{status, message, data})
}
