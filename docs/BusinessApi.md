# \BusinessApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataserviceSystemDeviceVedgesGet**](BusinessApi.md#DataserviceSystemDeviceVedgesGet) | **Get** /dataservice/system/device/vedges | Get  Device VEdges
[**DataserviceTemplatePolicyListSiteGet**](BusinessApi.md#DataserviceTemplatePolicyListSiteGet) | **Get** /dataservice/template/policy/list/site | Get  Prefix Template List



## DataserviceSystemDeviceVedgesGet

> string DataserviceSystemDeviceVedgesGet(ctx).Execute()

Get  Device VEdges

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
    resp, r, err := apiClient.BusinessApi.DataserviceSystemDeviceVedgesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceSystemDeviceVedgesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceSystemDeviceVedgesGet`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceSystemDeviceVedgesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceSystemDeviceVedgesGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataserviceTemplatePolicyListSiteGet

> string DataserviceTemplatePolicyListSiteGet(ctx).Execute()

Get  Prefix Template List

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
    resp, r, err := apiClient.BusinessApi.DataserviceTemplatePolicyListSiteGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceTemplatePolicyListSiteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceTemplatePolicyListSiteGet`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceTemplatePolicyListSiteGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceTemplatePolicyListSiteGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

