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
		c.JSON(200, types.LoginResponse{Code: types.ParamInvalid})
		return
	}

	// get member
	var member types.Member
	database.Db.Model(&types.Member{}).Where(&request).Find(&member)
	if member == (types.Member{}) {
		c.JSON(200, types.LoginResponse{Code: types.WrongPassword})
		return
	}

	// get uuid And Write Cookie
	sessionKey := uuid.NewV4().String()

	ctx := context.Background()
	datas := map[string]interface{}{
		"UserID":   member.UserID,
		"UserType": fmt.Sprint(member.UserType),
	}
	log.Println(datas)
	if err := database.Rdb.HMSet(ctx, sessionKey, datas).Err(); err != nil {
		log.Fatal(err)
	}
	c.SetCookie("camp-session", sessionKey, 3600, "/", "", false, true)

	response := types.LoginResponse{
		Code: types.OK,
		Data: struct{ UserID string }{
			UserID: member.UserID,
		},
	}

	c.JSON(200, response)
}
