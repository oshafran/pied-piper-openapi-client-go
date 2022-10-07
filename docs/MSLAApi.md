# \MSLAApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTemplate**](MSLAApi.md#GetAllTemplate) | **Get** /msla/template | 
[**GetLicenseAndDeviceCount**](MSLAApi.md#GetLicenseAndDeviceCount) | **Get** /msla/monitoring/licensedDeviceCount | 
[**GetLicenseAndDeviceCount1**](MSLAApi.md#GetLicenseAndDeviceCount1) | **Get** /msla/monitoring/licensedDistributionDetails | 
[**GetMSLADevices**](MSLAApi.md#GetMSLADevices) | **Get** /msla/devices | 
[**GetPackagingDistributionDetails**](MSLAApi.md#GetPackagingDistributionDetails) | **Get** /msla/monitoring/packagingDistributionDetails | 
[**GetSubscriptions**](MSLAApi.md#GetSubscriptions) | **Get** /msla/va/License | 
[**GetSubscriptions1**](MSLAApi.md#GetSubscriptions1) | **Post** /msla/template/licenses | 
[**SyncLicenses2**](MSLAApi.md#SyncLicenses2) | **Post** /msla/licenses/sync | 



## GetAllTemplate

> map[string]interface{} GetAllTemplate(ctx).Execute()





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
    resp, r, err := apiClient.MSLAApi.GetAllTemplate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetAllTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetAllTemplate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTemplateRequest struct via the builder pattern


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


## GetLicenseAndDeviceCount

> map[string]interface{} GetLicenseAndDeviceCount(ctx).Execute()





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
    resp, r, err := apiClient.MSLAApi.GetLicenseAndDeviceCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetLicenseAndDeviceCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseAndDeviceCount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetLicenseAndDeviceCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseAndDeviceCountRequest struct via the builder pattern


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


## GetLicenseAndDeviceCount1

> map[string]interface{} GetLicenseAndDeviceCount1(ctx).Execute()





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
    resp, r, err := apiClient.MSLAApi.GetLicenseAndDeviceCount1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetLicenseAndDeviceCount1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseAndDeviceCount1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetLicenseAndDeviceCount1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseAndDeviceCount1Request struct via the builder pattern


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


## GetMSLADevices

> map[string]interface{} GetMSLADevices(ctx).Execute()





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
    resp, r, err := apiClient.MSLAApi.GetMSLADevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetMSLADevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMSLADevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetMSLADevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMSLADevicesRequest struct via the builder pattern


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


## GetPackagingDistributionDetails

> map[string]interface{} GetPackagingDistributionDetails(ctx).Execute()





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
    resp, r, err := apiClient.MSLAApi.GetPackagingDistributionDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetPackagingDistributionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPackagingDistributionDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetPackagingDistributionDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagingDistributionDetailsRequest struct via the builder pattern


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


## GetSubscriptions

> map[string]interface{} GetSubscriptions(ctx).VirtualAccountId(virtualAccountId).LicenseType(licenseType).Execute()





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
    virtualAccountId := "virtualAccountId_example" // string | virtual_account_id (optional)
    licenseType := "licenseType_example" // string | licenseType (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MSLAApi.GetSubscriptions(context.Background()).VirtualAccountId(virtualAccountId).LicenseType(licenseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAccountId** | **string** | virtual_account_id | 
 **licenseType** | **string** | licenseType | 

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


## GetSubscriptions1

> map[string]interface{} GetSubscriptions1(ctx).RequestBody(requestBody).Execute()





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
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MSLAApi.GetSubscriptions1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.GetSubscriptions1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptions1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.GetSubscriptions1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptions1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]GetO365PreferredPathFromVAnalyticsRequestValue**](GetO365PreferredPathFromVAnalyticsRequestValue.md) |  | 

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


## SyncLicenses2

> map[string]interface{} SyncLicenses2(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Sync license (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MSLAApi.SyncLicenses2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSLAApi.SyncLicenses2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncLicenses2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MSLAApi.SyncLicenses2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncLicenses2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Sync license | 

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

