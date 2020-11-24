package router

import (
	"github.com/gin-gonic/gin"
	"goBlog/common/response"
	"goBlog/config/code"
	"goBlog/middleware/cross"
	"goBlog/router/users"
	"net/http"
)

func InitRouter(r *gin.Engine) {

	//跨域
	r.Use(cross.Cross())
	api := r.Group("/api")

	//用戶路由
	users.UsersRouter(api.Group("/users"))


}

/*
* description: 方法不存在
 */
func MethodNotFound(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, response.NewErrResponse(codeCfg.InvalidRequest))
	return
}

/*
* description: 路由不存在
 */
func RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, response.NewErrResponse(codeCfg.NotFound))
	return
}
