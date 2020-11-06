package users

import (
	"goBlog/models/base"
)

type UsersModel struct {
	base.BaseModel
	UserId     int    `json:"user_id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	AvatarUrl  string `json:"avatar_url"`
	ScreenName string `json:"screenName"`
}

func (UsersModel) TableName() string {
	return "blog_users"
}
