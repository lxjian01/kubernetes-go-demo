package kubernetes

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s2 "kubernetes-go-demo/global/k8s"
	"kubernetes-go-demo/httpd/utils"
	"kubernetes-go-demo/httpd/utils/k8s"
)

func GetServiceList(c *gin.Context){
	var resp utils.Response
	clientset := k8s2.GetClientset()
	serviceClient := k8s.ServiceClient{Name: "default"}
	serviceClient.InitServiceClient(clientset)
	serviceList,err := serviceClient.GetServiceList(metav1.ListOptions{})
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	serviceNameList := make([]string,0)
	for _,item := range serviceList.Items{
		serviceNameList = append(serviceNameList, item.Name)
	}
	resp.Data = serviceNameList
	resp.ToSuccess(c)
}
