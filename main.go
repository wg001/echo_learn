package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"echo_learn/dto"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/wg", wg)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
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
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!Wanggang")
}

func wg(c echo.Context) error {
	log.Info("jskajfdjasldfjsa")
	c.Logger().Infof("get token is:xx")
	fmt.Println("hello")
	header := c.Request().Header
	token := header.Get("token")
	c.Logger().Debug("get token is:", token)
	return nil
}
