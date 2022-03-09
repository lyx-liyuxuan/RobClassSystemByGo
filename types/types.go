package types

// 说明：
// 1. 所提到的「位数」均以字节长度为准
// 2. 所有的 ID 均为 int64（以 string 方式表现）

// 通用结构

type ErrNo int

const (
	OK                 ErrNo = 0
	ParamInvalid       ErrNo = 1   // 参数不合法
	UserHasExisted     ErrNo = 2   // 该 Username 已存在
	UserHasDeleted     ErrNo = 3   // 用户已删除
	UserNotExisted     ErrNo = 4   // 用户不存在
	WrongPassword      ErrNo = 5   // 密码错误
	LoginRequired      ErrNo = 6   // 用户未登录
	CourseNotAvailable ErrNo = 7   // 课程已满
	PermDenied         ErrNo = 10  // 没有操作权限
	UnknownError       ErrNo = 255 // 未知错误
)

type ResponseMeta struct {
	Code ErrNo
}

type UserType int

const (
	Admin UserType = 1
	//Student UserType = 2
	//Teacher UserType = 3
)
