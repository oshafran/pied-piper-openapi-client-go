# \PartnerDNACSDAAPIApi

All URIs are relative to *http://$VMANAGE_EXTERNAL_IP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSDAConfig**](PartnerDNACSDAAPIApi.md#CreateSDAConfig) | **Post** /partner/dnac/sda/config/{partnerId} | 
[**CreateSDAConfigFromNetconf**](PartnerDNACSDAAPIApi.md#CreateSDAConfigFromNetconf) | **Post** /partner/dnac/sda/netconfconfig/{partnerId} | 
[**GetDeviceDetails**](PartnerDNACSDAAPIApi.md#GetDeviceDetails) | **Get** /partner/dnac/sda/device/{partnerId}/{uuid} | 
[**GetOverlayVPNList**](PartnerDNACSDAAPIApi.md#GetOverlayVPNList) | **Get** /partner/dnac/sda/vpn | 
[**GetSDAEnabledDevices**](PartnerDNACSDAAPIApi.md#GetSDAEnabledDevices) | **Get** /partner/dnac/sda/device/{partnerId} | 
[**GetSitesForPartner**](PartnerDNACSDAAPIApi.md#GetSitesForPartner) | **Get** /partner/dnac/sda/site/{partnerId} | 



## CreateSDAConfig

> map[string]interface{} CreateSDAConfig(ctx, partnerId).Body(body).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Device SDA configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerDNACSDAAPIApi.CreateSDAConfig(context.Background(), partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDNACSDAAPIApi.CreateSDAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSDAConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerDNACSDAAPIApi.CreateSDAConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSDAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device SDA configuration | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSDAConfigFromNetconf

> map[string]interface{} CreateSDAConfigFromNetconf(ctx, partnerId).Body(body).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Device SDA configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerDNACSDAAPIApi.CreateSDAConfigFromNetconf(context.Background(), partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDNACSDAAPIApi.CreateSDAConfigFromNetconf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSDAConfigFromNetconf`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerDNACSDAAPIApi.CreateSDAConfigFromNetconf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSDAConfigFromNetconfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device SDA configuration | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceDetails

> []map[string]interface{} GetDeviceDetails(ctx, partnerId, uuid).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id
    uuid := "uuid_example" // string | Device uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerDNACSDAAPIApi.GetDeviceDetails(context.Background(), partnerId, uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDNACSDAAPIApi.GetDeviceDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceDetails`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerDNACSDAAPIApi.GetDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetOverlayVPNList

> []map[string]interface{} GetOverlayVPNList(ctx).Execute()





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
    resp, r, err := apiClient.PartnerDNACSDAAPIApi.GetOverlayVPNList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDNACSDAAPIApi.GetOverlayVPNList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOverlayVPNList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerDNACSDAAPIApi.GetOverlayVPNList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverlayVPNListRequest struct via the builder pattern


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


## GetSDAEnabledDevices

> []map[string]interface{} GetSDAEnabledDevices(ctx, partnerId).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerDNACSDAAPIApi.GetSDAEnabledDevices(context.Background(), partnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDNACSDAAPIApi.GetSDAEnabledDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSDAEnabledDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerDNACSDAAPIApi.GetSDAEnabledDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSDAEnabledDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSitesForPartner

> []map[string]interface{} GetSitesForPartner(ctx, partnerId).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerDNACSDAAPIApi.GetSitesForPartner(context.Background(), partnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDNACSDAAPIApi.GetSitesForPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSitesForPartner`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerDNACSDAAPIApi.GetSitesForPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesForPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

