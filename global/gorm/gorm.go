package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kubernetes-go-demo/config"
)

var (
	ormDB *gorm.DB
)

// 初始化
func InitDB() {
	conf := config.GetConfig().Mysql
	// dsn := "user:pass@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	ormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sqlDb, err := ormDB.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDb.Close()
	sqlDb.SetMaxIdleConns(conf.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(conf.MaxOpen) //设置最大的空闲连接数
}

func GetOrmDB() *gorm.DB {
	conf := config.GetConfig()
	if conf.Env != "prod" {
		return ormDB.Debug()
	}else{
		return ormDB
	}
}


