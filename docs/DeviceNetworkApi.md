# \DeviceNetworkApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkIssuesSummary**](DeviceNetworkApi.md#GetNetworkIssuesSummary) | **Get** /network/issues/summary | 
[**GetNetworkStatusSummary**](DeviceNetworkApi.md#GetNetworkStatusSummary) | **Get** /network/status | 
[**GetRebootCount**](DeviceNetworkApi.md#GetRebootCount) | **Get** /network/issues/rebootcount | 
[**GetVmanageControlStatus**](DeviceNetworkApi.md#GetVmanageControlStatus) | **Get** /network/connectionssummary | 



## GetNetworkIssuesSummary

> map[string]interface{} GetNetworkIssuesSummary(ctx).Execute()





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
    resp, r, err := apiClient.DeviceNetworkApi.GetNetworkIssuesSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceNetworkApi.GetNetworkIssuesSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkIssuesSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceNetworkApi.GetNetworkIssuesSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkIssuesSummaryRequest struct via the builder pattern


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


## GetNetworkStatusSummary

> []map[string]interface{} GetNetworkStatusSummary(ctx).Execute()





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
    resp, r, err := apiClient.DeviceNetworkApi.GetNetworkStatusSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceNetworkApi.GetNetworkStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkStatusSummary`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceNetworkApi.GetNetworkStatusSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkStatusSummaryRequest struct via the builder pattern


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


## GetRebootCount

> map[string]interface{} GetRebootCount(ctx).IsCached(isCached).Execute()





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
    isCached := true // bool | Is cached flag (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceNetworkApi.GetRebootCount(context.Background()).IsCached(isCached).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceNetworkApi.GetRebootCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRebootCount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceNetworkApi.GetRebootCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRebootCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isCached** | **bool** | Is cached flag | [default to false]

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


## GetVmanageControlStatus

> map[string]interface{} GetVmanageControlStatus(ctx).IsCached(isCached).VpnId(vpnId).Execute()





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
    isCached := true // bool | Is cached flag (optional) (default to false)
    vpnId := []openapiclient.VPNID{*openapiclient.NewVPNID()} // []VPNID | VPN Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceNetworkApi.GetVmanageControlStatus(context.Background()).IsCached(isCached).VpnId(vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceNetworkApi.GetVmanageControlStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVmanageControlStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceNetworkApi.GetVmanageControlStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVmanageControlStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isCached** | **bool** | Is cached flag | [default to false]
 **vpnId** | [**[]VPNID**](VPNID.md) | VPN Id | 

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

