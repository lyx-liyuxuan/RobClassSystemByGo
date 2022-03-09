package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// 权限检验
	if !CheckPermissions(c) {
		c.JSON(200, types.CreateMemberResponse{
			Code: types.PermDenied,
		})
		return
	}

	var request types.CreateMemberRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		c.JSON(200, types.CreateMemberResponse{
			Code: types.UnknownError,
		})
		return
	}

	// 参数检验
	if !CheckParameter(request) {
		c.JSON(200, types.CreateMemberResponse{
			Code: types.ParamInvalid,
		})
		return
	}

	// 完全查找数据行（包括删除) 比判断数据状态
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
		database.DB.Table("members").Create(&request)
		var member types.Member
		database.DB.Model(types.Members{}).Where("username=?", request.Username).Find(&member)
		response.Code = types.OK
		response.Data = struct{ UserID string }{UserID: member.UserID}
	}

	c.JSON(200, response)
}
