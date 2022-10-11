# \RealTimeMonitoringApplicationAwareRouteApi

All URIs are relative to *http://$VMANAGE_EXTERNAL_IP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppRouteSlaClassList**](RealTimeMonitoringApplicationAwareRouteApi.md#CreateAppRouteSlaClassList) | **Get** /device/app-route/sla-class | 
[**CreateAppRouteStatisticsList**](RealTimeMonitoringApplicationAwareRouteApi.md#CreateAppRouteStatisticsList) | **Get** /device/app-route/statistics | 



## CreateAppRouteSlaClassList

> map[string]interface{} CreateAppRouteSlaClassList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringApplicationAwareRouteApi.CreateAppRouteSlaClassList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringApplicationAwareRouteApi.CreateAppRouteSlaClassList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppRouteSlaClassList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringApplicationAwareRouteApi.CreateAppRouteSlaClassList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRouteSlaClassListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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


## CreateAppRouteStatisticsList

> map[string]interface{} CreateAppRouteStatisticsList(ctx).DeviceId(deviceId).RemoteSystemIp(remoteSystemIp).LocalColor(localColor).RemoteColor(remoteColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteSystemIp := "remoteSystemIp_example" // string | Remote system IP (optional)
    localColor := "localColor_example" // string | Local color (optional)
    remoteColor := "remoteColor_example" // string | Remote color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringApplicationAwareRouteApi.CreateAppRouteStatisticsList(context.Background()).DeviceId(deviceId).RemoteSystemIp(remoteSystemIp).LocalColor(localColor).RemoteColor(remoteColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringApplicationAwareRouteApi.CreateAppRouteStatisticsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppRouteStatisticsList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringApplicationAwareRouteApi.CreateAppRouteStatisticsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRouteStatisticsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteSystemIp** | **string** | Remote system IP | 
 **localColor** | **string** | Local color | 
 **remoteColor** | **string** | Remote color | 

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

