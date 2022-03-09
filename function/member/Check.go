package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// CheckPermission
// 权限校验
// 管理员 -> true
// 其余 -> false
func CheckPermissions(c *gin.Context) bool {

	sessionKey, err := c.Cookie("camp-session")
	if err != nil {
		return false
	}
	rets, err := database.RDB.HMGet(context.Background(), sessionKey, "UserType").Result()
	if err != nil || rets[0].(string) != "1" {
		return false
	} else {
		return true
	}
}

func CheckParameter(request types.CreateMemberRequest) bool {
	// 用户昵称
	for _, v := range request.Nickname {
		if !(unicode.IsLetter(v)) {
			return false
		}
	}
	len := strings.Count(request.Nickname, "")
	if len < 4 || len > 20 {
		return false
	}

	// 用户名
	for _, v := range request.Username {
		if !(unicode.IsLetter(v)) {
			return false
		}
	}
	len = strings.Count(request.Username, "")
	if len < 8 || len > 20 {
		return false
	}

	// 密码
	upper, lower, digit := false, false, false
	for _, v := range request.Password {
		if !(unicode.IsLetter(v) || unicode.IsDigit(v)) {
			return false
		}
		if unicode.IsUpper(v) {
			upper = true
		} else if unicode.IsDigit(v) {
			digit = true
		} else if unicode.IsLower(v) {
			lower = true
		}
	}
	if !(digit && lower && upper) {
		return false
	}
	len = strings.Count(request.Password, "")
	if len < 8 || len > 20 {
		return false
	}

	// 用户类型
	if request.UserType < 1 || request.UserType > 3 {
		return false
	}

	return true
}
