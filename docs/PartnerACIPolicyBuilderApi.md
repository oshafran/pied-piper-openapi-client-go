# \PartnerACIPolicyBuilderApi

All URIs are relative to *http://192.168.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDscpMappings**](PartnerACIPolicyBuilderApi.md#CreateDscpMappings) | **Post** /partner/aci/policy/dscpmapping/{partnerId} | 
[**DeleteDscpMappings**](PartnerACIPolicyBuilderApi.md#DeleteDscpMappings) | **Delete** /partner/aci/policy/dscpmapping/{partnerId} | 
[**GetACIDefinitions**](PartnerACIPolicyBuilderApi.md#GetACIDefinitions) | **Get** /partner/aci/policy | 
[**GetDataPrefixMappings**](PartnerACIPolicyBuilderApi.md#GetDataPrefixMappings) | **Get** /partner/aci/policy/prefixmapping/{partnerId} | 
[**GetDataPrefixSequences**](PartnerACIPolicyBuilderApi.md#GetDataPrefixSequences) | **Get** /partner/aci/policy/sequences | 
[**GetDscpMappings**](PartnerACIPolicyBuilderApi.md#GetDscpMappings) | **Get** /partner/aci/policy/dscpmapping/{partnerId} | 
[**GetEvents**](PartnerACIPolicyBuilderApi.md#GetEvents) | **Get** /partner/aci/policy/events/{partnerId} | 
[**SetDataPrefixMappings**](PartnerACIPolicyBuilderApi.md#SetDataPrefixMappings) | **Post** /partner/aci/policy/prefixmapping/{partnerId} | 



## CreateDscpMappings

> map[string]interface{} CreateDscpMappings(ctx, partnerId).Body(body).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id
    body := map[string]interface{}{ ... } // map[string]interface{} | ACI definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.CreateDscpMappings(context.Background(), partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.CreateDscpMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDscpMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.CreateDscpMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDscpMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | ACI definition | 

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


## DeleteDscpMappings

> map[string]interface{} DeleteDscpMappings(ctx, partnerId).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.DeleteDscpMappings(context.Background(), partnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.DeleteDscpMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDscpMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.DeleteDscpMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDscpMappingsRequest struct via the builder pattern


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


## GetACIDefinitions

> map[string]interface{} GetACIDefinitions(ctx).Execute()





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
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.GetACIDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.GetACIDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACIDefinitions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.GetACIDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetACIDefinitionsRequest struct via the builder pattern


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


## GetDataPrefixMappings

> map[string]interface{} GetDataPrefixMappings(ctx, partnerId).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.GetDataPrefixMappings(context.Background(), partnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.GetDataPrefixMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataPrefixMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.GetDataPrefixMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPrefixMappingsRequest struct via the builder pattern


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


## GetDataPrefixSequences

> []map[string]interface{} GetDataPrefixSequences(ctx).Execute()





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
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.GetDataPrefixSequences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.GetDataPrefixSequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataPrefixSequences`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.GetDataPrefixSequences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPrefixSequencesRequest struct via the builder pattern


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


## GetDscpMappings

> map[string]interface{} GetDscpMappings(ctx, partnerId).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.GetDscpMappings(context.Background(), partnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.GetDscpMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDscpMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.GetDscpMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDscpMappingsRequest struct via the builder pattern


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


## GetEvents

> []map[string]interface{} GetEvents(ctx, partnerId).Starttime(starttime).Endtime(endtime).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id
    starttime := int64(789) // int64 | Start time (optional)
    endtime := int64(789) // int64 | End time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.GetEvents(context.Background(), partnerId).Starttime(starttime).Endtime(endtime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **starttime** | **int64** | Start time | 
 **endtime** | **int64** | End time | 

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


## SetDataPrefixMappings

> map[string]interface{} SetDataPrefixMappings(ctx, partnerId).Body(body).Execute()





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
    partnerId := "partnerId_example" // string | Partner Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Prefix definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerACIPolicyBuilderApi.SetDataPrefixMappings(context.Background(), partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerACIPolicyBuilderApi.SetDataPrefixMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDataPrefixMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerACIPolicyBuilderApi.SetDataPrefixMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDataPrefixMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Prefix definition | 

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

