package kubernetes

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-go-demo/global/config"
	k8s2 "kubernetes-go-demo/global/k8s"
	"kubernetes-go-demo/httpd/utils"
	"kubernetes-go-demo/httpd/utils/k8s"
	"path/filepath"
)

func CreateDeploymentList(c *gin.Context){
	var resp utils.Response
	clientset := k8s2.GetClientset()
	deploymentClient := k8s.DeploymentClient{Name: "default"}
	deploymentClient.InitDeploymentClient(clientset)
	yamlFile := filepath.Join(config.GetConfig().YamlDir,"deployments/nginx-deployment.yaml")
	deployment, err:= deploymentClient.CreateDeployment(yamlFile)
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	resp.Data = deployment
	resp.ToSuccess(c)
}

func GetDeploymentList(c *gin.Context){
	var resp utils.Response
	clientset := k8s2.GetClientset()
	deploymentClient := k8s.DeploymentClient{Name: "default"}
	deploymentClient.InitDeploymentClient(clientset)
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