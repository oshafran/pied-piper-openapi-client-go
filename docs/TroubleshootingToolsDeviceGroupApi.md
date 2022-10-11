# \TroubleshootingToolsDeviceGroupApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeviceGroupList**](TroubleshootingToolsDeviceGroupApi.md#ListDeviceGroupList) | **Get** /group | 
[**ListDeviceGroups**](TroubleshootingToolsDeviceGroupApi.md#ListDeviceGroups) | **Get** /group/device | 
[**ListGroupDevices**](TroubleshootingToolsDeviceGroupApi.md#ListGroupDevices) | **Get** /group/devices | 
[**ListGroupDevicesForMap**](TroubleshootingToolsDeviceGroupApi.md#ListGroupDevicesForMap) | **Get** /group/map/devices | 
[**ListGroupLinksForMap**](TroubleshootingToolsDeviceGroupApi.md#ListGroupLinksForMap) | **Get** /group/map/devices/links | 



## ListDeviceGroupList

> []map[string]interface{} ListDeviceGroupList(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListDeviceGroupList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceGroupApi.ListDeviceGroupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceGroupList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceGroupApi.ListDeviceGroupList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceGroupListRequest struct via the builder pattern


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


## ListDeviceGroups

> []map[string]interface{} ListDeviceGroups(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListDeviceGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceGroupApi.ListDeviceGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceGroups`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceGroupApi.ListDeviceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceGroupsRequest struct via the builder pattern


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


## ListGroupDevices

> []map[string]interface{} ListGroupDevices(ctx).GroupId(groupId).Ssh(ssh).VpnId(vpnId).Execute()





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
    groupId := "groupId_example" // string | Group Id
    ssh := true // bool | SSH (optional) (default to false)
    vpnId := []openapiclient.VPNID{*openapiclient.NewVPNID()} // []VPNID | VPN Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListGroupDevices(context.Background()).GroupId(groupId).Ssh(ssh).VpnId(vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceGroupApi.ListGroupDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceGroupApi.ListGroupDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Group Id | 
 **ssh** | **bool** | SSH | [default to false]
 **vpnId** | [**[]VPNID**](VPNID.md) | VPN Id | 

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


## ListGroupDevicesForMap

> []map[string]interface{} ListGroupDevicesForMap(ctx).GroupId(groupId).VpnId(vpnId).Execute()





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
    groupId := "groupId_example" // string | Group Id
    vpnId := []openapiclient.VPNID{*openapiclient.NewVPNID()} // []VPNID | VPN Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListGroupDevicesForMap(context.Background()).GroupId(groupId).VpnId(vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceGroupApi.ListGroupDevicesForMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupDevicesForMap`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceGroupApi.ListGroupDevicesForMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupDevicesForMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Group Id | 
 **vpnId** | [**[]VPNID**](VPNID.md) | VPN Id | 

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


## ListGroupLinksForMap

> []map[string]interface{} ListGroupLinksForMap(ctx).GroupId(groupId).Execute()





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
    groupId := "groupId_example" // string | Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListGroupLinksForMap(context.Background()).GroupId(groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceGroupApi.ListGroupLinksForMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupLinksForMap`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceGroupApi.ListGroupLinksForMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupLinksForMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Group Id | 

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

