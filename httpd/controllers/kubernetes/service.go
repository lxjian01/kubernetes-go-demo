package kubernetes

import (
	"flag"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kubernetes-go-demo/global/log"
	"kubernetes-go-demo/httpd/utils"
	"kubernetes-go-demo/httpd/utils/k8s"
	"path/filepath"
)

func GetServiceList(c *gin.Context){
	var resp utils.Response

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	kubernetesConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(kubernetesConfig)
	if err != nil {
		panic(err)
	}

	serviceClient := k8s.ServiceClient{Name: "default"}
	serviceClient.InitServiceClient(clientset)
	ttt := make([]string,0)
	serviceList,err := serviceClient.GetServiceList(metav1.ListOptions{})
	for _,item := range serviceList.Items{
		ttt = append(ttt, item.Name)
		log.Infof("service name is %s \n",item.Name)
	}
	resp.Data = ttt
	resp.ToSuccess(c)
}
