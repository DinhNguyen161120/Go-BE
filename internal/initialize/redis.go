package initialize

import (
	"context"
	"fmt"

	"example.com/m/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	addr := fmt.Sprintf("%s:%v", r.Host, r.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Reids init error: ", zap.Error(err))
	}

	fmt.Println("InitRedis success!")

	global.Rdb = rdb
	// redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "example", "test redis", 0).Err()
	if err != nil {
		fmt.Println("Text redis error: ", err)
		return
	}

	value, err := global.Rdb.Get(ctx, "example").Result()

	if err != nil {
		fmt.Println("Redis Get error: ", err)
		return
	}

	fmt.Printf("Redis Get value success value = %s \n", value)
}
