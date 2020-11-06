package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"goBlog/controller/users"
	usersValidate "goBlog/validator/users"
)

func UsersRouter(router *gin.RouterGroup) {

	//router.Use(jwt.JWTAuth())
	//router.Use(sign.Sign())
	router.POST("/login", users.Login)
	router.POST("/wx_login", users.WxLogin)
	router.POST("/qq_login", users.QQLogin)

	//字段檢測
	UsersValidate()
}

func UsersValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("CheckUsername", usersValidate.CheckUsername)
	}
}
