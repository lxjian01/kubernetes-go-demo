package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kubernetes-go-demo/config"
	"kubernetes-go-demo/global/gorm"
	"kubernetes-go-demo/global/log"
	"kubernetes-go-demo/global/pools"
	"kubernetes-go-demo/global/redis"
	"kubernetes-go-demo/httpd"
	"os"
	"path/filepath"
)

var (
	daemon bool
	appConf config.AppConfig
	lockfile = "/var/run/jumpserver.pid"
)

//Execute方法触发init方法
func init() {
	//初始化配置文件转化成对应的结构体
	cobra.OnInitialize(initConfig)
	httpdCmd.AddCommand(versionCmd)
}

//项目启动调用的入口方法
func HttpdCmdExecute() error{
	//初始化Cobra
	err := httpdCmd.Execute()
	return err
}

var httpdCmd = &cobra.Command{
	Use:   "httpd",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Start http server error by ", r)
				os.Exit(1)
			}
		}()
		fmt.Println("Starting init system log")
		log.Init(appConf.Log)
		fmt.Println("Init system log ok")

		log.Info("Starting init pool")
		pools.InitPool(appConf.PoolNum)
		log.Info("Init pool ok")

		log.Info("Starting init mysql")
		gorm.InitDB(appConf.Mysql)
		log.Info("Init mysql ok")

		log.Info("Starting init redis")
		redis.InitRedis(appConf.Redis)
		defer redis.CloseRedis()
		log.Info("Init redis ok")

		// init gin server
		log.Info("Starting init gin server")
		httpd.StartHttpdServer(appConf.Httpd)
		log.Info("Start gin server ok")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show jumpserver version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jumpserver version is",appConf.Version)
		os.Exit(0)
	},
}

//通过viper初始化配置文件到结构体
func initConfig() {
	dir,_ := os.Getwd()
	env := os.Getenv("ENV")
	if env == ""{
		env = "dev"
	}
	configPath := filepath.Join(dir,"config/"+env)
	// 设置读取的文件路径
	viper.AddConfigPath(configPath)
	// 设置读取的文件名
	viper.SetConfigName("config")
	// 设置文件的类型
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Read config error by %v \n",err))
	}
	if err :=viper.Unmarshal(&appConf); err !=nil{
		panic(fmt.Sprintf("Unmarshal config error by %v \n",err))
	}
}