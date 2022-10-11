# \MonitoringAlarmsDetailsApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearStaleAlarm**](MonitoringAlarmsDetailsApi.md#ClearStaleAlarm) | **Post** /alarms/clear | 
[**CorrelAntiEntropy**](MonitoringAlarmsDetailsApi.md#CorrelAntiEntropy) | **Get** /alarms/reset | 
[**CreateAlarmQueryConfig**](MonitoringAlarmsDetailsApi.md#CreateAlarmQueryConfig) | **Get** /alarms/query/input | 
[**CreateNotificationRule**](MonitoringAlarmsDetailsApi.md#CreateNotificationRule) | **Post** /notifications/rule | 
[**DeleteNotificationRule**](MonitoringAlarmsDetailsApi.md#DeleteNotificationRule) | **Delete** /notifications/rules | 
[**DisableEnableAlarm**](MonitoringAlarmsDetailsApi.md#DisableEnableAlarm) | **Post** /alarms/disabled | 
[**DumpCorrelationEngineData**](MonitoringAlarmsDetailsApi.md#DumpCorrelationEngineData) | **Post** /alarms/dump | 
[**EnableDisableLinkStateAlarm**](MonitoringAlarmsDetailsApi.md#EnableDisableLinkStateAlarm) | **Post** /alarms/link-state-alarm | 
[**GetAlarmAggregationData**](MonitoringAlarmsDetailsApi.md#GetAlarmAggregationData) | **Get** /alarms/aggregation | 
[**GetAlarmDetails**](MonitoringAlarmsDetailsApi.md#GetAlarmDetails) | **Get** /alarms/uuid/{alarm_uuid} | 
[**GetAlarmSeverityCustomHistogram**](MonitoringAlarmsDetailsApi.md#GetAlarmSeverityCustomHistogram) | **Get** /alarms/severity/summary | 
[**GetAlarmSeverityMappings**](MonitoringAlarmsDetailsApi.md#GetAlarmSeverityMappings) | **Get** /alarms/severitymappings | 
[**GetAlarmTypesAsKeyValue**](MonitoringAlarmsDetailsApi.md#GetAlarmTypesAsKeyValue) | **Get** /alarms/rulenamedisplay/keyvalue | 
[**GetAlarms**](MonitoringAlarmsDetailsApi.md#GetAlarms) | **Get** /alarms | 
[**GetAlarmsBySeverity**](MonitoringAlarmsDetailsApi.md#GetAlarmsBySeverity) | **Get** /alarms/severity | 
[**GetCount1**](MonitoringAlarmsDetailsApi.md#GetCount1) | **Get** /alarms/doccount | 
[**GetCountPost1**](MonitoringAlarmsDetailsApi.md#GetCountPost1) | **Post** /alarms/doccount | 
[**GetDeviceTopic**](MonitoringAlarmsDetailsApi.md#GetDeviceTopic) | **Get** /alarms/topic | 
[**GetLinkStateAlarmConfig**](MonitoringAlarmsDetailsApi.md#GetLinkStateAlarmConfig) | **Get** /alarms/link-state-alarm | 
[**GetMasterManagerState**](MonitoringAlarmsDetailsApi.md#GetMasterManagerState) | **Get** /alarms/master | 
[**GetNonViewedActiveAlarmsCount**](MonitoringAlarmsDetailsApi.md#GetNonViewedActiveAlarmsCount) | **Get** /alarms/count | 
[**GetNonViewedAlarms**](MonitoringAlarmsDetailsApi.md#GetNonViewedAlarms) | **Get** /alarms/notviewed | 
[**GetNotificationRule**](MonitoringAlarmsDetailsApi.md#GetNotificationRule) | **Get** /notifications/rules | 
[**GetPostAlarmAggregationData**](MonitoringAlarmsDetailsApi.md#GetPostAlarmAggregationData) | **Post** /alarms/aggregation | 
[**GetPostStatBulkAlarmRawData**](MonitoringAlarmsDetailsApi.md#GetPostStatBulkAlarmRawData) | **Post** /alarms/page | 
[**GetRawAlarmData**](MonitoringAlarmsDetailsApi.md#GetRawAlarmData) | **Post** /alarms | 
[**GetStatBulkAlarmRawData**](MonitoringAlarmsDetailsApi.md#GetStatBulkAlarmRawData) | **Get** /alarms/page | 
[**GetStatDataFields1**](MonitoringAlarmsDetailsApi.md#GetStatDataFields1) | **Get** /alarms/fields | 
[**GetStatQueryFields1**](MonitoringAlarmsDetailsApi.md#GetStatQueryFields1) | **Get** /alarms/query/fields | 
[**GetStats**](MonitoringAlarmsDetailsApi.md#GetStats) | **Get** /alarms/stats | 
[**ListDisabledAlarm**](MonitoringAlarmsDetailsApi.md#ListDisabledAlarm) | **Get** /alarms/disabled | 
[**MarkAlarmsAsViewed**](MonitoringAlarmsDetailsApi.md#MarkAlarmsAsViewed) | **Post** /alarms/markviewed | 
[**MarkAllAlarmsAsViewed**](MonitoringAlarmsDetailsApi.md#MarkAllAlarmsAsViewed) | **Post** /alarms/markallasviewed | 
[**RestartCorrelationEngine**](MonitoringAlarmsDetailsApi.md#RestartCorrelationEngine) | **Get** /alarms/restart | 
[**SetPeriodicPurgeTimer**](MonitoringAlarmsDetailsApi.md#SetPeriodicPurgeTimer) | **Get** /alarms/purgefrequency | 
[**StartTracking**](MonitoringAlarmsDetailsApi.md#StartTracking) | **Post** /alarms/starttracking/{testName} | 
[**StopTracking**](MonitoringAlarmsDetailsApi.md#StopTracking) | **Post** /alarms/stoptracking/{testName} | 
[**UpdateNotificationRule**](MonitoringAlarmsDetailsApi.md#UpdateNotificationRule) | **Put** /notifications/rule | 



## ClearStaleAlarm

> ClearStaleAlarm(ctx).RequestBody(requestBody).Execute()





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
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | alarm_uuid (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.ClearStaleAlarm(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.ClearStaleAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearStaleAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]map[string]interface{}** | alarm_uuid | 

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


## CorrelAntiEntropy

> map[string]interface{} CorrelAntiEntropy(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.CorrelAntiEntropy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.CorrelAntiEntropy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorrelAntiEntropy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.CorrelAntiEntropy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCorrelAntiEntropyRequest struct via the builder pattern


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


## CreateAlarmQueryConfig

> map[string]interface{} CreateAlarmQueryConfig(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.CreateAlarmQueryConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.CreateAlarmQueryConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlarmQueryConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.CreateAlarmQueryConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlarmQueryConfigRequest struct via the builder pattern


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


## CreateNotificationRule

> map[string]interface{} CreateNotificationRule(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Notification rule (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.CreateNotificationRule(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.CreateNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationRule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.CreateNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Notification rule | 

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


## DeleteNotificationRule

> DeleteNotificationRule(ctx).RuleId(ruleId).Execute()





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
    ruleId := "ruleId_example" // string | Rule Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.DeleteNotificationRule(context.Background()).RuleId(ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.DeleteNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleId** | **string** | Rule Id | 

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


## DisableEnableAlarm

> DisableEnableAlarm(ctx).EventName(eventName).Disable(disable).Time(time).Body(body).Execute()





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
    eventName := "eventName_example" // string | Event name
    disable := true // bool | Disable
    time := int64(789) // int64 | time in hours [1, 72], -1 means infinite
    body := map[string]interface{}{ ... } // map[string]interface{} | alarm config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.DisableEnableAlarm(context.Background()).EventName(eventName).Disable(disable).Time(time).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.DisableEnableAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableEnableAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventName** | **string** | Event name | 
 **disable** | **bool** | Disable | 
 **time** | **int64** | time in hours [1, 72], -1 means infinite | 
 **body** | **map[string]interface{}** | alarm config | 

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


## DumpCorrelationEngineData

> DumpCorrelationEngineData(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.DumpCorrelationEngineData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.DumpCorrelationEngineData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDumpCorrelationEngineDataRequest struct via the builder pattern


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


## EnableDisableLinkStateAlarm

> EnableDisableLinkStateAlarm(ctx).LinkName(linkName).Enable(enable).Body(body).Execute()





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
    linkName := "linkName_example" // string | Link name (bgp, ospf)
    enable := true // bool | Enable
    body := map[string]interface{}{ ... } // map[string]interface{} | alarm config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.EnableDisableLinkStateAlarm(context.Background()).LinkName(linkName).Enable(enable).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.EnableDisableLinkStateAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableDisableLinkStateAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkName** | **string** | Link name (bgp, ospf) | 
 **enable** | **bool** | Enable | 
 **body** | **map[string]interface{}** | alarm config | 

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


## GetAlarmAggregationData

> map[string]interface{} GetAlarmAggregationData(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"active","type":"boolean","value":["true"],"operator":"equal"},"aggregation":{"metrics":[{"property":"severity","type":"count"}]}}" // string | Alarm query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarmAggregationData(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarmAggregationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmAggregationData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarmAggregationData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmAggregationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Alarm query string | 

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


## GetAlarmDetails

> map[string]interface{} GetAlarmDetails(ctx, alarmUuid).Execute()





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
    alarmUuid := "alarmUuid_example" // string | Alarm Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarmDetails(context.Background(), alarmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarmDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarmDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmUuid** | **string** | Alarm Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmDetailsRequest struct via the builder pattern


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


## GetAlarmSeverityCustomHistogram

> map[string]interface{} GetAlarmSeverityCustomHistogram(ctx).Query(query).Execute()





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
    query := "query_example" // string | Alarm histogram query string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarmSeverityCustomHistogram(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarmSeverityCustomHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmSeverityCustomHistogram`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarmSeverityCustomHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmSeverityCustomHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Alarm histogram query string | 

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


## GetAlarmSeverityMappings

> map[string]interface{} GetAlarmSeverityMappings(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarmSeverityMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarmSeverityMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmSeverityMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarmSeverityMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmSeverityMappingsRequest struct via the builder pattern


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


## GetAlarmTypesAsKeyValue

> map[string]interface{} GetAlarmTypesAsKeyValue(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarmTypesAsKeyValue(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarmTypesAsKeyValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmTypesAsKeyValue`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarmTypesAsKeyValue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmTypesAsKeyValueRequest struct via the builder pattern


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


## GetAlarms

> map[string]interface{} GetAlarms(ctx).Query(query).Execute()





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
    query := "{"query":{"field":"active","type":"boolean","value":["true"],"operator":"equal"}}" // string | Alarm query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarms(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarms`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Alarm query string | 

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


## GetAlarmsBySeverity

> map[string]interface{} GetAlarmsBySeverity(ctx).SeverityLevel(severityLevel).DeviceId(deviceId).Query(query).Execute()





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
    severityLevel := []string{"SeverityLevel_example"} // []string | Alarm severity
    deviceId := []string{"Inner_example"} // []string | Device Id
    query := "query_example" // string | Query filter

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetAlarmsBySeverity(context.Background()).SeverityLevel(severityLevel).DeviceId(deviceId).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetAlarmsBySeverity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlarmsBySeverity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetAlarmsBySeverity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlarmsBySeverityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **severityLevel** | **[]string** | Alarm severity | 
 **deviceId** | **[]string** | Device Id | 
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


## GetCount1

> map[string]interface{} GetCount1(ctx).Query(query).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetCount1(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetCount1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCount1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetCount1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCount1Request struct via the builder pattern


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


## GetCountPost1

> map[string]interface{} GetCountPost1(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetCountPost1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetCountPost1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountPost1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetCountPost1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountPost1Request struct via the builder pattern


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


## GetDeviceTopic

> GetDeviceTopic(ctx).Ip(ip).Execute()





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
    ip := "172.16.255.14" // string | Query topic

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetDeviceTopic(context.Background()).Ip(ip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetDeviceTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | Query topic | 

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


## GetLinkStateAlarmConfig

> GetLinkStateAlarmConfig(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetLinkStateAlarmConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetLinkStateAlarmConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkStateAlarmConfigRequest struct via the builder pattern


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


## GetMasterManagerState

> GetMasterManagerState(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetMasterManagerState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetMasterManagerState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterManagerStateRequest struct via the builder pattern


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


## GetNonViewedActiveAlarmsCount

> map[string]interface{} GetNonViewedActiveAlarmsCount(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetNonViewedActiveAlarmsCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetNonViewedActiveAlarmsCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonViewedActiveAlarmsCount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetNonViewedActiveAlarmsCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonViewedActiveAlarmsCountRequest struct via the builder pattern


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


## GetNonViewedAlarms

> map[string]interface{} GetNonViewedAlarms(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetNonViewedAlarms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetNonViewedAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonViewedAlarms`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetNonViewedAlarms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonViewedAlarmsRequest struct via the builder pattern


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


## GetNotificationRule

> map[string]interface{} GetNotificationRule(ctx).RuleId(ruleId).Execute()





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
    ruleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Rule Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetNotificationRule(context.Background()).RuleId(ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationRule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleId** | **string** | Rule Id | 

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


## GetPostAlarmAggregationData

> map[string]interface{} GetPostAlarmAggregationData(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Input query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetPostAlarmAggregationData(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetPostAlarmAggregationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAlarmAggregationData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetPostAlarmAggregationData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAlarmAggregationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Input query | 

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


## GetPostStatBulkAlarmRawData

> map[string]interface{} GetPostStatBulkAlarmRawData(ctx).ScrollId(scrollId).Count(count).Body(body).Execute()





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
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | Query offset
    count := int64(10) // int64 | Query size
    body := map[string]interface{}{ ... } // map[string]interface{} | Alarm query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetPostStatBulkAlarmRawData(context.Background()).ScrollId(scrollId).Count(count).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetPostStatBulkAlarmRawData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostStatBulkAlarmRawData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetPostStatBulkAlarmRawData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostStatBulkAlarmRawDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scrollId** | **string** | Query offset | 
 **count** | **int64** | Query size | 
 **body** | **map[string]interface{}** | Alarm query string | 

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


## GetRawAlarmData

> map[string]interface{} GetRawAlarmData(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Alarm query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetRawAlarmData(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetRawAlarmData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawAlarmData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetRawAlarmData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRawAlarmDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | page number | 
 **pageSize** | **int64** | page size | 
 **sortBy** | **string** | sort by | 
 **sortOrder** | **string** | sort order | 
 **body** | **map[string]interface{}** | Alarm query string | 

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


## GetStatBulkAlarmRawData

> map[string]interface{} GetStatBulkAlarmRawData(ctx).Query(query).ScrollId(scrollId).Count(count).Execute()





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
    query := "{"query":{"field":"active","type":"boolean","value":["true"],"operator":"equal"}}" // string | Alarm query string
    scrollId := "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAOIWZ1NQbXpvQ29Uc0stNzZ2UzlwTEREUQ==" // string | Query offset
    count := int64(10) // int64 | Query size

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetStatBulkAlarmRawData(context.Background()).Query(query).ScrollId(scrollId).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetStatBulkAlarmRawData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatBulkAlarmRawData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetStatBulkAlarmRawData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatBulkAlarmRawDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Alarm query string | 
 **scrollId** | **string** | Query offset | 
 **count** | **int64** | Query size | 

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


## GetStatDataFields1

> map[string]interface{} GetStatDataFields1(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetStatDataFields1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetStatDataFields1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatDataFields1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetStatDataFields1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatDataFields1Request struct via the builder pattern


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


## GetStatQueryFields1

> map[string]interface{} GetStatQueryFields1(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetStatQueryFields1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetStatQueryFields1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatQueryFields1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetStatQueryFields1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatQueryFields1Request struct via the builder pattern


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


## GetStats

> map[string]interface{} GetStats(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.GetStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.GetStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.GetStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


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


## ListDisabledAlarm

> []map[string]interface{} ListDisabledAlarm(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.ListDisabledAlarm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.ListDisabledAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDisabledAlarm`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.ListDisabledAlarm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDisabledAlarmRequest struct via the builder pattern


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


## MarkAlarmsAsViewed

> map[string]interface{} MarkAlarmsAsViewed(ctx).RequestBody(requestBody).Execute()





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
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | List of alarms (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.MarkAlarmsAsViewed(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.MarkAlarmsAsViewed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkAlarmsAsViewed`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.MarkAlarmsAsViewed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkAlarmsAsViewedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]map[string]interface{}** | List of alarms | 

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


## MarkAllAlarmsAsViewed

> MarkAllAlarmsAsViewed(ctx).Type_(type_).Body(body).Execute()





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
    type_ := "type__example" // string | Query filter, possible value are \"active\" \"cleared\" (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.MarkAllAlarmsAsViewed(context.Background()).Type_(type_).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.MarkAllAlarmsAsViewed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkAllAlarmsAsViewedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Query filter, possible value are \&quot;active\&quot; \&quot;cleared\&quot; | 
 **body** | **string** |  | 

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


## RestartCorrelationEngine

> map[string]interface{} RestartCorrelationEngine(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.RestartCorrelationEngine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.RestartCorrelationEngine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartCorrelationEngine`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.RestartCorrelationEngine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRestartCorrelationEngineRequest struct via the builder pattern


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


## SetPeriodicPurgeTimer

> map[string]interface{} SetPeriodicPurgeTimer(ctx).Interval(interval).ActiveTime(activeTime).Execute()





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
    interval := "interval_example" // string | Purge interval (optional)
    activeTime := "activeTime_example" // string | Purge activeTime (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.SetPeriodicPurgeTimer(context.Background()).Interval(interval).ActiveTime(activeTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.SetPeriodicPurgeTimer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPeriodicPurgeTimer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.SetPeriodicPurgeTimer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPeriodicPurgeTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | **string** | Purge interval | 
 **activeTime** | **string** | Purge activeTime | 

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


## StartTracking

> map[string]interface{} StartTracking(ctx, testName).Execute()





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
    testName := "testName_example" // string | test name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.StartTracking(context.Background(), testName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.StartTracking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartTracking`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.StartTracking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testName** | **string** | test name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTrackingRequest struct via the builder pattern


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


## StopTracking

> map[string]interface{} StopTracking(ctx, testName).Execute()





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
    testName := "testName_example" // string | test name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.StopTracking(context.Background(), testName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.StopTracking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopTracking`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAlarmsDetailsApi.StopTracking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testName** | **string** | test name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTrackingRequest struct via the builder pattern


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


## UpdateNotificationRule

> UpdateNotificationRule(ctx).RuleId(ruleId).Body(body).Execute()





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
    ruleId := "ruleId_example" // string | Rule Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Notification rule (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAlarmsDetailsApi.UpdateNotificationRule(context.Background()).RuleId(ruleId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAlarmsDetailsApi.UpdateNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleId** | **string** | Rule Id | 
 **body** | **map[string]interface{}** | Notification rule | 

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

