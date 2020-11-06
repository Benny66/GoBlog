package users

import (
	"goBlog/models/base"
)

type UserWxModel struct {
	base.BaseModel
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Openid    string `json:"openid"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Gender    string `json:"gender"`
}

func (UserWxModel) TableName() string {
	return "blog_users_wx"
}
