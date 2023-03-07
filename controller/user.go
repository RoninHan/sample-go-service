package controller

import (
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/utils"

	"github.com/gin-gonic/gin"

	"sample-go-service/forms"
)

// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}

	if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
		Response.Err(c, 400, 400, "otp failed", "")
	}

	Response.Success(c, 200, "success", "test")
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Current, UserListForm.PageSize)
	// 判断
	if (total + len(userlist)) == 0 {
		Response.Err(c, 400, 400, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	Response.Success(c, 200, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})
}
