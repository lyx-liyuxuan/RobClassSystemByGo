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

	// get request
	var request types.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	// get line
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

	ctx := context.Background()
	data := map[string]interface{}{
		"UserID":   line.UserID,
		"UserType": fmt.Sprint(line.UserType),
	}
	//log.Println(data)
	if err := database.RDB.HMSet(ctx, sessionKey, data).Err(); err != nil {
		log.Fatal(err)
		return
	}
	c.SetCookie("camp-session", sessionKey, 3600, "/", "", false, true)

	response := types.LoginResponse{
		Code: types.OK,
		Data: struct{ UserID string }{
			UserID: line.UserID,
		},
	}

	c.JSON(200, response)
}
