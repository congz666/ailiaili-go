package api

import (
	"ailiaili/serializer"
	"ailiaili/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AdminRegister 管理员注册接口
func AdminRegister(c *gin.Context) {
	var service service.AdminRegisterService
	if err := c.ShouldBind(&service); err == nil {
		if admin, err := service.AdminRegister(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildAdminResponse(admin)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminLogin 管理员登录接口
func AdminLogin(c *gin.Context) {
	var service service.AdminLoginService
	if err := c.ShouldBind(&service); err == nil {
		if admin, err := service.Login(); err == nil {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("admin_id", admin.ID)
			s.Save()
			//c.JSON(200, gin.H{"user_id": s.Get("user_id")})
			res := serializer.BuildAdminResponse(admin)
			c.JSON(200, res)
		} else {
			c.JSON(200, err)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminMe 管理员详情
func AdminMe(c *gin.Context) {
	admin, err := CurrentAdmin(c)
	if err != nil {
		c.JSON(200, err)
	} else {
		res := serializer.BuildAdminResponse(*admin)
		c.JSON(200, res)
	}
}

// AdminLogout 用户登出
func AdminLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}
