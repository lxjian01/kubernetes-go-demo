package main

import (
	"fmt"
	"kubernetes-go-demo/config"
	"kubernetes-go-demo/global/gorm"
	"kubernetes-go-demo/global/log"
	"kubernetes-go-demo/global/pools"
	"kubernetes-go-demo/global/redis"
	"kubernetes-go-demo/httpd"
	"os"
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
	log.Info("Init config ok")
	// init db
	gorm.InitDB()
	log.Info("Init db ok")
	// init redis pool
	redis.InitRedis()
	defer redis.CloseRedis()
	log.Info("Init redis ok")
	// init goroutine pool
	pools.InitPool()
	defer pools.ClosePool()
	log.Info("Init goroutine pool ok")

	// init gin server
	httpd.StartHttpdServer()
	log.Info("Start gin server ok")


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
	//serviceClient := utilsk8s.ServiceClient{Name: "default"}
	//serviceClient.InitServiceClient(clientset)
	//serviceList,err := serviceClient.GetServiceList(metav1.ListOptions{})
	//for _,item := range serviceList.Items{
	//	log.Infof("service name is %s \n",item.Name)
	//}
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





