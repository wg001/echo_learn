package utils

import "encoding/json"

func CheckParamsExist(params []byte,keys ...string)(string,bool) {
	var mapParam map[string]interface{}
	json.Unmarshal(params,&mapParam)
	for _,key:=range keys{
		if _,b :=mapParam[key];!b{
			return key,false
		}
	}
	return "",true
}