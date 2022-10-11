# \ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#CreatePolicyDefinition17) | **Post** /template/policy/definition/deviceaccesspolicyv6 | 
[**DeletePolicyDefinition17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#DeletePolicyDefinition17) | **Delete** /template/policy/definition/deviceaccesspolicyv6/{id} | 
[**EditMultiplePolicyDefinition17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#EditMultiplePolicyDefinition17) | **Put** /template/policy/definition/deviceaccesspolicyv6/multiple/{id} | 
[**EditPolicyDefinition17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#EditPolicyDefinition17) | **Put** /template/policy/definition/deviceaccesspolicyv6/{id} | 
[**GetDefinitions17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#GetDefinitions17) | **Get** /template/policy/definition/deviceaccesspolicyv6 | 
[**GetPolicyDefinition17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#GetPolicyDefinition17) | **Get** /template/policy/definition/deviceaccesspolicyv6/{id} | 
[**PreviewPolicyDefinition17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#PreviewPolicyDefinition17) | **Post** /template/policy/definition/deviceaccesspolicyv6/preview | 
[**PreviewPolicyDefinitionById17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#PreviewPolicyDefinitionById17) | **Get** /template/policy/definition/deviceaccesspolicyv6/preview/{id} | 
[**SavePolicyDefinitionInBulk17**](ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.md#SavePolicyDefinitionInBulk17) | **Put** /template/policy/definition/deviceaccesspolicyv6/bulk | 



## CreatePolicyDefinition17

> map[string]interface{} CreatePolicyDefinition17(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.CreatePolicyDefinition17(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.CreatePolicyDefinition17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.CreatePolicyDefinition17`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy definition | 

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


## DeletePolicyDefinition17

> DeletePolicyDefinition17(ctx, id).Execute()





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
    id := "id_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.DeletePolicyDefinition17(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.DeletePolicyDefinition17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyDefinition17Request struct via the builder pattern


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


## EditMultiplePolicyDefinition17

> map[string]interface{} EditMultiplePolicyDefinition17(ctx, id).Body(body).Execute()





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
    id := "id_example" // string | Policy Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.EditMultiplePolicyDefinition17(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.EditMultiplePolicyDefinition17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.EditMultiplePolicyDefinition17`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Policy definition | 

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


## EditPolicyDefinition17

> map[string]interface{} EditPolicyDefinition17(ctx, id).Body(body).Execute()





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
    id := "id_example" // string | Policy Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.EditPolicyDefinition17(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.EditPolicyDefinition17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.EditPolicyDefinition17`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Policy definition | 

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


## GetDefinitions17

> map[string]interface{} GetDefinitions17(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.GetDefinitions17(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.GetDefinitions17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.GetDefinitions17`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions17Request struct via the builder pattern


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


## GetPolicyDefinition17

> map[string]interface{} GetPolicyDefinition17(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.GetPolicyDefinition17(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.GetPolicyDefinition17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.GetPolicyDefinition17`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition17Request struct via the builder pattern


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


## PreviewPolicyDefinition17

> map[string]interface{} PreviewPolicyDefinition17(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.PreviewPolicyDefinition17(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.PreviewPolicyDefinition17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.PreviewPolicyDefinition17`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy definition | 

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


## PreviewPolicyDefinitionById17

> map[string]interface{} PreviewPolicyDefinitionById17(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.PreviewPolicyDefinitionById17(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.PreviewPolicyDefinitionById17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.PreviewPolicyDefinitionById17`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById17Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk17

> map[string]interface{} SavePolicyDefinitionInBulk17(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.SavePolicyDefinitionInBulk17(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.SavePolicyDefinitionInBulk17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk17`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyIPv6DeviceAccessListDefinitionBuilderApi.SavePolicyDefinitionInBulk17`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy definition | 

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

