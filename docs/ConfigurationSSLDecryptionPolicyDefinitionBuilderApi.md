# \ConfigurationSSLDecryptionPolicyDefinitionBuilderApi

All URIs are relative to *http://$VMANAGE_EXTERNAL_IP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyDefinition3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#CreatePolicyDefinition3) | **Post** /template/policy/definition/ssldecryption | 
[**DeletePolicyDefinition3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#DeletePolicyDefinition3) | **Delete** /template/policy/definition/ssldecryption/{id} | 
[**EditMultiplePolicyDefinition3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#EditMultiplePolicyDefinition3) | **Put** /template/policy/definition/ssldecryption/multiple/{id} | 
[**EditPolicyDefinition3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#EditPolicyDefinition3) | **Put** /template/policy/definition/ssldecryption/{id} | 
[**GetDefinitions3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#GetDefinitions3) | **Get** /template/policy/definition/ssldecryption | 
[**GetPolicyDefinition3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#GetPolicyDefinition3) | **Get** /template/policy/definition/ssldecryption/{id} | 
[**PreviewPolicyDefinition3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#PreviewPolicyDefinition3) | **Post** /template/policy/definition/ssldecryption/preview | 
[**PreviewPolicyDefinitionById3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#PreviewPolicyDefinitionById3) | **Get** /template/policy/definition/ssldecryption/preview/{id} | 
[**SavePolicyDefinitionInBulk3**](ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.md#SavePolicyDefinitionInBulk3) | **Put** /template/policy/definition/ssldecryption/bulk | 



## CreatePolicyDefinition3

> map[string]interface{} CreatePolicyDefinition3(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.CreatePolicyDefinition3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.CreatePolicyDefinition3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyDefinition3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.CreatePolicyDefinition3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyDefinition3Request struct via the builder pattern


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


## DeletePolicyDefinition3

> DeletePolicyDefinition3(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.DeletePolicyDefinition3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.DeletePolicyDefinition3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyDefinition3Request struct via the builder pattern


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


## EditMultiplePolicyDefinition3

> map[string]interface{} EditMultiplePolicyDefinition3(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.EditMultiplePolicyDefinition3(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.EditMultiplePolicyDefinition3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultiplePolicyDefinition3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.EditMultiplePolicyDefinition3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultiplePolicyDefinition3Request struct via the builder pattern


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


## EditPolicyDefinition3

> map[string]interface{} EditPolicyDefinition3(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.EditPolicyDefinition3(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.EditPolicyDefinition3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyDefinition3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.EditPolicyDefinition3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyDefinition3Request struct via the builder pattern


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


## GetDefinitions3

> map[string]interface{} GetDefinitions3(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.GetDefinitions3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.GetDefinitions3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefinitions3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.GetDefinitions3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefinitions3Request struct via the builder pattern


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


## GetPolicyDefinition3

> map[string]interface{} GetPolicyDefinition3(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.GetPolicyDefinition3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.GetPolicyDefinition3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyDefinition3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.GetPolicyDefinition3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyDefinition3Request struct via the builder pattern


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


## PreviewPolicyDefinition3

> map[string]interface{} PreviewPolicyDefinition3(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.PreviewPolicyDefinition3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.PreviewPolicyDefinition3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinition3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.PreviewPolicyDefinition3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinition3Request struct via the builder pattern


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


## PreviewPolicyDefinitionById3

> map[string]interface{} PreviewPolicyDefinitionById3(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.PreviewPolicyDefinitionById3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.PreviewPolicyDefinitionById3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyDefinitionById3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.PreviewPolicyDefinitionById3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyDefinitionById3Request struct via the builder pattern


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


## SavePolicyDefinitionInBulk3

> map[string]interface{} SavePolicyDefinitionInBulk3(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.SavePolicyDefinitionInBulk3(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.SavePolicyDefinitionInBulk3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SavePolicyDefinitionInBulk3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSSLDecryptionPolicyDefinitionBuilderApi.SavePolicyDefinitionInBulk3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePolicyDefinitionInBulk3Request struct via the builder pattern


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

