package utils

import (
	"regexp"
	"encoding/json"
)

//手机号检测
func PhoneCheck(phone string) (bool,error) {
	return regexp.MatchString("^1\\d{10}",phone)
}



func CheckParamsExistInBytes(params []byte,keys ...string)(string,bool) {
	var mapParam map[string]interface{}
	json.Unmarshal(params,&mapParam)
	for _,key:=range keys{
		if _,b :=mapParam[key];!b{
			return key,false
		}
	}
	return "",true
}
func CheckParamsExistInMap(params map[string]interface{},keys ...string)(string,bool) {
	for _,key:=range keys{
		if _,b :=params[key];!b{
			return key,false
		}
	}
	return "",true
}