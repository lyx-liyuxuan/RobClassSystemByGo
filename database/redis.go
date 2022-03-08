package database

import (
	"RobClassSystemByGo/types"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() {
	// 连接 Reids
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	// TODO 利用根Context创建一个父Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		fmt.Println("open redis fail")
		return
	}

	// 删除 redis 缓存
	res, err := Rdb.FlushDB(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("delete redis:", res)

	// 读取学生选课和课程列表，构建缓存
	course_cnt := make(map[string]int)
	var courses []struct {
		CourseID string
		Cap      int
	}
	Db.Table("courses").Find(&courses)
	for i := range courses {
		course_cnt[courses[i].CourseID] = courses[i].Cap
	}

	var data []types.SCourses
	Db.Table("s_courses").Find(&data)
	for i := range data {
		course_cnt[data[i].CourseID] -= 1
		err := Rdb.HSetNX(ctx, data[i].StudentID, data[i].CourseID, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	for k, v := range course_cnt {
		Rdb.Set(ctx, k+"cnt", v, 0)
	}

	fmt.Println("open redis success")
}
