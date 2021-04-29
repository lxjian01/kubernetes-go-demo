package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context){
	c.String(http.StatusOK, "Hello %s", "name")
}
