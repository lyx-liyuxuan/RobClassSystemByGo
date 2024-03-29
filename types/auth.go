package types

type LoginRequest struct {
	Username string `form:"Username" json:"Username" xml:"Username"  binding:"required"`
	Password string `form:"Password" json:"Password" xml:"Password"  binding:"required"`
}

// 登录成功后需要 Set-Cookie("camp-session", ${value})

type LoginResponse struct {
	// 密码错误范围密码错误状态码
	Code ErrNo
	Data struct {
		UserID string
	}
}

// 登出

type LogoutRequest struct{}

// 登出成功需要删除 Cookie

type LogoutResponse struct {
	Code ErrNo
}

// WhoAmI 接口，用来测试是否登录成功，只有此接口需要带上 Cookie

type WhoAmIRequest struct {
}

// 用户未登录请返回用户未登录状态码

type WhoAmIResponse struct {
	Code ErrNo
	Data Member
}
