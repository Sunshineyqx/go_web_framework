package api

import (
	"encoding/json"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if errs, ok := err.(validator.ValidationErrors); ok {
		return serializer.Response{
			Code:  serializer.CodeValidateErr,
			Error: util.RemoveTopStruct(errs.Translate(util.Translator())),
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err) // 直接参数错误....可以优化 根据binding。。
}
