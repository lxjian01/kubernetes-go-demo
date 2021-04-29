package main

import (
	"fmt"
	"kubernetes-go-demo/cmd"
	"os"
)

func main() {
	// start httpd server
	err := cmd.HttpdCmdExecute()
	if err != nil {
		fmt.Println("Start httpd server error by ",err)
		os.Exit(1)
	}

	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//
	//kubernetesConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err)
	//}
	//clientset, err := kubernetes.NewForConfig(kubernetesConfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// service

	//
	//// deployment
	//deploymentClient := utilsk8s.DeploymentClient{Name: "default"}
	//deploymentClient.InitDeploymentClient(clientset)
	//yamlFile := filepath.Join(config.GetConfig().YamlDir,"deployments/nginx-deployment.yaml")
	//nginxDeployment, err:= deploymentClient.CreateDeployment(yamlFile)
	//if err != nil {
	//	log.Errorf("Create deployment error by %v \n", err)
	//	return
	//}
	//log.Infof("Create deployment name is %s \n", nginxDeployment.Name)
	//deploymentList, err := deploymentClient.GetDeploymentList(metav1.ListOptions{})
	//for _,item := range deploymentList.Items{
	//	log.Infof("deployment name is %s \n",item.Name)
	//}
	//
	//// pod
	//podClient := utilsk8s.PodClient{Name: "default"}
	//podClient.InitPodClient(clientset)
	//podList,err := podClient.GetPodList(metav1.ListOptions{})
	//for _,item := range podList.Items{
	//	log.Infof("pod name is %s \n",item.Name)
	//}
	//podClient.CacheWatchPod()
	//for{
	//	time.Sleep(time.Minute * 6)
	//}
}





