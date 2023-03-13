package controller

import (
	"sample-go-service/Response"
	"sample-go-service/dao"
	"sample-go-service/models"
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

	// if !store.Verify(PasswordLoginForm.CaptchaId, PasswordLoginForm.Captcha, true) {
	// 	Response.Err(c, 400, 400, "otp failed", "")
	// 	return
	// }

	user, ok := dao.FindUserInfo(PasswordLoginForm.UserName, PasswordLoginForm.PassWord)
	if !ok {
		Response.Err(c, 401, 401, "未注册该用户", "")
		return
	}
	token := utils.CreateToken(c, int(user.ID), user.NickName, user.Role)
	userinfoMap := HandleUserModelToMap(user)
	userinfoMap["token"] = token
	Response.Success(c, 200, "success", userinfoMap)
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

func HandleUserModelToMap(user *models.User) map[string]interface{} {
	birthday := ""
	if user.Birthday == nil {
		birthday = ""
	} else {
		birthday = user.Birthday.Format("2006-01-02")
	}
	userItemMap := map[string]interface{}{
		"id":        user.ID,
		"nick_name": user.NickName,
		"head_url":  user.HeadUrl,
		"birthday":  birthday,
		"address":   user.Address,
		"desc":      user.Desc,
		"gender":    user.Gender,
		"role":      user.Role,
		"mobile":    user.Mobile,
	}
	return userItemMap
}
