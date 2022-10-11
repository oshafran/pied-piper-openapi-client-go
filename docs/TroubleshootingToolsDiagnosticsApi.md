# \TroubleshootingToolsDiagnosticsApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSession**](TroubleshootingToolsDiagnosticsApi.md#ClearSession) | **Get** /stream/device/log/sessions/clear/{sessionId} | 
[**DisableDeviceLog**](TroubleshootingToolsDiagnosticsApi.md#DisableDeviceLog) | **Get** /stream/device/log/disable/{sessionId} | 
[**DisablePacketCaptureSession**](TroubleshootingToolsDiagnosticsApi.md#DisablePacketCaptureSession) | **Get** /stream/device/capture/disable/{sessionId} | 
[**DisableSpeedTestSession**](TroubleshootingToolsDiagnosticsApi.md#DisableSpeedTestSession) | **Get** /stream/device/speed/disable/{sessionId} | 
[**DownloadDebugLog**](TroubleshootingToolsDiagnosticsApi.md#DownloadDebugLog) | **Get** /stream/device/log/download/{sessionId} | 
[**DownloadFile**](TroubleshootingToolsDiagnosticsApi.md#DownloadFile) | **Get** /stream/device/capture/download/{sessionId} | 
[**ForceStopPcapSession**](TroubleshootingToolsDiagnosticsApi.md#ForceStopPcapSession) | **Get** /stream/device/capture/forcedisbale/{sessionId} | 
[**FormPostPacketCapture**](TroubleshootingToolsDiagnosticsApi.md#FormPostPacketCapture) | **Post** /stream/device/capture/{deviceUUID}/{sessionId} | 
[**GetAggregationDataByQuery27**](TroubleshootingToolsDiagnosticsApi.md#GetAggregationDataByQuery27) | **Get** /stream/device/nwpi/aggregation | 
[**GetAggregationDataByQuery28**](TroubleshootingToolsDiagnosticsApi.md#GetAggregationDataByQuery28) | **Get** /statistics/speedtest/aggregation | 
[**GetConcurrentData**](TroubleshootingToolsDiagnosticsApi.md#GetConcurrentData) | **Get** /stream/device/nwpi/concurrentData | 
[**GetConcurrentDomainData**](TroubleshootingToolsDiagnosticsApi.md#GetConcurrentDomainData) | **Get** /stream/device/nwpi/concurrentDomainData | 
[**GetCount29**](TroubleshootingToolsDiagnosticsApi.md#GetCount29) | **Get** /stream/device/nwpi/doccount | 
[**GetCount30**](TroubleshootingToolsDiagnosticsApi.md#GetCount30) | **Get** /statistics/speedtest/doccount | 
[**GetCountPost29**](TroubleshootingToolsDiagnosticsApi.md#GetCountPost29) | **Post** /stream/device/nwpi/doccount | 
[**GetCountPost30**](TroubleshootingToolsDiagnosticsApi.md#GetCountPost30) | **Post** /statistics/speedtest/doccount | 
[**GetDBSchema**](TroubleshootingToolsDiagnosticsApi.md#GetDBSchema) | **Get** /diagnostics/dbschema | 
[**GetDeviceLog**](TroubleshootingToolsDiagnosticsApi.md#GetDeviceLog) | **Get** /stream/device/log/{sessionId} | 
[**GetDomainMetric**](TroubleshootingToolsDiagnosticsApi.md#GetDomainMetric) | **Get** /stream/device/nwpi/domainMetric | 
[**GetFileDownloadStatus**](TroubleshootingToolsDiagnosticsApi.md#GetFileDownloadStatus) | **Get** /stream/device/capture/status/{sessionId} | 
[**GetFinalizedData**](TroubleshootingToolsDiagnosticsApi.md#GetFinalizedData) | **Get** /stream/device/nwpi/finalizedData | 
[**GetFinalizedDomainData**](TroubleshootingToolsDiagnosticsApi.md#GetFinalizedDomainData) | **Get** /stream/device/nwpi/finalizedDomainData | 
[**GetFlowDetail**](TroubleshootingToolsDiagnosticsApi.md#GetFlowDetail) | **Get** /stream/device/nwpi/flowDetail | 
[**GetFlowMetric**](TroubleshootingToolsDiagnosticsApi.md#GetFlowMetric) | **Get** /stream/device/nwpi/flowMetric | 
[**GetInterfaceBandwidth**](TroubleshootingToolsDiagnosticsApi.md#GetInterfaceBandwidth) | **Get** /stream/device/speed/interface/bandwidth | 
[**GetLogType**](TroubleshootingToolsDiagnosticsApi.md#GetLogType) | **Get** /stream/device/log/type | 
[**GetMonitorState**](TroubleshootingToolsDiagnosticsApi.md#GetMonitorState) | **Get** /stream/device/nwpi/getMonitorState | 
[**GetNwpiDscp**](TroubleshootingToolsDiagnosticsApi.md#GetNwpiDscp) | **Get** /stream/device/nwpi/nwpiDSCP | 
[**GetNwpiNbarAppGroup**](TroubleshootingToolsDiagnosticsApi.md#GetNwpiNbarAppGroup) | **Get** /stream/device/nwpi/nwpiNbarAppGroup | 
[**GetNwpiProtocol**](TroubleshootingToolsDiagnosticsApi.md#GetNwpiProtocol) | **Get** /stream/device/nwpi/nwpiProtocol | 
[**GetPacketFeatures**](TroubleshootingToolsDiagnosticsApi.md#GetPacketFeatures) | **Get** /stream/device/nwpi/packetFeatures | 
[**GetPostAggregationAppDataByQuery26**](TroubleshootingToolsDiagnosticsApi.md#GetPostAggregationAppDataByQuery26) | **Post** /stream/device/nwpi/app-agg/aggregation | 
[**GetPostAggregationAppDataByQuery27**](TroubleshootingToolsDiagnosticsApi.md#GetPostAggregationAppDataByQuery27) | **Post** /statistics/speedtest/app-agg/aggregation | 
[**GetPostAggregationDataByQuery27**](TroubleshootingToolsDiagnosticsApi.md#GetPostAggregationDataByQuery27) | **Post** /stream/device/nwpi/aggregation | 
[**GetPostAggregationDataByQuery28**](TroubleshootingToolsDiagnosticsApi.md#GetPostAggregationDataByQuery28) | **Post** /statistics/speedtest/aggregation | 
[**GetPostStatBulkRawData27**](TroubleshootingToolsDiagnosticsApi.md#GetPostStatBulkRawData27) | **Post** /stream/device/nwpi/page | 
[**GetPostStatBulkRawData28**](TroubleshootingToolsDiagnosticsApi.md#GetPostStatBulkRawData28) | **Post** /statistics/speedtest/page | 
[**GetSession**](TroubleshootingToolsDiagnosticsApi.md#GetSession) | **Post** /stream/device/speed | 
[**GetSessionInfoCapture**](TroubleshootingToolsDiagnosticsApi.md#GetSessionInfoCapture) | **Post** /stream/device/capture | 
[**GetSessionInfoLog**](TroubleshootingToolsDiagnosticsApi.md#GetSessionInfoLog) | **Post** /stream/device/log | 
[**GetSessions**](TroubleshootingToolsDiagnosticsApi.md#GetSessions) | **Get** /stream/device/log/sessions | 
[**GetSpeedTest**](TroubleshootingToolsDiagnosticsApi.md#GetSpeedTest) | **Get** /stream/device/speed/{sessionId} | 
[**GetSpeedTestStatus**](TroubleshootingToolsDiagnosticsApi.md#GetSpeedTestStatus) | **Get** /stream/device/speed/status/{sessionId} | 
[**GetStatBulkRawData27**](TroubleshootingToolsDiagnosticsApi.md#GetStatBulkRawData27) | **Get** /stream/device/nwpi/page | 
[**GetStatBulkRawData28**](TroubleshootingToolsDiagnosticsApi.md#GetStatBulkRawData28) | **Get** /statistics/speedtest/page | 
[**GetStatDataFields29**](TroubleshootingToolsDiagnosticsApi.md#GetStatDataFields29) | **Get** /stream/device/nwpi/fields | 
[**GetStatDataFields30**](TroubleshootingToolsDiagnosticsApi.md#GetStatDataFields30) | **Get** /statistics/speedtest/fields | 
[**GetStatDataRawData26**](TroubleshootingToolsDiagnosticsApi.md#GetStatDataRawData26) | **Get** /stream/device/nwpi | 
[**GetStatDataRawData27**](TroubleshootingToolsDiagnosticsApi.md#GetStatDataRawData27) | **Get** /statistics/speedtest | 
[**GetStatDataRawDataAsCSV27**](TroubleshootingToolsDiagnosticsApi.md#GetStatDataRawDataAsCSV27) | **Get** /stream/device/nwpi/csv | 
[**GetStatDataRawDataAsCSV28**](TroubleshootingToolsDiagnosticsApi.md#GetStatDataRawDataAsCSV28) | **Get** /statistics/speedtest/csv | 
[**GetStatQueryFields29**](TroubleshootingToolsDiagnosticsApi.md#GetStatQueryFields29) | **Get** /stream/device/nwpi/query/fields | 
[**GetStatQueryFields30**](TroubleshootingToolsDiagnosticsApi.md#GetStatQueryFields30) | **Get** /statistics/speedtest/query/fields | 
[**GetStatsRawData27**](TroubleshootingToolsDiagnosticsApi.md#GetStatsRawData27) | **Post** /stream/device/nwpi | 
[**GetStatsRawData28**](TroubleshootingToolsDiagnosticsApi.md#GetStatsRawData28) | **Post** /statistics/speedtest | 
[**GetThreadPools**](TroubleshootingToolsDiagnosticsApi.md#GetThreadPools) | **Get** /diagnostics/threadpools | 
[**GetTraceFlow**](TroubleshootingToolsDiagnosticsApi.md#GetTraceFlow) | **Get** /stream/device/nwpi/traceFlow | 
[**GetTraceHistory**](TroubleshootingToolsDiagnosticsApi.md#GetTraceHistory) | **Get** /stream/device/nwpi/traceHistory | 
[**GetVnicInfoByVnfId**](TroubleshootingToolsDiagnosticsApi.md#GetVnicInfoByVnfId) | **Get** /stream/device/capture/vnicsInfo/{vnfId} | 
[**MonitorStart**](TroubleshootingToolsDiagnosticsApi.md#MonitorStart) | **Post** /stream/device/nwpi/monitor/start | 
[**MonitorStop**](TroubleshootingToolsDiagnosticsApi.md#MonitorStop) | **Post** /stream/device/nwpi/monitor/stop | 
[**ProcessDeviceStatus**](TroubleshootingToolsDiagnosticsApi.md#ProcessDeviceStatus) | **Post** /stream/device/status/{deviceUUID} | 
[**RenewSessionInfo**](TroubleshootingToolsDiagnosticsApi.md#RenewSessionInfo) | **Get** /stream/device/log/renew/{sessionId} | 
[**SaveSpeedTestResults**](TroubleshootingToolsDiagnosticsApi.md#SaveSpeedTestResults) | **Post** /stream/device/speed/{deviceUUID}/{sessionId} | 
[**SearchDeviceLog**](TroubleshootingToolsDiagnosticsApi.md#SearchDeviceLog) | **Post** /stream/device/log/search/{sessionId} | 
[**StartPcapSession**](TroubleshootingToolsDiagnosticsApi.md#StartPcapSession) | **Get** /stream/device/capture/start/{sessionId} | 
[**StartSpeedTest**](TroubleshootingToolsDiagnosticsApi.md#StartSpeedTest) | **Get** /stream/device/speed/start/{sessionId} | 
[**StopPcapSession**](TroubleshootingToolsDiagnosticsApi.md#StopPcapSession) | **Get** /stream/device/capture/stop/{sessionId} | 
[**StopSpeedTest**](TroubleshootingToolsDiagnosticsApi.md#StopSpeedTest) | **Get** /stream/device/speed/stop/{sessionId} | 
[**StreamLog**](TroubleshootingToolsDiagnosticsApi.md#StreamLog) | **Post** /stream/device/log/{logType}/{deviceUUID}/{sessionId} | 
[**TraceDelete**](TroubleshootingToolsDiagnosticsApi.md#TraceDelete) | **Delete** /stream/device/nwpi/trace/delete | 
[**TraceFinFlowWithQuery**](TroubleshootingToolsDiagnosticsApi.md#TraceFinFlowWithQuery) | **Get** /stream/device/nwpi/traceFinFlowWithQuery | 
[**TraceStart**](TroubleshootingToolsDiagnosticsApi.md#TraceStart) | **Post** /stream/device/nwpi/trace/start | 
[**TraceStop**](TroubleshootingToolsDiagnosticsApi.md#TraceStop) | **Post** /stream/device/nwpi/trace/stop/{traceId} | 



## ClearSession

> ClearSession(ctx, sessionId).Execute()



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
    sessionId := "sessionId_example" // string | Session Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.ClearSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.ClearSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSessionRequest struct via the builder pattern


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


## DisableDeviceLog

> DisableDeviceLog(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.DisableDeviceLog(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.DisableDeviceLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableDeviceLogRequest struct via the builder pattern


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


## DisablePacketCaptureSession

> DisablePacketCaptureSession(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.DisablePacketCaptureSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.DisablePacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePacketCaptureSessionRequest struct via the builder pattern


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


## DisableSpeedTestSession

> DisableSpeedTestSession(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.DisableSpeedTestSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.DisableSpeedTestSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableSpeedTestSessionRequest struct via the builder pattern


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


## DownloadDebugLog

> DownloadDebugLog(ctx, sessionId).Execute()



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
    sessionId := "sessionId_example" // string | Session Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.DownloadDebugLog(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.DownloadDebugLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDebugLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFile

> DownloadFile(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.DownloadFile(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.DownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceStopPcapSession

> ForceStopPcapSession(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.ForceStopPcapSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.ForceStopPcapSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiForceStopPcapSessionRequest struct via the builder pattern


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


## FormPostPacketCapture

> FormPostPacketCapture(ctx, deviceUUID, sessionId).Execute()



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
    deviceUUID := "deviceUUID_example" // string | 
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.FormPostPacketCapture(context.Background(), deviceUUID, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.FormPostPacketCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUUID** | **string** |  | 
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFormPostPacketCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAggregationDataByQuery27

> map[string]interface{} GetAggregationDataByQuery27(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["1"],"operator":"greater"},"fields":["latency"],"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery27(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query filter | 

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


## GetAggregationDataByQuery28

> map[string]interface{} GetAggregationDataByQuery28(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["1"],"operator":"greater"},"fields":["latency"],"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery28(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregationDataByQuery28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregationDataByQuery28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query filter | 

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


## GetConcurrentData

> GetConcurrentData(ctx).TraceId(traceId).Timestamp(timestamp).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetConcurrentData(context.Background()).TraceId(traceId).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetConcurrentData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConcurrentDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 

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


## GetConcurrentDomainData

> GetConcurrentDomainData(ctx).TraceId(traceId).Timestamp(timestamp).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetConcurrentDomainData(context.Background()).TraceId(traceId).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetConcurrentDomainData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConcurrentDomainDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 

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


## GetCount29

> map[string]interface{} GetCount29(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["1"],"operator":"greater"},"fields":["latency"]}" // string | Query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCount29(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetCount29``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount29`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetCount29`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount29Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

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


## GetCount30

> map[string]interface{} GetCount30(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["1"],"operator":"greater"},"fields":["latency"]}" // string | Query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCount30(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetCount30``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount30`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetCount30`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount30Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query | 

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


## GetCountPost29

> map[string]interface{} GetCountPost29(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCountPost29(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetCountPost29``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost29`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetCountPost29`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost29Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Query | 

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


## GetCountPost30

> map[string]interface{} GetCountPost30(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCountPost30(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetCountPost30``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost30`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetCountPost30`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost30Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Query | 

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


## GetDBSchema

> map[string]interface{} GetDBSchema(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetDBSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetDBSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDBSchema`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetDBSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDBSchemaRequest struct via the builder pattern


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


## GetDeviceLog

> GetDeviceLog(ctx, sessionId).LogId(logId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 
    logId := int64(789) // int64 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetDeviceLog(context.Background(), sessionId).LogId(logId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetDeviceLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logId** | **int64** |  | [default to -1]

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


## GetDomainMetric

> GetDomainMetric(ctx).TraceId(traceId).Timestamp(timestamp).Domain(domain).FirstTimestamp(firstTimestamp).LastTimestamp(lastTimestamp).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time
    domain := "domain_example" // string | domain name
    firstTimestamp := int64(789) // int64 | first timestamp of xAxis
    lastTimestamp := int64(789) // int64 | last timestamp of xAxis

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetDomainMetric(context.Background()).TraceId(traceId).Timestamp(timestamp).Domain(domain).FirstTimestamp(firstTimestamp).LastTimestamp(lastTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetDomainMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 
 **domain** | **string** | domain name | 
 **firstTimestamp** | **int64** | first timestamp of xAxis | 
 **lastTimestamp** | **int64** | last timestamp of xAxis | 

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


## GetFileDownloadStatus

> GetFileDownloadStatus(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFileDownloadStatus(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetFileDownloadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileDownloadStatusRequest struct via the builder pattern


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


## GetFinalizedData

> GetFinalizedData(ctx).TraceId(traceId).Timestamp(timestamp).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFinalizedData(context.Background()).TraceId(traceId).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetFinalizedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinalizedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 

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


## GetFinalizedDomainData

> GetFinalizedDomainData(ctx).TraceId(traceId).Timestamp(timestamp).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFinalizedDomainData(context.Background()).TraceId(traceId).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetFinalizedDomainData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinalizedDomainDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 

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


## GetFlowDetail

> GetFlowDetail(ctx).TraceId(traceId).Timestamp(timestamp).FlowId(flowId).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time
    flowId := int32(56) // int32 | flow id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFlowDetail(context.Background()).TraceId(traceId).Timestamp(timestamp).FlowId(flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetFlowDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 
 **flowId** | **int32** | flow id | 

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


## GetFlowMetric

> GetFlowMetric(ctx).TraceId(traceId).Timestamp(timestamp).FlowId(flowId).FirstTimestamp(firstTimestamp).LastTimestamp(lastTimestamp).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time
    flowId := int32(56) // int32 | flow id
    firstTimestamp := int64(789) // int64 | first timestamp of xAxis
    lastTimestamp := int64(789) // int64 | last timestamp of xAxis

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFlowMetric(context.Background()).TraceId(traceId).Timestamp(timestamp).FlowId(flowId).FirstTimestamp(firstTimestamp).LastTimestamp(lastTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetFlowMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 
 **flowId** | **int32** | flow id | 
 **firstTimestamp** | **int64** | first timestamp of xAxis | 
 **lastTimestamp** | **int64** | last timestamp of xAxis | 

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


## GetInterfaceBandwidth

> GetInterfaceBandwidth(ctx).Circuit(circuit).DeviceUUID(deviceUUID).Execute()



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
    circuit := "circuit_example" // string | 
    deviceUUID := map[string][]openapiclient.DeviceUuid{ ... } // DeviceUuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetInterfaceBandwidth(context.Background()).Circuit(circuit).DeviceUUID(deviceUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetInterfaceBandwidth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **circuit** | **string** |  | 
 **deviceUUID** | [**DeviceUuid**](DeviceUuid.md) |  | 

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


## GetLogType

> GetLogType(ctx).Uuid(uuid).Execute()



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
    uuid := "uuid_example" // string | Device uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetLogType(context.Background()).Uuid(uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetLogType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** | Device uuid | 

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


## GetMonitorState

> GetMonitorState(ctx).TraceId(traceId).State(state).Execute()



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
    traceId := int32(56) // int32 | trace id
    state := "state_example" // string | trace state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetMonitorState(context.Background()).TraceId(traceId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetMonitorState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **state** | **string** | trace state | 

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


## GetNwpiDscp

> GetNwpiDscp(ctx).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetNwpiDscp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetNwpiDscp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNwpiDscpRequest struct via the builder pattern


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


## GetNwpiNbarAppGroup

> GetNwpiNbarAppGroup(ctx).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetNwpiNbarAppGroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetNwpiNbarAppGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNwpiNbarAppGroupRequest struct via the builder pattern


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


## GetNwpiProtocol

> GetNwpiProtocol(ctx).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetNwpiProtocol(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetNwpiProtocol``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNwpiProtocolRequest struct via the builder pattern


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


## GetPacketFeatures

> GetPacketFeatures(ctx).TraceId(traceId).Timestamp(timestamp).FlowId(flowId).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time
    flowId := int32(56) // int32 | flow id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPacketFeatures(context.Background()).TraceId(traceId).Timestamp(timestamp).FlowId(flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPacketFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPacketFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 
 **flowId** | **int32** | flow id | 

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


## GetPostAggregationAppDataByQuery26

> map[string]interface{} GetPostAggregationAppDataByQuery26(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery26(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery26Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostAggregationAppDataByQuery27

> map[string]interface{} GetPostAggregationAppDataByQuery27(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery27(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationAppDataByQuery27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationAppDataByQuery27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostAggregationDataByQuery27

> map[string]interface{} GetPostAggregationDataByQuery27(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery27(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostAggregationDataByQuery28

> map[string]interface{} GetPostAggregationDataByQuery28(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery28(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostStatBulkRawData27

> map[string]interface{} GetPostStatBulkRawData27(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData27(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scrollId** | **string** | ES scroll Id | 
 **count** | **string** | Result size | 
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetPostStatBulkRawData28

> map[string]interface{} GetPostStatBulkRawData28(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData28(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkRawData28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkRawData28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scrollId** | **string** | ES scroll Id | 
 **count** | **string** | Result size | 
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetSession

> GetSession(ctx).Body(body).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSession(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionInfoCapture

> GetSessionInfoCapture(ctx).Body(body).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSessionInfoCapture(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetSessionInfoCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInfoCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionInfoLog

> GetSessionInfoLog(ctx).Body(body).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSessionInfoLog(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetSessionInfoLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInfoLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessions

> GetSessions(ctx).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern


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


## GetSpeedTest

> GetSpeedTest(ctx, sessionId).LogId(logId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 
    logId := int64(789) // int64 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSpeedTest(context.Background(), sessionId).LogId(logId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetSpeedTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpeedTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logId** | **int64** |  | [default to 0]

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


## GetSpeedTestStatus

> GetSpeedTestStatus(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSpeedTestStatus(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetSpeedTestStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpeedTestStatusRequest struct via the builder pattern


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


## GetStatBulkRawData27

> map[string]interface{} GetStatBulkRawData27(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["100"],"operator":"greater"},"size":1,"sort":[{"field":"latency","type":"long","order":"asc"}],"fields":["latency"]}" // string | Query string (optional)
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData27(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 
 **scrollId** | **string** | ES scroll Id | 
 **count** | **string** | Result size | 

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


## GetStatBulkRawData28

> map[string]interface{} GetStatBulkRawData28(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["100"],"operator":"greater"},"size":1,"sort":[{"field":"latency","type":"long","order":"asc"}],"fields":["latency"]}" // string | Query string (optional)
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | ES scroll Id (optional)
    count := "10" // string | Result size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData28(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkRawData28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkRawData28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 
 **scrollId** | **string** | ES scroll Id | 
 **count** | **string** | Result size | 

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


## GetStatDataFields29

> map[string]interface{} GetStatDataFields29(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataFields29(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatDataFields29``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields29`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatDataFields29`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields29Request struct via the builder pattern


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


## GetStatDataFields30

> map[string]interface{} GetStatDataFields30(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataFields30(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatDataFields30``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields30`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatDataFields30`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields30Request struct via the builder pattern


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


## GetStatDataRawData26

> map[string]interface{} GetStatDataRawData26(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["2020-05-10T01:00:00 UTC","2020-05-10T01:30:00 UTC"],"field":"entry_time","type":"date","operator":"between"}]},"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query string (optional)
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawData26(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatDataRawData26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawData26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatDataRawData26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawData26Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 

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


## GetStatDataRawData27

> map[string]interface{} GetStatDataRawData27(ctx).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
    query := "{"query":{"condition":"AND","rules":[{"value":["2020-05-10T01:00:00 UTC","2020-05-10T01:30:00 UTC"],"field":"entry_time","type":"date","operator":"between"}]},"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query string (optional)
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawData27(context.Background()).Query(query).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatDataRawData27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawData27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatDataRawData27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawData27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 

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


## GetStatDataRawDataAsCSV27

> string GetStatDataRawDataAsCSV27(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["100"],"operator":"greater"},"size":1000,"sort":[{"field":"latency","type":"long","order":"asc"}],"fields":["latency"],"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV27(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV27`: string
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatDataRawDataAsCSV28

> string GetStatDataRawDataAsCSV28(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"latency","type":"long","value":["100"],"operator":"greater"},"size":1000,"sort":[{"field":"latency","type":"long","order":"asc"}],"fields":["latency"],"aggregation":{"metrics":[{"property":"latency","type":"avg"}]}}" // string | Query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV28(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataRawDataAsCSV28`: string
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataRawDataAsCSV28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query string | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatQueryFields29

> map[string]interface{} GetStatQueryFields29(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatQueryFields29(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatQueryFields29``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields29`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatQueryFields29`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields29Request struct via the builder pattern


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


## GetStatQueryFields30

> map[string]interface{} GetStatQueryFields30(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatQueryFields30(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatQueryFields30``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields30`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatQueryFields30`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields30Request struct via the builder pattern


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


## GetStatsRawData27

> map[string]interface{} GetStatsRawData27(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatsRawData27(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatsRawData27``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData27`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatsRawData27`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetStatsRawData28

> map[string]interface{} GetStatsRawData28(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    page := int64(789) // int64 | page number (optional)
    pageSize := int64(789) // int64 | page size (optional)
    sortBy := "sortBy_example" // string | sort by (optional)
    sortOrder := "sortOrder_example" // string | sort order (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatsRawData28(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetStatsRawData28``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsRawData28`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetStatsRawData28`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRawData28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetThreadPools

> map[string]interface{} GetThreadPools(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetThreadPools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetThreadPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreadPools`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.GetThreadPools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreadPoolsRequest struct via the builder pattern


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


## GetTraceFlow

> GetTraceFlow(ctx).TraceId(traceId).Timestamp(timestamp).State(state).Execute()



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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time
    state := "state_example" // string | trace state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetTraceFlow(context.Background()).TraceId(traceId).Timestamp(timestamp).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetTraceFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 
 **state** | **string** | trace state | 

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


## GetTraceHistory

> GetTraceHistory(ctx).Execute()



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
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetTraceHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetTraceHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceHistoryRequest struct via the builder pattern


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


## GetVnicInfoByVnfId

> GetVnicInfoByVnfId(ctx, vnfId).Execute()



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
    vnfId := "vnfId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetVnicInfoByVnfId(context.Background(), vnfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.GetVnicInfoByVnfId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVnicInfoByVnfIdRequest struct via the builder pattern


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


## MonitorStart

> MonitorStart(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.MonitorStart(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.MonitorStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## MonitorStop

> MonitorStop(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.MonitorStop(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.MonitorStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## ProcessDeviceStatus

> ProcessDeviceStatus(ctx, deviceUUID).Body(body).Execute()





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
    deviceUUID := "deviceUUID_example" // string | Device uuid
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.ProcessDeviceStatus(context.Background(), deviceUUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.ProcessDeviceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUUID** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessDeviceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewSessionInfo

> RenewSessionInfo(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.RenewSessionInfo(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.RenewSessionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewSessionInfoRequest struct via the builder pattern


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


## SaveSpeedTestResults

> SaveSpeedTestResults(ctx, deviceUUID, sessionId).Body(body).Execute()



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
    deviceUUID := "deviceUUID_example" // string | 
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.SaveSpeedTestResults(context.Background(), deviceUUID, sessionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.SaveSpeedTestResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceUUID** | **string** |  | 
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveSpeedTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeviceLog

> SearchDeviceLog(ctx, sessionId).Body(body).Execute()



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
    sessionId := "sessionId_example" // string | Session Id
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.SearchDeviceLog(context.Background(), sessionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.SearchDeviceLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeviceLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPcapSession

> StartPcapSession(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.StartPcapSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.StartPcapSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartPcapSessionRequest struct via the builder pattern


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


## StartSpeedTest

> StartSpeedTest(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.StartSpeedTest(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.StartSpeedTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSpeedTestRequest struct via the builder pattern


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


## StopPcapSession

> StopPcapSession(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.StopPcapSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.StopPcapSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPcapSessionRequest struct via the builder pattern


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


## StopSpeedTest

> StopSpeedTest(ctx, sessionId).Execute()



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
    sessionId := map[string][]openapiclient.Uuid{ ... } // Uuid | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.StopSpeedTest(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.StopSpeedTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | [**Uuid**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSpeedTestRequest struct via the builder pattern


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


## StreamLog

> StreamLog(ctx, logType, deviceUUID, sessionId).Body(body).Execute()



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
    logType := "logType_example" // string | Log type
    deviceUUID := "deviceUUID_example" // string | Device uuid
    sessionId := "sessionId_example" // string | Session Id
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.StreamLog(context.Background(), logType, deviceUUID, sessionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.StreamLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logType** | **string** | Log type | 
**deviceUUID** | **string** | Device uuid | 
**sessionId** | **string** | Session Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

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


## TraceDelete

> TraceDelete(ctx).TraceId(traceId).Timestamp(timestamp).Execute()





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
    traceId := "traceId_example" // string | trace id
    timestamp := int64(789) // int64 | start time

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceDelete(context.Background()).TraceId(traceId).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.TraceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTraceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **string** | trace id | 
 **timestamp** | **int64** | start time | 

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


## TraceFinFlowWithQuery

> map[string]interface{} TraceFinFlowWithQuery(ctx).TraceId(traceId).Timestamp(timestamp).Query(query).Execute()





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
    traceId := int32(56) // int32 | trace id
    timestamp := int64(789) // int64 | start time
    query := "query_example" // string | Query filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceFinFlowWithQuery(context.Background()).TraceId(traceId).Timestamp(timestamp).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.TraceFinFlowWithQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TraceFinFlowWithQuery`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDiagnosticsApi.TraceFinFlowWithQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTraceFinFlowWithQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **int32** | trace id | 
 **timestamp** | **int64** | start time | 
 **query** | **string** | Query filter | 

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


## TraceStart

> TraceStart(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceStart(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.TraceStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTraceStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## TraceStop

> TraceStop(ctx, traceId).Execute()





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
    traceId := "traceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceStop(context.Background(), traceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDiagnosticsApi.TraceStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTraceStopRequest struct via the builder pattern


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

