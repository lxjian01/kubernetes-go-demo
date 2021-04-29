package kubernetes

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s2 "kubernetes-go-demo/global/k8s"
	"kubernetes-go-demo/httpd/utils"
	"kubernetes-go-demo/httpd/utils/k8s"
)

func GetPodList(c *gin.Context){
	var resp utils.Response
	clientset := k8s2.GetClientset()
	podClient := k8s.PodClient{Name: "default"}
	podClient.InitPodClient(clientset)
	podList,err := podClient.GetPodList(metav1.ListOptions{})
	if err != nil {
		resp.ToMsgBadRequest(c, err.Error())
		return
	}
	podNameList := make([]string,0)
	for _,item := range podList.Items{
		podNameList = append(podNameList, item.Name)
	}
	resp.Data = podNameList
	resp.ToSuccess(c)
}
