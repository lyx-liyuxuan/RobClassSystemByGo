package main

import (
	"RobClassSystemByGo/database"
	"fmt"
)

func main() {
	fmt.Print("hello world")

	database.InitDb(true)
}
