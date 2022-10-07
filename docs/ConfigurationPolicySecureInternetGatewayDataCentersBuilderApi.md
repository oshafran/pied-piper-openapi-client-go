# \ConfigurationPolicySecureInternetGatewayDataCentersBuilderApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSigDataCenterList**](ConfigurationPolicySecureInternetGatewayDataCentersBuilderApi.md#GetSigDataCenterList) | **Get** /sig/datacenters/{type} | 



## GetSigDataCenterList

> map[string]interface{} GetSigDataCenterList(ctx, type_).Execute()





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
    type_ := "type__example" // string | Provider type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicySecureInternetGatewayDataCentersBuilderApi.GetSigDataCenterList(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicySecureInternetGatewayDataCentersBuilderApi.GetSigDataCenterList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSigDataCenterList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicySecureInternetGatewayDataCentersBuilderApi.GetSigDataCenterList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Provider type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigDataCenterListRequest struct via the builder pattern


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

