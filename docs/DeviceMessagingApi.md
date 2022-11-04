# \DeviceMessagingApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceVmanageConnectionList**](DeviceMessagingApi.md#CreateDeviceVmanageConnectionList) | **Get** /messaging/device/vmanage | 



## CreateDeviceVmanageConnectionList

> []map[string]interface{} CreateDeviceVmanageConnectionList(ctx).Execute()





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
    resp, r, err := apiClient.DeviceMessagingApi.CreateDeviceVmanageConnectionList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceMessagingApi.CreateDeviceVmanageConnectionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceVmanageConnectionList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceMessagingApi.CreateDeviceVmanageConnectionList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceVmanageConnectionListRequest struct via the builder pattern


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

