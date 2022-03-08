package auth

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func WhoAmI(c *gin.Context) {
	sessionKey, err := c.Cookie("camp-session")

	if err != nil {
		c.JSON(200, types.WhoAmIResponse{Code: types.LoginRequired})
		return
	}
	rets, err := database.Rdb.HMGet(context.Background(), sessionKey, "UserID").Result()
	log.Println("rets:", rets, err)

	member := types.Member{
		UserID: string(rets[0].(string)),
	}
	database.Db.Table("members").Find(&member).Where((&member))
	if member == (types.Member{}) {
		c.JSON(200, types.WhoAmIResponse{Code: types.LoginRequired})
	} else {
		c.JSON(200, types.WhoAmIResponse{Code: types.OK, Data: member})
	}
}
