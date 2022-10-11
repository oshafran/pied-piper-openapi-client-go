# \MonitoringServerInfoApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerInfo**](MonitoringServerInfoApi.md#CreateServerInfo) | **Get** /server/info | 



## CreateServerInfo

> map[string]interface{} CreateServerInfo(ctx).Execute()





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
    resp, r, err := apiClient.MonitoringServerInfoApi.CreateServerInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringServerInfoApi.CreateServerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServerInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MonitoringServerInfoApi.CreateServerInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerInfoRequest struct via the builder pattern


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

