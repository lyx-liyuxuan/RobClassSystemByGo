package member

import (
	"RobClassSystemByGo/database"
	"RobClassSystemByGo/types"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CheckPermission
// 权限校验
// 管理员 -> true
// 其余 -> false
func CheckPermissions(c *gin.Context) bool {

	// TODO 修改
	sessionKey, err := c.Cookie("camp-session")
	if err != nil {
		return false
	}
	rets, err1 := database.RDB.HMGet(context.Background(), sessionKey, "UserType").Result()
	if err1 != nil || rets[0].(string) != fmt.Sprint(types.Admin) {
		return false
	} else {
		return true
	}

}

// TODO 参数校验
func CheckParameter(request types.CreateMemberRequest) bool {
	return true
}
