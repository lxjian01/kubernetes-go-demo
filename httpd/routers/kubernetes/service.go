package kubernetes

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers/kubernetes"
)

func KubernetesServiceRoutes(route *gin.Engine) {
	user := route.Group("/kubernetes/service")
	{
		user.GET("/list", kubernetes.GetServiceList)
	}
}
