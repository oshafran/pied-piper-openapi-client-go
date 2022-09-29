# \SDWANDeviceTemplateApi

All URIs are relative to *https://https*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExampleComportDataserviceTemplateFeatureGet**](SDWANDeviceTemplateApi.md#ExampleComportDataserviceTemplateFeatureGet) | **Get** //example.com:{port}/dataservice/template/feature | Template Feature
[**ExampleComportDataserviceTemplateFeatureTypesGet**](SDWANDeviceTemplateApi.md#ExampleComportDataserviceTemplateFeatureTypesGet) | **Get** //example.com:{port}/dataservice/template/feature/types | Template Feature Type



## ExampleComportDataserviceTemplateFeatureGet

> string ExampleComportDataserviceTemplateFeatureGet(ctx, port).XXSRFTOKEN(xXSRFTOKEN).Execute()

Template Feature

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
    resp, r, err := apiClient.SDWANDeviceTemplateApi.ExampleComportDataserviceTemplateFeatureGet(context.Background(), port).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANDeviceTemplateApi.ExampleComportDataserviceTemplateFeatureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceTemplateFeatureGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANDeviceTemplateApi.ExampleComportDataserviceTemplateFeatureGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceTemplateFeatureGetRequest struct via the builder pattern


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


## ExampleComportDataserviceTemplateFeatureTypesGet

> string ExampleComportDataserviceTemplateFeatureTypesGet(ctx, port).XXSRFTOKEN(xXSRFTOKEN).Execute()

Template Feature Type

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
    resp, r, err := apiClient.SDWANDeviceTemplateApi.ExampleComportDataserviceTemplateFeatureTypesGet(context.Background(), port).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDWANDeviceTemplateApi.ExampleComportDataserviceTemplateFeatureTypesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceTemplateFeatureTypesGet`: string
    fmt.Fprintf(os.Stdout, "Response from `SDWANDeviceTemplateApi.ExampleComportDataserviceTemplateFeatureTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceTemplateFeatureTypesGetRequest struct via the builder pattern


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

