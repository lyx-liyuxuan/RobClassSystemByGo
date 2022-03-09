package types

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

type GetCourseRequest struct {
	CourseID string
}

type GetCourseResponse struct {
	Code ErrNo
	Data Course
}
