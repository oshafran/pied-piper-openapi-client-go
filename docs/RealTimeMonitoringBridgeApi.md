# \RealTimeMonitoringBridgeApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBridgeInterfaceList**](RealTimeMonitoringBridgeApi.md#GetBridgeInterfaceList) | **Get** /device/bridge/interface | 
[**GetBridgeInterfaceMac**](RealTimeMonitoringBridgeApi.md#GetBridgeInterfaceMac) | **Get** /device/bridge/mac | 
[**GetBridgeInterfaceTable**](RealTimeMonitoringBridgeApi.md#GetBridgeInterfaceTable) | **Get** /device/bridge/table | 



## GetBridgeInterfaceList

> map[string]interface{} GetBridgeInterfaceList(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBridgeApi.GetBridgeInterfaceList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBridgeApi.GetBridgeInterfaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBridgeInterfaceList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBridgeApi.GetBridgeInterfaceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBridgeInterfaceListRequest struct via the builder pattern


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


## GetBridgeInterfaceMac

> map[string]interface{} GetBridgeInterfaceMac(ctx).DeviceId(deviceId).BridgeId(bridgeId).IfName(ifName).MacAddress(macAddress).Execute()





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
    bridgeId := "bridgeId_example" // string | Bridge ID (optional)
    ifName := "ifName_example" // string | Interface name (optional)
    macAddress := "macAddress_example" // string | MAC address (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBridgeApi.GetBridgeInterfaceMac(context.Background()).DeviceId(deviceId).BridgeId(bridgeId).IfName(ifName).MacAddress(macAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBridgeApi.GetBridgeInterfaceMac``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBridgeInterfaceMac`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBridgeApi.GetBridgeInterfaceMac`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBridgeInterfaceMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **bridgeId** | **string** | Bridge ID | 
 **ifName** | **string** | Interface name | 
 **macAddress** | **string** | MAC address | 

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


## GetBridgeInterfaceTable

> map[string]interface{} GetBridgeInterfaceTable(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBridgeApi.GetBridgeInterfaceTable(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBridgeApi.GetBridgeInterfaceTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBridgeInterfaceTable`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBridgeApi.GetBridgeInterfaceTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBridgeInterfaceTableRequest struct via the builder pattern


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

