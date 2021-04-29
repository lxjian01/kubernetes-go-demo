package kubernetes

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers/kubernetes"
)

func ServiceRoutes(route *gin.Engine) {
	service := route.Group("/kubernetes/service")
	{
		service.GET("/list", kubernetes.GetServiceList)
	}
}
