package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kubernetes-go-demo/utils"
	"path/filepath"
)


func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deploymentsClient := clientset.AppsV1().Deployments("nginx-deployment")
	fmt.Println(deploymentsClient)

	// pod
	podClient := utils.PodClient{Name: "default"}
	podClient.InitPodClient(clientset)
	podList,err := podClient.GetPodList(metav1.ListOptions{})
	for _,item := range podList.Items{
		fmt.Println(item.Name)
	}
}





