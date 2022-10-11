# \ConfigurationPolicyRewriteRuleDefinitionBuilderApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#CreatePolicyDefinition19) | **Post** /template/policy/definition/rewriterule | 
[**DeletePolicyDefinition19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#DeletePolicyDefinition19) | **Delete** /template/policy/definition/rewriterule/{id} | 
[**EditMultiplePolicyDefinition19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#EditMultiplePolicyDefinition19) | **Put** /template/policy/definition/rewriterule/multiple/{id} | 
[**EditPolicyDefinition19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#EditPolicyDefinition19) | **Put** /template/policy/definition/rewriterule/{id} | 
[**GetDefinitions19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#GetDefinitions19) | **Get** /template/policy/definition/rewriterule | 
[**GetPolicyDefinition19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#GetPolicyDefinition19) | **Get** /template/policy/definition/rewriterule/{id} | 
[**PreviewPolicyDefinition19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#PreviewPolicyDefinition19) | **Post** /template/policy/definition/rewriterule/preview | 
[**PreviewPolicyDefinitionById19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#PreviewPolicyDefinitionById19) | **Get** /template/policy/definition/rewriterule/preview/{id} | 
[**SavePolicyDefinitionInBulk19**](ConfigurationPolicyRewriteRuleDefinitionBuilderApi.md#SavePolicyDefinitionInBulk19) | **Put** /template/policy/definition/rewriterule/bulk | 



## CreatePolicyDefinition19

> map[string]interface{} CreatePolicyDefinition19(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.CreatePolicyDefinition19(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.CreatePolicyDefinition19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.CreatePolicyDefinition19`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition19Request struct via the builder pattern


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


## DeletePolicyDefinition19

> DeletePolicyDefinition19(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.DeletePolicyDefinition19(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.DeletePolicyDefinition19``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyDefinition19Request struct via the builder pattern


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


## EditMultiplePolicyDefinition19

> map[string]interface{} EditMultiplePolicyDefinition19(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.EditMultiplePolicyDefinition19(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.EditMultiplePolicyDefinition19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.EditMultiplePolicyDefinition19`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition19Request struct via the builder pattern


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


## EditPolicyDefinition19

> map[string]interface{} EditPolicyDefinition19(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.EditPolicyDefinition19(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.EditPolicyDefinition19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.EditPolicyDefinition19`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition19Request struct via the builder pattern


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


## GetDefinitions19

> map[string]interface{} GetDefinitions19(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.GetDefinitions19(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.GetDefinitions19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.GetDefinitions19`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions19Request struct via the builder pattern


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


## GetPolicyDefinition19

> map[string]interface{} GetPolicyDefinition19(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.GetPolicyDefinition19(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.GetPolicyDefinition19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.GetPolicyDefinition19`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition19Request struct via the builder pattern


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


## PreviewPolicyDefinition19

> map[string]interface{} PreviewPolicyDefinition19(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.PreviewPolicyDefinition19(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.PreviewPolicyDefinition19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.PreviewPolicyDefinition19`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition19Request struct via the builder pattern


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


## PreviewPolicyDefinitionById19

> map[string]interface{} PreviewPolicyDefinitionById19(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.PreviewPolicyDefinitionById19(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.PreviewPolicyDefinitionById19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.PreviewPolicyDefinitionById19`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById19Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk19

> map[string]interface{} SavePolicyDefinitionInBulk19(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyRewriteRuleDefinitionBuilderApi.SavePolicyDefinitionInBulk19(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.SavePolicyDefinitionInBulk19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk19`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyRewriteRuleDefinitionBuilderApi.SavePolicyDefinitionInBulk19`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk19Request struct via the builder pattern


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

