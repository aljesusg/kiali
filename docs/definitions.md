
<a name="definitions"></a>
## Definitions

<a name="istioconfiglist"></a>
### IstioConfigList
This is used for returning a response of IstioConfigList


|Name|Schema|
|---|---|
|**destinationPolicies**  <br>*optional*|[destinationPolicies](#destinationpolicies)|
|**destinationRules**  <br>*optional*|[destinationRules](#destinationrules)|
|**namespace**  <br>*required*|[namespace](#namespace)|
|**routeRules**  <br>*optional*|[routeRules](#routerules)|
|**rules**  <br>*optional*|[istioRules](#istiorules)|
|**virtualServices**  <br>*optional*|[virtualServices](#virtualservices)|


<a name="statusinfo"></a>
### StatusInfo
This is used for returning a response of Kiali Status


|Name|Description|Schema|
|---|---|---|
|**externalServices**  <br>*required*|An array of external services installed|< [externalServiceInfo](#externalserviceinfo) > array|
|**status**  <br>*required*|The state of Kiali<br>A hash of key,values with versions of Kiali and state|< string, string > map|
|**warningMessages**  <br>*optional*|An array of warningMessages|< string > array|


<a name="tokengenerated"></a>
### TokenGenerated
This is used for returning the token


|Name|Description|Schema|
|---|---|---|
|**expired_at**  <br>*required*|The expired time for the token<br>A string with the Datetime when the token will be expired  <br>**Example** : `"2018-06-20 19:40:54.116369887 +0000 UTC m=+43224.838320603"`|string|
|**token**  <br>*required*|The authentication token<br>A string with the authentication token for the user  <br>**Example** : `"zI1NiIsIsR5cCI6IkpXVCJ9.ezJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNTI5NTIzNjU0fQ.PPZvRGnR6VA4v7FmgSfQcGQr-VD"`|string|


<a name="destinationpolicies"></a>
### destinationPolicies
This is used for returning an array of DestinationPolicies

*Type* : < [destinationPolicy](#destinationpolicy) > array


<a name="destinationpolicy"></a>
### destinationPolicy
This is used for returning a DestinationPolicy


|Name|Description|Schema|
|---|---|---|
|**circuitBreaker**  <br>*optional*||object|
|**createdAt**  <br>*required*|The creation date of the destinationPolicy|string|
|**destination**  <br>*optional*||object|
|**loadbalancing**  <br>*optional*||object|
|**name**  <br>*required*|The name of the destinationPolicy|string|
|**resourceVersion**  <br>*required*|The resource version of the destinationPolicy|string|
|**source**  <br>*optional*||object|


<a name="destinationrule"></a>
### destinationRule
This is used for returning a DestinationRule


|Name|Description|Schema|
|---|---|---|
|**createdAt**  <br>*required*|The creation date of the destinationRule|string|
|**host**  <br>*optional*||object|
|**name**  <br>*required*|The name of the destinationRule|string|
|**resourceVersion**  <br>*required*|The resource version of the destinationRule|string|
|**subsets**  <br>*optional*||object|
|**trafficPolicy**  <br>*optional*||object|


<a name="destinationrules"></a>
### destinationRules
This is used for returning an array of DestinationRules

*Type* : < [destinationRule](#destinationrule) > array


<a name="externalserviceinfo"></a>
### externalServiceInfo
This is used for returning a response of Kiali Status


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*required*|The name of the service  <br>**Example** : `"Istio"`|string|
|**version**  <br>*required*|The installed version of the service  <br>**Example** : `"0.8.0"`|string|


<a name="istiorule"></a>
### istioRule
This is used for returning a IstioRule


|Name|Description|Schema|
|---|---|---|
|**actions**  <br>*optional*||object|
|**match**  <br>*optional*||object|
|**name**  <br>*required*|The name of the istioRule|string|


<a name="istiorules"></a>
### istioRules
This is used for returning an array of IstioRules

*Type* : < [istioRule](#istiorule) > array


<a name="namespace"></a>
### namespace
It is used to describe a set of objects.


|Name|Description|Schema|
|---|---|---|
|**name**  <br>*required*|The id of the namespace.  <br>**Example** : `"istio-system"`|string|


<a name="routerule"></a>
### routeRule
This is used for returning a RouteRule


|Name|Description|Schema|
|---|---|---|
|**appendHeaders**  <br>*optional*||object|
|**corsPolicy**  <br>*optional*||object|
|**createdAt**  <br>*required*|The created time  <br>**Example** : `"2018-06-20T07:39:52Z"`|string|
|**destination**  <br>*optional*||object|
|**httpFault**  <br>*optional*||object|
|**httpReqRetries**  <br>*optional*||object|
|**httpReqTimeout**  <br>*optional*||object|
|**l4Fault**  <br>*optional*||object|
|**match**  <br>*optional*||object|
|**mirror**  <br>*optional*||object|
|**name**  <br>*required*|The name of the routeRule  <br>**Example** : `"details-default"`|string|
|**precedence**  <br>*optional*||object|
|**redirect**  <br>*optional*||object|
|**resourceVersion**  <br>*required*|**Example** : `"1507"`|string|
|**rewrite**  <br>*optional*||object|
|**route**  <br>*optional*||object|
|**routeWarning**  <br>*optional*||string|
|**websocketUpgrade**  <br>*optional*||object|


<a name="routerules"></a>
### routeRules
This is used for returning an array of RouteRule

*Type* : < [routeRule](#routerule) > array


<a name="virtualservice"></a>
### virtualService
This is used for returning a VirtualService


|Name|Description|Schema|
|---|---|---|
|**createdAt**  <br>*required*|The creation date of the virtualService|string|
|**gateways**  <br>*optional*||object|
|**hosts**  <br>*optional*||object|
|**http**  <br>*optional*||object|
|**name**  <br>*required*|The name of the virtualService|string|
|**resourceVersion**  <br>*required*|The resource version of the virtualService|string|
|**tcp**  <br>*optional*||object|


<a name="virtualservices"></a>
### virtualServices
This is used for returning an array of VirtualServices

*Type* : < [virtualService](#virtualservice) > array



