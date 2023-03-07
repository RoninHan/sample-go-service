package initialize

import (
	"context"
	"fmt"
	"sample-go-service/global"
	"time"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() {
	color.Blue(fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port))
	addr := fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port)
	// 生成redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 链接redis
	_, err := client.Ping(ctx).Result()
	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		color.Yellow(err.Error())
	}

	client.Set(ctx, "test", "testvalue", time.Second)
	value := client.Get(ctx, "test")
	color.Blue(value.Val())

	//*表示获取所有的key
	keys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

}
