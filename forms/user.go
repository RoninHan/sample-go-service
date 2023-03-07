package forms

type PasswordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`       // 验证码id
}

type UserListForm struct {
	ID       string
	Password string
	NickName string
	Head_url string
	Birthday string
	Address  string
	Gender   string
	mobile   string
	Current  int
	PageSize int
}
