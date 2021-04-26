package k8s

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
	"kubernetes-go-demo/log"
	"time"
)

type PodClient struct {
	Name string
	podInterface corev1.PodInterface
}

func (client *PodClient) InitPodClient(clientset *kubernetes.Clientset) {
	podInterface := clientset.CoreV1().Pods(client.Name)
	client.podInterface = podInterface
}

func (client *PodClient) CreatePod(pod *v1.Pod) (*v1.Pod,error){
	podInfo,err := client.podInterface.Create(pod)
	return podInfo,err
}

func (client *PodClient) UpdatePod(pod *v1.Pod) (*v1.Pod,error){
	podInfo,err := client.podInterface.Update(pod)
	return podInfo,err
}

func (client *PodClient) DeletePod(podName string,options *metav1.DeleteOptions) error{
	err := client.podInterface.Delete(podName,options)
	return err
}

func (client *PodClient) GetPod(podName string, options metav1.GetOptions) (*v1.Pod,error){
	podInfo,err := client.podInterface.Get(podName,options)
	return podInfo,err
}

func (client *PodClient) GetPodList(opts metav1.ListOptions) (*v1.PodList,error){
	podList,err := client.podInterface.List(opts)
	return podList,err
}

//监听Pod变化
func (client *PodClient) WatchPod() {
	w, _ := client.podInterface.Watch(metav1.ListOptions{})
	for {
		select {
		case e, ok := <-w.ResultChan():
			if !ok {
				// 说明该通道已经被close掉了
				fmt.Println("pod watch chan has been close!!!!")
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

func (client *PodClient) CacheWatchPod() cache.Store {
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
		time.Second * 10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				pod := obj.(*v1.Pod)
				log.Infof("pod added: podName=%s, time is %s \n", pod.Name, time.Now())
			},
			UpdateFunc:func(oldObj, newObj interface{}) {
				oldPod := oldObj.(*v1.Pod)
				newPod := newObj.(*v1.Pod)
				log.Infof("pod changed: oldPodName=%s, newPodName=%s, time is %s \n", oldPod.Name, newPod.Name,time.Now())
			},
			DeleteFunc: func(obj interface{}) {
				pod := obj.(*v1.Pod)
				log.Infof("pod deleted: podName=%s, time is %s \n", pod.Name, time.Now())
			},
		},
	)
	go podController.Run(wait.NeverStop)
	return podStore
}