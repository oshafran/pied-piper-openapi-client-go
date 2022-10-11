# \ConfigurationPolicyExpandedCommunityListBuilderApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyList11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#CreatePolicyList11) | **Post** /template/policy/list/expandedcommunity | 
[**DeletePolicyList11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#DeletePolicyList11) | **Delete** /template/policy/list/expandedcommunity/{id} | 
[**DeletePolicyListsWithInfoTag11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#DeletePolicyListsWithInfoTag11) | **Delete** /template/policy/list/expandedcommunity | 
[**EditPolicyList11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#EditPolicyList11) | **Put** /template/policy/list/expandedcommunity/{id} | 
[**GetListsById11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#GetListsById11) | **Get** /template/policy/list/expandedcommunity/{id} | 
[**GetPolicyLists10**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#GetPolicyLists10) | **Get** /template/policy/list/expandedcommunity | 
[**GetPolicyListsWithInfoTag11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#GetPolicyListsWithInfoTag11) | **Get** /template/policy/list/expandedcommunity/filtered | 
[**PreviewPolicyList11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#PreviewPolicyList11) | **Post** /template/policy/list/expandedcommunity/preview | 
[**PreviewPolicyListById11**](ConfigurationPolicyExpandedCommunityListBuilderApi.md#PreviewPolicyListById11) | **Get** /template/policy/list/expandedcommunity/preview/{id} | 



## CreatePolicyList11

> map[string]interface{} CreatePolicyList11(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.CreatePolicyList11(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.CreatePolicyList11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyList11`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.CreatePolicyList11`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyList11Request struct via the builder pattern


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


## DeletePolicyList11

> DeletePolicyList11(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.DeletePolicyList11(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.DeletePolicyList11``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyList11Request struct via the builder pattern


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


## DeletePolicyListsWithInfoTag11

> []map[string]interface{} DeletePolicyListsWithInfoTag11(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.DeletePolicyListsWithInfoTag11(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.DeletePolicyListsWithInfoTag11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyListsWithInfoTag11`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.DeletePolicyListsWithInfoTag11`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyListsWithInfoTag11Request struct via the builder pattern


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


## EditPolicyList11

> map[string]interface{} EditPolicyList11(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.EditPolicyList11(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.EditPolicyList11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyList11`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.EditPolicyList11`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyList11Request struct via the builder pattern


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


## GetListsById11

> map[string]interface{} GetListsById11(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.GetListsById11(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.GetListsById11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListsById11`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.GetListsById11`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsById11Request struct via the builder pattern


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


## GetPolicyLists10

> []map[string]interface{} GetPolicyLists10(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.GetPolicyLists10(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.GetPolicyLists10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyLists10`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.GetPolicyLists10`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyLists10Request struct via the builder pattern


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


## GetPolicyListsWithInfoTag11

> []map[string]interface{} GetPolicyListsWithInfoTag11(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.GetPolicyListsWithInfoTag11(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.GetPolicyListsWithInfoTag11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyListsWithInfoTag11`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.GetPolicyListsWithInfoTag11`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyListsWithInfoTag11Request struct via the builder pattern


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


## PreviewPolicyList11

> map[string]interface{} PreviewPolicyList11(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.PreviewPolicyList11(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.PreviewPolicyList11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyList11`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.PreviewPolicyList11`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyList11Request struct via the builder pattern


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


## PreviewPolicyListById11

> map[string]interface{} PreviewPolicyListById11(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyExpandedCommunityListBuilderApi.PreviewPolicyListById11(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyExpandedCommunityListBuilderApi.PreviewPolicyListById11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyListById11`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyExpandedCommunityListBuilderApi.PreviewPolicyListById11`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyListById11Request struct via the builder pattern


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

