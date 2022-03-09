package course

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var request types.CreateCourseRequest
	if err := c.ShouldBind(request); err != nil {
		log.Println(err)
		return
	}

	database.DB.Model(types.Courses{}).Create(&request)
	line := types.Course{}
	database.DB.Model(types.Courses{}).Where(&request).Find(&line)
	response := types.CreateCourseResponse{
		Code: types.OK,
		Data: struct{ CourseID string }{CourseID: line.CourseID},
	}
	c.JSON(200, response)
}
