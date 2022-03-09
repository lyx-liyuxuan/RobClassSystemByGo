package auth

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func WhoAmI(c *gin.Context) {
	// 从 cookie 获取 sessionKey
	sessionKey, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(200, types.WhoAmIResponse{Code: types.LoginRequired})
		return
	}
	// 从 redis 获取 UserID
	rets, err1 := database.RDB.HMGet(context.Background(), sessionKey, "UserID").Result()
	if err1 != nil {
		log.Println(err)
	}
	userID := rets[0].(string)
	// 从数据库中提出数据
	var member types.Member
	database.DB.Table("members").Where("user_id = ?", userID).Find(&member)
	c.JSON(200, types.WhoAmIResponse{
		Code: types.OK,
		Data: member,
	})
}
