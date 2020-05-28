package service

import (
	"ailiaili/model"
	"ailiaili/serializer"
)

// ListNoticesService 公告列表服务
type ListNoticesService struct {
}

// List 公告列表
func (service *ListNoticesService) List() serializer.Response {
	notices := []model.Notice{}
	if err := model.DB.Order("ID desc").Find(&notices).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildNotices(notices),
	}
}
