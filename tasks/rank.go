package tasks

import (
	"ailiaili/cache"
	"ailiaili/model"
)

// RestartDailyRank 重启一天的排名(误)
func RestartDailyRank() error {
	return cache.RedisClient.Del("rank:daily").Err()
}

// RestartPucnt 重启一天投稿数
func RestartPucnt() error {
	// var user model.User
	err := model.DB.Table("users").Updates(map[string]interface{}{"upcnt": 0}).Error
	return err
}
