package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&Comment{})
	DB.AutoMigrate(&VideoLike{})
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Notice{})
	//创建外键
	DB.Model(&VideoLike{}).AddForeignKey("video_id", "videos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&VideoLike{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
