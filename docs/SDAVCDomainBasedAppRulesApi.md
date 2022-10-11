# \SDAVCDomainBasedAppRulesApi

All URIs are relative to *http://192.168.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudConnectorDomainAppRules**](SDAVCDomainBasedAppRulesApi.md#GetCloudConnectorDomainAppRules) | **Get** /monitor/sdavccloudconnector/domain | 



## GetCloudConnectorDomainAppRules

> map[string]interface{} GetCloudConnectorDomainAppRules(ctx).Execute()





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
    resp, r, err := apiClient.SDAVCDomainBasedAppRulesApi.GetCloudConnectorDomainAppRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDAVCDomainBasedAppRulesApi.GetCloudConnectorDomainAppRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudConnectorDomainAppRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SDAVCDomainBasedAppRulesApi.GetCloudConnectorDomainAppRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudConnectorDomainAppRulesRequest struct via the builder pattern


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

