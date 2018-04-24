package main

import (
	"testing"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

func Test_Wg(t *testing.T) {
	req := req.New()
	response,err:=req.Get("http://wgtest.wg113.com/")
	if err==nil{
		logrus.Fatal("error",response)
	}
}