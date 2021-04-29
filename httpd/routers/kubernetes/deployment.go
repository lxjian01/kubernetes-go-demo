package kubernetes

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers/kubernetes"
)

func DeploymentRoutes(route *gin.Engine) {
	deployment := route.Group("/kubernetes/deployment")
	{
		deployment.GET("/list", kubernetes.GetDeploymentList)
	}
}
