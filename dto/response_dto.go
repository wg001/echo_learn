package dto

const (
	RIGHT_CODE = 10000
	ERROR_CODE = 10001
)

type RightResponse struct {
	Status  int16       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
