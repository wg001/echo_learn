package dto

type Register struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}
