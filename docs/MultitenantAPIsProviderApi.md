# \MultitenantAPIsProviderApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](MultitenantAPIsProviderApi.md#CreateTenant) | **Post** /tenant | 
[**CreateTenantAsync**](MultitenantAPIsProviderApi.md#CreateTenantAsync) | **Post** /tenant/async | 
[**CreateTenantAsyncBulk**](MultitenantAPIsProviderApi.md#CreateTenantAsyncBulk) | **Post** /tenant/bulk/async | 
[**DeleteTenant**](MultitenantAPIsProviderApi.md#DeleteTenant) | **Post** /tenant/{tenantId}/delete | 
[**DeleteTenantAsyncBulk**](MultitenantAPIsProviderApi.md#DeleteTenantAsyncBulk) | **Delete** /tenant/bulk/async | 
[**DeleteTenantBackup**](MultitenantAPIsProviderApi.md#DeleteTenantBackup) | **Delete** /tenantbackup/delete | 
[**DownloadExistingBackupFile**](MultitenantAPIsProviderApi.md#DownloadExistingBackupFile) | **Get** /tenantbackup/download/{path} | 
[**ExportTenantBackup**](MultitenantAPIsProviderApi.md#ExportTenantBackup) | **Get** /tenantbackup/export | 
[**ForceStatusCollection**](MultitenantAPIsProviderApi.md#ForceStatusCollection) | **Post** /tenantstatus/force | 
[**GetAllTenantStatuses**](MultitenantAPIsProviderApi.md#GetAllTenantStatuses) | **Get** /tenantstatus | 
[**GetAllTenants**](MultitenantAPIsProviderApi.md#GetAllTenants) | **Get** /tenant | 
[**GetTenant**](MultitenantAPIsProviderApi.md#GetTenant) | **Get** /tenant/{tenantId} | 
[**GetTenantHostingCapacityOnvSmarts**](MultitenantAPIsProviderApi.md#GetTenantHostingCapacityOnvSmarts) | **Get** /tenant/vsmart/capacity | 
[**GetTenantvSmartMapping**](MultitenantAPIsProviderApi.md#GetTenantvSmartMapping) | **Get** /tenant/vsmart | 
[**ImportTenantBackup**](MultitenantAPIsProviderApi.md#ImportTenantBackup) | **Post** /tenantbackup/import | 
[**ListTenantBackup**](MultitenantAPIsProviderApi.md#ListTenantBackup) | **Get** /tenantbackup/list | 
[**SwitchTenant**](MultitenantAPIsProviderApi.md#SwitchTenant) | **Post** /tenant/{tenantId}/switch | 
[**TenantvSmartMtMigrate**](MultitenantAPIsProviderApi.md#TenantvSmartMtMigrate) | **Post** /tenant/vsmart-mt/migrate | 
[**UpdateTenant**](MultitenantAPIsProviderApi.md#UpdateTenant) | **Put** /tenant/{tenantId} | 
[**VSessionId**](MultitenantAPIsProviderApi.md#VSessionId) | **Post** /tenant/{tenantId}/vsessionid | 



## CreateTenant

> map[string]interface{} CreateTenant(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Tenant model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.CreateTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.CreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Tenant model | 

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


## CreateTenantAsync

> map[string]interface{} CreateTenantAsync(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Tenant model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.CreateTenantAsync(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.CreateTenantAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenantAsync`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.CreateTenantAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Tenant model | 

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


## CreateTenantAsyncBulk

> map[string]interface{} CreateTenantAsyncBulk(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Tenant model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.CreateTenantAsyncBulk(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.CreateTenantAsyncBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenantAsyncBulk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.CreateTenantAsyncBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantAsyncBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Tenant model | 

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


## DeleteTenant

> DeleteTenant(ctx, tenantId).Body(body).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Tenant model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.DeleteTenant(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.DeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Tenant model | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantAsyncBulk

> map[string]interface{} DeleteTenantAsyncBulk(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Tenant model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.DeleteTenantAsyncBulk(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.DeleteTenantAsyncBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenantAsyncBulk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.DeleteTenantAsyncBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantAsyncBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Tenant model | 

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


## DeleteTenantBackup

> map[string]interface{} DeleteTenantBackup(ctx).FileName(fileName).Execute()





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
    fileName := "fileName_example" // string | File name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.DeleteTenantBackup(context.Background()).FileName(fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.DeleteTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.DeleteTenantBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **string** | File name | 

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


## DownloadExistingBackupFile

> map[string]interface{} DownloadExistingBackupFile(ctx, path).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.DownloadExistingBackupFile(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.DownloadExistingBackupFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadExistingBackupFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.DownloadExistingBackupFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | File path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadExistingBackupFileRequest struct via the builder pattern


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


## ExportTenantBackup

> map[string]interface{} ExportTenantBackup(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.ExportTenantBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.ExportTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.ExportTenantBackup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportTenantBackupRequest struct via the builder pattern


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


## ForceStatusCollection

> ForceStatusCollection(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.ForceStatusCollection(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.ForceStatusCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiForceStatusCollectionRequest struct via the builder pattern


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


## GetAllTenantStatuses

> []map[string]interface{} GetAllTenantStatuses(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.GetAllTenantStatuses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.GetAllTenantStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTenantStatuses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.GetAllTenantStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTenantStatusesRequest struct via the builder pattern


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


## GetAllTenants

> []map[string]interface{} GetAllTenants(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | List all tenants associated with a vSmart (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.GetAllTenants(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.GetAllTenants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTenants`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.GetAllTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | List all tenants associated with a vSmart | 

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


## GetTenant

> map[string]interface{} GetTenant(ctx, tenantId).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.GetTenant(context.Background(), tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.GetTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


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


## GetTenantHostingCapacityOnvSmarts

> []map[string]interface{} GetTenantHostingCapacityOnvSmarts(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.GetTenantHostingCapacityOnvSmarts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.GetTenantHostingCapacityOnvSmarts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantHostingCapacityOnvSmarts`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.GetTenantHostingCapacityOnvSmarts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantHostingCapacityOnvSmartsRequest struct via the builder pattern


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


## GetTenantvSmartMapping

> []map[string]interface{} GetTenantvSmartMapping(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.GetTenantvSmartMapping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.GetTenantvSmartMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantvSmartMapping`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.GetTenantvSmartMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantvSmartMappingRequest struct via the builder pattern


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


## ImportTenantBackup

> map[string]interface{} ImportTenantBackup(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.ImportTenantBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.ImportTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.ImportTenantBackup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiImportTenantBackupRequest struct via the builder pattern


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


## ListTenantBackup

> map[string]interface{} ListTenantBackup(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.ListTenantBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.ListTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.ListTenantBackup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTenantBackupRequest struct via the builder pattern


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


## SwitchTenant

> map[string]interface{} SwitchTenant(ctx, tenantId).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.SwitchTenant(context.Background(), tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.SwitchTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.SwitchTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchTenantRequest struct via the builder pattern


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


## TenantvSmartMtMigrate

> []map[string]interface{} TenantvSmartMtMigrate(ctx).Execute()





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
    resp, r, err := apiClient.MultitenantAPIsProviderApi.TenantvSmartMtMigrate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.TenantvSmartMtMigrate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantvSmartMtMigrate`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.TenantvSmartMtMigrate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantvSmartMtMigrateRequest struct via the builder pattern


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


## UpdateTenant

> map[string]interface{} UpdateTenant(ctx, tenantId).Body(body).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Tenant model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.UpdateTenant(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Tenant model | 

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


## VSessionId

> map[string]interface{} VSessionId(ctx, tenantId).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultitenantAPIsProviderApi.VSessionId(context.Background(), tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultitenantAPIsProviderApi.VSessionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VSessionId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultitenantAPIsProviderApi.VSessionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVSessionIdRequest struct via the builder pattern


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

