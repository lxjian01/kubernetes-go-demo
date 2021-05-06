package redis

import (
	"github.com/gomodule/redigo/redis"
	"kubernetes-go-demo/config"
	"kubernetes-go-demo/global/log"
	"reflect"
	"strconv"
	"time"
	"encoding/json"
	"fmt"
)

var (
	redisPool *redis.Pool
)

// 初始化redis数据库
func InitRedis(conf *config.RedisConfig) {
	redisPool = &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", conf.Host)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return conn, err
		},
	}
}

func Serialization(value interface{}) ([]byte, error) {
	if bytes, ok := value.([]byte); ok {
		return bytes, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []byte(strconv.FormatUint(v.Uint(), 10)), nil
	case reflect.Map:
	}
	k, err := json.Marshal(value)
	return k, err
}

func Deserialization(byt []byte, ptr interface{}) (err error) {
	if bytes, ok := ptr.(*[]byte); ok {
		*bytes = byt
		return
	}
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		switch p := v.Elem(); p.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			i, err = strconv.ParseInt(string(byt), 10, 64)
			if err != nil {
				fmt.Printf("Deserialization: failed to parse int '%s': %s", string(byt), err)
			} else {
				p.SetInt(i)
			}
			return

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			i, err = strconv.ParseUint(string(byt), 10, 64)
			if err != nil {
				fmt.Printf("Deserialization: failed to parse uint '%s': %s", string(byt), err)
			} else {
				p.SetUint(i)
			}
			return
		}
	}
	err = json.Unmarshal(byt, &ptr)
	return
}

// string 类型 添加, v 可以是任意类型
func StringSet(name string, v interface{}) error {
	conn := redisPool.Get()

	defer conn.Close()
	_, err := conn.Do("SET", name, v)
	return err
}

// 获取 字符串类型的值
func StringGet(name string) (interface{},error) {
	conn := redisPool.Get()
	defer conn.Close()
	temp, err := conn.Do("Get", name)
	return temp,err
}

func CloseRedis(){
	if err := redisPool.Close();err != nil{
		log.Error("Close redis error by",err)
	}
}
