package utils

import (
	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
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