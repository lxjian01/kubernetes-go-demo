package k8s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/cache"
	"kubernetes-go-demo/log"
	"time"
)

type DeploymentClient struct {
	Name string
	deploymentInterface appsv1.DeploymentInterface
}

func (client *DeploymentClient) InitDeploymentClient(clientset *kubernetes.Clientset) {
	deploymentInterface := clientset.AppsV1().Deployments(client.Name)
	client.deploymentInterface = deploymentInterface
}

func (client *DeploymentClient) CreateDeployment(yamlFile string) (*v1.Deployment,error){
	deploymentBytes,err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Errorf("Read deployment file error by %v \n", err)
		return nil, err
	}
	deployment := &v1.Deployment{}
	deploymentJson,err := yaml.ToJSON(deploymentBytes)
	if err != nil {
		log.Errorf("Deployment bytes to json error by %v \n", err)
		return nil, err
	}
	err = json.Unmarshal(deploymentJson,deployment)
	if err != nil {
		log.Errorf("Unmarshal deployment error by %v \n", err)
		return nil, err
	}
	log.Infof("Starting create deployment %s \n", deployment.Name)
	deploymentInfo,err := client.deploymentInterface.Create(deployment)
	return deploymentInfo,err
}

func (client *DeploymentClient) UpdateDeployment(deployment *v1.Deployment) (*v1.Deployment,error){
	deploymentInfo,err := client.deploymentInterface.Update(deployment)
	return deploymentInfo,err
}

func (client *DeploymentClient) DeleteDeployment(deploymentName string,deployment *metav1.DeleteOptions) error{
	err := client.deploymentInterface.Delete(deploymentName,deployment)
	return err
}

func (client *DeploymentClient) GetDeploymentList(opts metav1.ListOptions) (*v1.DeploymentList,error){
	deploymentList,err := client.deploymentInterface.List(opts)
	return deploymentList,err
}

//监听Deployment变化
func (client *DeploymentClient) WatchDeployment() {
	w, _ := client.deploymentInterface.Watch(metav1.ListOptions{})
	for {
		select {
		case e, ok := <-w.ResultChan():
			if !ok {
				// 说明该通道已经被close掉了
				fmt.Println("deployment watch chan has been close!!!!")
				fmt.Println("clean chan over!")
				time.Sleep(time.Second * 5)
			}
			if e.Object != nil {
				fmt.Println("chan is ok")
				fmt.Println(e.Type)
				deployment := e.Object.(*v1.Deployment)
				fmt.Println(deployment.Name)
			}
		}
	}
}

func (client *DeploymentClient) CacheWatchDeployment() cache.Store {
	deploymentStore, deploymentController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return client.deploymentInterface.List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return client.deploymentInterface.Watch(lo)
			},
		},
		&v1.Deployment{},
		time.Second * 10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				deployment := obj.(*v1.Deployment)
				log.Infof("deployment add: deploymentName=%s, time is %s \n", deployment.Name, time.Now())
			},
			UpdateFunc:func(oldObj, newObj interface{}) {
				oldDeployment := oldObj.(*v1.Deployment)
				newDeployment := newObj.(*v1.Deployment)
				log.Infof("deployment update: oldDeploymentName=%s, newDeploymentName=%s, time is %s \n", oldDeployment.Name, newDeployment.Name,time.Now())
			},
			DeleteFunc: func(obj interface{}) {
				deployment := obj.(*v1.Deployment)
				log.Infof("deployment delete: deploymentName=%s, time is %s \n", deployment.Name, time.Now())
			},
		},
	)
	go deploymentController.Run(wait.NeverStop)
	return deploymentStore
}