# \SystemCloudServiceWebexApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebexDataCenters**](SystemCloudServiceWebexApi.md#DeleteWebexDataCenters) | **Delete** /webex/datacenter | 
[**GetWebexDataCenters**](SystemCloudServiceWebexApi.md#GetWebexDataCenters) | **Post** /webex/datacenter | 
[**GetWebexDataCentersSyncStatus**](SystemCloudServiceWebexApi.md#GetWebexDataCentersSyncStatus) | **Get** /webex/datacenter/syncstatus | 
[**SetWebexDataCentersSyncStatus**](SystemCloudServiceWebexApi.md#SetWebexDataCentersSyncStatus) | **Put** /webex/datacenter/syncstatus | 
[**UpdateWebexDataCenters**](SystemCloudServiceWebexApi.md#UpdateWebexDataCenters) | **Post** /webex/datacenter/sync | 



## DeleteWebexDataCenters

> map[string]interface{} DeleteWebexDataCenters(ctx).Execute()





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
    resp, r, err := apiClient.SystemCloudServiceWebexApi.DeleteWebexDataCenters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemCloudServiceWebexApi.DeleteWebexDataCenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWebexDataCenters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemCloudServiceWebexApi.DeleteWebexDataCenters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebexDataCentersRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebexDataCenters

> map[string]interface{} GetWebexDataCenters(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemCloudServiceWebexApi.GetWebexDataCenters(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemCloudServiceWebexApi.GetWebexDataCenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebexDataCenters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemCloudServiceWebexApi.GetWebexDataCenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebexDataCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebexDataCentersSyncStatus

> map[string]interface{} GetWebexDataCentersSyncStatus(ctx).Execute()





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
    resp, r, err := apiClient.SystemCloudServiceWebexApi.GetWebexDataCentersSyncStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemCloudServiceWebexApi.GetWebexDataCentersSyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebexDataCentersSyncStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemCloudServiceWebexApi.GetWebexDataCentersSyncStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebexDataCentersSyncStatusRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWebexDataCentersSyncStatus

> map[string]interface{} SetWebexDataCentersSyncStatus(ctx).Execute()





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
    resp, r, err := apiClient.SystemCloudServiceWebexApi.SetWebexDataCentersSyncStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemCloudServiceWebexApi.SetWebexDataCentersSyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetWebexDataCentersSyncStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemCloudServiceWebexApi.SetWebexDataCentersSyncStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetWebexDataCentersSyncStatusRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebexDataCenters

> map[string]interface{} UpdateWebexDataCenters(ctx).Execute()





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
    resp, r, err := apiClient.SystemCloudServiceWebexApi.UpdateWebexDataCenters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemCloudServiceWebexApi.UpdateWebexDataCenters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebexDataCenters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemCloudServiceWebexApi.UpdateWebexDataCenters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebexDataCentersRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

