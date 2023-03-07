package utils

import (
	"net/http"
	"sample-go-service/Response"
	"sample-go-service/global"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	//如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		Response.Err(c, http.StatusBadRequest, 500, "字段較嚴錯誤", err.Error())
	}
	msg := removeTopStruct(errs.Translate(global.Trans))
	Response.Err(c, http.StatusBadRequest, 400, "字段較嚴錯誤", msg)
	return
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		// 从文本的逗号开始切分   处理后"mobile": "mobile为必填字段"  处理前: "PasswordLoginForm.mobile": "mobile为必填字段"
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
