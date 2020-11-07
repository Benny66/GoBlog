package mysqlCfg

type dbConfig struct {
	DBHost         string
	DBPort         int
	DBUser         string
	DBPass         string
	DBName         string
	DBMaxIdleConns int
	DBMaxOpenConns int
}

var DBCfg = dbConfig{
	//数据库配置
	DBHost:         "127.0.0.1",
	DBPort:         3306,
	DBUser:         "root",
	DBPass:         "test",
	DBName:         "test",
	DBMaxIdleConns: 1000,
	DBMaxOpenConns: 0,
}