# \RealTimeMonitoringBGPApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBGPNeighborsList**](RealTimeMonitoringBGPApi.md#CreateBGPNeighborsList) | **Get** /device/bgp/neighbors | 
[**CreateBGPRoutesList**](RealTimeMonitoringBGPApi.md#CreateBGPRoutesList) | **Get** /device/bgp/routes | 
[**CreateBGPSummary**](RealTimeMonitoringBGPApi.md#CreateBGPSummary) | **Get** /device/bgp/summary | 



## CreateBGPNeighborsList

> []map[string]interface{} CreateBGPNeighborsList(ctx).DeviceId(deviceId).VpnId(vpnId).PeerAddr(peerAddr).As(as).Execute()





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
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    peerAddr := "peerAddr_example" // string | Peer address (optional)
    as := "as_example" // string | AS number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBGPApi.CreateBGPNeighborsList(context.Background()).DeviceId(deviceId).VpnId(vpnId).PeerAddr(peerAddr).As(as).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBGPApi.CreateBGPNeighborsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBGPNeighborsList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBGPApi.CreateBGPNeighborsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBGPNeighborsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **peerAddr** | **string** | Peer address | 
 **as** | **string** | AS number | 

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


## CreateBGPRoutesList

> []map[string]interface{} CreateBGPRoutesList(ctx).DeviceId(deviceId).VpnId(vpnId).Prefix(prefix).Nexthop(nexthop).Execute()





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
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    prefix := "prefix_example" // string | IP prefix (optional)
    nexthop := "nexthop_example" // string | Next hop (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBGPApi.CreateBGPRoutesList(context.Background()).DeviceId(deviceId).VpnId(vpnId).Prefix(prefix).Nexthop(nexthop).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBGPApi.CreateBGPRoutesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBGPRoutesList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBGPApi.CreateBGPRoutesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBGPRoutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **prefix** | **string** | IP prefix | 
 **nexthop** | **string** | Next hop | 

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


## CreateBGPSummary

> map[string]interface{} CreateBGPSummary(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBGPApi.CreateBGPSummary(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBGPApi.CreateBGPSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBGPSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBGPApi.CreateBGPSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBGPSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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

