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

type ServiceClient struct {
	Name string
	serviceInterface corev1.ServiceInterface
}

func (client *ServiceClient) InitServiceClient(clientset *kubernetes.Clientset) {
	serviceInterface := clientset.CoreV1().Services(client.Name)
	client.serviceInterface = serviceInterface
}

func (client *ServiceClient) CreateService(service *v1.Service) (*v1.Service,error){
	serviceInfo,err := client.serviceInterface.Create(service)
	return serviceInfo,err
}

func (client *ServiceClient) UpdateService(service *v1.Service) (*v1.Service,error){
	serviceInfo,err := client.serviceInterface.Update(service)
	return serviceInfo,err
}

func (client *ServiceClient) DeleteService(serviceName string,options *metav1.DeleteOptions) error{
	err := client.serviceInterface.Delete(serviceName,options)
	return err
}

func (client *ServiceClient) GetService(serviceName string, options metav1.GetOptions) (*v1.Service,error){
	serviceInfo,err := client.serviceInterface.Get(serviceName,options)
	return serviceInfo,err
}

func (client *ServiceClient) GetServiceList(opts metav1.ListOptions) (*v1.ServiceList,error){
	serviceList,err := client.serviceInterface.List(opts)
	return serviceList,err
}

//监听Deployment变化
func (client *ServiceClient) WatchDeployment() {
	w, _ := client.serviceInterface.Watch(metav1.ListOptions{})
	for {
		select {
		case e, ok := <-w.ResultChan():
			if !ok {
				// 说明该通道已经被close掉了
				fmt.Println("service watch chan has been close!!!!")
				fmt.Println("clean chan over!")
				time.Sleep(time.Second * 5)
			}
			if e.Object != nil {
				fmt.Println("chan is ok")
				fmt.Println(e.Type)
				deployment := e.Object.(*v1.Service)
				fmt.Println(deployment.Name)
			}
		}
	}
}

func (client *ServiceClient) CacheWatchService() cache.Store {
	serviceStore, serviceController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return client.serviceInterface.List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return client.serviceInterface.Watch(lo)
			},
		},
		&v1.Service{},
		time.Second * 10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				service := obj.(*v1.Service)
				log.Infof("service add: serviceName=%s, time is %s \n", service.Name, time.Now())
			},
			UpdateFunc:func(oldObj, newObj interface{}) {
				oldService := oldObj.(*v1.Service)
				newService := newObj.(*v1.Service)
				log.Infof("service update: oldServiceName=%s, newServiceName=%s, time is %s \n", oldService.Name, newService.Name,time.Now())
			},
			DeleteFunc: func(obj interface{}) {
				service := obj.(*v1.Service)
				log.Infof("service delete: serviceName=%s, time is %s \n", service.Name, time.Now())
			},
		},
	)
	go serviceController.Run(wait.NeverStop)
	return serviceStore
}