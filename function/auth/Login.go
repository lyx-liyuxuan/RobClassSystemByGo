package auth

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Login(c *gin.Context) {
	// 获取请求
	var request types.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	// 获取数据行
	var line types.Members
	database.DB.Model(&types.Member{}).Where(&request).Find(&line)
	if line == (types.Members{}) {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
		return
	}

	// 获取唯一标识符 uuid 作为该数据行的键
	sessionKey := uuid.NewV4().String()

	// redis 记录 sessionKey 对应的 UserID， UserType
	ctx := context.Background()
	// TODO 优化 UserType 写法
	data := map[string]string{
		"UserID":   line.UserID,
		"UserType": fmt.Sprint(line.UserType),
	}
	if err := database.RDB.HMSet(ctx, sessionKey, data).Err(); err != nil {
		log.Fatal(err)
		return
	}

	// 设置 cookie
	c.SetCookie("camp-session", sessionKey, 3600, "/", "", false, true)

	response := types.LoginResponse{
		Code: types.OK,
		Data: struct{ UserID string }{
			UserID: line.UserID,
		},
	}

	c.JSON(200, response)
}
