# \SystemReverseProxyApi

All URIs are relative to *http://192.168.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReverseProxyMappings**](SystemReverseProxyApi.md#CreateReverseProxyMappings) | **Post** /system/reverseproxy/{uuid} | 
[**GetReverseProxyMappings**](SystemReverseProxyApi.md#GetReverseProxyMappings) | **Get** /system/reverseproxy/{uuid} | 



## CreateReverseProxyMappings

> CreateReverseProxyMappings(ctx, uuid).Body(body).Execute()





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
    uuid := "uuid_example" // string | Device uuid
    body := map[string]interface{}{ ... } // map[string]interface{} | Device reverse proxy mappings (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemReverseProxyApi.CreateReverseProxyMappings(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemReverseProxyApi.CreateReverseProxyMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReverseProxyMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device reverse proxy mappings | 

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


## GetReverseProxyMappings

> map[string]interface{} GetReverseProxyMappings(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | Device uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemReverseProxyApi.GetReverseProxyMappings(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemReverseProxyApi.GetReverseProxyMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReverseProxyMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemReverseProxyApi.GetReverseProxyMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReverseProxyMappingsRequest struct via the builder pattern


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

