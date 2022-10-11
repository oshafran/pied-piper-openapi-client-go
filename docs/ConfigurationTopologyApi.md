# \ConfigurationTopologyApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceTopology**](ConfigurationTopologyApi.md#CreateDeviceTopology) | **Get** /topology/device | 
[**CreateFullTopology**](ConfigurationTopologyApi.md#CreateFullTopology) | **Get** /topology | 
[**CreatePhysicalTopology**](ConfigurationTopologyApi.md#CreatePhysicalTopology) | **Get** /topology/physical | 
[**GetSiteTopology**](ConfigurationTopologyApi.md#GetSiteTopology) | **Get** /topology/device/site/{siteId} | 



## CreateDeviceTopology

> map[string]interface{} CreateDeviceTopology(ctx).DeviceId(deviceId).Execute()





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
    deviceId := []openapiclient.DeviceIP{*openapiclient.NewDeviceIP()} // []DeviceIP | Device Id list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTopologyApi.CreateDeviceTopology(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTopologyApi.CreateDeviceTopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceTopology`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTopologyApi.CreateDeviceTopology`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | [**[]DeviceIP**](DeviceIP.md) | Device Id list | 

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


## CreateFullTopology

> []map[string]interface{} CreateFullTopology(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationTopologyApi.CreateFullTopology(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTopologyApi.CreateFullTopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFullTopology`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTopologyApi.CreateFullTopology`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFullTopologyRequest struct via the builder pattern


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


## CreatePhysicalTopology

> map[string]interface{} CreatePhysicalTopology(ctx).DeviceId(deviceId).Execute()





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
    deviceId := []openapiclient.DeviceIP{*openapiclient.NewDeviceIP()} // []DeviceIP | Device Id list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTopologyApi.CreatePhysicalTopology(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTopologyApi.CreatePhysicalTopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhysicalTopology`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTopologyApi.CreatePhysicalTopology`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhysicalTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | [**[]DeviceIP**](DeviceIP.md) | Device Id list | 

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


## GetSiteTopology

> map[string]interface{} GetSiteTopology(ctx, siteId).Execute()





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
    siteId := "siteId_example" // string | Site Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTopologyApi.GetSiteTopology(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTopologyApi.GetSiteTopology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSiteTopology`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTopologyApi.GetSiteTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteTopologyRequest struct via the builder pattern


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

