package main

import (
	"net/http"
	_ "net/http/pprof"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"echo_learn/dto"
	"echo_learn/routers"
	"github.com/labstack/gommon/log"
	"echo_learn/utils"
	"os"
	"github.com/sevenNt/echo-pprof"
	"html/template"
	"io"
)


// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	// Echo instance
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routers.GetRouters(e)
	// Routes
	e.GET("/", hello)
	e.POST("/wg", wg)
	echopprof.Wrap(e)
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
	logrus.SetOutput(os.Stdout)
	utils.ConfigLocalFilesystemLogger(conf.Log.LogPath,conf.Log.LogName,7*24*60*60,24*60*60)
	dto.SetGlobalConf(&conf)
}

// Handler
func hello(c echo.Context) error {
	logrus.Info("get params: ",c.ParamNames())
	return c.String(http.StatusOK, "Hello, World!Wanggang")
}

func wg(c echo.Context) error {
	header := c.Request().Header
	token := header.Get("token")
	logrus.Info("get token is:", token)
	logrus.Debug("get token is :",token)
	return c.JSON(http.StatusOK,dto.Response{dto.RIGHT_CODE,"success","token is :"+token})
}