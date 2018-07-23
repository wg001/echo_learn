package utils

import "regexp"

//手机号检测
func PhoneCheck(phone string) (bool,error) {
	return regexp.MatchString("^1\\d{10}",phone)
}
