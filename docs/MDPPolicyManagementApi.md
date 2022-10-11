# \MDPPolicyManagementApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveMDPPolicies**](MDPPolicyManagementApi.md#RetrieveMDPPolicies) | **Get** /mdp/policies/{nmsId} | 
[**UpdatePolicyStatus**](MDPPolicyManagementApi.md#UpdatePolicyStatus) | **Put** /mdp/policies/{nmsId} | 



## RetrieveMDPPolicies

> []map[string]interface{} RetrieveMDPPolicies(ctx, nmsId).Execute()





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
    nmsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPPolicyManagementApi.RetrieveMDPPolicies(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPPolicyManagementApi.RetrieveMDPPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPPolicies`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPPolicyManagementApi.RetrieveMDPPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPPoliciesRequest struct via the builder pattern


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


## UpdatePolicyStatus

> map[string]interface{} UpdatePolicyStatus(ctx, nmsId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | policyList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MDPPolicyManagementApi.UpdatePolicyStatus(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MDPPolicyManagementApi.UpdatePolicyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MDPPolicyManagementApi.UpdatePolicyStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | policyList | 

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

