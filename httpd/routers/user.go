package routers

import (
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/httpd/controllers"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/list", controllers.GetUserList)
		//user.POST("/test", controllers.Test)
	}
}