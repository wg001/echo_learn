package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

type config struct{}

func init() {
	configpath, confErr := ioutil.ReadFile("resource/config.yaml")
	if confErr != nil {
		fmt.Printf("err info is %v", confErr)
	}
	conf := make(map[string]string)

	yaml.Unmarshal(configpath, &conf)
	print(conf)
	//baseLogPath := path.Join(GlobalConfig.LogConf.Logdir,
	//	GlobalConfig.LogConf.Filename)
	//writer, err := rotatelogs.New(
	//	baseLogPath+".%Y%m%d%H%M",
	//	rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
	//	rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
	//	rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	//)
	//if err != nil {
	//	log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	//}
	//
	////log.SetFormatter(&log.TextFormatter{})
	//switch level := GlobalConfig.LogConf.LogLevel; level {
	///*
	//如果日志级别不是debug就不要打印日志到控制台了
	// */
	//case "debug":
	//	log.SetLevel(log.DebugLevel)
	//	log.SetOutput(os.Stderr)
	//case "info":
	//	setNull()
	//	log.SetLevel(log.InfoLevel)
	//case "warn":
	//	setNull()
	//	log.SetLevel(log.WarnLevel)
	//case "error":
	//	setNull()
	//	log.SetLevel(log.ErrorLevel)
	//default:
	//	setNull()
	//	log.SetLevel(log.InfoLevel)
	//}
	//
	//lfHook := lfshook.NewHook(lfshook.WriterMap{
	//	log.DebugLevel: writer, // 为不同级别设置不同的输出目的
	//	log.InfoLevel:  writer,
	//	log.WarnLevel:  writer,
	//	log.ErrorLevel: writer,
	//	log.FatalLevel: writer,
	//	log.PanicLevel: writer,
	//})
	//log.AddHook(lfHook)
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
