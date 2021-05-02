package controllers

import (
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gin-gonic/gin"
	"kubernetes-go-demo/global/log"
	"kubernetes-go-demo/global/machinery"
	"time"
)

func DelayedTask(c *gin.Context){
	signature := &tasks.Signature{
		Name: "multiply",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 2,
			},
			{
				Type:  "int64",
				Value: 3,
			},
		},
	}
	eta := time.Now().UTC().Add(time.Second * 5)
	signature.ETA = &eta
	asyncResult, err := machinery.GetServer().SendTask(signature)
	if err != nil {
		log.Error("Machinery send delayed task add error by ", err)
		c.JSON(400, gin.H{"add": err, "result": asyncResult})
		return
	}
	c.JSON(200, gin.H{"add": err, "result": asyncResult})
}

func SendTask(c *gin.Context){
	signature := &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	asyncResult, err := machinery.GetServer().SendTask(signature)
	if err != nil {
		log.Error("Machinery send task add error by ", err)
		c.JSON(400, gin.H{"add": err, "result": asyncResult})
		return
	}
	c.JSON(200, gin.H{"add": err, "result": asyncResult})
}

func PeriodicTask(c *gin.Context){
	signature := &tasks.Signature{
		UUID: "22222222222",
		Name: "multiply",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	err := machinery.GetServer().RegisterPeriodicTask("*/1 * * * ?", "periodic-task", signature)
	if err != nil {
		log.Error("Machinery send task add error by ", err)
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"add": err})
}

func GetTaskList(c *gin.Context){
	signature := &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	asyncResult, err := machinery.GetServer().SendTask(signature)
	if err != nil {
		log.Error("Machinery send task add error by ", err)
		c.JSON(400, gin.H{"add": err, "result": asyncResult})
		return
	}
	c.JSON(200, gin.H{"add": err, "result": asyncResult})
}
