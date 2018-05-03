package main_test

import (
	"testing"
	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"fmt"
	"echo_learn/utils"
	"net/http"
	"os"
	"echo_learn/dto"
	"gopkg.in/yaml.v2"
	"github.com/labstack/gommon/log"
	"io/ioutil"
)

type Data struct {
	Serires    []Serires      `json:"series"`
	Title      string         `json:"title"`
	Categories []string       `json:"categories"`
	Y          map[string]int `json:"y"`
}

type Serires struct {
	LabelName  string `json:"label_name"`
	ValueLabel []int  `json:"value_label"`
	LabelData  string `json:"label_data"`
}

type ResponseObj struct {
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
	Status int    `json:"status"`
}

type ErrorResObj struct {
	Msg  string        `json:"msg"`
	Code int           `json:"code"`
	Data []interface{} `json:"data"`
}

type Score struct {
	Id           uint
	UserId       uint
	Name         string
	TypePosition string
	CourseId     uint
	Score        uint
	StudentId    uint
}

func init() {
	configpath, confErr := ioutil.ReadFile("config/log_conf.yaml")
	if confErr != nil || len(configpath) == 0 {
		log.Panicf("err info is %v", confErr)
	}
	conf := dto.GlobalConfig{}
	yamlError := yaml.Unmarshal(configpath, &conf)
	if yamlError != nil {
		log.Debugf("config err")
		panic(yamlError.Error())
	}
	dto.SetGlobalConf(&conf)
	db := dto.GetDbStruct().GetDBObj()
	if db == nil{
		fmt.Println("xxooxl")
	}
}

func Test_Wg(t *testing.T) {

	db := dto.GetDbStruct().GetDBObj()
	if db == nil{
		fmt.Println("xxooxl")
	}
	var result []Score
	db.Raw("select * from score").Scan(&result)
	//if err1!=nil{
	//	panic(err1.Error())
	//}
	for k,v:=range result{
		fmt.Printf("key:%v, value:%v\n",k,v)
	}
	//da.Scan(&result)
	os.Exit(0)
	req := req.New()
	response, err := req.Get("http://www.test.com/go.php")
	if response.Response().StatusCode != http.StatusOK {
		rqer, _ := response.ToString()
		logrus.Fatal("something error--", response.Response().StatusCode, rqer)
	}
	var resObj ResponseObj
	if err != nil {
		logrus.Fatal("error", response.Dump())
	}
	str, _ := response.ToString()
	resBytes, errByte := response.ToBytes()
	if errByte != nil {
		logrus.Fatal("something error")
	}
	if checkKey, checkres := utils.CheckParamsExist(resBytes, "code", "msg"); !checkres {
		logrus.Fatalf("key %s not exist", checkKey)
	}
	fmt.Println(str)
	response.ToJSON(&resObj)
	if resObj.Status == 0 {
		fmt.Println("**********")
	}
	fmt.Println("----", resObj.Msg)
}
