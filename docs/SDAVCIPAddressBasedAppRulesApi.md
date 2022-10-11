# \SDAVCIPAddressBasedAppRulesApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudConnectorIpAddressAppRules**](SDAVCIPAddressBasedAppRulesApi.md#GetCloudConnectorIpAddressAppRules) | **Get** /monitor/sdavccloudconnector/ipaddress | 



## GetCloudConnectorIpAddressAppRules

> map[string]interface{} GetCloudConnectorIpAddressAppRules(ctx).Execute()





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
    resp, r, err := apiClient.SDAVCIPAddressBasedAppRulesApi.GetCloudConnectorIpAddressAppRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDAVCIPAddressBasedAppRulesApi.GetCloudConnectorIpAddressAppRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudConnectorIpAddressAppRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SDAVCIPAddressBasedAppRulesApi.GetCloudConnectorIpAddressAppRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudConnectorIpAddressAppRulesRequest struct via the builder pattern


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

