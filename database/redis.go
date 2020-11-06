package database

import (
	"github.com/garyburd/redigo/redis"
	"goBlog/config/defaultCfg"
	"time"
)

var RedisPool Redis

func init() {
	//初始化
	RedisPool = Redis{
		RedisHost:     defaultCfg.Cfg.RedisHost,
		RedisPassword: defaultCfg.Cfg.RedisPass,
		RedisDB:       defaultCfg.Cfg.RedisDB,
		RedisTimeout:  defaultCfg.Cfg.RedisTimeout,
		RedisPool:     nil,
	}
	RedisPool.RedisPool = RedisPool.NewPool()
}

type Redis struct {
	RedisHost     string
	RedisDB       string
	RedisPassword string
	RedisTimeout  int64
	RedisPool     *redis.Pool
}

func (r *Redis) NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     16,  //最初的连接数量
		MaxActive:   0,   //连接池最大的连接数量，0按需分配
		IdleTimeout: 120, //连接关闭时间（秒）
		Wait:        true,
		Dial:        r.RedisConnect,
	}
}

func (r *Redis) RedisConnect() (redis.Conn, error) {
	c, err := redis.Dial("tcp", r.RedisHost)
	if err != nil {
		return nil, err
	}
	_, err = c.Do("AUTH", r.RedisPassword)
	if err != nil {
		return nil, err
	}
	_, err = c.Do("SELECT", r.RedisDB)
	if err != nil {
		return nil, err
	}
	redis.DialConnectTimeout(time.Duration(r.RedisTimeout) * time.Second)
	redis.DialReadTimeout(time.Duration(r.RedisTimeout) * time.Second)
	redis.DialWriteTimeout(time.Duration(r.RedisTimeout) * time.Second)
	return c, nil
}

func (r *Redis) Get(k string) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	v, err := c.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r *Redis) Set(k, value string) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, value)
	return err
}

func (r *Redis) SetExp(k, value string, expire int64) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, value, "EX", expire)
	return err
}
