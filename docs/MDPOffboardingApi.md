# \MDPOffboardingApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisconnectFromMdp**](MDPOffboardingApi.md#DisconnectFromMdp) | **Get** /mdp/disconnect/{nmsId} | 
[**Offboard**](MDPOffboardingApi.md#Offboard) | **Delete** /mdp/onboard/{nmsId} | 



## DisconnectFromMdp

> []map[string]interface{} DisconnectFromMdp(ctx, nmsId).Execute()





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
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPOffboardingApi.DisconnectFromMdp(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPOffboardingApi.DisconnectFromMdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisconnectFromMdp`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPOffboardingApi.DisconnectFromMdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectFromMdpRequest struct via the builder pattern


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


## Offboard

> Offboard(ctx, nmsId).Execute()





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
    nmsId := "nmsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPOffboardingApi.Offboard(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPOffboardingApi.Offboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOffboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

