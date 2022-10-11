# \SystemContainerApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateContainer**](SystemContainerApi.md#ActivateContainer) | **Post** /sdavc/task/{taskId} | 
[**ActivateContainerOnRemoteHost**](SystemContainerApi.md#ActivateContainerOnRemoteHost) | **Post** /container-manager/activate/{containerName} | 
[**DeActivateContainer**](SystemContainerApi.md#DeActivateContainer) | **Post** /container-manager/deactivate/{containerName} | 
[**DoesValidImageExist**](SystemContainerApi.md#DoesValidImageExist) | **Get** /container-manager/doesValidImageExist/{containerName} | 
[**GetContainerInspectData**](SystemContainerApi.md#GetContainerInspectData) | **Get** /container-manager/inspect/{containerName} | 
[**GetContainerSettings**](SystemContainerApi.md#GetContainerSettings) | **Get** /container-manager/settings/{containerName} | 
[**GetCustomApp**](SystemContainerApi.md#GetCustomApp) | **Get** /sdavc/customapps | 
[**TestLoadBalancer**](SystemContainerApi.md#TestLoadBalancer) | **Post** /sdavc/test | 



## ActivateContainer

> ActivateContainer(ctx, taskId).Body(body).Execute()





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
    taskId := "taskId_example" // string | Task Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Container task config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemContainerApi.ActivateContainer(context.Background(), taskId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.ActivateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Task Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Container task config | 

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


## ActivateContainerOnRemoteHost

> ActivateContainerOnRemoteHost(ctx, containerName).Url(url).HostIp(hostIp).Checksum(checksum).Execute()





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
    containerName := "containerName_example" // string | Container name
    url := "url_example" // string | Container image URL (optional)
    hostIp := "hostIp_example" // string | Container host IP (optional)
    checksum := "checksum_example" // string | Container image checksum (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemContainerApi.ActivateContainerOnRemoteHost(context.Background(), containerName).Url(url).HostIp(hostIp).Checksum(checksum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.ActivateContainerOnRemoteHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerName** | **string** | Container name | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateContainerOnRemoteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **string** | Container image URL | 
 **hostIp** | **string** | Container host IP | 
 **checksum** | **string** | Container image checksum | 

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


## DeActivateContainer

> DeActivateContainer(ctx, containerName).HostIp(hostIp).Execute()





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
    containerName := "containerName_example" // string | Container name
    hostIp := "hostIp_example" // string | Container host IP (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemContainerApi.DeActivateContainer(context.Background(), containerName).HostIp(hostIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.DeActivateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerName** | **string** | Container name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeActivateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostIp** | **string** | Container host IP | 

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


## DoesValidImageExist

> map[string]interface{} DoesValidImageExist(ctx, containerName).Execute()





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
    containerName := "containerName_example" // string | Container name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemContainerApi.DoesValidImageExist(context.Background(), containerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.DoesValidImageExist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DoesValidImageExist`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemContainerApi.DoesValidImageExist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerName** | **string** | Container name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoesValidImageExistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetContainerInspectData

> string GetContainerInspectData(ctx, containerName).HostIp(hostIp).Execute()





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
    containerName := "containerName_example" // string | Container name
    hostIp := "hostIp_example" // string | Container host IP (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemContainerApi.GetContainerInspectData(context.Background(), containerName).HostIp(hostIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.GetContainerInspectData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerInspectData`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemContainerApi.GetContainerInspectData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerName** | **string** | Container name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerInspectDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostIp** | **string** | Container host IP | 

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


## GetContainerSettings

> map[string]interface{} GetContainerSettings(ctx, containerName).HostIp(hostIp).Execute()





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
    containerName := "containerName_example" // string | Container name
    hostIp := "hostIp_example" // string | Container host IP (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemContainerApi.GetContainerSettings(context.Background(), containerName).HostIp(hostIp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.GetContainerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemContainerApi.GetContainerSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerName** | **string** | Container name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostIp** | **string** | Container host IP | 

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


## GetCustomApp

> []map[string]interface{} GetCustomApp(ctx).Execute()





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
    resp, r, err := apiClient.SystemContainerApi.GetCustomApp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.GetCustomApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomApp`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemContainerApi.GetCustomApp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAppRequest struct via the builder pattern


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


## TestLoadBalancer

> TestLoadBalancer(ctx).Execute()





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
    resp, r, err := apiClient.SystemContainerApi.TestLoadBalancer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemContainerApi.TestLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestLoadBalancerRequest struct via the builder pattern


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

