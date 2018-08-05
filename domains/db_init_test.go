package domains_test

import (
	"echo_learn/dto"
	"gopkg.in/yaml.v2"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	configpath, confErr := ioutil.ReadFile("../config/log_conf.yaml")
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
	db,err := dto.GetDbStruct()
	if err!=nil{

	}
	if db == nil{
		fmt.Println("xxooxl")
	}
	db.GetDBObj().LogMode(true)
}