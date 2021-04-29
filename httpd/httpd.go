package httpd

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/config"
	"kubernetes-go-demo/global/log"
	"kubernetes-go-demo/httpd/middlewares"
	"kubernetes-go-demo/httpd/routers"
	"kubernetes-go-demo/httpd/routers/kubernetes"
	"net"
	"strconv"
)

func StartHttpdServer(conf *config.HttpdConfig) {
	router := gin.Default()
	// 添加自定义的 logger 间件
	router.Use(middlewares.Logger(), gin.Recovery())
	router.Use(middlewares.Auth(), gin.Recovery())
	// 添加路由
	routers.UserRoutes(router)      //Added all user routers
	kubernetes.KubernetesServiceRoutes(router)      //Added all user routers
	// 拼接host
	Host := conf.Host
	Port := strconv.Itoa(conf.Port)
	addr := net.JoinHostPort(Host, Port)
	log.Info("Start HTTP server at", addr)
	err := router.Run(addr)
	if err != nil{
		log.Error("Start server error by",err)
		panic(err)
	}
	log.Info("Start server ok")
}
