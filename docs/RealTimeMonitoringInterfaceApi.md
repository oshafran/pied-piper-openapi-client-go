# \RealTimeMonitoringInterfaceApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateDeviceInterfaceVPN**](RealTimeMonitoringInterfaceApi.md#GenerateDeviceInterfaceVPN) | **Get** /device/interface/vpn | 
[**GetDeviceInterface**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterface) | **Get** /device/interface | 
[**GetDeviceInterfaceARPStats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfaceARPStats) | **Get** /device/interface/arp_stats | 
[**GetDeviceInterfaceErrorStats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfaceErrorStats) | **Get** /device/interface/error_stats | 
[**GetDeviceInterfaceIpv6Stats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfaceIpv6Stats) | **Get** /device/interface/ipv6Stats | 
[**GetDeviceInterfacePktSizes**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfacePktSizes) | **Get** /device/interface/pkt_size | 
[**GetDeviceInterfacePortStats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfacePortStats) | **Get** /device/interface/port_stats | 
[**GetDeviceInterfaceQosStats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfaceQosStats) | **Get** /device/interface/qosStats | 
[**GetDeviceInterfaceQueueStats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfaceQueueStats) | **Get** /device/interface/queue_stats | 
[**GetDeviceInterfaceStats**](RealTimeMonitoringInterfaceApi.md#GetDeviceInterfaceStats) | **Get** /device/interface/stats | 
[**GetDeviceSerialInterface**](RealTimeMonitoringInterfaceApi.md#GetDeviceSerialInterface) | **Get** /device/interface/serial | 
[**GetSyncedDeviceInterface**](RealTimeMonitoringInterfaceApi.md#GetSyncedDeviceInterface) | **Get** /device/interface/synced | 
[**Trustsec**](RealTimeMonitoringInterfaceApi.md#Trustsec) | **Get** /device/interface/trustsec | 



## GenerateDeviceInterfaceVPN

> map[string]interface{} GenerateDeviceInterfaceVPN(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GenerateDeviceInterfaceVPN(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GenerateDeviceInterfaceVPN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDeviceInterfaceVPN`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GenerateDeviceInterfaceVPN`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceInterfaceVPNRequest struct via the builder pattern


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


## GetDeviceInterface

> map[string]interface{} GetDeviceInterface(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterface(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfaceARPStats

> map[string]interface{} GetDeviceInterfaceARPStats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfaceARPStats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceARPStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfaceARPStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceARPStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceARPStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfaceErrorStats

> map[string]interface{} GetDeviceInterfaceErrorStats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfaceErrorStats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceErrorStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfaceErrorStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceErrorStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceErrorStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfaceIpv6Stats

> map[string]interface{} GetDeviceInterfaceIpv6Stats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfaceIpv6Stats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceIpv6Stats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfaceIpv6Stats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceIpv6Stats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceIpv6StatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfacePktSizes

> map[string]interface{} GetDeviceInterfacePktSizes(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfacePktSizes(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfacePktSizes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfacePktSizes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfacePktSizes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfacePktSizesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfacePortStats

> map[string]interface{} GetDeviceInterfacePortStats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfacePortStats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfacePortStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfacePortStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfacePortStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfacePortStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfaceQosStats

> map[string]interface{} GetDeviceInterfaceQosStats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfaceQosStats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceQosStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfaceQosStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceQosStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceQosStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfaceQueueStats

> map[string]interface{} GetDeviceInterfaceQueueStats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfaceQueueStats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceQueueStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfaceQueueStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceQueueStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceQueueStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceInterfaceStats

> map[string]interface{} GetDeviceInterfaceStats(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceInterfaceStats(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInterfaceStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceInterfaceStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInterfaceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetDeviceSerialInterface

> map[string]interface{} GetDeviceSerialInterface(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetDeviceSerialInterface(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetDeviceSerialInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSerialInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetDeviceSerialInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSerialInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## GetSyncedDeviceInterface

> map[string]interface{} GetSyncedDeviceInterface(ctx).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()





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
    afType := "afType_example" // string | AF Type
    deviceId := "deviceId_example" // string | Device Id
    vpnId := "vpnId_example" // string | VPN Id (optional)
    ifname := "ifname_example" // string | IF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.GetSyncedDeviceInterface(context.Background()).AfType(afType).DeviceId(deviceId).VpnId(vpnId).Ifname(ifname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.GetSyncedDeviceInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyncedDeviceInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.GetSyncedDeviceInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncedDeviceInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afType** | **string** | AF Type | 
 **deviceId** | **string** | Device Id | 
 **vpnId** | **string** | VPN Id | 
 **ifname** | **string** | IF Name | 

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


## Trustsec

> map[string]interface{} Trustsec(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringInterfaceApi.Trustsec(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringInterfaceApi.Trustsec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Trustsec`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringInterfaceApi.Trustsec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrustsecRequest struct via the builder pattern


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

