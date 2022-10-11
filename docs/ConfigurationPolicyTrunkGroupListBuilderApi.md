# \ConfigurationPolicyTrunkGroupListBuilderApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyList35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#CreatePolicyList35) | **Post** /template/policy/list/trunkgroup | 
[**DeletePolicyList35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#DeletePolicyList35) | **Delete** /template/policy/list/trunkgroup/{id} | 
[**DeletePolicyListsWithInfoTag35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#DeletePolicyListsWithInfoTag35) | **Delete** /template/policy/list/trunkgroup | 
[**EditPolicyList35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#EditPolicyList35) | **Put** /template/policy/list/trunkgroup/{id} | 
[**GetListsById35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#GetListsById35) | **Get** /template/policy/list/trunkgroup/{id} | 
[**GetPolicyLists32**](ConfigurationPolicyTrunkGroupListBuilderApi.md#GetPolicyLists32) | **Get** /template/policy/list/trunkgroup | 
[**GetPolicyListsWithInfoTag35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#GetPolicyListsWithInfoTag35) | **Get** /template/policy/list/trunkgroup/filtered | 
[**PreviewPolicyList35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#PreviewPolicyList35) | **Post** /template/policy/list/trunkgroup/preview | 
[**PreviewPolicyListById35**](ConfigurationPolicyTrunkGroupListBuilderApi.md#PreviewPolicyListById35) | **Get** /template/policy/list/trunkgroup/preview/{id} | 



## CreatePolicyList35

> map[string]interface{} CreatePolicyList35(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.CreatePolicyList35(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.CreatePolicyList35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyList35`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.CreatePolicyList35`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyList35Request struct via the builder pattern


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


## DeletePolicyList35

> DeletePolicyList35(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.DeletePolicyList35(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.DeletePolicyList35``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyList35Request struct via the builder pattern


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


## DeletePolicyListsWithInfoTag35

> []map[string]interface{} DeletePolicyListsWithInfoTag35(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.DeletePolicyListsWithInfoTag35(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.DeletePolicyListsWithInfoTag35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyListsWithInfoTag35`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.DeletePolicyListsWithInfoTag35`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyListsWithInfoTag35Request struct via the builder pattern


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


## EditPolicyList35

> map[string]interface{} EditPolicyList35(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.EditPolicyList35(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.EditPolicyList35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyList35`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.EditPolicyList35`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyList35Request struct via the builder pattern


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


## GetListsById35

> map[string]interface{} GetListsById35(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.GetListsById35(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.GetListsById35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListsById35`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.GetListsById35`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsById35Request struct via the builder pattern


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


## GetPolicyLists32

> []map[string]interface{} GetPolicyLists32(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.GetPolicyLists32(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.GetPolicyLists32``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyLists32`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.GetPolicyLists32`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyLists32Request struct via the builder pattern


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


## GetPolicyListsWithInfoTag35

> []map[string]interface{} GetPolicyListsWithInfoTag35(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.GetPolicyListsWithInfoTag35(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.GetPolicyListsWithInfoTag35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyListsWithInfoTag35`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.GetPolicyListsWithInfoTag35`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyListsWithInfoTag35Request struct via the builder pattern


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


## PreviewPolicyList35

> map[string]interface{} PreviewPolicyList35(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.PreviewPolicyList35(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.PreviewPolicyList35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyList35`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.PreviewPolicyList35`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyList35Request struct via the builder pattern


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


## PreviewPolicyListById35

> map[string]interface{} PreviewPolicyListById35(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyTrunkGroupListBuilderApi.PreviewPolicyListById35(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyTrunkGroupListBuilderApi.PreviewPolicyListById35``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyListById35`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyTrunkGroupListBuilderApi.PreviewPolicyListById35`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyListById35Request struct via the builder pattern


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

