package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var request types.DeleteMemberRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var line types.Members
	var response types.DeleteMemberResponse
	database.DB.Model(types.Members{}).Unscoped().Where(&request).Find(&line)
	if line == (types.Members{}) {
		response.Code = types.UserNotExisted
	} else if line.Deleted.Valid {
		response.Code = types.UserHasDeleted
	} else {
		database.DB.Model(types.Members{}).Where("user_id=?", request.UserID).Delete(&types.Members{})
		response.Code = types.OK
	}
	c.JSON(200, response)
}
