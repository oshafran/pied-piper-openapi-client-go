# \ConfigurationPolicyDataPrefixListBuilderApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyList10**](ConfigurationPolicyDataPrefixListBuilderApi.md#CreatePolicyList10) | **Post** /template/policy/list/dataprefix | 
[**DeletePolicyList10**](ConfigurationPolicyDataPrefixListBuilderApi.md#DeletePolicyList10) | **Delete** /template/policy/list/dataprefix/{id} | 
[**DeletePolicyListsWithInfoTag10**](ConfigurationPolicyDataPrefixListBuilderApi.md#DeletePolicyListsWithInfoTag10) | **Delete** /template/policy/list/dataprefix | 
[**EditPolicyList10**](ConfigurationPolicyDataPrefixListBuilderApi.md#EditPolicyList10) | **Put** /template/policy/list/dataprefix/{id} | 
[**GetListsById10**](ConfigurationPolicyDataPrefixListBuilderApi.md#GetListsById10) | **Get** /template/policy/list/dataprefix/{id} | 
[**GetPolicyLists9**](ConfigurationPolicyDataPrefixListBuilderApi.md#GetPolicyLists9) | **Get** /template/policy/list/dataprefix | 
[**GetPolicyListsWithInfoTag10**](ConfigurationPolicyDataPrefixListBuilderApi.md#GetPolicyListsWithInfoTag10) | **Get** /template/policy/list/dataprefix/filtered | 
[**PreviewPolicyList10**](ConfigurationPolicyDataPrefixListBuilderApi.md#PreviewPolicyList10) | **Post** /template/policy/list/dataprefix/preview | 
[**PreviewPolicyListById10**](ConfigurationPolicyDataPrefixListBuilderApi.md#PreviewPolicyListById10) | **Get** /template/policy/list/dataprefix/preview/{id} | 



## CreatePolicyList10

> map[string]interface{} CreatePolicyList10(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.CreatePolicyList10(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.CreatePolicyList10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyList10`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.CreatePolicyList10`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyList10Request struct via the builder pattern


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


## DeletePolicyList10

> DeletePolicyList10(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.DeletePolicyList10(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.DeletePolicyList10``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyList10Request struct via the builder pattern


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


## DeletePolicyListsWithInfoTag10

> []map[string]interface{} DeletePolicyListsWithInfoTag10(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.DeletePolicyListsWithInfoTag10(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.DeletePolicyListsWithInfoTag10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyListsWithInfoTag10`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.DeletePolicyListsWithInfoTag10`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyListsWithInfoTag10Request struct via the builder pattern


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


## EditPolicyList10

> map[string]interface{} EditPolicyList10(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.EditPolicyList10(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.EditPolicyList10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyList10`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.EditPolicyList10`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyList10Request struct via the builder pattern


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


## GetListsById10

> map[string]interface{} GetListsById10(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.GetListsById10(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.GetListsById10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListsById10`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.GetListsById10`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsById10Request struct via the builder pattern


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


## GetPolicyLists9

> []map[string]interface{} GetPolicyLists9(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.GetPolicyLists9(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.GetPolicyLists9``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyLists9`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.GetPolicyLists9`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyLists9Request struct via the builder pattern


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


## GetPolicyListsWithInfoTag10

> []map[string]interface{} GetPolicyListsWithInfoTag10(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.GetPolicyListsWithInfoTag10(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.GetPolicyListsWithInfoTag10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyListsWithInfoTag10`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.GetPolicyListsWithInfoTag10`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyListsWithInfoTag10Request struct via the builder pattern


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


## PreviewPolicyList10

> map[string]interface{} PreviewPolicyList10(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.PreviewPolicyList10(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.PreviewPolicyList10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyList10`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.PreviewPolicyList10`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyList10Request struct via the builder pattern


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


## PreviewPolicyListById10

> map[string]interface{} PreviewPolicyListById10(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyDataPrefixListBuilderApi.PreviewPolicyListById10(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyDataPrefixListBuilderApi.PreviewPolicyListById10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyListById10`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyDataPrefixListBuilderApi.PreviewPolicyListById10`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyListById10Request struct via the builder pattern


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

