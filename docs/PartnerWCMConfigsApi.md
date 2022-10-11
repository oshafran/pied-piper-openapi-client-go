# \PartnerWCMConfigsApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushNetconfConfigs**](PartnerWCMConfigsApi.md#PushNetconfConfigs) | **Post** /partner/wcm/netconf/{nmsId} | 



## PushNetconfConfigs

> map[string]interface{} PushNetconfConfigs(ctx, nmsId).RequestBody(requestBody).Execute()





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
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NMS Id
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Netconf configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerWCMConfigsApi.PushNetconfConfigs(context.Background(), nmsId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerWCMConfigsApi.PushNetconfConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushNetconfConfigs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerWCMConfigsApi.PushNetconfConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushNetconfConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | Netconf configuration | 

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

