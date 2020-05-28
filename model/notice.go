package model

import (
	"github.com/jinzhu/gorm"
)

// Notice 视频模型
type Notice struct {
	gorm.Model
	Title string
	Info  string
}
