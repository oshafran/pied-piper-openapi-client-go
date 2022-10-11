# \ConfigurationSecurityTemplatePolicyApi

All URIs are relative to *http://1.1.1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityTemplate**](ConfigurationSecurityTemplatePolicyApi.md#CreateSecurityTemplate) | **Post** /template/policy/security | 
[**DeleteSecurityTemplate**](ConfigurationSecurityTemplatePolicyApi.md#DeleteSecurityTemplate) | **Delete** /template/policy/security/{policyId} | 
[**EditSecurityTemplate**](ConfigurationSecurityTemplatePolicyApi.md#EditSecurityTemplate) | **Put** /template/policy/security/{policyId} | 
[**EditTemplateWithLenientLock**](ConfigurationSecurityTemplatePolicyApi.md#EditTemplateWithLenientLock) | **Put** /template/policy/security/staging/{policyId} | 
[**GenerateSecurityPolicySummary**](ConfigurationSecurityTemplatePolicyApi.md#GenerateSecurityPolicySummary) | **Get** /template/policy/security/summary | 
[**GenerateSecurityTemplateList**](ConfigurationSecurityTemplatePolicyApi.md#GenerateSecurityTemplateList) | **Get** /template/policy/security | 
[**GetDeviceListById**](ConfigurationSecurityTemplatePolicyApi.md#GetDeviceListById) | **Get** /template/policy/security/devices/{policyId} | 
[**GetSecurityPolicyDeviceList**](ConfigurationSecurityTemplatePolicyApi.md#GetSecurityPolicyDeviceList) | **Get** /template/policy/security/devices | 
[**GetSecurityTemplate**](ConfigurationSecurityTemplatePolicyApi.md#GetSecurityTemplate) | **Get** /template/policy/security/definition/{policyId} | 
[**GetSecurityTemplatesForDevice**](ConfigurationSecurityTemplatePolicyApi.md#GetSecurityTemplatesForDevice) | **Get** /template/policy/security/{deviceModel} | 



## CreateSecurityTemplate

> CreateSecurityTemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.CreateSecurityTemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.CreateSecurityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Policy template | 

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


## DeleteSecurityTemplate

> DeleteSecurityTemplate(ctx, policyId).Execute()





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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.DeleteSecurityTemplate(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.DeleteSecurityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityTemplateRequest struct via the builder pattern


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


## EditSecurityTemplate

> map[string]interface{} EditSecurityTemplate(ctx, policyId).Body(body).Execute()





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
    policyId := "policyId_example" // string | Policy Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.EditSecurityTemplate(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.EditSecurityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSecurityTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.EditSecurityTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSecurityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Policy template | 

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


## EditTemplateWithLenientLock

> map[string]interface{} EditTemplateWithLenientLock(ctx, policyId).Body(body).Execute()





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
    policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Policy Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Policy template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.EditTemplateWithLenientLock(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.EditTemplateWithLenientLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTemplateWithLenientLock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.EditTemplateWithLenientLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTemplateWithLenientLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Policy template | 

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


## GenerateSecurityPolicySummary

> map[string]interface{} GenerateSecurityPolicySummary(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.GenerateSecurityPolicySummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.GenerateSecurityPolicySummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSecurityPolicySummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.GenerateSecurityPolicySummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSecurityPolicySummaryRequest struct via the builder pattern


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


## GenerateSecurityTemplateList

> []map[string]interface{} GenerateSecurityTemplateList(ctx).Mode(mode).Execute()





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
    mode := "mode_example" // string | Mode (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.GenerateSecurityTemplateList(context.Background()).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.GenerateSecurityTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSecurityTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.GenerateSecurityTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSecurityTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | Mode | 

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


## GetDeviceListById

> []map[string]interface{} GetDeviceListById(ctx, policyId).Execute()





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
    policyId := "policyId_example" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.GetDeviceListById(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.GetDeviceListById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceListById`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.GetDeviceListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceListByIdRequest struct via the builder pattern


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


## GetSecurityPolicyDeviceList

> []map[string]interface{} GetSecurityPolicyDeviceList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.GetSecurityPolicyDeviceList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.GetSecurityPolicyDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityPolicyDeviceList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.GetSecurityPolicyDeviceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPolicyDeviceListRequest struct via the builder pattern


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


## GetSecurityTemplate

> map[string]interface{} GetSecurityTemplate(ctx, policyId).Execute()





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
    policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Policy Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.GetSecurityTemplate(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.GetSecurityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.GetSecurityTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityTemplateRequest struct via the builder pattern


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


## GetSecurityTemplatesForDevice

> map[string]interface{} GetSecurityTemplatesForDevice(ctx, deviceModel).Execute()





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
    deviceModel := map[string][]openapiclient.DeviceModel{ ... } // DeviceModel | Device model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSecurityTemplatePolicyApi.GetSecurityTemplatesForDevice(context.Background(), deviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSecurityTemplatePolicyApi.GetSecurityTemplatesForDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityTemplatesForDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSecurityTemplatePolicyApi.GetSecurityTemplatesForDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceModel** | [**DeviceModel**](.md) | Device model | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityTemplatesForDeviceRequest struct via the builder pattern


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

