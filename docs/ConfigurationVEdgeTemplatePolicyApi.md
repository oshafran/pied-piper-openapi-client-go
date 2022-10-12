# \ConfigurationVEdgeTemplatePolicyApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePolicyResourceGroup**](ConfigurationVEdgeTemplatePolicyApi.md#ChangePolicyResourceGroup) | **Post** /template/policy/vedge/{resourceGroupName}/{policyId} | 
[**CreateVEdgeTemplate**](ConfigurationVEdgeTemplatePolicyApi.md#CreateVEdgeTemplate) | **Post** /template/policy/vedge | 
[**DeleteVEdgeTemplate**](ConfigurationVEdgeTemplatePolicyApi.md#DeleteVEdgeTemplate) | **Delete** /template/policy/vedge/{policyId} | 
[**EditVEdgeTemplate**](ConfigurationVEdgeTemplatePolicyApi.md#EditVEdgeTemplate) | **Put** /template/policy/vedge/{policyId} | 
[**GeneratePolicyTemplateList**](ConfigurationVEdgeTemplatePolicyApi.md#GeneratePolicyTemplateList) | **Get** /template/policy/vedge | 
[**GetDeviceListByPolicy**](ConfigurationVEdgeTemplatePolicyApi.md#GetDeviceListByPolicy) | **Get** /template/policy/vedge/devices/{policyId} | 
[**GetVEdgePolicyDeviceList**](ConfigurationVEdgeTemplatePolicyApi.md#GetVEdgePolicyDeviceList) | **Get** /template/policy/vedge/devices | 
[**GetVEdgeTemplate**](ConfigurationVEdgeTemplatePolicyApi.md#GetVEdgeTemplate) | **Get** /template/policy/vedge/definition/{policyId} | 



## ChangePolicyResourceGroup

> ChangePolicyResourceGroup(ctx, policyId, resourceGroupName).Execute()





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
    resourceGroupName := "resourceGroupName_example" // string | Resrouce group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.ChangePolicyResourceGroup(context.Background(), policyId, resourceGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.ChangePolicyResourceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 
**resourceGroupName** | **string** | Resrouce group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePolicyResourceGroupRequest struct via the builder pattern


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


## CreateVEdgeTemplate

> map[string]interface{} CreateVEdgeTemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.CreateVEdgeTemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.CreateVEdgeTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVEdgeTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVEdgeTemplatePolicyApi.CreateVEdgeTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVEdgeTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Template policy | 

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


## DeleteVEdgeTemplate

> DeleteVEdgeTemplate(ctx, policyId).Execute()





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
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.DeleteVEdgeTemplate(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.DeleteVEdgeTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVEdgeTemplateRequest struct via the builder pattern


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


## EditVEdgeTemplate

> map[string]interface{} EditVEdgeTemplate(ctx, policyId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.EditVEdgeTemplate(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.EditVEdgeTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditVEdgeTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVEdgeTemplatePolicyApi.EditVEdgeTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVEdgeTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Template policy | 

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


## GeneratePolicyTemplateList

> map[string]interface{} GeneratePolicyTemplateList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.GeneratePolicyTemplateList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.GeneratePolicyTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePolicyTemplateList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVEdgeTemplatePolicyApi.GeneratePolicyTemplateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePolicyTemplateListRequest struct via the builder pattern


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


## GetDeviceListByPolicy

> []map[string]interface{} GetDeviceListByPolicy(ctx, policyId).Execute()





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
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.GetDeviceListByPolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.GetDeviceListByPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceListByPolicy`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVEdgeTemplatePolicyApi.GetDeviceListByPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceListByPolicyRequest struct via the builder pattern


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


## GetVEdgePolicyDeviceList

> []map[string]interface{} GetVEdgePolicyDeviceList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.GetVEdgePolicyDeviceList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.GetVEdgePolicyDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVEdgePolicyDeviceList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVEdgeTemplatePolicyApi.GetVEdgePolicyDeviceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVEdgePolicyDeviceListRequest struct via the builder pattern


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


## GetVEdgeTemplate

> map[string]interface{} GetVEdgeTemplate(ctx, policyId).Execute()





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
    resp, r, err := apiClient.ConfigurationVEdgeTemplatePolicyApi.GetVEdgeTemplate(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVEdgeTemplatePolicyApi.GetVEdgeTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVEdgeTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVEdgeTemplatePolicyApi.GetVEdgeTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVEdgeTemplateRequest struct via the builder pattern


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

