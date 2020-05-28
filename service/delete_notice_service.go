package service

import (
	"ailiaili/model"
	"ailiaili/serializer"
)

// DeleteNoticeService 删除视频的服务
type DeleteNoticeService struct { //将前端的数据绑定到结构体内
}

// Delete 删除视频
func (service *DeleteNoticeService) Delete(id string) serializer.Response {
	var notice model.Notice
	err := model.DB.First(&notice, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "公告不存在",
			Error:  err.Error(),
		}
	}

	//删除数据库内容
	err = model.DB.Delete(&notice).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "公告删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Msg: "公告删除成功",
	}
}
