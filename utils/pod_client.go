package utils

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"time"
)

type PodClient struct {
	Name string
	podInterface corev1.PodInterface
}

func (client *PodClient) InitClient(clientset *kubernetes.Clientset) {
	pod := clientset.CoreV1().Pods(client.Name)
	client.podInterface = pod
}

func (client *PodClient) GetPodList() (*v1.PodList,error){
	podList,err := client.podInterface.List(metav1.ListOptions{})
	return podList,err
}

func (client *PodClient) WatchPod() cache.Store {
	podStore, podController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return client.podInterface.List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return client.podInterface.Watch(lo)
			},
		},
		&v1.Pod{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)
	go podController.Run(wait.NeverStop)
	return podStore
}

//监听Pod变化
func (client *PodClient) StartWatchPod() {
	w, _ := client.podInterface.Watch(metav1.ListOptions{})
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
				pod := e.Object.(*v1.Pod)
				fmt.Println(pod.Name)
			}
		}
	}
}