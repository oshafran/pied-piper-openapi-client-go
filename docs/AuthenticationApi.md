# \AuthenticationApi

All URIs are relative to *https://https*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExampleComportDataserviceClientTokenGet**](AuthenticationApi.md#ExampleComportDataserviceClientTokenGet) | **Get** //example.com:{port}/dataservice/client/token | Token
[**ExampleComportJSecurityCheckPost**](AuthenticationApi.md#ExampleComportJSecurityCheckPost) | **Post** //example.com:{port}/j_security_check | Authentication



## ExampleComportDataserviceClientTokenGet

> string ExampleComportDataserviceClientTokenGet(ctx, port).Execute()

Token

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
    port := "port_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.ExampleComportDataserviceClientTokenGet(context.Background(), port).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.ExampleComportDataserviceClientTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportDataserviceClientTokenGet`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.ExampleComportDataserviceClientTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportDataserviceClientTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExampleComportJSecurityCheckPost

> string ExampleComportJSecurityCheckPost(ctx, port).ContentType(contentType).JUsername(jUsername).JPassword(jPassword).Execute()

Authentication

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
    port := "port_example" // string | 
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    jUsername := "jUsername_example" // string |  (optional)
    jPassword := "jPassword_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.ExampleComportJSecurityCheckPost(context.Background(), port).ContentType(contentType).JUsername(jUsername).JPassword(jPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.ExampleComportJSecurityCheckPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExampleComportJSecurityCheckPost`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.ExampleComportJSecurityCheckPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExampleComportJSecurityCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **jUsername** | **string** |  | 
 **jPassword** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

