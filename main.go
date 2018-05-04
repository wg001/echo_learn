package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"echo_learn/dto"
	"echo_learn/routers"
	"github.com/labstack/gommon/log"
	utils "echo_learn/utils"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routers.GetRouters(e)
	// Routes
	e.GET("/", hello)
	e.POST("/wg", wg)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	configpath, confErr := ioutil.ReadFile("config/log_conf.yaml")
	if confErr != nil || len(configpath) == 0 {
		logrus.Panicf("err info is %v", confErr)
	}
	conf := dto.GlobalConfig{}
	yamlError := yaml.Unmarshal(configpath, &conf)
	if yamlError != nil {
		log.Debugf("config err")
		panic(yamlError.Error())
	}
	utils.ConfigLocalFilesystemLogger(conf.Log.LogPath,conf.Log.LogName,600,600)
	dto.SetGlobalConf(&conf)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!Wanggang")
}

func wg(c echo.Context) error {
	fmt.Println("hello")
	header := c.Request().Header
	token := header.Get("token")
	logrus.Debug("get token is:", token)
	return nil
}