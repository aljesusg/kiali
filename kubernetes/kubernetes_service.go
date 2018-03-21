package kubernetes

import (
	"strings"

	"github.com/kiali/swscore/config"

	"k8s.io/api/apps/v1beta1"
	autoscalingV1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DeploymentFilterLabelName = "app"
)

type serviceResponse struct {
	service *v1.Service
	err     error
}

type endpointsResponse struct {
	endpoints *v1.Endpoints
	err       error
}

type deploymentsResponse struct {
	deployments *v1beta1.DeploymentList
	err         error
}

type autoscalersResponse struct {
	autoscalers *autoscalingV1.HorizontalPodAutoscalerList
	err         error
}

type podResponse struct {
	pod *v1.Pod
	err error
}

// GetNamespaces returns a list of all namespaces of the cluster.
// It returns a list of all namespaces of the cluster.
// It returns an error on any problem.
func (in *IstioClient) GetNamespaces() (*v1.NamespaceList, error) {
	namespaces, err := in.k8s.CoreV1().Namespaces().List(emptyListOptions)
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

// GetServices returns a list of services for a given namespace.
// It returns an error on any problem.
func (in *IstioClient) GetServices(namespaceName string) (*v1.ServiceList, error) {
	services, err := in.k8s.CoreV1().Services(namespaceName).List(emptyListOptions)
	if err != nil {
		return nil, err
	}

	return services, nil
}

// GetServiceDetails returns full details for a given service, consisting on service description, endpoints and pods.
// A service is defined by the namespace and the service name.
// It returns an error on any problem.
func (in *IstioClient) GetServiceDetails(namespace string, serviceName string) (*ServiceDetails, error) {
	serviceChan := make(chan serviceResponse)
	endpointsChan := make(chan endpointsResponse)
	deploymentsChan := make(chan deploymentsResponse)
	autoscalersChan := make(chan autoscalersResponse)

	go func() {
		service, err := in.k8s.CoreV1().Services(namespace).Get(serviceName, emptyGetOptions)
		serviceChan <- serviceResponse{service: service, err: err}
	}()

	go func() {
		endpoints, err := in.k8s.CoreV1().Endpoints(namespace).Get(serviceName, emptyGetOptions)
		endpointsChan <- endpointsResponse{endpoints: endpoints, err: err}
	}()

	go func() {
		deployments, err := in.k8s.AppsV1beta1().Deployments(namespace).List(*getDeploymentFilterListOptions(serviceName))
		deploymentsChan <- deploymentsResponse{deployments: deployments, err: err}
	}()

	go func() {
		autoscalers, err := in.k8s.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(emptyListOptions)
		autoscalersChan <- autoscalersResponse{autoscalers: autoscalers, err: err}
	}()

	serviceDetails := ServiceDetails{}

	serviceResponse := <-serviceChan
	if serviceResponse.err != nil {
		return nil, serviceResponse.err
	}
	serviceDetails.Service = serviceResponse.service

	endpointsResponse := <-endpointsChan
	if endpointsResponse.err != nil {
		return nil, endpointsResponse.err
	}
	serviceDetails.Endpoints = endpointsResponse.endpoints

	deploymentsResponse := <-deploymentsChan
	if deploymentsResponse.err != nil {
		return nil, deploymentsResponse.err
	}
	serviceDetails.Deployments = deploymentsResponse.deployments

	autoscalersResponse := <-autoscalersChan
	if autoscalersResponse.err != nil {
		return nil, autoscalersResponse.err
	}
	serviceDetails.Autoscalers = autoscalersResponse.autoscalers

	return &serviceDetails, nil
}

func filterAutoscalersByDeployments(deploymentNames []string, al *autoscalingV1.HorizontalPodAutoscalerList) *autoscalingV1.HorizontalPodAutoscalerList {
	autoscalers := make([]autoscalingV1.HorizontalPodAutoscaler, 0, len(al.Items))

	for _, autoscaler := range al.Items {
		for _, deploymentName := range deploymentNames {
			if deploymentName == autoscaler.Spec.ScaleTargetRef.Name {
				autoscalers = append(autoscalers, autoscaler)
			}
		}
	}

	al.Items = autoscalers
	return al
}

func getDeploymentNames(deployments *v1beta1.DeploymentList) []string {
	deploymentNames := make([]string, len(deployments.Items))
	for _, deployment := range deployments.Items {
		deploymentNames = append(deploymentNames, deployment.Name)
	}

	return deploymentNames
}

func getDeploymentFilterListOptions(serviceName string) *meta_v1.ListOptions {
	filterLabelName := config.Get().ServiceFilterLabelName
	return GetLabeledListOptions(strings.Join([]string{filterLabelName, serviceName}, "="))
}
