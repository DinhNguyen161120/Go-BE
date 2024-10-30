package initialize

import (
	"fmt"

	"example.com/m/global"
	"go.uber.org/zap"
)

func Run() {
	// alway load config first
	LoadConfig()
	fmt.Println("Load config Mysql: ", global.Config.Mysql)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
