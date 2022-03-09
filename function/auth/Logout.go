package auth

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	sessionKey, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(200, types.LogoutResponse{
			Code: types.LoginRequired,
		})
		return
	}
	ctx := context.Background()
	database.RDB.Del(ctx, sessionKey)
	c.SetCookie("camp-session", sessionKey, -1, "/", "", false, true)

	c.JSON(200, types.LogoutResponse{
		Code: types.OK,
	})
}
