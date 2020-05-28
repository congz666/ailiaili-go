package service

import (
	"ailiaili/model"
	"ailiaili/serializer"
)

// UpdateNoticeService 更新视频的服务
type UpdateNoticeService struct { //将前端的数据绑定到结构体内
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新视频
func (service *UpdateNoticeService) Update(id string) serializer.Response {
	var notice model.Notice
	//找到视频
	err := model.DB.First(&notice, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "公告不存在",
			Error:  err.Error(),
		}
	}

	notice.Title = service.Title
	notice.Info = service.Info
	err = model.DB.Save(&notice).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "公告修改失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildNotice(notice),
	}
}
