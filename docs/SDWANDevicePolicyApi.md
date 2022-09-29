# \SDWANDevicePolicyApi

All URIs are relative to *https://https*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExampleComportDataserviceTemplatePolicyListGet**](SDWANDevicePolicyApi.md#ExampleComportDataserviceTemplatePolicyListGet) | **Get** //example.com:{port}/dataservice/template/policy/list | Policy List
[**ExampleComportDataserviceTemplatePolicyVedgeDevicesGet**](SDWANDevicePolicyApi.md#ExampleComportDataserviceTemplatePolicyVedgeDevicesGet) | **Get** //example.com:{port}/dataservice/template/policy/vedge/devices | vEdge Template Policy



## ExampleComportDataserviceTemplatePolicyListGet

> string ExampleComportDataserviceTemplatePolicyListGet(ctx, port).XXSRFTOKEN(xXSRFTOKEN).Execute()

Policy List

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
    port := "port_example" // string | 
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANDevicePolicyApi.ExampleComportDataserviceTemplatePolicyListGet(context.Background(), port).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANDevicePolicyApi.ExampleComportDataserviceTemplatePolicyListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceTemplatePolicyListGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANDevicePolicyApi.ExampleComportDataserviceTemplatePolicyListGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceTemplatePolicyListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xXSRFTOKEN** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExampleComportDataserviceTemplatePolicyVedgeDevicesGet

> string ExampleComportDataserviceTemplatePolicyVedgeDevicesGet(ctx, port).Execute()

vEdge Template Policy

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
    port := "port_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDWANDevicePolicyApi.ExampleComportDataserviceTemplatePolicyVedgeDevicesGet(context.Background(), port).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANDevicePolicyApi.ExampleComportDataserviceTemplatePolicyVedgeDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceTemplatePolicyVedgeDevicesGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANDevicePolicyApi.ExampleComportDataserviceTemplatePolicyVedgeDevicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceTemplatePolicyVedgeDevicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

