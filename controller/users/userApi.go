package users

import (
	"github.com/gin-gonic/gin"
	"goBlog/common/response"
	"goBlog/config/codeCfg"
	"goBlog/validator/users"
	"net/http"
)

func Register(c *gin.Context) {

}

func Login(c *gin.Context) {
	var loginParams = users.LoginParams{}
	if c.BindJSON(&loginParams) == nil {
		//检测用户名密码

		//生成token

		//返回登录成功信息
		c.JSON(http.StatusOK, response.NewSuccessResponse(loginParams))
	} else {
		c.JSON(http.StatusOK, response.NewErrResponse(codeCfg.ParamError))
	}
}

func WxLogin(c *gin.Context)  {

	c.JSON(http.StatusOK, response.NewSuccessResponse(true))

}

func QQLogin(c *gin.Context)  {
	c.JSON(http.StatusOK, response.NewSuccessResponse(true))

}

func ForgetPassword(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
