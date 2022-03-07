package types

type BookCourseRequest struct {
	StudentID string
	CourseID  string
}

// 课程已满返回 CourseNotAvailable

type BookCourseResponse struct {
	Code ErrNo
}

type GetStudentCourseRequest struct {
	StudentID string
}

type GetStudentCourseResponse struct {
	Code ErrNo
	Data struct {
		CourseList []Course
	}
}
