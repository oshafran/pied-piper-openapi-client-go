# \MonitoringHealthDevicesApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevicesHealth**](MonitoringHealthDevicesApi.md#GetDevicesHealth) | **Get** /health/devices | 
[**GetDevicesHealthOverview**](MonitoringHealthDevicesApi.md#GetDevicesHealthOverview) | **Get** /health/devices/overview | 



## GetDevicesHealth

> map[string]interface{} GetDevicesHealth(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).StartingDeviceId(startingDeviceId).SiteId(siteId).GroupId(groupId).GroupId2(groupId2).VpnId(vpnId).Reachable(reachable).ControlStatus(controlStatus).Personality(personality).Health(health).Execute()





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
    page := int64(789) // int64 | Page Number (optional)
    pageSize := int64(789) // int64 | Page Size (optional)
    sortBy := "sortBy_example" // string | Sort By Property (optional)
    sortOrder := "sortOrder_example" // string | Sort Order (optional)
    startingDeviceId := "startingDeviceId_example" // string | Optional device ID to start first page (optional)
    siteId := "siteId_example" // string | Optional site ID to filter devices (optional)
    groupId := "groupId_example" // string | Optional group ID to filter devices (optional)
    groupId2 := "groupId_example" // string | Optional group ID to filter devices (optional)
    vpnId := "vpnId_example" // string | Optional vpn ID to filter devices (optional)
    reachable := true // bool |  (optional)
    controlStatus := "controlStatus_example" // string |  (optional)
    personality := "personality_example" // string |  (optional)
    health := "health_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringHealthDevicesApi.GetDevicesHealth(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).StartingDeviceId(startingDeviceId).SiteId(siteId).GroupId(groupId).GroupId2(groupId2).VpnId(vpnId).Reachable(reachable).ControlStatus(controlStatus).Personality(personality).Health(health).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringHealthDevicesApi.GetDevicesHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicesHealth`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringHealthDevicesApi.GetDevicesHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page Number | 
 **pageSize** | **int64** | Page Size | 
 **sortBy** | **string** | Sort By Property | 
 **sortOrder** | **string** | Sort Order | 
 **startingDeviceId** | **string** | Optional device ID to start first page | 
 **siteId** | **string** | Optional site ID to filter devices | 
 **groupId** | **string** | Optional group ID to filter devices | 
 **groupId2** | **string** | Optional group ID to filter devices | 
 **vpnId** | **string** | Optional vpn ID to filter devices | 
 **reachable** | **bool** |  | 
 **controlStatus** | **string** |  | 
 **personality** | **string** |  | 
 **health** | **string** |  | 

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


## GetDevicesHealthOverview

> map[string]interface{} GetDevicesHealthOverview(ctx).VpnId(vpnId).VpnId2(vpnId2).Execute()





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
    vpnId := "vpnId_example" // string | Optional vpn ID to filter devices (optional)
    vpnId2 := "vpnId_example" // string | Optional vpn ID to filter devices (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringHealthDevicesApi.GetDevicesHealthOverview(context.Background()).VpnId(vpnId).VpnId2(vpnId2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringHealthDevicesApi.GetDevicesHealthOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicesHealthOverview`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringHealthDevicesApi.GetDevicesHealthOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesHealthOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnId** | **string** | Optional vpn ID to filter devices | 
 **vpnId2** | **string** | Optional vpn ID to filter devices | 

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

