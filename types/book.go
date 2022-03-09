package types

type BookCourseRequest struct {
	StudentID string
	CourseID  string
}

// 课程已满返回 CourseNotAvailable

type BookCourseResponse struct {
	Code ErrNo
}

// TODO 添加异步检测，直到完成写入再读取
