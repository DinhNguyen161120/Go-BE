package global

import (
	"example.com/m/pkg/logger"
	"example.com/m/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config // Mysql, Redis, ...
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
