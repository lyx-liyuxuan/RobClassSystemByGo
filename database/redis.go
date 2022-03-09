package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis() {
	// 连接 Reids
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	// 利用根Context创建一个父Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := RDB.Ping(ctx).Result()

	if err != nil {
		fmt.Println("open redis fail")
		return
	}

	// 删除 redis 缓存
	res, err := RDB.FlushDB(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("delete redis:", res)

	// TODO 从数据库中提取剩余 Cap

	fmt.Println("open redis success")
}
