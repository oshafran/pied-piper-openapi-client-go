# \ConfigurationSecureInternetGatewayTunnelsApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSigTunnelList**](ConfigurationSecureInternetGatewayTunnelsApi.md#GetSigTunnelList) | **Get** /template/cloudx/sig_tunnels | 



## GetSigTunnelList

> GetSigTunnelList(ctx).DeviceId(deviceId).Execute()





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
    resp, r, err := apiClient.ConfigurationSecureInternetGatewayTunnelsApi.GetSigTunnelList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecureInternetGatewayTunnelsApi.GetSigTunnelList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSigTunnelListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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

