# \ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#CreatePolicyDefinition7) | **Post** /template/policy/definition/zonebasedfw | 
[**DeletePolicyDefinition7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#DeletePolicyDefinition7) | **Delete** /template/policy/definition/zonebasedfw/{id} | 
[**EditMultiplePolicyDefinition7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#EditMultiplePolicyDefinition7) | **Put** /template/policy/definition/zonebasedfw/multiple/{id} | 
[**EditPolicyDefinition7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#EditPolicyDefinition7) | **Put** /template/policy/definition/zonebasedfw/{id} | 
[**GetDefinitions7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#GetDefinitions7) | **Get** /template/policy/definition/zonebasedfw | 
[**GetPolicyDefinition7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#GetPolicyDefinition7) | **Get** /template/policy/definition/zonebasedfw/{id} | 
[**PreviewPolicyDefinition7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#PreviewPolicyDefinition7) | **Post** /template/policy/definition/zonebasedfw/preview | 
[**PreviewPolicyDefinitionById7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#PreviewPolicyDefinitionById7) | **Get** /template/policy/definition/zonebasedfw/preview/{id} | 
[**SavePolicyDefinitionInBulk7**](ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.md#SavePolicyDefinitionInBulk7) | **Put** /template/policy/definition/zonebasedfw/bulk | 



## CreatePolicyDefinition7

> map[string]interface{} CreatePolicyDefinition7(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.CreatePolicyDefinition7(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.CreatePolicyDefinition7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.CreatePolicyDefinition7`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition7Request struct via the builder pattern


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


## DeletePolicyDefinition7

> DeletePolicyDefinition7(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.DeletePolicyDefinition7(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.DeletePolicyDefinition7``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyDefinition7Request struct via the builder pattern


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


## EditMultiplePolicyDefinition7

> map[string]interface{} EditMultiplePolicyDefinition7(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.EditMultiplePolicyDefinition7(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.EditMultiplePolicyDefinition7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.EditMultiplePolicyDefinition7`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition7Request struct via the builder pattern


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


## EditPolicyDefinition7

> map[string]interface{} EditPolicyDefinition7(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.EditPolicyDefinition7(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.EditPolicyDefinition7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.EditPolicyDefinition7`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition7Request struct via the builder pattern


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


## GetDefinitions7

> map[string]interface{} GetDefinitions7(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.GetDefinitions7(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.GetDefinitions7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.GetDefinitions7`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions7Request struct via the builder pattern


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


## GetPolicyDefinition7

> map[string]interface{} GetPolicyDefinition7(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.GetPolicyDefinition7(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.GetPolicyDefinition7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.GetPolicyDefinition7`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition7Request struct via the builder pattern


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


## PreviewPolicyDefinition7

> map[string]interface{} PreviewPolicyDefinition7(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.PreviewPolicyDefinition7(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.PreviewPolicyDefinition7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.PreviewPolicyDefinition7`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition7Request struct via the builder pattern


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


## PreviewPolicyDefinitionById7

> map[string]interface{} PreviewPolicyDefinitionById7(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.PreviewPolicyDefinitionById7(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.PreviewPolicyDefinitionById7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.PreviewPolicyDefinitionById7`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById7Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk7

> map[string]interface{} SavePolicyDefinitionInBulk7(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.SavePolicyDefinitionInBulk7(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.SavePolicyDefinitionInBulk7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk7`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyZoneBasedFirewallDefinitionBuilderApi.SavePolicyDefinitionInBulk7`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk7Request struct via the builder pattern


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

