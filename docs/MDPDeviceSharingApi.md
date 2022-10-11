# \MDPDeviceSharingApi

All URIs are relative to *http://192.168.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachDevices**](MDPDeviceSharingApi.md#AttachDevices) | **Post** /mdp/attachDevices/{nmsId} | 
[**DetachDevices**](MDPDeviceSharingApi.md#DetachDevices) | **Post** /mdp/detachDevices/{nmsId} | 
[**EditAttachedDevices**](MDPDeviceSharingApi.md#EditAttachedDevices) | **Put** /mdp/attachDevices/{nmsId} | 
[**RetrieveMDPAttachedDevices**](MDPDeviceSharingApi.md#RetrieveMDPAttachedDevices) | **Get** /mdp/attachDevices/{nmsId} | 
[**RetrieveMDPSupportedDevices**](MDPDeviceSharingApi.md#RetrieveMDPSupportedDevices) | **Get** /mdp/devices/{nmsId} | 



## AttachDevices

> map[string]interface{} AttachDevices(ctx, nmsId).Body(body).Execute()





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
    nmsId := "nmsId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | deviceList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPDeviceSharingApi.AttachDevices(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPDeviceSharingApi.AttachDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPDeviceSharingApi.AttachDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | deviceList | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachDevices

> map[string]interface{} DetachDevices(ctx, nmsId).Body(body).Execute()





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
    nmsId := "nmsId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | deviceList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPDeviceSharingApi.DetachDevices(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPDeviceSharingApi.DetachDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPDeviceSharingApi.DetachDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | deviceList | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAttachedDevices

> map[string]interface{} EditAttachedDevices(ctx, nmsId).Body(body).Execute()





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
    nmsId := "nmsId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | deviceList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPDeviceSharingApi.EditAttachedDevices(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPDeviceSharingApi.EditAttachedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAttachedDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPDeviceSharingApi.EditAttachedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAttachedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | deviceList | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMDPAttachedDevices

> []map[string]interface{} RetrieveMDPAttachedDevices(ctx, nmsId).Execute()





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
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPDeviceSharingApi.RetrieveMDPAttachedDevices(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPDeviceSharingApi.RetrieveMDPAttachedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPAttachedDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPDeviceSharingApi.RetrieveMDPAttachedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPAttachedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RetrieveMDPSupportedDevices

> []map[string]interface{} RetrieveMDPSupportedDevices(ctx, nmsId).Execute()





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
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPDeviceSharingApi.RetrieveMDPSupportedDevices(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPDeviceSharingApi.RetrieveMDPSupportedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPSupportedDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPDeviceSharingApi.RetrieveMDPSupportedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPSupportedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

