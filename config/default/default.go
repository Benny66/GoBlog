package defaultCfg

import (
	"github.com/joho/godotenv"
	"goBlog/common/alarm"
	_func "goBlog/common/func"
	mysqlCfg "goBlog/config/mysql"
	"os"
	"strconv"
)


type config struct {
	IPAddress      string
	Port           string
	Mode           string
	SignKey        string
	SignExpire     int64
}

var Cfg = config{
	IPAddress: "0.0.0.0",
	Port:      ":8000",
	Mode:         "release",
	SignKey:      "goBlog",
	SignExpire:   60,
}

func LoadENV() {
	if !_func.IsFileExists(_func.GetAbsPath(".env")) {
		return
	}
	alarm.New(_func.GetAbsPath(".env"))
	err := godotenv.Load(_func.GetAbsPath(".env"))
	if err != nil {
		alarm.New(err.Error())
	}
	if v := os.Getenv("IPAddress"); v != "" {
		Cfg.IPAddress = v
	}
	if v := os.Getenv("Port"); v != "" {
		Cfg.Port = v
	}
	if v := os.Getenv("DBHost"); v != "" {
		mysqlCfg.DBCfg.DBHost = v
	}
	if v := os.Getenv("DBPort"); v != "" {
		mysqlCfg.DBCfg.DBPort, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("DBUser"); v != "" {
		mysqlCfg.DBCfg.DBUser = v
	}
	if v := os.Getenv("DBPass"); v != "" {
		mysqlCfg.DBCfg.DBPass = v
	}
	if v := os.Getenv("DBName"); v != "" {
		mysqlCfg.DBCfg.DBName = v
	}
	if v := os.Getenv("DBMaxOpenConns"); v != "" {
		mysqlCfg.DBCfg.DBMaxOpenConns, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("DBMaxIdleConns"); v != "" {
		mysqlCfg.DBCfg.DBMaxIdleConns, _ = strconv.Atoi(v)
	}
	if v := os.Getenv("Mode"); v != "" {
		Cfg.Mode = v
	}
}
