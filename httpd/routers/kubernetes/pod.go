package kubernetes

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers/kubernetes"
)

func PodRoutes(route *gin.Engine) {
	pod := route.Group("/kubernetes/pod")
	{
		pod.GET("/list", kubernetes.GetPodList)
	}
}
