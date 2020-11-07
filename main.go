package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goBlog/common/alarm"
	_func "goBlog/common/func"
	"goBlog/config/default"
	"goBlog/database"
	"goBlog/middleware/logger"
	recover2 "goBlog/middleware/recover"
	"goBlog/router"
)

var Db *gorm.DB

func init() {
	defaultCfg.LoadENV()
	var err error
	Db, err = database.Connect(Db)
	if err != nil {
		alarm.Panic("mysql connect error:" + err.Error())
		return
	}
	alarm.New("mysql connect success")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			alarm.Panic("system error:" + _func.InterfaceToString(err))
		}
	}()
	defer Db.Close()

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
