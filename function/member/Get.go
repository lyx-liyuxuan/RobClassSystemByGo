package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {

	var request types.GetMemberRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var line types.Members
	var response types.GetMemberResponse
	database.DB.Model(types.Members{}).Unscoped().Where(&request).Find(&line)

	if line == (types.Members{}) {
		response.Code = types.UserNotExisted
	} else if line.Deleted.Valid {
		response.Code = types.UserHasDeleted
	} else {
		response.Code = types.OK
		response.Data = types.Member{
			UserID:   line.UserID,
			Nickname: line.Nickname,
			Username: line.Username,
			UserType: line.UserType,
		}
	}

	c.JSON(200, response)
}
