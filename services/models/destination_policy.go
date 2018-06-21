package models

import (
	"github.com/kiali/kiali/kubernetes"
)

type DestinationPolicies []DestinationPolicy
type DestinationPolicy struct {
	Name            string      `json:"name"`
	CreatedAt       string      `json:"createdAt"`
	ResourceVersion string      `json:"resourceVersion"`
	Source          interface{} `json:"source"`
	Destination     interface{} `json:"destination"`
	LoadBalancing   interface{} `json:"loadbalancing"`
	CircuitBreaker  interface{} `json:"circuitBreaker"`
}

func (policies *DestinationPolicies) Parse(destinationPolicies []kubernetes.IstioObject) {
	for _, dp := range destinationPolicies {
		destinationPolicy := DestinationPolicy{}
		destinationPolicy.Parse(dp)
		*policies = append(*policies, destinationPolicy)
	}
}

func (policy *DestinationPolicy) Parse(destinationPolicy kubernetes.IstioObject) {
	policy.Name = destinationPolicy.GetObjectMeta().Name
	policy.CreatedAt = formatTime(destinationPolicy.GetObjectMeta().CreationTimestamp.Time)
	policy.ResourceVersion = destinationPolicy.GetObjectMeta().ResourceVersion
	policy.Source = destinationPolicy.GetSpec()["source"]
	policy.Destination = destinationPolicy.GetSpec()["destination"]
	policy.LoadBalancing = destinationPolicy.GetSpec()["loadBalancing"]
	policy.CircuitBreaker = destinationPolicy.GetSpec()["circuitBreaker"]
}
