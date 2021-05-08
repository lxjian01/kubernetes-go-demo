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
	pool *redis.Pool
)

// 初始化redis数据库
func InitRedis(conf *config.RedisConfig) {
	pool = &redis.Pool{
		MaxIdle:     conf.MaxIdle, //  最大连接数，即最多的tcp连接数，一般建议往大的配置，但不要超过操作系统文件句柄个数（centos下可以ulimit -n查看）
		MaxActive:   conf.MaxActive, // 最大空闲连接数，即会有这么多个连接提前等待着，但过了超时时间也会关闭。
		IdleTimeout: 60 * time.Second, // 空闲连接超时时间，但应该设置比redis服务器超时时间短。否则服务端超时了，客户端保持着连接也没用
		Wait:        true,             // 当超过最大连接数 是报错还是等待， true 等待 false 报错
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", conf.Host, redis.DialDatabase(conf.Db))
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
	conn := pool.Get()
	defer conn.Close()
	s, _ := Serialization(v) // 序列化
	_, err := conn.Do("SET", name, s)
	return err
}

// 获取 字符串类型的值
func StringGet(name string, v interface{}) error {
	conn := pool.Get()
	defer conn.Close()
	temp, _ := redis.Bytes(conn.Do("Get", name))
	err := Deserialization(temp, &v) // 反序列化
	return err
}

// 判断所在的 key 是否存在
func Exist(name string) (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.Bool(conn.Do("EXISTS", name))
	return v, err
}

// 自增
func StringIncr(name string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.Int(conn.Do("INCR", name))
	return v, err
}

// 设置过期时间 （单位 秒）
func Expire(name string, newSecondsLifeTime int64) error {
	// 设置key 的过期时间
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", name, newSecondsLifeTime)
	return err
}

// 删除指定的键
func Delete(keys ...interface{}) (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.Bool(conn.Do("DEL", keys...))
	return v, err
}

// 查看指定的长度
func StrLen(name string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.Int(conn.Do("STRLEN", name))
	return v, err
}

// //  hash ///
// 删除指定的 hash 键
func Hdel(name, key string) (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	var err error
	v, err := redis.Bool(conn.Do("HDEL", name, key))
	return v, err
}

// 查看hash 中指定是否存在
func HExists(name, field string) (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	var err error
	v, err := redis.Bool(conn.Do("HEXISTS", name, field))
	return v, err
}

// 获取hash 的键的个数
func HLen(name string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.Int(conn.Do("HLEN", name))
	return v, err
}

// 传入的 字段列表获得对应的值
func HMget(name string, fields ...string) ([]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	args := []interface{}{name}
	for _, field := range fields {
		args = append(args, field)
	}
	value, err := redis.Values(conn.Do("HMGET", args...))

	return value, err
}

// 设置单个值, value 还可以是一个 map slice 等
func HSet(name string, key string, value interface{}) (err error) {
	conn := pool.Get()
	defer conn.Close()
	v, _ := Serialization(value)
	_, err = conn.Do("HSET", name, key, v)
	return
}

// 设置多个值 , obj 可以是指针 slice map struct
func HMSet(name string, obj interface{}) (err error) {
	conn := pool.Get()
	defer conn.Close()
	_, err = conn.Do("HSET", redis.Args{}.Add(name).AddFlat(&obj)...)
	return
}

// 获取单个hash 中的值
func HGet(name, field string, v interface{}) (err error) {
	conn := pool.Get()
	defer conn.Close()
	temp, _ := redis.Bytes(conn.Do("Get", name, field))
	err = Deserialization(temp, &v) // 反序列化
	return
}

// set 集合

// 获取 set 集合中所有的元素, 想要什么类型的自己指定
func Smembers(name string, v interface{}) (err error) {
	conn := pool.Get()
	defer conn.Close()
	temp, _ := redis.Bytes(conn.Do("smembers", name))
	err = Deserialization(temp, &v)
	return err
}

// 获取集合中元素的个数
func ScardInt64s(name string) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	v, err := redis.Int64(conn.Do("SCARD", name))
	return v, err
}


func CloseRedis(){
	if err := pool.Close();err != nil{
		log.Error("Close redis error by",err)
	}
}
