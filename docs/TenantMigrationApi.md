# \TenantMigrationApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadTenantData**](TenantMigrationApi.md#DownloadTenantData) | **Get** /tenantmigration/download/{path} | 
[**ExportTenantData**](TenantMigrationApi.md#ExportTenantData) | **Post** /tenantmigration/export | 
[**GetMigrationToken**](TenantMigrationApi.md#GetMigrationToken) | **Get** /tenantmigration/migrationToken | 
[**ImportTenantData**](TenantMigrationApi.md#ImportTenantData) | **Post** /tenantmigration/import | 
[**MigrateNetwork**](TenantMigrationApi.md#MigrateNetwork) | **Post** /tenantmigration/networkMigration | 
[**ReTriggerNetworkMigration**](TenantMigrationApi.md#ReTriggerNetworkMigration) | **Get** /tenantmigration/networkMigration | 



## DownloadTenantData

> map[string]interface{} DownloadTenantData(ctx, path).Execute()





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
    path := "path_example" // string | File path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantMigrationApi.DownloadTenantData(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMigrationApi.DownloadTenantData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadTenantData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantMigrationApi.DownloadTenantData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | File path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTenantDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTenantData

> map[string]interface{} ExportTenantData(ctx).CreateTenantModel(createTenantModel).Execute()





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
    createTenantModel := *openapiclient.NewCreateTenantModel() // CreateTenantModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantMigrationApi.ExportTenantData(context.Background()).CreateTenantModel(createTenantModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMigrationApi.ExportTenantData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTenantData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantMigrationApi.ExportTenantData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTenantDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTenantModel** | [**CreateTenantModel**](CreateTenantModel.md) |  | 

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


## GetMigrationToken

> map[string]interface{} GetMigrationToken(ctx).MigrationId(migrationId).Execute()





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
    migrationId := "migrationId_example" // string | Migration Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantMigrationApi.GetMigrationToken(context.Background()).MigrationId(migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMigrationApi.GetMigrationToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMigrationToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantMigrationApi.GetMigrationToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMigrationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationId** | **string** | Migration Id | 

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


## ImportTenantData

> map[string]interface{} ImportTenantData(ctx).Execute()





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
    resp, r, err := apiClient.TenantMigrationApi.ImportTenantData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMigrationApi.ImportTenantData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportTenantData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantMigrationApi.ImportTenantData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiImportTenantDataRequest struct via the builder pattern


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


## MigrateNetwork

> map[string]interface{} MigrateNetwork(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Network migration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantMigrationApi.MigrateNetwork(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMigrationApi.MigrateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrateNetwork`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantMigrationApi.MigrateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Network migration | 

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


## ReTriggerNetworkMigration

> ReTriggerNetworkMigration(ctx).Execute()





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
    resp, r, err := apiClient.TenantMigrationApi.ReTriggerNetworkMigration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMigrationApi.ReTriggerNetworkMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReTriggerNetworkMigrationRequest struct via the builder pattern


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

