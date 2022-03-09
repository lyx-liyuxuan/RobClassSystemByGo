package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	if !CheckPermissions(c) {
		c.JSON(200, types.CreateMemberResponse{
			Code: types.PermDenied,
		})
		return
	}

	var request types.CreateMemberRequest
	if err := c.ShouldBind(request); err != nil {
		log.Println(err)
		return
	}

	if !CheckParameter(request) {
		c.JSON(200, types.CreateMemberResponse{
			Code: types.ParamInvalid,
		})
		return
	}

	var line types.Members
	database.DB.Model(types.Members{}).Unscoped().Where("username = ?", request.Username).Find(&line)
	var response types.CreateMemberResponse
	if line != (types.Members{}) {
		if line.Deleted.Valid {
			response.Code = types.UserHasDeleted
		} else {
			response.Code = types.UserHasExisted
		}
	} else {
		database.DB.Model(types.Members{}).Create(&request)
		var member types.Member
		database.DB.Model(types.Members{}).Where("username=?", request.Username).Find(&member)
		response.Code = types.OK
		response.Data = struct{ UserID string }{UserID: member.UserID}
	}
	c.JSON(200, response)
}
