
<a name="paths"></a>
## Paths

<a name="root"></a>
### GET /

#### Description
Handler to get the status of Kiali


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|HTTP status code 200 and statusInfo model in data|[Response 200](#root-response-200)|

<a name="root-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**code**  <br>*optional*|HTTP status code  <br>**Default** : `200`|integer (int64)|
|**data**  <br>*optional*||[StatusInfo](#statusinfo)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="istioconfiglist"></a>
### GET /namespaces/{namespace}/istio

#### Description
Handler to get the list of Istio Config of a namespace


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**namespace**  <br>*required*|The id of the namespace.|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|HTTP status code 200 and IstioConfigList model in data|[Response 200](#istioconfiglist-response-200)|

<a name="istioconfiglist-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**code**  <br>*optional*|HTTP status code  <br>**Default** : `200`|integer (int64)|
|**data**  <br>*optional*||[IstioConfigList](#istioconfiglist)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="getstatus"></a>
### GET /status

#### Description
Handler to get the status of Kiali


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|HTTP status code 200 and statusInfo model in data|[Response 200](#getstatus-response-200)|

<a name="getstatus-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**code**  <br>*optional*|HTTP status code  <br>**Default** : `200`|integer (int64)|
|**data**  <br>*optional*||[StatusInfo](#statusinfo)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


<a name="gettoken"></a>
### GET /token

#### Description
Handler to get the authentication token


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|HTTP status code 200 and tokenGenerated model in data|[Response 200](#gettoken-response-200)|

<a name="gettoken-response-200"></a>
**Response 200**

|Name|Description|Schema|
|---|---|---|
|**code**  <br>*optional*|HTTP status code  <br>**Default** : `200`|integer (int64)|
|**data**  <br>*optional*||[TokenGenerated](#tokengenerated)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### Security

|Type|Name|Scopes|
|---|---|---|
|**Unknown**|**[authorization](#authorization)**|user,password|



