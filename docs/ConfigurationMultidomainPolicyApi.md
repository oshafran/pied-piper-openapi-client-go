# \ConfigurationMultidomainPolicyApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInternalPolicy**](ConfigurationMultidomainPolicyApi.md#AddInternalPolicy) | **Put** /mdp/policies/mdpconfig | 
[**AttachDevices**](ConfigurationMultidomainPolicyApi.md#AttachDevices) | **Post** /mdp/attachDevices/{nmsId} | 
[**DetachDevices**](ConfigurationMultidomainPolicyApi.md#DetachDevices) | **Post** /mdp/detachDevices/{nmsId} | 
[**DisconnectFromMdp**](ConfigurationMultidomainPolicyApi.md#DisconnectFromMdp) | **Get** /mdp/disconnect/{nmsId} | 
[**EditAttachedDevices**](ConfigurationMultidomainPolicyApi.md#EditAttachedDevices) | **Put** /mdp/attachDevices/{nmsId} | 
[**GetMDPOnboardingStatus**](ConfigurationMultidomainPolicyApi.md#GetMDPOnboardingStatus) | **Get** /mdp/onboard/status | 
[**Offboard**](ConfigurationMultidomainPolicyApi.md#Offboard) | **Delete** /mdp/onboard/{nmsId} | 
[**OnboardMDP**](ConfigurationMultidomainPolicyApi.md#OnboardMDP) | **Post** /mdp/onboard | 
[**RetrieveMDPAttachedDevices**](ConfigurationMultidomainPolicyApi.md#RetrieveMDPAttachedDevices) | **Get** /mdp/attachDevices/{nmsId} | 
[**RetrieveMDPConfigObject**](ConfigurationMultidomainPolicyApi.md#RetrieveMDPConfigObject) | **Get** /mdp/policies/mdpconfig/{deviceId} | 
[**RetrieveMDPPolicies**](ConfigurationMultidomainPolicyApi.md#RetrieveMDPPolicies) | **Get** /mdp/policies/{nmsId} | 
[**RetrieveMDPSupportedDevices**](ConfigurationMultidomainPolicyApi.md#RetrieveMDPSupportedDevices) | **Get** /mdp/devices/{nmsId} | 
[**UpdateOnboardingPayload**](ConfigurationMultidomainPolicyApi.md#UpdateOnboardingPayload) | **Put** /mdp/onboard/{nmsId} | 
[**UpdatePolicyStatus**](ConfigurationMultidomainPolicyApi.md#UpdatePolicyStatus) | **Put** /mdp/policies/{nmsId} | 



## AddInternalPolicy

> map[string]interface{} AddInternalPolicy(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | addInternalPolicy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.AddInternalPolicy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.AddInternalPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInternalPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.AddInternalPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInternalPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | addInternalPolicy | 

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


## AttachDevices

> map[string]interface{} AttachDevices(ctx, nmsId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | deviceList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.AttachDevices(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.AttachDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.AttachDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | deviceList | 

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


## DetachDevices

> map[string]interface{} DetachDevices(ctx, nmsId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | deviceList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.DetachDevices(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.DetachDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.DetachDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | deviceList | 

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


## DisconnectFromMdp

> []map[string]interface{} DisconnectFromMdp(ctx, nmsId).Execute()





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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.DisconnectFromMdp(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.DisconnectFromMdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisconnectFromMdp`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.DisconnectFromMdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectFromMdpRequest struct via the builder pattern


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


## EditAttachedDevices

> map[string]interface{} EditAttachedDevices(ctx, nmsId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | deviceList (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.EditAttachedDevices(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.EditAttachedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAttachedDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.EditAttachedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAttachedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | deviceList | 

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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.GetMDPOnboardingStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.GetMDPOnboardingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMDPOnboardingStatus`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.GetMDPOnboardingStatus`: %v\n", resp)
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


## Offboard

> Offboard(ctx, nmsId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.Offboard(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.Offboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOffboardRequest struct via the builder pattern


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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.OnboardMDP(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.OnboardMDP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnboardMDP`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.OnboardMDP`: %v\n", resp)
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


## RetrieveMDPAttachedDevices

> []map[string]interface{} RetrieveMDPAttachedDevices(ctx, nmsId).Execute()





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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPAttachedDevices(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.RetrieveMDPAttachedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPAttachedDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.RetrieveMDPAttachedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPAttachedDevicesRequest struct via the builder pattern


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


## RetrieveMDPConfigObject

> []map[string]interface{} RetrieveMDPConfigObject(ctx, deviceId).Execute()





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
    deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPConfigObject(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.RetrieveMDPConfigObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPConfigObject`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.RetrieveMDPConfigObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPConfigObjectRequest struct via the builder pattern


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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPPolicies(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.RetrieveMDPPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPPolicies`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.RetrieveMDPPolicies`: %v\n", resp)
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


## RetrieveMDPSupportedDevices

> []map[string]interface{} RetrieveMDPSupportedDevices(ctx, nmsId).Execute()





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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPSupportedDevices(context.Background(), nmsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.RetrieveMDPSupportedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMDPSupportedDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.RetrieveMDPSupportedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nmsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMDPSupportedDevicesRequest struct via the builder pattern


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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.UpdateOnboardingPayload(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.UpdateOnboardingPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOnboardingPayload`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.UpdateOnboardingPayload`: %v\n", resp)
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
    resp, r, err := apiClient.ConfigurationMultidomainPolicyApi.UpdatePolicyStatus(context.Background(), nmsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultidomainPolicyApi.UpdatePolicyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultidomainPolicyApi.UpdatePolicyStatus`: %v\n", resp)
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

