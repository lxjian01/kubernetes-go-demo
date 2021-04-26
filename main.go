package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	utilsk8s "kubernetes-go-demo/utils/k8s"
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

	// service
	serviceClient := utilsk8s.ServiceClient{Name: "default"}
	serviceClient.InitServiceClient(clientset)
	serviceList,err := serviceClient.GetServiceList(metav1.ListOptions{})
	for _,item := range serviceList.Items{
		fmt.Println(item.Name)
	}

	// deployment
	deploymentClient := utilsk8s.DeploymentClient{Name: "default"}
	deploymentClient.InitDeploymentClient(clientset)
	deploymentList,err := deploymentClient.GetDeploymentList(metav1.ListOptions{})
	for _,item := range deploymentList.Items{
		fmt.Println(item.Name)
	}

	// pod
	podClient := utilsk8s.PodClient{Name: "default"}
	podClient.InitPodClient(clientset)
	podList,err := podClient.GetPodList(metav1.ListOptions{})
	for _,item := range podList.Items{
		fmt.Println(item.Name)
	}
}





