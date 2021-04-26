package k8s

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
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

func (client *DeploymentClient) CreateDeployment(deployment *v1.Deployment) (*v1.Deployment,error){
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