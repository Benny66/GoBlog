package redisCfg

type config struct {
	RedisHost    string
	RedisPass    string
	RedisDB      string
	RedisTimeout int64
}

var RedisCfg = config{
	RedisHost:    "127.0.0.1",
	RedisPass:    "test",
	RedisDB:      "1",
	RedisTimeout: 300,
}
