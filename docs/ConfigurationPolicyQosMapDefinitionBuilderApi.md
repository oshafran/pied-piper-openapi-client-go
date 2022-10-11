# \ConfigurationPolicyQosMapDefinitionBuilderApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#CreatePolicyDefinition1) | **Post** /template/policy/definition/qosmap | 
[**DeletePolicyDefinition1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#DeletePolicyDefinition1) | **Delete** /template/policy/definition/qosmap/{id} | 
[**EditMultiplePolicyDefinition1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#EditMultiplePolicyDefinition1) | **Put** /template/policy/definition/qosmap/multiple/{id} | 
[**EditPolicyDefinition1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#EditPolicyDefinition1) | **Put** /template/policy/definition/qosmap/{id} | 
[**GetDefinitions1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#GetDefinitions1) | **Get** /template/policy/definition/qosmap | 
[**GetPolicyDefinition1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#GetPolicyDefinition1) | **Get** /template/policy/definition/qosmap/{id} | 
[**PreviewPolicyDefinition1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#PreviewPolicyDefinition1) | **Post** /template/policy/definition/qosmap/preview | 
[**PreviewPolicyDefinitionById1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#PreviewPolicyDefinitionById1) | **Get** /template/policy/definition/qosmap/preview/{id} | 
[**SavePolicyDefinitionInBulk1**](ConfigurationPolicyQosMapDefinitionBuilderApi.md#SavePolicyDefinitionInBulk1) | **Put** /template/policy/definition/qosmap/bulk | 



## CreatePolicyDefinition1

> map[string]interface{} CreatePolicyDefinition1(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.CreatePolicyDefinition1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.CreatePolicyDefinition1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.CreatePolicyDefinition1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition1Request struct via the builder pattern


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


## DeletePolicyDefinition1

> DeletePolicyDefinition1(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.DeletePolicyDefinition1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.DeletePolicyDefinition1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyDefinition1Request struct via the builder pattern


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


## EditMultiplePolicyDefinition1

> map[string]interface{} EditMultiplePolicyDefinition1(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.EditMultiplePolicyDefinition1(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.EditMultiplePolicyDefinition1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.EditMultiplePolicyDefinition1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition1Request struct via the builder pattern


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


## EditPolicyDefinition1

> map[string]interface{} EditPolicyDefinition1(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.EditPolicyDefinition1(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.EditPolicyDefinition1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.EditPolicyDefinition1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition1Request struct via the builder pattern


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


## GetDefinitions1

> map[string]interface{} GetDefinitions1(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.GetDefinitions1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.GetDefinitions1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.GetDefinitions1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions1Request struct via the builder pattern


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


## GetPolicyDefinition1

> map[string]interface{} GetPolicyDefinition1(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.GetPolicyDefinition1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.GetPolicyDefinition1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.GetPolicyDefinition1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition1Request struct via the builder pattern


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


## PreviewPolicyDefinition1

> map[string]interface{} PreviewPolicyDefinition1(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinition1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinition1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinition1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition1Request struct via the builder pattern


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


## PreviewPolicyDefinitionById1

> map[string]interface{} PreviewPolicyDefinitionById1(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinitionById1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinitionById1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinitionById1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById1Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk1

> map[string]interface{} SavePolicyDefinitionInBulk1(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.SavePolicyDefinitionInBulk1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyQosMapDefinitionBuilderApi.SavePolicyDefinitionInBulk1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyQosMapDefinitionBuilderApi.SavePolicyDefinitionInBulk1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk1Request struct via the builder pattern


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

