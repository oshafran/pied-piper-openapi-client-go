# \ConfigurationVSmartTemplatePolicyApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePolicy**](ConfigurationVSmartTemplatePolicyApi.md#ActivatePolicy) | **Post** /template/policy/vsmart/activate/{policyId} | 
[**ActivatePolicyForCloudServices**](ConfigurationVSmartTemplatePolicyApi.md#ActivatePolicyForCloudServices) | **Post** /template/policy/vsmart/activate/central/{policyId} | 
[**CheckVSmartConnectivityStatus**](ConfigurationVSmartTemplatePolicyApi.md#CheckVSmartConnectivityStatus) | **Get** /template/policy/vsmart/connectivity/status | 
[**CreateVSmartTemplate**](ConfigurationVSmartTemplatePolicyApi.md#CreateVSmartTemplate) | **Post** /template/policy/vsmart | 
[**DeActivatePolicy**](ConfigurationVSmartTemplatePolicyApi.md#DeActivatePolicy) | **Post** /template/policy/vsmart/deactivate/{policyId} | 
[**DeleteVSmartTemplate**](ConfigurationVSmartTemplatePolicyApi.md#DeleteVSmartTemplate) | **Delete** /template/policy/vsmart/{policyId} | 
[**EditTemplateWithoutLockChecks**](ConfigurationVSmartTemplatePolicyApi.md#EditTemplateWithoutLockChecks) | **Put** /template/policy/vsmart/central/{policyId} | 
[**EditVSmartTemplate**](ConfigurationVSmartTemplatePolicyApi.md#EditVSmartTemplate) | **Put** /template/policy/vsmart/{policyId} | 
[**GenerateVSmartPolicyTemplateList**](ConfigurationVSmartTemplatePolicyApi.md#GenerateVSmartPolicyTemplateList) | **Get** /template/policy/vsmart | 
[**GetTemplateByPolicyId**](ConfigurationVSmartTemplatePolicyApi.md#GetTemplateByPolicyId) | **Get** /template/policy/vsmart/definition/{policyId} | 
[**QosmosNbarMigrationWarning**](ConfigurationVSmartTemplatePolicyApi.md#QosmosNbarMigrationWarning) | **Get** /template/policy/vsmart/qosmos_nbar_migration_warning | 



## ActivatePolicy

> map[string]interface{} ActivatePolicy(ctx, policyId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.ActivatePolicy(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.ActivatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivatePolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.ActivatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Template policy | 

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


## ActivatePolicyForCloudServices

> map[string]interface{} ActivatePolicyForCloudServices(ctx, policyId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.ActivatePolicyForCloudServices(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.ActivatePolicyForCloudServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivatePolicyForCloudServices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.ActivatePolicyForCloudServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyForCloudServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Template policy | 

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


## CheckVSmartConnectivityStatus

> map[string]interface{} CheckVSmartConnectivityStatus(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.CheckVSmartConnectivityStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.CheckVSmartConnectivityStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckVSmartConnectivityStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.CheckVSmartConnectivityStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckVSmartConnectivityStatusRequest struct via the builder pattern


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


## CreateVSmartTemplate

> map[string]interface{} CreateVSmartTemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.CreateVSmartTemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.CreateVSmartTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVSmartTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.CreateVSmartTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVSmartTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Template policy | 

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


## DeActivatePolicy

> map[string]interface{} DeActivatePolicy(ctx, policyId).Execute()





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
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.DeActivatePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.DeActivatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeActivatePolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.DeActivatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeActivatePolicyRequest struct via the builder pattern


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


## DeleteVSmartTemplate

> DeleteVSmartTemplate(ctx, policyId).Execute()





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
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.DeleteVSmartTemplate(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.DeleteVSmartTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVSmartTemplateRequest struct via the builder pattern


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


## EditTemplateWithoutLockChecks

> []map[string]interface{} EditTemplateWithoutLockChecks(ctx, policyId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.EditTemplateWithoutLockChecks(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.EditTemplateWithoutLockChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTemplateWithoutLockChecks`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.EditTemplateWithoutLockChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTemplateWithoutLockChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Template policy | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditVSmartTemplate

> []map[string]interface{} EditVSmartTemplate(ctx, policyId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Template policy (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.EditVSmartTemplate(context.Background(), policyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.EditVSmartTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditVSmartTemplate`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.EditVSmartTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVSmartTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Template policy | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateVSmartPolicyTemplateList

> []map[string]interface{} GenerateVSmartPolicyTemplateList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.GenerateVSmartPolicyTemplateList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.GenerateVSmartPolicyTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateVSmartPolicyTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.GenerateVSmartPolicyTemplateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateVSmartPolicyTemplateListRequest struct via the builder pattern


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


## GetTemplateByPolicyId

> map[string]interface{} GetTemplateByPolicyId(ctx, policyId).Execute()





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
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.GetTemplateByPolicyId(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.GetTemplateByPolicyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateByPolicyId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.GetTemplateByPolicyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Policy Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByPolicyIdRequest struct via the builder pattern


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


## QosmosNbarMigrationWarning

> map[string]interface{} QosmosNbarMigrationWarning(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationVSmartTemplatePolicyApi.QosmosNbarMigrationWarning(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVSmartTemplatePolicyApi.QosmosNbarMigrationWarning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QosmosNbarMigrationWarning`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationVSmartTemplatePolicyApi.QosmosNbarMigrationWarning`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQosmosNbarMigrationWarningRequest struct via the builder pattern


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

