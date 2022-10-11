# \RealTimeMonitoringCellularEIOLTEConnectionServiceApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEiolteConnectionInfo**](RealTimeMonitoringCellularEIOLTEConnectionServiceApi.md#GetEiolteConnectionInfo) | **Get** /device/cellularEiolte/connections | 
[**GetEiolteRadioInfo**](RealTimeMonitoringCellularEIOLTEConnectionServiceApi.md#GetEiolteRadioInfo) | **Get** /device/cellularEiolte/radio | 



## GetEiolteConnectionInfo

> map[string]interface{} GetEiolteConnectionInfo(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCellularEIOLTEConnectionServiceApi.GetEiolteConnectionInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCellularEIOLTEConnectionServiceApi.GetEiolteConnectionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEiolteConnectionInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCellularEIOLTEConnectionServiceApi.GetEiolteConnectionInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEiolteConnectionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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


## GetEiolteRadioInfo

> map[string]interface{} GetEiolteRadioInfo(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringCellularEIOLTEConnectionServiceApi.GetEiolteRadioInfo(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringCellularEIOLTEConnectionServiceApi.GetEiolteRadioInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEiolteRadioInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringCellularEIOLTEConnectionServiceApi.GetEiolteRadioInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEiolteRadioInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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

