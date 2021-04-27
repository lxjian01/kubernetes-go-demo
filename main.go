package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kubernetes-go-demo/config"
	"kubernetes-go-demo/log"
	utilsk8s "kubernetes-go-demo/utils/k8s"
	"os"
	"path/filepath"
	"time"
)


func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	// init config
	config.InitConfig()
	log.Info("111111111111111111111")

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

	// service
	serviceClient := utilsk8s.ServiceClient{Name: "default"}
	serviceClient.InitServiceClient(clientset)
	serviceList,err := serviceClient.GetServiceList(metav1.ListOptions{})
	for _,item := range serviceList.Items{
		log.Infof("service name is %s \n",item.Name)
	}

	// deployment
	deploymentClient := utilsk8s.DeploymentClient{Name: "default"}
	deploymentClient.InitDeploymentClient(clientset)
	yamlFile := filepath.Join(config.GetConfig().YamlDir,"deployments/nginx-deployment.yaml")
	nginxDeployment, err:= deploymentClient.CreateDeployment(yamlFile)
	if err != nil {
		log.Errorf("Create deployment error by %v \n", err)
		return
	}
	log.Infof("Create deployment name is %s \n", nginxDeployment.Name)
	deploymentList, err := deploymentClient.GetDeploymentList(metav1.ListOptions{})
	for _,item := range deploymentList.Items{
		log.Infof("deployment name is %s \n",item.Name)
	}

	// pod
	podClient := utilsk8s.PodClient{Name: "default"}
	podClient.InitPodClient(clientset)
	podList,err := podClient.GetPodList(metav1.ListOptions{})
	for _,item := range podList.Items{
		log.Infof("pod name is %s \n",item.Name)
	}
	podClient.CacheWatchPod()
	for{
		time.Sleep(time.Minute * 6)
	}
}





