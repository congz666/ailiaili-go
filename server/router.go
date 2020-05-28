package server

import (
	"ailiaili/api"
	"ailiaili/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 中间件, 顺序不能改
		v1.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
		v1.Use(middleware.Cors()) //跨域
		// 获取用户身份
		v1.Use(middleware.CurrentUser())

		v1.POST("ping", api.Ping)

		//视频详情
		v1.GET("video/:id", api.ShowVideo)
		//用户详情
		v1.GET("users/:id", api.ShowUser)
		//查询视频列表
		v1.GET("videos", api.ListVideo)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 用户注册
		v1.POST("user/register", api.UserRegister)

		//头像上传
		v1.POST("upload/token", api.UploadToken)
		//视频上传
		v1.POST("upload/tack", api.UploadTack)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//查询视频评论列表
		v1.GET("videos/comments/:id", api.ListComment)
		//查询用户视频列表
		v1.GET("user/videos/:id", api.UserVideos)

		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("user/me", api.UserMe)
			//用户退出
			v1.DELETE("user/logout", api.UserLogout)
			//用户投稿
			v1.POST("videos", api.CreateVideo)
			//更新视频
			v1.PUT("video/:id", api.UpdateVideo)
			//删除视频
			v1.DELETE("video/:id", api.DeleteVideo)
			//查询自己视频列表
			v1.GET("my/videos", api.UserMeVideos)
			//更新用户信息
			v1.PUT("user/account", api.UserChange)
			//发表评论(视频id))
			v1.POST("video/comment/:id", api.VideoComment)
			//删除评论(评论id)
			v1.DELETE("comment/:id", api.DeleteComment)
		}
		//管理员组
		v2 := r.Group("api/v2")
		{
			// 中间件, 顺序不能改
			v2.Use(middleware.AdminSession(os.Getenv("SESSION_SECRET")))
			v2.Use(middleware.Cors()) //跨域
			// 获取用户身份
			v2.Use(middleware.CurrentAdmin())

			// 管理员注册
			v2.POST("admin/register", api.AdminRegister)
			// 管理员登录
			v2.POST("admin/login", api.AdminLogin)
			// Admin Routing
			v2.GET("admin/me", api.AdminMe)

			//查询公告列表
			v2.GET("notices", api.ListNotices)

			v2.Use(middleware.AdminRequired())
			{
				//删除视频
				v2.DELETE("video/:id", api.DeleteVideo)
				//新建公告
				v2.POST("notice", api.CreateNotice)
				//更新公告
				v2.PUT("notice/:id", api.UpdateNotice)
				//删除公告
				v2.DELETE("notice/:id", api.DeleteNotice)
				//管理员退出
				v2.DELETE("admin/logout", api.AdminLogout)
			}
		}

		//

	}
	return r
}
