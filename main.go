package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
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

	tt := WatchResources(podsClient)
	a := tt.ListKeys()
	fmt.Println(a)
	for _,item := range a{
		fmt.Println("1111111111111")
		fmt.Println(item)
		fmt.Println("22222222222")
	}
	//startWatchPod(podsClient)
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
				yyyy := e.Object.(*v1.Pod)
				fmt.Println(yyyy.Name)
			}
		}
	}
}

func WatchResources(podsClient corev1.PodInterface) cache.Store {
	podStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return podsClient.List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return podsClient.Watch(lo)
			},
		},
		&v1.Pod{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return podStore
}
