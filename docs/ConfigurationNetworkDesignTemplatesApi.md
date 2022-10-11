# \ConfigurationNetworkDesignTemplatesApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditDeviceProfileTemplate**](ConfigurationNetworkDesignTemplatesApi.md#EditDeviceProfileTemplate) | **Put** /networkdesign/profile/template/{templateId} | 
[**EditGlobalTemplate**](ConfigurationNetworkDesignTemplatesApi.md#EditGlobalTemplate) | **Put** /networkdesign/global/template/{templateId} | 
[**GenerateProfileTemplateList**](ConfigurationNetworkDesignTemplatesApi.md#GenerateProfileTemplateList) | **Get** /networkdesign/profile/template | 
[**GetDeviceProfileFeatureTemplateList**](ConfigurationNetworkDesignTemplatesApi.md#GetDeviceProfileFeatureTemplateList) | **Get** /networkdesign/profile/feature | 
[**GetDeviceProfileTemplate**](ConfigurationNetworkDesignTemplatesApi.md#GetDeviceProfileTemplate) | **Get** /networkdesign/profile/template/{templateId} | 
[**GetGlobalTemplate**](ConfigurationNetworkDesignTemplatesApi.md#GetGlobalTemplate) | **Get** /networkdesign/global/template/{templateId} | 



## EditDeviceProfileTemplate

> EditDeviceProfileTemplate(ctx, templateId).Body(body).Execute()





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
    templateId := "templateId_example" // string | Template Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Global template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignTemplatesApi.EditDeviceProfileTemplate(context.Background(), templateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignTemplatesApi.EditDeviceProfileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDeviceProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Global template | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditGlobalTemplate

> EditGlobalTemplate(ctx, templateId).Body(body).Execute()





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
    templateId := "templateId_example" // string | Template Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Global template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignTemplatesApi.EditGlobalTemplate(context.Background(), templateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignTemplatesApi.EditGlobalTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditGlobalTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Global template | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateProfileTemplateList

> []map[string]interface{} GenerateProfileTemplateList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignTemplatesApi.GenerateProfileTemplateList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignTemplatesApi.GenerateProfileTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateProfileTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignTemplatesApi.GenerateProfileTemplateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateProfileTemplateListRequest struct via the builder pattern


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


## GetDeviceProfileFeatureTemplateList

> []map[string]interface{} GetDeviceProfileFeatureTemplateList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignTemplatesApi.GetDeviceProfileFeatureTemplateList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignTemplatesApi.GetDeviceProfileFeatureTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileFeatureTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignTemplatesApi.GetDeviceProfileFeatureTemplateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileFeatureTemplateListRequest struct via the builder pattern


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


## GetDeviceProfileTemplate

> map[string]interface{} GetDeviceProfileTemplate(ctx, templateId).Execute()





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
    templateId := "templateId_example" // string | Template Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignTemplatesApi.GetDeviceProfileTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignTemplatesApi.GetDeviceProfileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignTemplatesApi.GetDeviceProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileTemplateRequest struct via the builder pattern


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


## GetGlobalTemplate

> map[string]interface{} GetGlobalTemplate(ctx, templateId).Execute()





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
    templateId := "templateId_example" // string | Template Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignTemplatesApi.GetGlobalTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignTemplatesApi.GetGlobalTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignTemplatesApi.GetGlobalTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalTemplateRequest struct via the builder pattern


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

