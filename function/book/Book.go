package book

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Book(c *gin.Context) {
	// TODO 压测

	// 获取请求
	var request types.BookCourseRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	// 获取剩余容量
	ctx := context.Background()
	result, _ := database.RDB.Get(ctx, "CID"+request.CourseID).Result()
	val, _ := strconv.Atoi(result)

	if val <= 0 {
		c.JSON(http.StatusOK, types.BookCourseResponse{Code: types.CourseNotAvailable})
	} else {
		val -= 1 // 消耗容量
		// 异步写入
		copyContext := c.Copy()
		database.RDB.Set(ctx, "CID"+request.CourseID, val, 0)
		go func() {
			AsyncMysql(copyContext)
		}()
		c.JSON(200, types.BookCourseResponse{Code: types.OK})
	}
}

func AsyncMysql(c *gin.Context) {
	// TODO 选课与否， 添加失败返还redis数据
	var request types.BookCourseRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}
	database.DB.Table("s_courses").Create(&request)
}
