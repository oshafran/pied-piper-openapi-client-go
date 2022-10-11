# \ScheduledBackupRestoreApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSchduledBackup**](ScheduledBackupRestoreApi.md#DeleteSchduledBackup) | **Delete** /backup/backupinfo | 
[**DeleteSchedule**](ScheduledBackupRestoreApi.md#DeleteSchedule) | **Delete** /schedule/{schedulerId} | 
[**DownloadBackupFile**](ScheduledBackupRestoreApi.md#DownloadBackupFile) | **Get** /backup/download/{path} | 
[**ExportBackup**](ScheduledBackupRestoreApi.md#ExportBackup) | **Post** /backup/export | 
[**GetLocalBackupInfo**](ScheduledBackupRestoreApi.md#GetLocalBackupInfo) | **Get** /backup/backupinfo/{localBackupInfoId} | 
[**GetScheduleRecordForBackup**](ScheduledBackupRestoreApi.md#GetScheduleRecordForBackup) | **Get** /schedule/{schedulerId} | 
[**ImportScheduledBackup**](ScheduledBackupRestoreApi.md#ImportScheduledBackup) | **Post** /restore/import | 
[**ListBackup**](ScheduledBackupRestoreApi.md#ListBackup) | **Get** /backup/list | 
[**ListSchedules**](ScheduledBackupRestoreApi.md#ListSchedules) | **Get** /schedule/list | 
[**RemoteImportBackup**](ScheduledBackupRestoreApi.md#RemoteImportBackup) | **Post** /restore/remoteimport | 
[**ScheduleBackup**](ScheduledBackupRestoreApi.md#ScheduleBackup) | **Post** /schedule/create | 



## DeleteSchduledBackup

> map[string]interface{} DeleteSchduledBackup(ctx).TaskId(taskId).BackupInfoId(backupInfoId).Execute()





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
    taskId := "taskId_example" // string | task id (optional)
    backupInfoId := "backupInfoId_example" // string | Local Backup Info Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.DeleteSchduledBackup(context.Background()).TaskId(taskId).BackupInfoId(backupInfoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.DeleteSchduledBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchduledBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.DeleteSchduledBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSchduledBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **string** | task id | 
 **backupInfoId** | **string** | Local Backup Info Id | 

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


## DeleteSchedule

> map[string]interface{} DeleteSchedule(ctx, schedulerId).Execute()





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
    schedulerId := "schedulerId_example" // string | scheduler id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.DeleteSchedule(context.Background(), schedulerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.DeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.DeleteSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schedulerId** | **string** | scheduler id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


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


## DownloadBackupFile

> map[string]interface{} DownloadBackupFile(ctx, path).Execute()





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
    resp, r, err := apiClient.ScheduledBackupRestoreApi.DownloadBackupFile(context.Background(), path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.DownloadBackupFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadBackupFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.DownloadBackupFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | File path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBackupFileRequest struct via the builder pattern


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


## ExportBackup

> map[string]interface{} ExportBackup(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | backup request information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.ExportBackup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.ExportBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.ExportBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | backup request information | 

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


## GetLocalBackupInfo

> map[string]interface{} GetLocalBackupInfo(ctx, localBackupInfoId).Execute()





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
    localBackupInfoId := "localBackupInfoId_example" // string | localBackupInfo Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.GetLocalBackupInfo(context.Background(), localBackupInfoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.GetLocalBackupInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalBackupInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.GetLocalBackupInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localBackupInfoId** | **string** | localBackupInfo Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalBackupInfoRequest struct via the builder pattern


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


## GetScheduleRecordForBackup

> map[string]interface{} GetScheduleRecordForBackup(ctx, schedulerId).Execute()





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
    schedulerId := "schedulerId_example" // string | scheduler id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.GetScheduleRecordForBackup(context.Background(), schedulerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.GetScheduleRecordForBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduleRecordForBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.GetScheduleRecordForBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schedulerId** | **string** | scheduler id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRecordForBackupRequest struct via the builder pattern


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


## ImportScheduledBackup

> map[string]interface{} ImportScheduledBackup(ctx).Execute()





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
    resp, r, err := apiClient.ScheduledBackupRestoreApi.ImportScheduledBackup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.ImportScheduledBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportScheduledBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.ImportScheduledBackup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiImportScheduledBackupRequest struct via the builder pattern


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


## ListBackup

> map[string]interface{} ListBackup(ctx).Size(size).Execute()





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
    size := "size_example" // string | size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.ListBackup(context.Background()).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.ListBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.ListBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **string** | size | 

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


## ListSchedules

> map[string]interface{} ListSchedules(ctx).Limit(limit).Execute()





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
    limit := int64(789) // int64 | size (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.ListSchedules(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.ListSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.ListSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | size | [default to 100]

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


## RemoteImportBackup

> map[string]interface{} RemoteImportBackup(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.RemoteImportBackup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.RemoteImportBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoteImportBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.RemoteImportBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoteImportBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## ScheduleBackup

> map[string]interface{} ScheduleBackup(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | schedule request information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduledBackupRestoreApi.ScheduleBackup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledBackupRestoreApi.ScheduleBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduledBackupRestoreApi.ScheduleBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | schedule request information | 

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

