package serializer

import "ailiaili/model"

// Admin 用户序列化器
type Admin struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
}

// AdminResponse 单个用户序列化
type AdminResponse struct {
	Response
	Data Admin `json:"data"`
}

// BuildAdmin 序列化用户
func BuildAdmin(admin model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		UserName:  admin.UserName,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt.Unix(),
	}
}

// BuildAdminResponse 序列化用户响应
func BuildAdminResponse(admin model.Admin) AdminResponse {
	return AdminResponse{
		Data: BuildAdmin(admin),
	}
}
