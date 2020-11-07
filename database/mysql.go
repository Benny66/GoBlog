package database

import (
	"fmt"
	mysqlCfg "goBlog/config/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
	连接mysql数据库
 */
func Connect(Db *gorm.DB) (*gorm.DB, error) {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlCfg.DBCfg.DBUser,
		mysqlCfg.DBCfg.DBPass,
		mysqlCfg.DBCfg.DBHost,
		mysqlCfg.DBCfg.DBPort,
		mysqlCfg.DBCfg.DBName,
	)
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		return Db, err
	}
	//设置闲置的连接数
	Db.DB().SetMaxIdleConns(mysqlCfg.DBCfg.DBMaxIdleConns)
	//设置最大打开的连接数，默认为0表示无限制
	Db.DB().SetMaxOpenConns(mysqlCfg.DBCfg.DBMaxOpenConns)
	return Db, nil
}
