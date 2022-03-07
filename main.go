package main

import (
	"RobClassSystemByGo/database"
	"fmt"
)

func main() {
	fmt.Print("hello world\n")
	// TODO Init Database

	// connect table
	database.ConnectDb()

	// init table
	// database.InitDb()
}
