# \PartnerRegistrationApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeviceMapping**](PartnerRegistrationApi.md#DeleteDeviceMapping) | **Post** /partner/{partnerType}/unmap/{nmsId} | 
[**DeletePartner**](PartnerRegistrationApi.md#DeletePartner) | **Delete** /partner/{partnerType}/{nmsId} | 
[**GetDataChangeInfo**](PartnerRegistrationApi.md#GetDataChangeInfo) | **Get** /serverlongpoll/event/poll/{partnerId} | 
[**GetPartner**](PartnerRegistrationApi.md#GetPartner) | **Get** /partner/{partnerType}/{nmsId} | 
[**GetPartnerDevices**](PartnerRegistrationApi.md#GetPartnerDevices) | **Get** /partner/{partnerType}/map/{nmsId} | 
[**GetPartners**](PartnerRegistrationApi.md#GetPartners) | **Get** /partner | 
[**GetPartnersByPartnerType**](PartnerRegistrationApi.md#GetPartnersByPartnerType) | **Get** /partner/{partnerType} | 
[**GetVPNList**](PartnerRegistrationApi.md#GetVPNList) | **Get** /partner/vpn | 
[**MapDevices**](PartnerRegistrationApi.md#MapDevices) | **Post** /partner/{partnerType}/map/{nmsId} | 
[**RegisterPartner**](PartnerRegistrationApi.md#RegisterPartner) | **Post** /partner/{partnerType} | 
[**UnmapDevices**](PartnerRegistrationApi.md#UnmapDevices) | **Delete** /partner/{partnerType}/map/{nmsId} | 
[**UpdatePartner**](PartnerRegistrationApi.md#UpdatePartner) | **Put** /partner/{partnerType}/{nmsId} | 



## DeleteDeviceMapping

> map[string]interface{} DeleteDeviceMapping(ctx, partnerType, nmsId).Body(body).Execute()





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
    partnerType := "partnerType_example" // string | Partner type
    nmsId := "nmsId_example" // string | NMS Id
    body := map[string]interface{}{ ... } // map[string]interface{} | List of devices (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.DeleteDeviceMapping(context.Background(), partnerType, nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.DeleteDeviceMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeviceMapping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.DeleteDeviceMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | **string** | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | List of devices | 

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


## DeletePartner

> map[string]interface{} DeletePartner(ctx, partnerType, nmsId).Execute()





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
    partnerType := map[string][]openapiclient.PartnerType{ ... } // PartnerType | Partner type
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NMS Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.DeletePartner(context.Background(), partnerType, nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.DeletePartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePartner`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.DeletePartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | [**PartnerType**](.md) | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartnerRequest struct via the builder pattern


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


## GetDataChangeInfo

> GetDataChangeInfo(ctx, partnerId).EventId(eventId).EventNames(eventNames).WaitTime(waitTime).Execute()





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
    eventId := "eventId_example" // string | Continuation token of ongoing event-polling session (optional)
    eventNames := []openapiclient.EventName{*openapiclient.NewEventName()} // []EventName | Names of type of events to filter on (optional)
    waitTime := int64(789) // int64 | Maximum polling wait time in seconds (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.GetDataChangeInfo(context.Background(), partnerId).EventId(eventId).EventNames(eventNames).WaitTime(waitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.GetDataChangeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | Partner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataChangeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventId** | **string** | Continuation token of ongoing event-polling session | 
 **eventNames** | [**[]EventName**](EventName.md) | Names of type of events to filter on | 
 **waitTime** | **int64** | Maximum polling wait time in seconds | [default to 0]

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


## GetPartner

> map[string]interface{} GetPartner(ctx, partnerType, nmsId).Execute()





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
    partnerType := "partnerType_example" // string | Partner type
    nmsId := "nmsId_example" // string | NMS Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.GetPartner(context.Background(), partnerType, nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.GetPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartner`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.GetPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | **string** | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerRequest struct via the builder pattern


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


## GetPartnerDevices

> map[string]interface{} GetPartnerDevices(ctx, partnerType, nmsId).Execute()





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
    partnerType := map[string][]openapiclient.PartnerType{ ... } // PartnerType | Partner type
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NMS Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.GetPartnerDevices(context.Background(), partnerType, nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.GetPartnerDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartnerDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.GetPartnerDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | [**PartnerType**](.md) | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerDevicesRequest struct via the builder pattern


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


## GetPartners

> []map[string]interface{} GetPartners(ctx).Execute()





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
    resp, r, err := apiClient.PartnerRegistrationApi.GetPartners(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.GetPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartners`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.GetPartners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnersRequest struct via the builder pattern


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


## GetPartnersByPartnerType

> []map[string]interface{} GetPartnersByPartnerType(ctx, partnerType).Execute()





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
    partnerType := "partnerType_example" // string | Partner type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.GetPartnersByPartnerType(context.Background(), partnerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.GetPartnersByPartnerType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartnersByPartnerType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.GetPartnersByPartnerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | **string** | Partner type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnersByPartnerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetVPNList

> []map[string]interface{} GetVPNList(ctx).Execute()





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
    resp, r, err := apiClient.PartnerRegistrationApi.GetVPNList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.GetVPNList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVPNList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.GetVPNList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVPNListRequest struct via the builder pattern


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


## MapDevices

> map[string]interface{} MapDevices(ctx, partnerType, nmsId).Body(body).Execute()





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
    partnerType := "partnerType_example" // string | Partner type
    nmsId := "nmsId_example" // string | NMS Id
    body := map[string]interface{}{ ... } // map[string]interface{} | List of devices (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.MapDevices(context.Background(), partnerType, nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.MapDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MapDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.MapDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | **string** | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMapDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | List of devices | 

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


## RegisterPartner

> map[string]interface{} RegisterPartner(ctx, partnerType).Body(body).Execute()





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
    partnerType := "partnerType_example" // string | Partner type
    body := map[string]interface{}{ ... } // map[string]interface{} | Partner (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.RegisterPartner(context.Background(), partnerType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.RegisterPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterPartner`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.RegisterPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | **string** | Partner type | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Partner | 

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


## UnmapDevices

> map[string]interface{} UnmapDevices(ctx, partnerType, nmsId).Execute()





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
    partnerType := map[string][]openapiclient.PartnerType{ ... } // PartnerType | Partner type
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NMS Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.UnmapDevices(context.Background(), partnerType, nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.UnmapDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmapDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartnerRegistrationApi.UnmapDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | [**PartnerType**](.md) | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmapDevicesRequest struct via the builder pattern


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


## UpdatePartner

> UpdatePartner(ctx, partnerType, nmsId).Body(body).Execute()





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
    partnerType := "partnerType_example" // string | Partner type
    nmsId := "nmsId_example" // string | NMS Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Partner (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartnerRegistrationApi.UpdatePartner(context.Background(), partnerType, nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerRegistrationApi.UpdatePartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerType** | **string** | Partner type | 
**nmsId** | **string** | NMS Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | Partner | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

