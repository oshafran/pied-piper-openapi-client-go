# \MonitoringColocationClusterApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterConfigByClusterId**](MonitoringColocationClusterApi.md#GetClusterConfigByClusterId) | **Get** /colocation/monitor/cluster/config | 
[**GetClusterDetailsByClusterId**](MonitoringColocationClusterApi.md#GetClusterDetailsByClusterId) | **Get** /colocation/monitor/cluster | 
[**GetClusterPortMappingByClusterId**](MonitoringColocationClusterApi.md#GetClusterPortMappingByClusterId) | **Get** /colocation/monitor/cluster/portView | 
[**GetDeviceDetailByDeviceId**](MonitoringColocationClusterApi.md#GetDeviceDetailByDeviceId) | **Get** /colocation/monitor/device | 
[**GetPNFConfig**](MonitoringColocationClusterApi.md#GetPNFConfig) | **Get** /colocation/monitor/pnf/configuration | 
[**GetServiceChainDetails**](MonitoringColocationClusterApi.md#GetServiceChainDetails) | **Get** /colocation/monitor/servicechain | 
[**GetServiceGroupByClusterId**](MonitoringColocationClusterApi.md#GetServiceGroupByClusterId) | **Get** /colocation/monitor/servicegroup | 
[**GetSystemStatusByDeviceId**](MonitoringColocationClusterApi.md#GetSystemStatusByDeviceId) | **Get** /colocation/monitor/device/system | 
[**GetVNFAlarmCount**](MonitoringColocationClusterApi.md#GetVNFAlarmCount) | **Get** /colocation/monitor/vnf/alarms/count | 
[**GetVNFEventsCountDetail**](MonitoringColocationClusterApi.md#GetVNFEventsCountDetail) | **Get** /colocation/monitor/vnf/alarms | 
[**GetVNFEventsDetail**](MonitoringColocationClusterApi.md#GetVNFEventsDetail) | **Get** /colocation/monitor/vnf/events | 
[**GetVNFInterfaceDetail**](MonitoringColocationClusterApi.md#GetVNFInterfaceDetail) | **Get** /colocation/monitor/vnf/interface | 
[**GetpnfDetails**](MonitoringColocationClusterApi.md#GetpnfDetails) | **Get** /colocation/monitor/pnf | 
[**GetvnfByDeviceId**](MonitoringColocationClusterApi.md#GetvnfByDeviceId) | **Get** /colocation/monitor/device/vnf | 
[**GetvnfDetails**](MonitoringColocationClusterApi.md#GetvnfDetails) | **Get** /colocation/monitor/vnf | 
[**ListNetworkFunctionMap**](MonitoringColocationClusterApi.md#ListNetworkFunctionMap) | **Get** /colocation/monitor/networkfunction/listmap | 
[**VnfActions**](MonitoringColocationClusterApi.md#VnfActions) | **Post** /colocation/monitor/vnf/action | 



## GetClusterConfigByClusterId

> []map[string]interface{} GetClusterConfigByClusterId(ctx).ClusterId(clusterId).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetClusterConfigByClusterId(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetClusterConfigByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterConfigByClusterId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetClusterConfigByClusterId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterConfigByClusterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

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


## GetClusterDetailsByClusterId

> map[string]interface{} GetClusterDetailsByClusterId(ctx).ClusterId(clusterId).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetClusterDetailsByClusterId(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetClusterDetailsByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterDetailsByClusterId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetClusterDetailsByClusterId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterDetailsByClusterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

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


## GetClusterPortMappingByClusterId

> []map[string]interface{} GetClusterPortMappingByClusterId(ctx).ClusterId(clusterId).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetClusterPortMappingByClusterId(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetClusterPortMappingByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterPortMappingByClusterId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetClusterPortMappingByClusterId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPortMappingByClusterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

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


## GetDeviceDetailByDeviceId

> map[string]interface{} GetDeviceDetailByDeviceId(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetDeviceDetailByDeviceId(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetDeviceDetailByDeviceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceDetailByDeviceId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetDeviceDetailByDeviceId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDetailByDeviceIdRequest struct via the builder pattern


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


## GetPNFConfig

> []map[string]interface{} GetPNFConfig(ctx).PnfSerialNumber(pnfSerialNumber).ClusterId(clusterId).Execute()





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
    pnfSerialNumber := "pnfSerialNumber_example" // string | PNF serial number (optional)
    clusterId := "clusterId_example" // string | Cluster Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetPNFConfig(context.Background()).PnfSerialNumber(pnfSerialNumber).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetPNFConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPNFConfig`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetPNFConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPNFConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pnfSerialNumber** | **string** | PNF serial number | 
 **clusterId** | **string** | Cluster Id | 

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


## GetServiceChainDetails

> map[string]interface{} GetServiceChainDetails(ctx).ClusterId(clusterId).UserGroupName(userGroupName).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id (optional)
    userGroupName := "userGroupName_example" // string | UserGroup Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetServiceChainDetails(context.Background()).ClusterId(clusterId).UserGroupName(userGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetServiceChainDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceChainDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetServiceChainDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceChainDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 
 **userGroupName** | **string** | UserGroup Name | 

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


## GetServiceGroupByClusterId

> []map[string]interface{} GetServiceGroupByClusterId(ctx).ClusterId(clusterId).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetServiceGroupByClusterId(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetServiceGroupByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceGroupByClusterId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetServiceGroupByClusterId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceGroupByClusterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

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


## GetSystemStatusByDeviceId

> []map[string]interface{} GetSystemStatusByDeviceId(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetSystemStatusByDeviceId(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetSystemStatusByDeviceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemStatusByDeviceId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetSystemStatusByDeviceId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatusByDeviceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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


## GetVNFAlarmCount

> []map[string]interface{} GetVNFAlarmCount(ctx).UserGroup(userGroup).Execute()





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
    userGroup := "userGroup_example" // string | user group name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetVNFAlarmCount(context.Background()).UserGroup(userGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetVNFAlarmCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVNFAlarmCount`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetVNFAlarmCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVNFAlarmCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userGroup** | **string** | user group name | 

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


## GetVNFEventsCountDetail

> []map[string]interface{} GetVNFEventsCountDetail(ctx).UserGroup(userGroup).Execute()





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
    userGroup := "userGroup_example" // string | user group name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetVNFEventsCountDetail(context.Background()).UserGroup(userGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetVNFEventsCountDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVNFEventsCountDetail`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetVNFEventsCountDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVNFEventsCountDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userGroup** | **string** | user group name | 

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


## GetVNFEventsDetail

> []map[string]interface{} GetVNFEventsDetail(ctx).VnfName(vnfName).Execute()





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
    vnfName := "vnfName_example" // string | VNF name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetVNFEventsDetail(context.Background()).VnfName(vnfName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetVNFEventsDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVNFEventsDetail`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetVNFEventsDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVNFEventsDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vnfName** | **string** | VNF name | 

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


## GetVNFInterfaceDetail

> []map[string]interface{} GetVNFInterfaceDetail(ctx).VnfName(vnfName).DeviceIp(deviceIp).DeviceClass(deviceClass).Execute()





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
    vnfName := "vnfName_example" // string | VNF name (optional)
    deviceIp := "deviceIp_example" // string | Device IP (optional)
    deviceClass := "deviceClass_example" // string | Device class (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetVNFInterfaceDetail(context.Background()).VnfName(vnfName).DeviceIp(deviceIp).DeviceClass(deviceClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetVNFInterfaceDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVNFInterfaceDetail`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetVNFInterfaceDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVNFInterfaceDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vnfName** | **string** | VNF name | 
 **deviceIp** | **string** | Device IP | 
 **deviceClass** | **string** | Device class | 

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


## GetpnfDetails

> []map[string]interface{} GetpnfDetails(ctx).ClusterId(clusterId).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetpnfDetails(context.Background()).ClusterId(clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetpnfDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetpnfDetails`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetpnfDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetpnfDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 

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


## GetvnfByDeviceId

> []map[string]interface{} GetvnfByDeviceId(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetvnfByDeviceId(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetvnfByDeviceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetvnfByDeviceId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetvnfByDeviceId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetvnfByDeviceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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


## GetvnfDetails

> map[string]interface{} GetvnfDetails(ctx).ClusterId(clusterId).UserGroupName(userGroupName).VnfName(vnfName).Execute()





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
    clusterId := "clusterId_example" // string | Cluster Id (optional)
    userGroupName := "userGroupName_example" // string | UserGroup Name (optional)
    vnfName := "vnfName_example" // string | VNF Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.GetvnfDetails(context.Background()).ClusterId(clusterId).UserGroupName(userGroupName).VnfName(vnfName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.GetvnfDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetvnfDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.GetvnfDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetvnfDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Cluster Id | 
 **userGroupName** | **string** | UserGroup Name | 
 **vnfName** | **string** | VNF Name | 

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


## ListNetworkFunctionMap

> []map[string]interface{} ListNetworkFunctionMap(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringColocationClusterApi.ListNetworkFunctionMap(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.ListNetworkFunctionMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkFunctionMap`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringColocationClusterApi.ListNetworkFunctionMap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkFunctionMapRequest struct via the builder pattern


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


## VnfActions

> VnfActions(ctx).VmName(vmName).DeviceId(deviceId).Action(action).Execute()





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
    vmName := "vmName_example" // string | VM Name (optional)
    deviceId := "deviceId_example" // string | Device Id (optional)
    action := "action_example" // string | Action (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringColocationClusterApi.VnfActions(context.Background()).VmName(vmName).DeviceId(deviceId).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringColocationClusterApi.VnfActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVnfActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vmName** | **string** | VM Name | 
 **deviceId** | **string** | Device Id | 
 **action** | **string** | Action | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

