package api

import (
	"ailiaili/service"

	"github.com/gin-gonic/gin"
)

// CreateNotice 新建公告
func CreateNotice(c *gin.Context) {
	service := service.CreateNoticeService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListNotices 公告列表详情接口
func ListNotices(c *gin.Context) {
	service := service.ListNoticesService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateNotice 更新公告的接口
func UpdateNotice(c *gin.Context) {
	service := service.UpdateNoticeService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteNotice 删除视频的接口
func DeleteNotice(c *gin.Context) {
	service := service.DeleteNoticeService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
