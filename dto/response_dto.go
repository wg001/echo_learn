package dto

import (
	"github.com/labstack/echo"
	"net/http"
)

const (
	RIGHT_CODE = 10000
	ERROR_CODE = 10001
)

type Response struct {
	Status  int16       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GenResponse(c echo.Context,status int16,message string,data interface{}) error {
	return c.JSON(http.StatusOK,Response{status,message,data})
}
