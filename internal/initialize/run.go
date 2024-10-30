package initialize

import (
	"fmt"

	"example.com/m/global"
)

func Run() {
	// alway load config first
	LoadConfig()
	fmt.Println("Load config Mysql: ", global.Config.Mysql)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
