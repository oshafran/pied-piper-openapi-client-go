# \MDPOnboardingApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMDPOnboardingStatus**](MDPOnboardingApi.md#GetMDPOnboardingStatus) | **Get** /mdp/onboard/status | 
[**OnboardMDP**](MDPOnboardingApi.md#OnboardMDP) | **Post** /mdp/onboard | 
[**UpdateOnboardingPayload**](MDPOnboardingApi.md#UpdateOnboardingPayload) | **Put** /mdp/onboard/{nmsId} | 



## GetMDPOnboardingStatus

> []map[string]interface{} GetMDPOnboardingStatus(ctx).Execute()





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
    resp, r, err := apiClient.MDPOnboardingApi.GetMDPOnboardingStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPOnboardingApi.GetMDPOnboardingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMDPOnboardingStatus`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPOnboardingApi.GetMDPOnboardingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMDPOnboardingStatusRequest struct via the builder pattern


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


## OnboardMDP

> map[string]interface{} OnboardMDP(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Onboard (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPOnboardingApi.OnboardMDP(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPOnboardingApi.OnboardMDP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnboardMDP`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPOnboardingApi.OnboardMDP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnboardMDPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Onboard | 

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


## UpdateOnboardingPayload

> map[string]interface{} UpdateOnboardingPayload(ctx, nmsId).Body(body).Execute()





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
    nmsId := "nmsId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | Onboard (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPOnboardingApi.UpdateOnboardingPayload(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPOnboardingApi.UpdateOnboardingPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOnboardingPayload`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPOnboardingApi.UpdateOnboardingPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOnboardingPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Onboard | 

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

