# \ConfigurationVoiceTemplatePolicyApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoiceTemplate**](ConfigurationVoiceTemplatePolicyApi.md#CreateVoiceTemplate) | **Post** /template/policy/voice | 
[**DeleteVoiceTemplate**](ConfigurationVoiceTemplatePolicyApi.md#DeleteVoiceTemplate) | **Delete** /template/policy/voice/{policyId} | 
[**EditVoiceTemplate**](ConfigurationVoiceTemplatePolicyApi.md#EditVoiceTemplate) | **Put** /template/policy/voice/{policyId} | 
[**GenerateVoicePolicySummary**](ConfigurationVoiceTemplatePolicyApi.md#GenerateVoicePolicySummary) | **Get** /template/policy/voice/summary | 
[**GenerateVoiceTemplateList**](ConfigurationVoiceTemplatePolicyApi.md#GenerateVoiceTemplateList) | **Get** /template/policy/voice | 
[**GetDeviceListByPolicyId**](ConfigurationVoiceTemplatePolicyApi.md#GetDeviceListByPolicyId) | **Get** /template/policy/voice/devices/{policyId} | 
[**GetTemplateById**](ConfigurationVoiceTemplatePolicyApi.md#GetTemplateById) | **Get** /template/policy/voice/definition/{policyId} | 
[**GetVoicePolicyDeviceList**](ConfigurationVoiceTemplatePolicyApi.md#GetVoicePolicyDeviceList) | **Get** /template/policy/voice/devices | 
[**GetVoiceTemplatesForDevice**](ConfigurationVoiceTemplatePolicyApi.md#GetVoiceTemplatesForDevice) | **Get** /template/policy/voice/{deviceModel} | 



## CreateVoiceTemplate

> map[string]interface{} CreateVoiceTemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.CreateVoiceTemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.CreateVoiceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVoiceTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.CreateVoiceTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoiceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy template | 

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


## DeleteVoiceTemplate

> DeleteVoiceTemplate(ctx, policyId).Execute()





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
    policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.DeleteVoiceTemplate(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.DeleteVoiceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoiceTemplateRequest struct via the builder pattern


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


## EditVoiceTemplate

> map[string]interface{} EditVoiceTemplate(ctx, policyId).Body(body).Execute()





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
    policyId := "policyId_example" // string | Policy Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.EditVoiceTemplate(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.EditVoiceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditVoiceTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.EditVoiceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVoiceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Policy template | 

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


## GenerateVoicePolicySummary

> []map[string]interface{} GenerateVoicePolicySummary(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.GenerateVoicePolicySummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.GenerateVoicePolicySummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateVoicePolicySummary`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.GenerateVoicePolicySummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateVoicePolicySummaryRequest struct via the builder pattern


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


## GenerateVoiceTemplateList

> []map[string]interface{} GenerateVoiceTemplateList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.GenerateVoiceTemplateList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.GenerateVoiceTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateVoiceTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.GenerateVoiceTemplateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateVoiceTemplateListRequest struct via the builder pattern


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


## GetDeviceListByPolicyId

> []map[string]interface{} GetDeviceListByPolicyId(ctx, policyId).Execute()





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
    policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.GetDeviceListByPolicyId(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.GetDeviceListByPolicyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceListByPolicyId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.GetDeviceListByPolicyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceListByPolicyIdRequest struct via the builder pattern


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


## GetTemplateById

> map[string]interface{} GetTemplateById(ctx, policyId).Execute()





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
    policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.GetTemplateById(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.GetTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.GetTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByIdRequest struct via the builder pattern


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


## GetVoicePolicyDeviceList

> []map[string]interface{} GetVoicePolicyDeviceList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.GetVoicePolicyDeviceList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.GetVoicePolicyDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoicePolicyDeviceList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.GetVoicePolicyDeviceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoicePolicyDeviceListRequest struct via the builder pattern


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


## GetVoiceTemplatesForDevice

> []map[string]interface{} GetVoiceTemplatesForDevice(ctx, deviceModel).Execute()





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
    deviceModel := map[string][]openapiclient.DeviceModel{ ... } // DeviceModel | Device model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVoiceTemplatePolicyApi.GetVoiceTemplatesForDevice(context.Background(), deviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVoiceTemplatePolicyApi.GetVoiceTemplatesForDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVoiceTemplatesForDevice`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVoiceTemplatePolicyApi.GetVoiceTemplatesForDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceModel** | [**DeviceModel**](.md) | Device model | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVoiceTemplatesForDeviceRequest struct via the builder pattern


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

