# \MonitoringStatsDownloadApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Download1**](MonitoringStatsDownloadApi.md#Download1) | **Get** /statistics/download/{processType}/file/{fileType}/{queue}/{deviceIp}/{token}/{fileName} | 
[**DownloadList**](MonitoringStatsDownloadApi.md#DownloadList) | **Post** /statistics/download/{processType}/filelist | 
[**FetchList**](MonitoringStatsDownloadApi.md#FetchList) | **Get** /statistics/download/{processType}/fetchvManageList | 



## Download1

> map[string]interface{} Download1(ctx, processType, fileType, queue, deviceIp, token, fileName).Execute()





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
    processType := "processType_example" // string | Process type
    fileType := "fileType_example" // string | File type
    queue := "queue_example" // string | Queue name
    deviceIp := "deviceIp_example" // string | Device IP
    token := "token_example" // string | Token
    fileName := "fileName_example" // string | File name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatsDownloadApi.Download1(context.Background(), processType, fileType, queue, deviceIp, token, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsDownloadApi.Download1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringStatsDownloadApi.Download1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processType** | **string** | Process type | 
**fileType** | **string** | File type | 
**queue** | **string** | Queue name | 
**deviceIp** | **string** | Device IP | 
**token** | **string** | Token | 
**fileName** | **string** | File name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownload1Request struct via the builder pattern


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


## DownloadList

> DownloadList(ctx, processType).RequestBody(requestBody).Execute()





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
    processType := "processType_example" // string | Possible types are: remoteprocessing, dr
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatsDownloadApi.DownloadList(context.Background(), processType).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsDownloadApi.DownloadList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processType** | **string** | Possible types are: remoteprocessing, dr | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**map[string]GetO365PreferredPathFromVAnalyticsRequestValue**](GetO365PreferredPathFromVAnalyticsRequestValue.md) |  | 

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


## FetchList

> FetchList(ctx, processType).Execute()



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
    processType := "processType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringStatsDownloadApi.FetchList(context.Background(), processType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringStatsDownloadApi.FetchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**processType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

