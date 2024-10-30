package initialize

import (
	"example.com/m/global"
	"example.com/m/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
