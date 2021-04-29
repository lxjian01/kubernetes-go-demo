package redis

import (
	"github.com/gomodule/redigo/redis"
	"kubernetes-go-demo/config"
	"kubernetes-go-demo/global/log"
	"time"
)

var (
	redisPool *redis.Pool
)

// 初始化redis数据库
func InitRedis() {
	conf := config.GetConfig().Redis
	var err error
	redisPool = &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			var conn redis.Conn
			conn, err = redis.Dial("tcp", conf.Host)
			if _, err := conn.Do("AUTH", conf.Password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
	}
	if err != nil {
		panic(err)
	}
}

func CloseRedis(){
	if err := redisPool.Close();err != nil{
		log.Error("Close redis error by",err)
	}
}
