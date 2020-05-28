package api

import (
	"ailiaili/conf"
	"ailiaili/model"
	"ailiaili/serializer"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "Pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	user, _ := c.Get("user")
	if user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// CurrentAdmin 获取当前管理员
func CurrentAdmin(c *gin.Context) (*model.Admin, *serializer.Response) {
	admin, _ := c.Get("admin")
	if admin != nil {
		if a, ok := admin.(*model.Admin); ok {
			return a, nil
		}
	}
	return nil, &serializer.Response{
		Status: 40002,
		Msg:    "需要登录",
	}
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
