# \ColocationClusterApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcitvateCloudDockCluster**](ColocationClusterApi.md#AcitvateCloudDockCluster) | **Post** /colocation/cluster/activate | 
[**CloudDockClusterPreview**](ColocationClusterApi.md#CloudDockClusterPreview) | **Get** /colocation/cluster/config | 
[**CreateCloudDockCluster**](ColocationClusterApi.md#CreateCloudDockCluster) | **Post** /colocation/cluster | 
[**DeAcitvateCloudDockCluster**](ColocationClusterApi.md#DeAcitvateCloudDockCluster) | **Post** /colocation/cluster/deactivate | 
[**DeleteCloudDockClusterByName**](ColocationClusterApi.md#DeleteCloudDockClusterByName) | **Delete** /colocation/cluster/{clustername} | 
[**Dummyccm**](ColocationClusterApi.md#Dummyccm) | **Get** /colocation/cluster/activateClusterDummy | 
[**DummycspState**](ColocationClusterApi.md#DummycspState) | **Get** /colocation/cluster/activateClusterDummyState | 
[**GetCloudDockClusterDetail**](ColocationClusterApi.md#GetCloudDockClusterDetail) | **Get** /colocation/cluster | 
[**GetCloudDockClusterDetailById**](ColocationClusterApi.md#GetCloudDockClusterDetailById) | **Get** /colocation/cluster/id | 
[**RmaCloudDockCsp**](ColocationClusterApi.md#RmaCloudDockCsp) | **Post** /colocation/cluster/rma | 
[**UpdateCloudDockCluster**](ColocationClusterApi.md#UpdateCloudDockCluster) | **Put** /colocation/cluster | 
[**UpdateCspToCluster**](ColocationClusterApi.md#UpdateCspToCluster) | **Put** /colocation/cluster/attached/csp | 



## AcitvateCloudDockCluster

> map[string]interface{} AcitvateCloudDockCluster(ctx).ClusterName(clusterName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.AcitvateCloudDockCluster(context.Background()).ClusterName(clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.AcitvateCloudDockCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcitvateCloudDockCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationClusterApi.AcitvateCloudDockCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcitvateCloudDockClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** | Cluster name | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDockClusterPreview

> string CloudDockClusterPreview(ctx).SerialNumber(serialNumber).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serialNumber := "serialNumber_example" // string | Serial number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.CloudDockClusterPreview(context.Background()).SerialNumber(serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.CloudDockClusterPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudDockClusterPreview`: string
    fmt.Fprintf(os.Stdout, "Response from `ColocationClusterApi.CloudDockClusterPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudDockClusterPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serialNumber** | **string** | Serial number | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudDockCluster

> CreateCloudDockCluster(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | Cluster config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.CreateCloudDockCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.CreateCloudDockCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudDockClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Cluster config | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeAcitvateCloudDockCluster

> DeAcitvateCloudDockCluster(ctx).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | Cluster Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.DeAcitvateCloudDockCluster(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.DeAcitvateCloudDockCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeAcitvateCloudDockClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudDockClusterByName

> DeleteCloudDockClusterByName(ctx, clustername).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clustername := "clustername_example" // string | Cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.DeleteCloudDockClusterByName(context.Background(), clustername).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.DeleteCloudDockClusterByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clustername** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudDockClusterByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Dummyccm

> Dummyccm(ctx).ClusterName(clusterName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.Dummyccm(context.Background()).ClusterName(clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.Dummyccm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDummyccmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** | Cluster name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DummycspState

> DummycspState(ctx).ClusterName(clusterName).State(state).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Cluster name
    state := "state_example" // string | Cluster state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.DummycspState(context.Background()).ClusterName(clusterName).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.DummycspState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDummycspStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** | Cluster name | 
 **state** | **string** | Cluster state | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDockClusterDetail

> map[string]interface{} GetCloudDockClusterDetail(ctx).ClusterName(clusterName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.GetCloudDockClusterDetail(context.Background()).ClusterName(clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.GetCloudDockClusterDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDockClusterDetail`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationClusterApi.GetCloudDockClusterDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDockClusterDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** | Cluster name | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudDockClusterDetailById

> map[string]interface{} GetCloudDockClusterDetailById(ctx).ClusterId(clusterId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | Cluster Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.GetCloudDockClusterDetailById(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.GetCloudDockClusterDetailById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDockClusterDetailById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationClusterApi.GetCloudDockClusterDetailById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDockClusterDetailByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RmaCloudDockCsp

> map[string]interface{} RmaCloudDockCsp(ctx).ClusterName(clusterName).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterName := "clusterName_example" // string | Cluster name
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.RmaCloudDockCsp(context.Background()).ClusterName(clusterName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.RmaCloudDockCsp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RmaCloudDockCsp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationClusterApi.RmaCloudDockCsp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRmaCloudDockCspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterName** | **string** | Cluster name | 
 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudDockCluster

> UpdateCloudDockCluster(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | Cluster config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.UpdateCloudDockCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.UpdateCloudDockCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudDockClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Cluster config | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCspToCluster

> UpdateCspToCluster(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | CSP config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationClusterApi.UpdateCspToCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationClusterApi.UpdateCspToCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCspToClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | CSP config | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

