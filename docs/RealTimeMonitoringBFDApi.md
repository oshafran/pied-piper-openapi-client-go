# \RealTimeMonitoringBFDApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBFDHistoryList**](RealTimeMonitoringBFDApi.md#CreateBFDHistoryList) | **Get** /device/bfd/history | 
[**CreateBFDLinkList**](RealTimeMonitoringBFDApi.md#CreateBFDLinkList) | **Get** /device/bfd/links | 
[**CreateBFDSessions**](RealTimeMonitoringBFDApi.md#CreateBFDSessions) | **Get** /device/bfd/sessions | 
[**CreateBFDSummary**](RealTimeMonitoringBFDApi.md#CreateBFDSummary) | **Get** /device/bfd/summary | 
[**CreateSyncedBFDSession**](RealTimeMonitoringBFDApi.md#CreateSyncedBFDSession) | **Get** /device/bfd/synced/sessions | 
[**CreateTLOCSummary**](RealTimeMonitoringBFDApi.md#CreateTLOCSummary) | **Get** /device/bfd/tloc | 
[**GetBFDSiteStateDetail**](RealTimeMonitoringBFDApi.md#GetBFDSiteStateDetail) | **Get** /device/bfd/sites/detail | 
[**GetBFDSitesSummary**](RealTimeMonitoringBFDApi.md#GetBFDSitesSummary) | **Get** /device/bfd/sites/summary | 
[**GetDeviceBFDStateSummary**](RealTimeMonitoringBFDApi.md#GetDeviceBFDStateSummary) | **Get** /device/bfd/state/device | 
[**GetDeviceBFDStateSummaryTloc**](RealTimeMonitoringBFDApi.md#GetDeviceBFDStateSummaryTloc) | **Get** /device/bfd/state/device/tloc | 
[**GetDeviceBFDStatus**](RealTimeMonitoringBFDApi.md#GetDeviceBFDStatus) | **Get** /device/bfd/status | 
[**GetDeviceBFDStatusSummary**](RealTimeMonitoringBFDApi.md#GetDeviceBFDStatusSummary) | **Get** /device/bfd/summary/device | 



## CreateBFDHistoryList

> []map[string]interface{} CreateBFDHistoryList(ctx).DeviceId(deviceId).SystemIp(systemIp).Color(color).Execute()





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
    systemIp := "systemIp_example" // string | System IP (optional)
    color := "color_example" // string | Remote color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.CreateBFDHistoryList(context.Background()).DeviceId(deviceId).SystemIp(systemIp).Color(color).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.CreateBFDHistoryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBFDHistoryList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.CreateBFDHistoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBFDHistoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **systemIp** | **string** | System IP | 
 **color** | **string** | Remote color | 

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


## CreateBFDLinkList

> []map[string]interface{} CreateBFDLinkList(ctx).State(state).Execute()





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
    state := "state_example" // string | Device state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.CreateBFDLinkList(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.CreateBFDLinkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBFDLinkList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.CreateBFDLinkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBFDLinkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Device state | 

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


## CreateBFDSessions

> []map[string]interface{} CreateBFDSessions(ctx).DeviceId(deviceId).SystemIp(systemIp).Color(color).LocalColor(localColor).RegionType(regionType).Execute()





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
    systemIp := "systemIp_example" // string | System IP (optional)
    color := "color_example" // string | Remote color (optional)
    localColor := "localColor_example" // string | Source color (optional)
    regionType := "regionType_example" // string | Region type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.CreateBFDSessions(context.Background()).DeviceId(deviceId).SystemIp(systemIp).Color(color).LocalColor(localColor).RegionType(regionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.CreateBFDSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBFDSessions`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.CreateBFDSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBFDSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **systemIp** | **string** | System IP | 
 **color** | **string** | Remote color | 
 **localColor** | **string** | Source color | 
 **regionType** | **string** | Region type | 

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


## CreateBFDSummary

> map[string]interface{} CreateBFDSummary(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.CreateBFDSummary(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.CreateBFDSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBFDSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.CreateBFDSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBFDSummaryRequest struct via the builder pattern


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


## CreateSyncedBFDSession

> []map[string]interface{} CreateSyncedBFDSession(ctx).DeviceId(deviceId).SystemIp(systemIp).Color(color).LocalColor(localColor).Execute()





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
    systemIp := "systemIp_example" // string | System IP (optional)
    color := "color_example" // string | Remote color (optional)
    localColor := "localColor_example" // string | Source color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.CreateSyncedBFDSession(context.Background()).DeviceId(deviceId).SystemIp(systemIp).Color(color).LocalColor(localColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.CreateSyncedBFDSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncedBFDSession`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.CreateSyncedBFDSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyncedBFDSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 
 **systemIp** | **string** | System IP | 
 **color** | **string** | Remote color | 
 **localColor** | **string** | Source color | 

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


## CreateTLOCSummary

> map[string]interface{} CreateTLOCSummary(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.CreateTLOCSummary(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.CreateTLOCSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTLOCSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.CreateTLOCSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTLOCSummaryRequest struct via the builder pattern


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


## GetBFDSiteStateDetail

> map[string]interface{} GetBFDSiteStateDetail(ctx).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.GetBFDSiteStateDetail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.GetBFDSiteStateDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBFDSiteStateDetail`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.GetBFDSiteStateDetail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBFDSiteStateDetailRequest struct via the builder pattern


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


## GetBFDSitesSummary

> map[string]interface{} GetBFDSitesSummary(ctx).VpnId(vpnId).IsCached(isCached).Execute()





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
    vpnId := []openapiclient.VPNID{*openapiclient.NewVPNID()} // []VPNID | Filter VPN
    isCached := true // bool | Flag for caching (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.GetBFDSitesSummary(context.Background()).VpnId(vpnId).IsCached(isCached).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.GetBFDSitesSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBFDSitesSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.GetBFDSitesSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBFDSitesSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnId** | [**[]VPNID**](VPNID.md) | Filter VPN | 
 **isCached** | **bool** | Flag for caching | [default to false]

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


## GetDeviceBFDStateSummary

> map[string]interface{} GetDeviceBFDStateSummary(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.GetDeviceBFDStateSummary(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.GetDeviceBFDStateSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceBFDStateSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.GetDeviceBFDStateSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceBFDStateSummaryRequest struct via the builder pattern


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


## GetDeviceBFDStateSummaryTloc

> map[string]interface{} GetDeviceBFDStateSummaryTloc(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.GetDeviceBFDStateSummaryTloc(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.GetDeviceBFDStateSummaryTloc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceBFDStateSummaryTloc`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.GetDeviceBFDStateSummaryTloc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceBFDStateSummaryTlocRequest struct via the builder pattern


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


## GetDeviceBFDStatus

> map[string]interface{} GetDeviceBFDStatus(ctx).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.GetDeviceBFDStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.GetDeviceBFDStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceBFDStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.GetDeviceBFDStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceBFDStatusRequest struct via the builder pattern


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


## GetDeviceBFDStatusSummary

> map[string]interface{} GetDeviceBFDStatusSummary(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.RealTimeMonitoringBFDApi.GetDeviceBFDStatusSummary(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringBFDApi.GetDeviceBFDStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceBFDStatusSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringBFDApi.GetDeviceBFDStatusSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceBFDStatusSummaryRequest struct via the builder pattern


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

