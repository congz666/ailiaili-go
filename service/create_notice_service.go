package service

import (
	"ailiaili/model"
	"ailiaili/serializer"
)

// CreateNoticeService 视频投稿的服务
type CreateNoticeService struct { //将前端的数据绑定到结构体内
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Create 创建视频
func (service *CreateNoticeService) Create() serializer.Response {
	notice := model.Notice{
		Title: service.Title,
		Info:  service.Info,
	}

	//判断是否有标题
	if notice.Title == "" {
		return serializer.Response{
			Status: 50002,
			Msg:    "请输入标题",
			//Error:  err.Error(),
		}
	}
	//判断是否有内容
	if notice.Info == "" {
		return serializer.Response{
			Status: 50002,
			Msg:    "请输入内容",
			//Error:  err.Error(),
		}
	}

	err := model.DB.Create(&notice).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "公告创建失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildNotice(notice),
	}
}
