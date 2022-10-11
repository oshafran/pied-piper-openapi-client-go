# \UtilityLoggingApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DebugLog**](UtilityLoggingApi.md#DebugLog) | **Post** /util/logging/debuglog | 
[**ListLogFileDetails**](UtilityLoggingApi.md#ListLogFileDetails) | **Get** /util/logfile/appserver | 
[**ListLoggers**](UtilityLoggingApi.md#ListLoggers) | **Get** /util/logging/loggers | 
[**ListVManageServerLogLastNLines**](UtilityLoggingApi.md#ListVManageServerLogLastNLines) | **Get** /util/logfile/appserver/lastnlines | 
[**SetLogLevel**](UtilityLoggingApi.md#SetLogLevel) | **Post** /util/logging/level | 



## DebugLog

> DebugLog(ctx).LoggerName(loggerName).LogMessage(logMessage).Execute()





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
    loggerName := "loggerName_example" // string | 
    logMessage := "logMessage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilityLoggingApi.DebugLog(context.Background()).LoggerName(loggerName).LogMessage(logMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityLoggingApi.DebugLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDebugLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loggerName** | **string** |  | 
 **logMessage** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogFileDetails

> string ListLogFileDetails(ctx).Execute()





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
    resp, r, err := apiClient.UtilityLoggingApi.ListLogFileDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityLoggingApi.ListLogFileDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogFileDetails`: string
    fmt.Fprintf(os.Stdout, "Response from `UtilityLoggingApi.ListLogFileDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogFileDetailsRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoggers

> []map[string]interface{} ListLoggers(ctx).Execute()





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
    resp, r, err := apiClient.UtilityLoggingApi.ListLoggers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityLoggingApi.ListLoggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoggers`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UtilityLoggingApi.ListLoggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoggersRequest struct via the builder pattern


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


## ListVManageServerLogLastNLines

> string ListVManageServerLogLastNLines(ctx).Lines(lines).Execute()





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
    lines := int64(789) // int64 | Number of lines (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilityLoggingApi.ListVManageServerLogLastNLines(context.Background()).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityLoggingApi.ListVManageServerLogLastNLines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVManageServerLogLastNLines`: string
    fmt.Fprintf(os.Stdout, "Response from `UtilityLoggingApi.ListVManageServerLogLastNLines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVManageServerLogLastNLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lines** | **int64** | Number of lines | [default to 100]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLogLevel

> SetLogLevel(ctx).LoggerName(loggerName).LogLevel(logLevel).Execute()





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
    loggerName := "loggerName_example" // string | 
    logLevel := "logLevel_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilityLoggingApi.SetLogLevel(context.Background()).LoggerName(loggerName).LogLevel(logLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilityLoggingApi.SetLogLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetLogLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loggerName** | **string** |  | 
 **logLevel** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

