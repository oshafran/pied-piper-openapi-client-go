# \MonitoringDPIOnDemandTroubleshootingApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQueueEntry**](MonitoringDPIOnDemandTroubleshootingApi.md#CreateQueueEntry) | **Post** /statistics/on-demand/queue | 
[**DeleteQueueEntry**](MonitoringDPIOnDemandTroubleshootingApi.md#DeleteQueueEntry) | **Delete** /statistics/on-demand/queue/{entryId} | 
[**GetQueueEntries**](MonitoringDPIOnDemandTroubleshootingApi.md#GetQueueEntries) | **Get** /statistics/on-demand/queue | 
[**GetQueueProperties**](MonitoringDPIOnDemandTroubleshootingApi.md#GetQueueProperties) | **Get** /statistics/on-demand/queue/properties | 
[**UpdateQueueEntry**](MonitoringDPIOnDemandTroubleshootingApi.md#UpdateQueueEntry) | **Put** /statistics/on-demand/queue/{entryId} | 



## CreateQueueEntry

> map[string]interface{} CreateQueueEntry(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | On-demand queue entry (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDPIOnDemandTroubleshootingApi.CreateQueueEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDPIOnDemandTroubleshootingApi.CreateQueueEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQueueEntry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDPIOnDemandTroubleshootingApi.CreateQueueEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueueEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | On-demand queue entry | 

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


## DeleteQueueEntry

> DeleteQueueEntry(ctx, entryId).Execute()





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
    entryId := "entryId_example" // string | Entry Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDPIOnDemandTroubleshootingApi.DeleteQueueEntry(context.Background(), entryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDPIOnDemandTroubleshootingApi.DeleteQueueEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | Entry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetQueueEntries

> map[string]interface{} GetQueueEntries(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDPIOnDemandTroubleshootingApi.GetQueueEntries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDPIOnDemandTroubleshootingApi.GetQueueEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueEntries`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDPIOnDemandTroubleshootingApi.GetQueueEntries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueEntriesRequest struct via the builder pattern


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


## GetQueueProperties

> map[string]interface{} GetQueueProperties(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringDPIOnDemandTroubleshootingApi.GetQueueProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDPIOnDemandTroubleshootingApi.GetQueueProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueueProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDPIOnDemandTroubleshootingApi.GetQueueProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueuePropertiesRequest struct via the builder pattern


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


## UpdateQueueEntry

> map[string]interface{} UpdateQueueEntry(ctx, entryId).Body(body).Execute()





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
    entryId := "entryId_example" // string | Entry Id
    body := map[string]interface{}{ ... } // map[string]interface{} | On-demand queue entry (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringDPIOnDemandTroubleshootingApi.UpdateQueueEntry(context.Background(), entryId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringDPIOnDemandTroubleshootingApi.UpdateQueueEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQueueEntry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringDPIOnDemandTroubleshootingApi.UpdateQueueEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | Entry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQueueEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | On-demand queue entry | 

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

