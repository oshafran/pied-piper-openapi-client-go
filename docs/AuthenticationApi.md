# \AuthenticationApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataserviceClientTokenGet**](AuthenticationApi.md#DataserviceClientTokenGet) | **Get** /dataservice/client/token | Token
[**JSecurityCheckPost**](AuthenticationApi.md#JSecurityCheckPost) | **Post** /j_security_check | Authentication



## DataserviceClientTokenGet

> string DataserviceClientTokenGet(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.DataserviceClientTokenGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.DataserviceClientTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataserviceClientTokenGet`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.DataserviceClientTokenGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDataserviceClientTokenGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Cookie](../README.md#Cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JSecurityCheckPost

> string JSecurityCheckPost(ctx).ContentType(contentType).JUsername(jUsername).JPassword(jPassword).Execute()

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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    jUsername := "jUsername_example" // string |  (optional)
    jPassword := "jPassword_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.JSecurityCheckPost(context.Background()).ContentType(contentType).JUsername(jUsername).JPassword(jPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.JSecurityCheckPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JSecurityCheckPost`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.JSecurityCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJSecurityCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **jUsername** | **string** |  | 
 **jPassword** | **string** |  | 

### Return type

**string**

### Authorization

[Cookie](../README.md#Cookie)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

