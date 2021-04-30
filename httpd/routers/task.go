package routers

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers"
)

func TaskRoutes(route *gin.Engine) {
	task := route.Group("/task")
	{
		task.GET("/send_task", controllers.SendTask)
		task.GET("/delayed_task", controllers.DelayedTask)
		task.GET("/list", controllers.GetTaskList)
		//user.POST("/test", controllers.Test)
	}
}