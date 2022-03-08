package main

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("hello world\n")
	// TODO Init Database

	// connect table
	database.ConnectDb()

	// init table
	// database.InitDb()

	database.InitRedis()

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	router.RegisterRouter(g)
	g.Run(":80")
}
