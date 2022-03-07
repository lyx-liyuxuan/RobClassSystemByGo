package types

import "gorm.io/gorm"

/// 系统内置管理员账号
// 账号名：JudgeAdmin 密码：JudgePassword2022

// 创建成员
// 参数不合法返回 ParamInvalid

// 只有管理员才能添加

type Member struct {
	UserID   string
	Nickname string
	Username string
	UserType UserType
}

type Members struct {
	UserID   string         `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	Nickname string         `gorm:"type:varchar(32) not null"`
	Username string         `gorm:"type:varchar(32) not null;uniqueIndex:udx_name"`
	UserType UserType       `gorm:"type:int not null"`
	Password string         `gorm:"type:char(32) not null"`
	Deleted  gorm.DeletedAt `gorm:"uniqueIndex:udx_name;"`
}

type Course struct {
	CourseID  string
	Name      string
	TeacherID string
}

type Courses struct {
	CourseID  string `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(32) not null"`
	Cap       int    `gorm:"type:int not null"`
	TeacherID string `gorm:"type:varchar(32);index"`
}

type SCourses struct {
	SCID      string `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	CourseID  string `gorm:"type:bigint UNSIGNED not null"`
	StudentID string `gorm:"type:bigint UNSIGNED not null;index"`
}
