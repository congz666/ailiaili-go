package serializer

import (
	"ailiaili/model"
)

// Notice 视频序列化器
type Notice struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"Info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildNotice 序列化公告
func BuildNotice(item model.Notice) Notice {
	return Notice{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(), //time.Time转换为int64（时间戳）
	}
}

// BuildNotices 序列化公告列表
func BuildNotices(items []model.Notice) []Notice {
	var notices []Notice

	for _, item := range items {
		notice := BuildNotice(item)
		//把video给videos然后赋给videos 因为不是传指针（ps:有点傻逼）
		notices = append(notices, notice)
	}
	return notices
}
