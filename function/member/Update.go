package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var request types.UpdateMemberRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var line types.Members
	var response types.UpdateMemberResponse
	database.DB.Model(types.Members{}).Unscoped().Where(&request).Find(&line)
	if line == (types.Members{}) {
		response.Code = types.UserNotExisted
	} else if line.Deleted.Valid {
		response.Code = types.UserHasDeleted
	} else {
		database.DB.Model(types.Members{}).Where("user_id=?", request.UserID).Update("Nickname", request.Nickname)
		response.Code = types.OK
	}
	c.JSON(200, response)
}
