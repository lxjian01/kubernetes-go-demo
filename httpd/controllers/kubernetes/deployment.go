package kubernetes

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-go-demo/global/config"
	"kubernetes-go-demo/httpd/utils"
	"kubernetes-go-demo/httpd/utils/kubeutil"
	"path/filepath"
)

func CreateDeploymentList(c *gin.Context){
	var resp utils.Response
	deploymentClient := kubeutil.NewDeploymentClient("default")
	yamlFile := filepath.Join(config.GetAppConfig().YamlDir,"deployments/nginx-deployment.yaml")
	deployment, err:= deploymentClient.CreateDeployment(yamlFile)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = deployment
	resp.ToSuccess(c)
}

func UpdateDeploymentList(c *gin.Context){
	var resp utils.Response
	deploymentClient := kubeutil.NewDeploymentClient("default")
	yamlFile := filepath.Join(config.GetAppConfig().YamlDir,"deployments/nginx-deployment.yaml")
	deployment, err:= deploymentClient.UpdateDeployment(yamlFile)
	opt := metav1.ListOptions{}
	deploymentClient.WatchDeployment(opt)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = deployment
	resp.ToSuccess(c)
}


func GetDeploymentList(c *gin.Context){
	var resp utils.Response
	deploymentClient := kubeutil.NewDeploymentClient("default")
	deploymentList, err := deploymentClient.GetDeploymentList(metav1.ListOptions{})
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	deploymentNameList := make([]string,0)
	for _,item := range deploymentList.Items{
		deploymentNameList = append(deploymentNameList, item.Name)
	}
	resp.Data = deploymentNameList
	resp.ToSuccess(c)
}