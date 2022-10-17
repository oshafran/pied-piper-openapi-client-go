# \ConfigurationSettingsApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalyticsDataFile**](ConfigurationSettingsApi.md#CreateAnalyticsDataFile) | **Post** /settings/configuration/analytics/dca | 
[**EditCertConfiguration**](ConfigurationSettingsApi.md#EditCertConfiguration) | **Put** /settings/configuration/certificate/{settingType} | 
[**EditConfiguration**](ConfigurationSettingsApi.md#EditConfiguration) | **Put** /settings/configuration/{settingType} | 
[**GetBanner**](ConfigurationSettingsApi.md#GetBanner) | **Get** /settings/banner | 
[**GetCertConfiguration**](ConfigurationSettingsApi.md#GetCertConfiguration) | **Get** /settings/configuration/certificate/{settingType} | 
[**GetConfigurationBySettingType**](ConfigurationSettingsApi.md#GetConfigurationBySettingType) | **Get** /settings/configuration/{settingType} | 
[**GetGoogleMapKey**](ConfigurationSettingsApi.md#GetGoogleMapKey) | **Get** /settings/configuration/googleMapKey | 
[**GetMaintenanceWindow**](ConfigurationSettingsApi.md#GetMaintenanceWindow) | **Get** /settings/configuration/maintenanceWindow | 
[**GetSessionTimout**](ConfigurationSettingsApi.md#GetSessionTimout) | **Get** /settings/clientSessionTimeout | 
[**NewCertConfiguration**](ConfigurationSettingsApi.md#NewCertConfiguration) | **Post** /settings/configuration/certificate/{settingType} | 
[**NewConfiguration**](ConfigurationSettingsApi.md#NewConfiguration) | **Post** /settings/configuration/{settingType} | 



## CreateAnalyticsDataFile

> map[string]interface{} CreateAnalyticsDataFile(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSettingsApi.CreateAnalyticsDataFile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.CreateAnalyticsDataFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnalyticsDataFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.CreateAnalyticsDataFile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyticsDataFileRequest struct via the builder pattern


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


## EditCertConfiguration

> map[string]interface{} EditCertConfiguration(ctx, settingType).Body(body).Execute()





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
    settingType := "settingType_example" // string | Setting type
    body := map[string]interface{}{ ... } // map[string]interface{} | Certificate config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSettingsApi.EditCertConfiguration(context.Background(), settingType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.EditCertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCertConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.EditCertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingType** | **string** | Setting type | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Certificate config | 

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


## EditConfiguration

> map[string]interface{} EditConfiguration(ctx, settingType).Body(body).Execute()





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
    settingType := "settingType_example" // string | Setting type
    body := map[string]interface{}{ ... } // map[string]interface{} | Configuration setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSettingsApi.EditConfiguration(context.Background(), settingType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.EditConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.EditConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingType** | **string** | Setting type | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Configuration setting | 

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


## GetBanner

> map[string]interface{} GetBanner(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSettingsApi.GetBanner(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.GetBanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBanner`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.GetBanner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannerRequest struct via the builder pattern


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


## GetCertConfiguration

> map[string]interface{} GetCertConfiguration(ctx, settingType).Execute()





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
    settingType := "settingType_example" // string | Setting type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSettingsApi.GetCertConfiguration(context.Background(), settingType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.GetCertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.GetCertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingType** | **string** | Setting type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetConfigurationBySettingType

> map[string]interface{} GetConfigurationBySettingType(ctx, settingType).Execute()





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
    settingType := "settingType_example" // string | Setting type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSettingsApi.GetConfigurationBySettingType(context.Background(), settingType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.GetConfigurationBySettingType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationBySettingType`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.GetConfigurationBySettingType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingType** | **string** | Setting type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationBySettingTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetGoogleMapKey

> map[string]interface{} GetGoogleMapKey(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSettingsApi.GetGoogleMapKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.GetGoogleMapKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGoogleMapKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.GetGoogleMapKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoogleMapKeyRequest struct via the builder pattern


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


## GetMaintenanceWindow

> map[string]interface{} GetMaintenanceWindow(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSettingsApi.GetMaintenanceWindow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.GetMaintenanceWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenanceWindow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.GetMaintenanceWindow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindowRequest struct via the builder pattern


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


## GetSessionTimout

> map[string]interface{} GetSessionTimout(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSettingsApi.GetSessionTimout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.GetSessionTimout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTimout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.GetSessionTimout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionTimoutRequest struct via the builder pattern


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


## NewCertConfiguration

> map[string]interface{} NewCertConfiguration(ctx, settingType).Body(body).Execute()





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
    settingType := "settingType_example" // string | Setting type
    body := map[string]interface{}{ ... } // map[string]interface{} | Certificate config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSettingsApi.NewCertConfiguration(context.Background(), settingType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.NewCertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewCertConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.NewCertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingType** | **string** | Setting type | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewCertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Certificate config | 

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


## NewConfiguration

> map[string]interface{} NewConfiguration(ctx, settingType).Body(body).Execute()





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
    settingType := "settingType_example" // string | Setting type
    body := map[string]interface{}{ ... } // map[string]interface{} | Configuration setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSettingsApi.NewConfiguration(context.Background(), settingType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSettingsApi.NewConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSettingsApi.NewConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingType** | **string** | Setting type | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Configuration setting | 

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

