# \ConfigurationPolicyFaxProtocolListBuilderApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyList13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#CreatePolicyList13) | **Post** /template/policy/list/faxprotocol | 
[**DeletePolicyList13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#DeletePolicyList13) | **Delete** /template/policy/list/faxprotocol/{id} | 
[**DeletePolicyListsWithInfoTag13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#DeletePolicyListsWithInfoTag13) | **Delete** /template/policy/list/faxprotocol | 
[**EditPolicyList13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#EditPolicyList13) | **Put** /template/policy/list/faxprotocol/{id} | 
[**GetListsById13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#GetListsById13) | **Get** /template/policy/list/faxprotocol/{id} | 
[**GetPolicyLists12**](ConfigurationPolicyFaxProtocolListBuilderApi.md#GetPolicyLists12) | **Get** /template/policy/list/faxprotocol | 
[**GetPolicyListsWithInfoTag13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#GetPolicyListsWithInfoTag13) | **Get** /template/policy/list/faxprotocol/filtered | 
[**PreviewPolicyList13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#PreviewPolicyList13) | **Post** /template/policy/list/faxprotocol/preview | 
[**PreviewPolicyListById13**](ConfigurationPolicyFaxProtocolListBuilderApi.md#PreviewPolicyListById13) | **Get** /template/policy/list/faxprotocol/preview/{id} | 



## CreatePolicyList13

> map[string]interface{} CreatePolicyList13(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.CreatePolicyList13(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.CreatePolicyList13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyList13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.CreatePolicyList13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyList13Request struct via the builder pattern


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


## DeletePolicyList13

> DeletePolicyList13(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.DeletePolicyList13(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.DeletePolicyList13``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyList13Request struct via the builder pattern


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


## DeletePolicyListsWithInfoTag13

> []map[string]interface{} DeletePolicyListsWithInfoTag13(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.DeletePolicyListsWithInfoTag13(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.DeletePolicyListsWithInfoTag13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyListsWithInfoTag13`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.DeletePolicyListsWithInfoTag13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyListsWithInfoTag13Request struct via the builder pattern


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


## EditPolicyList13

> map[string]interface{} EditPolicyList13(ctx, id).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.EditPolicyList13(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.EditPolicyList13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditPolicyList13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.EditPolicyList13`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPolicyList13Request struct via the builder pattern


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


## GetListsById13

> map[string]interface{} GetListsById13(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.GetListsById13(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.GetListsById13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListsById13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.GetListsById13`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsById13Request struct via the builder pattern


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


## GetPolicyLists12

> []map[string]interface{} GetPolicyLists12(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.GetPolicyLists12(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.GetPolicyLists12``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyLists12`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.GetPolicyLists12`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyLists12Request struct via the builder pattern


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


## GetPolicyListsWithInfoTag13

> []map[string]interface{} GetPolicyListsWithInfoTag13(ctx).InfoTag(infoTag).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.GetPolicyListsWithInfoTag13(context.Background()).InfoTag(infoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.GetPolicyListsWithInfoTag13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyListsWithInfoTag13`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.GetPolicyListsWithInfoTag13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyListsWithInfoTag13Request struct via the builder pattern


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


## PreviewPolicyList13

> map[string]interface{} PreviewPolicyList13(ctx).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.PreviewPolicyList13(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.PreviewPolicyList13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyList13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.PreviewPolicyList13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyList13Request struct via the builder pattern


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


## PreviewPolicyListById13

> map[string]interface{} PreviewPolicyListById13(ctx, id).Execute()





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
    resp, r, err := apiClient.ConfigurationPolicyFaxProtocolListBuilderApi.PreviewPolicyListById13(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPolicyFaxProtocolListBuilderApi.PreviewPolicyListById13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewPolicyListById13`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationPolicyFaxProtocolListBuilderApi.PreviewPolicyListById13`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewPolicyListById13Request struct via the builder pattern


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

