package service

import (
	"ailiaili/model"
	"ailiaili/serializer"
)

// AdminLoginService 管理用户登录的服务
type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *AdminLoginService) Login() (model.Admin, *serializer.Response) {
	var admin model.Admin

	if err := model.DB.Where("user_name = ?", service.UserName).First(&admin).Error; err != nil {
		return admin, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if admin.AdminCheckPassword(service.Password) == false {
		return admin, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return admin, nil
}
