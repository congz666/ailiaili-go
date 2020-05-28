package service

import (
	"ailiaili/model"
	"ailiaili/serializer"
)

// AdminRegisterService 管理用户注册服务
type AdminRegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// AdminValid 验证表单
func (service *AdminRegisterService) AdminValid() *serializer.Response {
	//验证注册时的前后密码
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不相同",
		}
	}
	count := 0
	//去账号相同（）
	model.DB.Model(&model.Admin{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "用户账号已经注册",
		}
	}

	return nil
}

// AdminRegister 用户注册
func (service *AdminRegisterService) AdminRegister() (model.Admin, *serializer.Response) {
	admin := model.Admin{
		UserName: service.UserName,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.AdminValid(); err != nil {
		return admin, err
	}

	// 加密密码
	if err := admin.AdminSetPassword(service.Password); err != nil {
		return admin, &serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&admin).Error; err != nil {
		return admin, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return admin, nil
}
