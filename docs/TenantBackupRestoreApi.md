# \TenantBackupRestoreApi

All URIs are relative to *http://$VMANAGE_EXTERNAL_IP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTenantBackup**](TenantBackupRestoreApi.md#DeleteTenantBackup) | **Delete** /tenantbackup/delete | 
[**DownloadExistingBackupFile**](TenantBackupRestoreApi.md#DownloadExistingBackupFile) | **Get** /tenantbackup/download/{path} | 
[**ExportTenantBackup**](TenantBackupRestoreApi.md#ExportTenantBackup) | **Get** /tenantbackup/export | 
[**ImportTenantBackup**](TenantBackupRestoreApi.md#ImportTenantBackup) | **Post** /tenantbackup/import | 
[**ListTenantBackup**](TenantBackupRestoreApi.md#ListTenantBackup) | **Get** /tenantbackup/list | 



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
    resp, r, err := apiClient.TenantBackupRestoreApi.DeleteTenantBackup(context.Background()).FileName(fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBackupRestoreApi.DeleteTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantBackupRestoreApi.DeleteTenantBackup`: %v\n", resp)
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
    resp, r, err := apiClient.TenantBackupRestoreApi.DownloadExistingBackupFile(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBackupRestoreApi.DownloadExistingBackupFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadExistingBackupFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantBackupRestoreApi.DownloadExistingBackupFile`: %v\n", resp)
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
    resp, r, err := apiClient.TenantBackupRestoreApi.ExportTenantBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBackupRestoreApi.ExportTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantBackupRestoreApi.ExportTenantBackup`: %v\n", resp)
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
    resp, r, err := apiClient.TenantBackupRestoreApi.ImportTenantBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBackupRestoreApi.ImportTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantBackupRestoreApi.ImportTenantBackup`: %v\n", resp)
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
    resp, r, err := apiClient.TenantBackupRestoreApi.ListTenantBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBackupRestoreApi.ListTenantBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TenantBackupRestoreApi.ListTenantBackup`: %v\n", resp)
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

