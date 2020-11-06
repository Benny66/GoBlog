package database

import (
	"goBlog/common/alarm"
	"goBlog/config/defaultCfg"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB
var err error

func init() {
	//设置全局变量
	defaultCfg.LoadENV()

	Db, err = gorm.Open("mysql",
		defaultCfg.Cfg.DBUser+":"+defaultCfg.Cfg.DBPass+"@tcp("+defaultCfg.Cfg.DBHost+":"+defaultCfg.Cfg.DBPort+")/"+defaultCfg.Cfg.DBName+
			"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		alarm.Panic("mysql connect error:" + err.Error())
	} else {
		alarm.New("mysql connect success:" + defaultCfg.Cfg.DBHost)
	}
}
