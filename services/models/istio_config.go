package models

// HTTP status code 200 and IstioConfigList model in data
// swagger:response istioConfigList
type swaggIstioConfigList struct {
	// in:body
	Body struct {
		// HTTP status code
		// default: 200
		Code int `json:"code"`
		// IstioConfigList model
		Data IstioConfigList `json:"data"`
	}
}

// IstioConfigList istioConfigList
//
// This is used for returning a response of IstioConfigList
//
// swagger:model IstioConfigList
type IstioConfigList struct {
	// The namespace of istioConfiglist
	//
	// required: true
	Namespace           Namespace           `json:"namespace"`
	RouteRules          RouteRules          `json:"routeRules"`
	DestinationPolicies DestinationPolicies `json:"destinationPolicies"`
	VirtualServices     VirtualServices     `json:"virtualServices"`
	DestinationRules    DestinationRules    `json:"destinationRules"`
	Rules               IstioRules          `json:"rules"`
}

type IstioConfigDetails struct {
	Namespace         Namespace          `json:"namespace"`
	ObjectType        string             `json:"objectType"`
	RouteRule         *RouteRule         `json:"routeRule"`
	DestinationPolicy *DestinationPolicy `json:"destinationPolicy"`
	VirtualService    *VirtualService    `json:"virtualService"`
	DestinationRule   *DestinationRule   `json:"destinationRule"`
	Rule              *IstioRuleDetails  `json:"rule"`
}
