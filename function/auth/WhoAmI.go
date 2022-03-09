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
	rets, err := database.RDB.HMGet(context.Background(), sessionKey, "UserID").Result()
	if err != nil {
		log.Println(err)
	}
	userID := rets[0].(string)

	var member types.Member
	database.DB.Table("members").Where("user_id = ?", userID).Find(&member)
	c.JSON(200, types.WhoAmIResponse{
		Code: types.OK,
		Data: member,
	})
}
