package router

import (
	"RobClassSystemByGo/function/auth"
	"RobClassSystemByGo/function/course"
	"RobClassSystemByGo/function/member"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 登录
	g.POST("/auth/login", auth.Login)
	g.POST("/auth/logout", auth.Logout)
	g.GET("/auth/whoami", auth.WhoAmI)

	// // 成员管理
	g.POST("/member/create", member.Create)
	g.GET("/member/get", member.Get)
	g.POST("/member/update", member.Update)
	g.POST("/member/delete", member.Delete)

	// // 排课
	g.POST("/course/create", course.Create)
	g.GET("/course/get", course.Get)

	// g.POST("/teacher/bind_course", controller.Teacher_bind_course)
	// g.POST("/teacher/unbind_course", controller.Teacher_unbind_course)
	// g.GET("/teacher/get_course", controller.Teacher_get_course)
	// g.POST("/course/schedule", controller.Course_schedule)

	// // 抢课
	// g.POST("/student/book_course", controller.Student_book_course)
	// g.GET("/student/course", controller.Student_course)

}
