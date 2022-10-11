# \ConfigurationFQDNListBuilderApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyList16**](ConfigurationFQDNListBuilderApi.md#CreatePolicyList16) | **Post** /template/policy/list/fqdn | 
[**DeletePolicyList16**](ConfigurationFQDNListBuilderApi.md#DeletePolicyList16) | **Delete** /template/policy/list/fqdn/{id} | 
[**DeletePolicyListsWithInfoTag16**](ConfigurationFQDNListBuilderApi.md#DeletePolicyListsWithInfoTag16) | **Delete** /template/policy/list/fqdn | 
[**EditPolicyList16**](ConfigurationFQDNListBuilderApi.md#EditPolicyList16) | **Put** /template/policy/list/fqdn/{id} | 
[**GetListsById16**](ConfigurationFQDNListBuilderApi.md#GetListsById16) | **Get** /template/policy/list/fqdn/{id} | 
[**GetPolicyLists14**](ConfigurationFQDNListBuilderApi.md#GetPolicyLists14) | **Get** /template/policy/list/fqdn | 
[**GetPolicyListsWithInfoTag16**](ConfigurationFQDNListBuilderApi.md#GetPolicyListsWithInfoTag16) | **Get** /template/policy/list/fqdn/filtered | 
[**PreviewPolicyList16**](ConfigurationFQDNListBuilderApi.md#PreviewPolicyList16) | **Post** /template/policy/list/fqdn/preview | 
[**PreviewPolicyListById16**](ConfigurationFQDNListBuilderApi.md#PreviewPolicyListById16) | **Get** /template/policy/list/fqdn/preview/{id} | 



## CreatePolicyList16

> map[string]interface{} CreatePolicyList16(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.CreatePolicyList16(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.CreatePolicyList16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyList16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.CreatePolicyList16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyList16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy list | 

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


## DeletePolicyList16

> DeletePolicyList16(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.DeletePolicyList16(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.DeletePolicyList16``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyList16Request struct via the builder pattern


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


## DeletePolicyListsWithInfoTag16

> []map[string]interface{} DeletePolicyListsWithInfoTag16(ctx).InfoTag(infoTag).Execute()





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
    infoTag := "infoTag_example" // string | InfoTag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.DeletePolicyListsWithInfoTag16(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.DeletePolicyListsWithInfoTag16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyListsWithInfoTag16`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.DeletePolicyListsWithInfoTag16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyListsWithInfoTag16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infoTag** | **string** | InfoTag | 

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


## EditPolicyList16

> map[string]interface{} EditPolicyList16(ctx, id).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.EditPolicyList16(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.EditPolicyList16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyList16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.EditPolicyList16`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyList16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Policy list | 

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


## GetListsById16

> map[string]interface{} GetListsById16(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.GetListsById16(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.GetListsById16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListsById16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.GetListsById16`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsById16Request struct via the builder pattern


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


## GetPolicyLists14

> []map[string]interface{} GetPolicyLists14(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.GetPolicyLists14(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.GetPolicyLists14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyLists14`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.GetPolicyLists14`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyLists14Request struct via the builder pattern


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


## GetPolicyListsWithInfoTag16

> []map[string]interface{} GetPolicyListsWithInfoTag16(ctx).InfoTag(infoTag).Execute()





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
    infoTag := "infoTag_example" // string | InfoTag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.GetPolicyListsWithInfoTag16(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.GetPolicyListsWithInfoTag16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyListsWithInfoTag16`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.GetPolicyListsWithInfoTag16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyListsWithInfoTag16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infoTag** | **string** | InfoTag | 

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


## PreviewPolicyList16

> map[string]interface{} PreviewPolicyList16(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.PreviewPolicyList16(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.PreviewPolicyList16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyList16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.PreviewPolicyList16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyList16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy list | 

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


## PreviewPolicyListById16

> map[string]interface{} PreviewPolicyListById16(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationFQDNListBuilderApi.PreviewPolicyListById16(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFQDNListBuilderApi.PreviewPolicyListById16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyListById16`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFQDNListBuilderApi.PreviewPolicyListById16`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyListById16Request struct via the builder pattern


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

