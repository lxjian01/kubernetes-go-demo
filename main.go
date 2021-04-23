package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"reflect"
	"time"
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
	podsClient := clientset.CoreV1().Pods("default")
	//startWatchDeployment(deploymentsClient)
	pods,err := podsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _,item := range pods.Items{
		fmt.Println(item.Name)
	}
	startWatchPod(podsClient)
	//deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
}

//监听Pod变化
func startWatchPod(podsClient corev1.PodInterface) {
	w, _ := podsClient.Watch(metav1.ListOptions{})

	for {
		select {
		case e, ok := <-w.ResultChan():
			if !ok {
				// 说明该通道已经被close掉了
				fmt.Println("!!!!!podWatch chan has been close!!!!")
				fmt.Println("clean chan over!")
				time.Sleep(time.Second * 5)
			}
			if e.Object != nil {
				fmt.Println("chan is ok")
				fmt.Println(e.Type)

				v := reflect.ValueOf(e.Object)
				for i := 0; i < v.NumField(); i++ {
					switch v.Field(i).Kind() {
					case reflect.Int:
						if i == 0 {
							fmt.Sprintf("%d", v.Field(i).Int())
						} else {
							fmt.Sprintf("%d", v.Field(i).Int())
						}
					case reflect.String:
						if i == 0 {
							fmt.Sprintf("%s", v.Field(i).String())
						} else {
							fmt.Sprintf("%s", v.Field(i).String())
						}
					default:
						fmt.Println("Unsupported type")
						return
					}
				}
			}
		}
	}
}

//监听Deployment变化
func startWatchDeployment(deploymentsClient appsv1.DeploymentInterface) {
	w, _ := deploymentsClient.Watch(metav1.ListOptions{})
	for {
		select {
		case e, _ := <-w.ResultChan():
			fmt.Println(e.Type, e.Object)
		}
	}
}
