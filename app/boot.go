package app

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/fvbock/endless"
	"github.com/spf13/viper"
	"github.com/zqhong/albedo/util"
	"log"
	"net/http"
	"os"
	"time"
)

func InitWeb() {
	InitConfig()

	InitLogger(util.GetRootDir() + "/runtime/log/albedo-web.log")

	InitEnv()

	InitDb()

	InitRedis()

	InitMemCache()

	InitGin()
}

func InitCli() {
	InitConfig()

	InitLogger(util.GetRootDir() + "/runtime/log/albedo-cli.log")

	InitEnv()

	InitDb()

	InitRedis()

	InitMemCache()
}

func InitEnv() {
	loc, err := util.GetLocation()

	if err != nil {
		log.Println("时区设置失败：" + err.Error())
		os.Exit(1)
	}

	time.Local = loc
	logs.Debug(fmt.Sprintf("默认时区（%s）设置成功", loc.String()))
}

func RunWeb() {
	logs.Debug(fmt.Sprintf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")))
	logs.Info(endless.ListenAndServe(viper.GetString("addr"), Engine).Error())

	CheckServer()
}

func CheckServer() {
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logs.Error(fmt.Sprintf("The router has no response, or it might took too long to start up, err: %s", err.Error()))
		}
		logs.Debug("The router has been deployed successfully")
	}()
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logs.Debug("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("can not connect to the router")
}
