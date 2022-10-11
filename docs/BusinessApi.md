# \BusinessApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataserviceSystemDeviceVedgesGet**](BusinessApi.md#DataserviceSystemDeviceVedgesGet) | **Get** /dataservice/system/device/vedges | Get Device VEdges
[**DataserviceTemplatePolicyListSiteDelete**](BusinessApi.md#DataserviceTemplatePolicyListSiteDelete) | **Delete** /dataservice/template/policy/list/site/ | Delete VPN Template List
[**DataserviceTemplatePolicyListSiteGet**](BusinessApi.md#DataserviceTemplatePolicyListSiteGet) | **Get** /dataservice/template/policy/list/site | Get Prefix Template List
[**DataserviceTemplatePolicyListSiteIdDelete**](BusinessApi.md#DataserviceTemplatePolicyListSiteIdDelete) | **Delete** /dataservice/template/policy/list/site/{id} | Delete Site Template List
[**DataserviceTemplatePolicyListSitePost**](BusinessApi.md#DataserviceTemplatePolicyListSitePost) | **Post** /dataservice/template/policy/list/site | Create Site Template List
[**DataserviceTemplatePolicyListVpnPost**](BusinessApi.md#DataserviceTemplatePolicyListVpnPost) | **Post** /dataservice/template/policy/list/vpn | Create VPN Template List



## DataserviceSystemDeviceVedgesGet

> string DataserviceSystemDeviceVedgesGet(ctx).Execute()

Get Device VEdges

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
    resp, r, err := apiClient.BusinessApi.DataserviceSystemDeviceVedgesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceSystemDeviceVedgesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceSystemDeviceVedgesGet`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceSystemDeviceVedgesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceSystemDeviceVedgesGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataserviceTemplatePolicyListSiteDelete

> string DataserviceTemplatePolicyListSiteDelete(ctx).XXSRFTOKEN(xXSRFTOKEN).Execute()

Delete VPN Template List

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
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessApi.DataserviceTemplatePolicyListSiteDelete(context.Background()).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceTemplatePolicyListSiteDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceTemplatePolicyListSiteDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceTemplatePolicyListSiteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceTemplatePolicyListSiteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXSRFTOKEN** | **string** |  | 

### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataserviceTemplatePolicyListSiteGet

> string DataserviceTemplatePolicyListSiteGet(ctx).Execute()

Get Prefix Template List

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
    resp, r, err := apiClient.BusinessApi.DataserviceTemplatePolicyListSiteGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceTemplatePolicyListSiteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceTemplatePolicyListSiteGet`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceTemplatePolicyListSiteGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceTemplatePolicyListSiteGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataserviceTemplatePolicyListSiteIdDelete

> string DataserviceTemplatePolicyListSiteIdDelete(ctx, id).XXSRFTOKEN(xXSRFTOKEN).Execute()

Delete Site Template List

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
    id := "id_example" // string | 
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessApi.DataserviceTemplatePolicyListSiteIdDelete(context.Background(), id).XXSRFTOKEN(xXSRFTOKEN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceTemplatePolicyListSiteIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceTemplatePolicyListSiteIdDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceTemplatePolicyListSiteIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceTemplatePolicyListSiteIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xXSRFTOKEN** | **string** |  | 

### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataserviceTemplatePolicyListSitePost

> string DataserviceTemplatePolicyListSitePost(ctx).XXSRFTOKEN(xXSRFTOKEN).Body(body).Execute()

Create Site Template List

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
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessApi.DataserviceTemplatePolicyListSitePost(context.Background()).XXSRFTOKEN(xXSRFTOKEN).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceTemplatePolicyListSitePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceTemplatePolicyListSitePost`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceTemplatePolicyListSitePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceTemplatePolicyListSitePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXSRFTOKEN** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataserviceTemplatePolicyListVpnPost

> string DataserviceTemplatePolicyListVpnPost(ctx).XXSRFTOKEN(xXSRFTOKEN).Body(body).Execute()

Create VPN Template List

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
    xXSRFTOKEN := "{{X-XSRF-TOKEN}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BusinessApi.DataserviceTemplatePolicyListVpnPost(context.Background()).XXSRFTOKEN(xXSRFTOKEN).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessApi.DataserviceTemplatePolicyListVpnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceTemplatePolicyListVpnPost`: string
    fmt.Fprintf(os.Stdout, "Response from `BusinessApi.DataserviceTemplatePolicyListVpnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceTemplatePolicyListVpnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xXSRFTOKEN** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

