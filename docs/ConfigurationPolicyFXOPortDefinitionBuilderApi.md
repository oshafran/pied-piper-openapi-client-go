# \ConfigurationPolicyFXOPortDefinitionBuilderApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#CreatePolicyDefinition26) | **Post** /template/policy/definition/fxoport | 
[**DeletePolicyDefinition26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#DeletePolicyDefinition26) | **Delete** /template/policy/definition/fxoport/{id} | 
[**EditMultiplePolicyDefinition26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#EditMultiplePolicyDefinition26) | **Put** /template/policy/definition/fxoport/multiple/{id} | 
[**EditPolicyDefinition26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#EditPolicyDefinition26) | **Put** /template/policy/definition/fxoport/{id} | 
[**GetDefinitions26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#GetDefinitions26) | **Get** /template/policy/definition/fxoport | 
[**GetPolicyDefinition26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#GetPolicyDefinition26) | **Get** /template/policy/definition/fxoport/{id} | 
[**PreviewPolicyDefinition26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#PreviewPolicyDefinition26) | **Post** /template/policy/definition/fxoport/preview | 
[**PreviewPolicyDefinitionById26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#PreviewPolicyDefinitionById26) | **Get** /template/policy/definition/fxoport/preview/{id} | 
[**SavePolicyDefinitionInBulk26**](ConfigurationPolicyFXOPortDefinitionBuilderApi.md#SavePolicyDefinitionInBulk26) | **Put** /template/policy/definition/fxoport/bulk | 



## CreatePolicyDefinition26

> map[string]interface{} CreatePolicyDefinition26(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.CreatePolicyDefinition26(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.CreatePolicyDefinition26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.CreatePolicyDefinition26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition26Request struct via the builder pattern


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


## DeletePolicyDefinition26

> DeletePolicyDefinition26(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.DeletePolicyDefinition26(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.DeletePolicyDefinition26``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyDefinition26Request struct via the builder pattern


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


## EditMultiplePolicyDefinition26

> map[string]interface{} EditMultiplePolicyDefinition26(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.EditMultiplePolicyDefinition26(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.EditMultiplePolicyDefinition26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.EditMultiplePolicyDefinition26`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition26Request struct via the builder pattern


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


## EditPolicyDefinition26

> map[string]interface{} EditPolicyDefinition26(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.EditPolicyDefinition26(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.EditPolicyDefinition26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.EditPolicyDefinition26`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition26Request struct via the builder pattern


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


## GetDefinitions26

> map[string]interface{} GetDefinitions26(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.GetDefinitions26(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.GetDefinitions26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.GetDefinitions26`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions26Request struct via the builder pattern


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


## GetPolicyDefinition26

> map[string]interface{} GetPolicyDefinition26(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.GetPolicyDefinition26(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.GetPolicyDefinition26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.GetPolicyDefinition26`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition26Request struct via the builder pattern


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


## PreviewPolicyDefinition26

> map[string]interface{} PreviewPolicyDefinition26(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.PreviewPolicyDefinition26(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.PreviewPolicyDefinition26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.PreviewPolicyDefinition26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition26Request struct via the builder pattern


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


## PreviewPolicyDefinitionById26

> map[string]interface{} PreviewPolicyDefinitionById26(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.PreviewPolicyDefinitionById26(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.PreviewPolicyDefinitionById26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.PreviewPolicyDefinitionById26`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById26Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk26

> map[string]interface{} SavePolicyDefinitionInBulk26(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFXOPortDefinitionBuilderApi.SavePolicyDefinitionInBulk26(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFXOPortDefinitionBuilderApi.SavePolicyDefinitionInBulk26``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk26`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFXOPortDefinitionBuilderApi.SavePolicyDefinitionInBulk26`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk26Request struct via the builder pattern


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

