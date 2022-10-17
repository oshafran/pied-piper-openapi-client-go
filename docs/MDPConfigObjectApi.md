# \MDPConfigObjectApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveMDPConfigObject**](MDPConfigObjectApi.md#RetrieveMDPConfigObject) | **Get** /mdp/policies/mdpconfig/{deviceId} | 



## RetrieveMDPConfigObject

> []map[string]interface{} RetrieveMDPConfigObject(ctx, deviceId).Execute()





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
    deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPConfigObjectApi.RetrieveMDPConfigObject(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPConfigObjectApi.RetrieveMDPConfigObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPConfigObject`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPConfigObjectApi.RetrieveMDPConfigObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPConfigObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

