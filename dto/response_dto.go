package dto

import (
	"net/http"
)

const (
	HTTP_OK    = http.StatusOK
	RIGHT_CODE = 10000
	ERROR_CODE = 10001
)

const (
	STATUS_RIGHT               = "SUCCESS"
	STATUS_ERROR               = "ERROR"
	STATUS_REQUEST_NOT_ALLOWED = "非法请求"
)

type Response struct {
	Status  int16       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
