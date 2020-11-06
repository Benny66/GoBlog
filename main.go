package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goBlog/common/alarm"
	"goBlog/config/defaultCfg"
	orm "goBlog/database"
	"goBlog/middleware/logger"
	recover2 "goBlog/middleware/recover"
	"goBlog/router"
)


func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer orm.Db.Close()

	gin.SetMode(defaultCfg.Cfg.Mode)

	engine := gin.Default()
	engine.NoRoute(router.RouteNotFound)
	engine.NoMethod(router.MethodNotFound)

	//日志
	engine.Use(logger.LoggerToFile(), recover2.Recover())

	////签名验证，需放在路由配置前
	//engine.Use(sign.Sign())

	//路由
	router.InitRouter(engine)

	err := engine.Run(defaultCfg.Cfg.Port)
	if err != nil {
		alarm.Panic(err.Error())
	}
}
