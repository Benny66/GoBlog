package defaultCfg

import (
	"github.com/joho/godotenv"
	"goBlog/common/alarm"
	_func "goBlog/common/func"
	"os"
)

const (
	LOG_FILE_NAME = "system.log"
)

type config struct {
	IPAddress    string
	Port         string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	Mode         string
	SignKey      string
	SignExpire   int64
	RedisHost    string
	RedisPass    string
	RedisDB      string
	RedisTimeout int64
}

var Cfg = config{
	IPAddress:    "0.0.0.0",
	Port:         ":8000",
	DBHost:       "127.0.0.1",
	DBPort:       "3306",
	DBUser:       "root",
	DBPass:       "test",
	DBName:       "test",
	Mode:         "release",
	SignKey:      "goBlog",
	SignExpire:   60,
	RedisHost:    "127.0.0.1",
	RedisPass:    "test",
	RedisDB:      "1",
	RedisTimeout: 300,
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
		Cfg.DBHost = v
	}
	if v := os.Getenv("DBPort"); v != "" {
		Cfg.DBPort = v
	}
	if v := os.Getenv("DBUser"); v != "" {
		Cfg.DBUser = v
	}
	if v := os.Getenv("DBPass"); v != "" {
		Cfg.DBPass = v
	}
	if v := os.Getenv("DBName"); v != "" {
		Cfg.DBName = v
	}
	if v := os.Getenv("Mode"); v != "" {
		Cfg.Mode = v
	}
}
