package course

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

// TODO 判断课程存在与否
func Create(c *gin.Context) {
	var request types.CreateCourseRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	database.DB.Table("courses").Create(&request)
	line := types.Courses{}
	database.DB.Model(types.Courses{}).Where(&request).Find(&line)
	ctx := context.Background()
	database.RDB.Set(ctx, "CID"+line.CourseID, line.Cap, 0)
	response := types.CreateCourseResponse{
		Code: types.OK,
		Data: struct{ CourseID string }{CourseID: line.CourseID},
	}
	c.JSON(200, response)
}
