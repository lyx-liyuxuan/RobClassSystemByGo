package types

type CreateMemberRequest struct {
	Nickname string   `form:"Nickname" json:"Nickname" xml:"Nickname"  binding:"required"` // required，不小于 4 位 不超过 20 位
	Username string   `form:"Username" json:"Username" xml:"Username"  binding:"required"` // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string   `form:"Password" json:"Password" xml:"Password"  binding:"required"` // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType UserType `form:"UserType" json:"UserType" xml:"UserType"  binding:"required"` // required, 枚举值
}

type CreateMemberResponse struct {
	Code ErrNo
	Data struct {
		UserID string // int64 范围
	}
}

// 获取成员信息

type GetMemberRequest struct {
	UserID string `form:"UserID" json:"UserID" xml:"UserID"  binding:"required"`
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type GetMemberResponse struct {
	Code ErrNo
	Data Member
}

// 批量获取成员信息

// TODO 批量获取成员信息
//type GetMemberListRequest struct {
//	Offset int
//	Limit  int
//}
//
//type GetMemberListResponse struct {
//	Code ErrNo
//	Data struct {
//		MemberList []Member
//	}
//}

// 更新成员信息

type UpdateMemberRequest struct {
	UserID   string `form:"UserID" json:"UserID" xml:"UserID"  binding:"required"`
	Nickname string `form:"Nickname" json:"Nickname" xml:"Nickname"  binding:"required"`
}

type UpdateMemberResponse struct {
	Code ErrNo
}

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type DeleteMemberRequest struct {
	UserID string `form:"UserID" json:"UserID" xml:"UserID"  binding:"required"`
}

type DeleteMemberResponse struct {
	Code ErrNo
}
