package kubernetes

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers/kubernetes"
)

func DeploymentRoutes(route *gin.Engine) {
	deployment := route.Group("/kubernetes")
	{
		deployment.GET("/deployment/list", kubernetes.GetDeploymentList)
		deployment.POST("/deployment", kubernetes.CreateDeploymentList)
		deployment.PUT("/deployment", kubernetes.UpdateDeploymentList)
	}
}
