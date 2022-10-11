# \ColocationServiceGroupApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceGroupCluster**](ColocationServiceGroupApi.md#CreateServiceGroupCluster) | **Post** /colocation/servicegroup | 
[**DeleteServiceGroupCluster**](ColocationServiceGroupApi.md#DeleteServiceGroupCluster) | **Delete** /colocation/servicegroup/{name} | 
[**GetAvailableChains**](ColocationServiceGroupApi.md#GetAvailableChains) | **Get** /colocation/servicegroup/servicechains | 
[**GetDefaultChain**](ColocationServiceGroupApi.md#GetDefaultChain) | **Get** /colocation/servicegroup/servicechain/default | 
[**GetServiceChain**](ColocationServiceGroupApi.md#GetServiceChain) | **Get** /colocation/servicegroup | 
[**GetServiceGroupInCluster**](ColocationServiceGroupApi.md#GetServiceGroupInCluster) | **Get** /colocation/servicegroup/attached | 
[**UpdateServiceGroupCluster**](ColocationServiceGroupApi.md#UpdateServiceGroupCluster) | **Put** /colocation/servicegroup | 



## CreateServiceGroupCluster

> CreateServiceGroupCluster(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Service group (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.CreateServiceGroupCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.CreateServiceGroupCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceGroupClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Service group | 

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


## DeleteServiceGroupCluster

> DeleteServiceGroupCluster(ctx, name).Execute()





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
    name := "name_example" // string | Service group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.DeleteServiceGroupCluster(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.DeleteServiceGroupCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Service group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceGroupClusterRequest struct via the builder pattern


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


## GetAvailableChains

> []map[string]interface{} GetAvailableChains(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.GetAvailableChains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.GetAvailableChains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableChains`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationServiceGroupApi.GetAvailableChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableChainsRequest struct via the builder pattern


### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultChain

> map[string]interface{} GetDefaultChain(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.GetDefaultChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.GetDefaultChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultChain`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationServiceGroupApi.GetDefaultChain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultChainRequest struct via the builder pattern


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


## GetServiceChain

> map[string]interface{} GetServiceChain(ctx).ServiceGroupName(serviceGroupName).Execute()





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
    serviceGroupName := "serviceGroupName_example" // string | Service group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.GetServiceChain(context.Background()).ServiceGroupName(serviceGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.GetServiceChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceChain`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationServiceGroupApi.GetServiceChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceGroupName** | **string** | Service group name | 

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


## GetServiceGroupInCluster

> map[string]interface{} GetServiceGroupInCluster(ctx).ClusterId(clusterId).UserGroupName(userGroupName).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id (optional)
    userGroupName := "userGroupName_example" // string | UserGroup Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.GetServiceGroupInCluster(context.Background()).ClusterId(clusterId).UserGroupName(userGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.GetServiceGroupInCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceGroupInCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ColocationServiceGroupApi.GetServiceGroupInCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceGroupInClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 
 **userGroupName** | **string** | UserGroup Name | 

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


## UpdateServiceGroupCluster

> UpdateServiceGroupCluster(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Service group (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ColocationServiceGroupApi.UpdateServiceGroupCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ColocationServiceGroupApi.UpdateServiceGroupCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceGroupClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Service group | 

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

