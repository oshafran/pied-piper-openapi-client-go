# \ConfigurationPolicyVEdgeRouteDefinitionBuilderApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#CreatePolicyDefinition24) | **Post** /template/policy/definition/vedgeroute | 
[**DeletePolicyDefinition24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#DeletePolicyDefinition24) | **Delete** /template/policy/definition/vedgeroute/{id} | 
[**EditMultiplePolicyDefinition24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#EditMultiplePolicyDefinition24) | **Put** /template/policy/definition/vedgeroute/multiple/{id} | 
[**EditPolicyDefinition24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#EditPolicyDefinition24) | **Put** /template/policy/definition/vedgeroute/{id} | 
[**GetDefinitions24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#GetDefinitions24) | **Get** /template/policy/definition/vedgeroute | 
[**GetPolicyDefinition24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#GetPolicyDefinition24) | **Get** /template/policy/definition/vedgeroute/{id} | 
[**PreviewPolicyDefinition24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#PreviewPolicyDefinition24) | **Post** /template/policy/definition/vedgeroute/preview | 
[**PreviewPolicyDefinitionById24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#PreviewPolicyDefinitionById24) | **Get** /template/policy/definition/vedgeroute/preview/{id} | 
[**SavePolicyDefinitionInBulk24**](ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.md#SavePolicyDefinitionInBulk24) | **Put** /template/policy/definition/vedgeroute/bulk | 



## CreatePolicyDefinition24

> map[string]interface{} CreatePolicyDefinition24(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.CreatePolicyDefinition24(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.CreatePolicyDefinition24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.CreatePolicyDefinition24`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition24Request struct via the builder pattern


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


## DeletePolicyDefinition24

> DeletePolicyDefinition24(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.DeletePolicyDefinition24(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.DeletePolicyDefinition24``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyDefinition24Request struct via the builder pattern


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


## EditMultiplePolicyDefinition24

> map[string]interface{} EditMultiplePolicyDefinition24(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.EditMultiplePolicyDefinition24(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.EditMultiplePolicyDefinition24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.EditMultiplePolicyDefinition24`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition24Request struct via the builder pattern


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


## EditPolicyDefinition24

> map[string]interface{} EditPolicyDefinition24(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.EditPolicyDefinition24(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.EditPolicyDefinition24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.EditPolicyDefinition24`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition24Request struct via the builder pattern


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


## GetDefinitions24

> map[string]interface{} GetDefinitions24(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.GetDefinitions24(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.GetDefinitions24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.GetDefinitions24`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions24Request struct via the builder pattern


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


## GetPolicyDefinition24

> map[string]interface{} GetPolicyDefinition24(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.GetPolicyDefinition24(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.GetPolicyDefinition24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.GetPolicyDefinition24`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition24Request struct via the builder pattern


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


## PreviewPolicyDefinition24

> map[string]interface{} PreviewPolicyDefinition24(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.PreviewPolicyDefinition24(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.PreviewPolicyDefinition24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.PreviewPolicyDefinition24`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition24Request struct via the builder pattern


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


## PreviewPolicyDefinitionById24

> map[string]interface{} PreviewPolicyDefinitionById24(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.PreviewPolicyDefinitionById24(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.PreviewPolicyDefinitionById24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.PreviewPolicyDefinitionById24`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById24Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk24

> map[string]interface{} SavePolicyDefinitionInBulk24(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.SavePolicyDefinitionInBulk24(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.SavePolicyDefinitionInBulk24``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk24`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyVEdgeRouteDefinitionBuilderApi.SavePolicyDefinitionInBulk24`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk24Request struct via the builder pattern


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

