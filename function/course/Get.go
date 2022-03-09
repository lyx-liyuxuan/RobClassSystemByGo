package course

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"log"

	"github.com/gin-gonic/gin"
)

// TODO 判断课程存在与否
func Get(c *gin.Context) {
	var request types.GetCourseRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var line types.Course
	database.DB.Model(types.Courses{}).Where(&request).Find(&line)
	c.JSON(200, types.GetCourseResponse{
		Code: types.OK,
		Data: line,
	})
}
