package main

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("hello world\n")

	// TODO 判断是否需要重置数据库
	// 仅连接数据库
	//database.ConnectDb()
	// 重置数据库并连接
	database.InitDb()

	database.InitRedis()

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	router.RegisterRouter(g)
	err := g.Run(":80")
	if err != nil {
		return
	}
}

// TODO 功能完善
// TODO 压测
// TODO 并发优化
// TODO Mysql优化
// TODO Redis优化
