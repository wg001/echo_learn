package utils

import (
	"testing"
	"fmt"
	"os"
)

func TestPhoneCheck(t *testing.T) {
	phone := "1582553844"
	checkRes, err :=PhoneCheck(phone)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if	checkRes{
		fmt.Println("ok")
	}else{
		fmt.Println("no")
	}
}
