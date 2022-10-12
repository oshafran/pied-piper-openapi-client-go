# \ResourcePoolApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResources**](ResourcePoolApi.md#CreateResources) | **Put** /resourcepool/resource/vpn | 
[**DeleteResources**](ResourcePoolApi.md#DeleteResources) | **Delete** /resourcepool/resource/vpn | 
[**GetResources**](ResourcePoolApi.md#GetResources) | **Get** /resourcepool/resource/vpn | 



## CreateResources

> map[string]interface{} CreateResources(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | create resources from resource pool (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcePoolApi.CreateResources(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolApi.CreateResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResources`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolApi.CreateResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | create resources from resource pool | 

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


## DeleteResources

> DeleteResources(ctx).TenantId(tenantId).TenantVpn(tenantVpn).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id
    tenantVpn := int64(789) // int64 | Tenant Vpn Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcePoolApi.DeleteResources(context.Background()).TenantId(tenantId).TenantVpn(tenantVpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolApi.DeleteResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant Id | 
 **tenantVpn** | **int64** | Tenant Vpn Number | 

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


## GetResources

> map[string]interface{} GetResources(ctx).TenantId(tenantId).TenantVpn(tenantVpn).Execute()





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
    tenantId := "tenantId_example" // string | Tenant Id
    tenantVpn := int64(789) // int64 | Tenant Vpn Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcePoolApi.GetResources(context.Background()).TenantId(tenantId).TenantVpn(tenantVpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcePoolApi.GetResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResources`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ResourcePoolApi.GetResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Tenant Id | 
 **tenantVpn** | **int64** | Tenant Vpn Number | 

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

