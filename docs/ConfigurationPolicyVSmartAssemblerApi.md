# \ConfigurationPolicyVSmartAssemblerApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Preview3**](ConfigurationPolicyVSmartAssemblerApi.md#Preview3) | **Post** /template/policy/assembly/vsmart | 
[**PreviewById3**](ConfigurationPolicyVSmartAssemblerApi.md#PreviewById3) | **Get** /template/policy/assembly/vsmart/{id} | 



## Preview3

> map[string]interface{} Preview3(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy assembly (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyVSmartAssemblerApi.Preview3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVSmartAssemblerApi.Preview3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Preview3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVSmartAssemblerApi.Preview3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreview3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy assembly | 

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


## PreviewById3

> map[string]interface{} PreviewById3(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyVSmartAssemblerApi.PreviewById3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVSmartAssemblerApi.PreviewById3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewById3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVSmartAssemblerApi.PreviewById3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewById3Request struct via the builder pattern


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

