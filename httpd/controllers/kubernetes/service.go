package kubernetes

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-go-demo/httpd/utils"
	"kubernetes-go-demo/httpd/utils/k8s"
)

func GetServiceList(c *gin.Context){
	var resp utils.Response
	serviceClient := k8s.NewServiceClient("default")
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
