package types

// -------------------------------------
// 排课

// 创建课程
// Method: Post
type CreateCourseRequest struct {
	CourseName string
	Cap        int
}

type CreateCourseResponse struct {
	Code ErrNo
	Data struct {
		CourseID string
	}
}

// 获取课程
// Method: Get
type GetCourseRequest struct {
	CourseID string
}

type GetCourseResponse struct {
	Code ErrNo
	Data Course
}

// TODO 教师课程绑定
// 老师绑定课程
// Method： Post
// 注：这里的 teacherID 不需要做已落库校验
// 一个老师可以绑定多个课程 , 不过，一个课程只能绑定在一个老师下面
//type BindCourseRequest struct {
//	CourseID  string
//	TeacherID string
//}
//
//type BindCourseResponse struct {
//	Code ErrNo
//}
//
//// 老师解绑课程
//// Method： Post
//type UnbindCourseRequest struct {
//	CourseID  string
//	TeacherID string
//}
//
//type UnbindCourseResponse struct {
//	Code ErrNo
//}
//
//// 获取老师下所有课程
//// Method：Get
//type GetTeacherCourseRequest struct {
//	TeacherID string
//}
//
//type GetTeacherCourseResponse struct {
//	Code ErrNo
//	Data struct {
//		CourseList []*Course
//	}
//}
//
//// 排课求解器，使老师绑定课程的最优解， 老师有且只能绑定一个课程
//// Method： Post
//type ScheduleCourseRequest struct {
//	TeacherCourseRelationShip map[string][]string // key 为 teacherID , val 为老师期望绑定的课程 courseID 数组
//}
//
//type ScheduleCourseResponse struct {
//	Code ErrNo
//	Data map[string]string // key 为 teacherID , val 为老师最终绑定的课程 courseID
//}
