# \ConfigurationConfigurationGroupApi

All URIs are relative to *http://$VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigGroup**](ConfigurationConfigurationGroupApi.md#CreateConfigGroup) | **Post** /v1/config-group | 
[**CreateConfigGroupAssociation**](ConfigurationConfigurationGroupApi.md#CreateConfigGroupAssociation) | **Post** /v1/config-group/{configGroupId}/device/associate | 
[**CreateConfigGroupDeviceVariables**](ConfigurationConfigurationGroupApi.md#CreateConfigGroupDeviceVariables) | **Put** /v1/config-group/{configGroupId}/device/variables | 
[**CreateConfigGroupDeviceVariables1**](ConfigurationConfigurationGroupApi.md#CreateConfigGroupDeviceVariables1) | **Get** /v1/config-group/{configGroupId}/device/variables/schema | 
[**DeleteConfigGroup**](ConfigurationConfigurationGroupApi.md#DeleteConfigGroup) | **Delete** /v1/config-group/{configGroupId} | 
[**DeleteConfigGroupAssociation**](ConfigurationConfigurationGroupApi.md#DeleteConfigGroupAssociation) | **Delete** /v1/config-group/{configGroupId}/device/associate | 
[**DeployConfigGroup**](ConfigurationConfigurationGroupApi.md#DeployConfigGroup) | **Post** /v1/config-group/{configGroupId}/device/deploy | 
[**EditConfigGroup**](ConfigurationConfigurationGroupApi.md#EditConfigGroup) | **Put** /v1/config-group/{configGroupId} | 
[**GetCedgeConfigGroupSchemaBySchemaType**](ConfigurationConfigurationGroupApi.md#GetCedgeConfigGroupSchemaBySchemaType) | **Get** /v1/config-group/schema/sdwan | 
[**GetConfigGroup**](ConfigurationConfigurationGroupApi.md#GetConfigGroup) | **Get** /v1/config-group/{configGroupId} | 
[**GetConfigGroupAssociation**](ConfigurationConfigurationGroupApi.md#GetConfigGroupAssociation) | **Get** /v1/config-group/{configGroupId}/device/associate | 
[**GetConfigGroupBySolution**](ConfigurationConfigurationGroupApi.md#GetConfigGroupBySolution) | **Get** /v1/config-group | 
[**GetConfigGroupDeviceConfigurationPreview**](ConfigurationConfigurationGroupApi.md#GetConfigGroupDeviceConfigurationPreview) | **Post** /v1/config-group/{configGroupId}/device/{deviceId}/preview | 
[**GetConfigGroupDeviceVariables**](ConfigurationConfigurationGroupApi.md#GetConfigGroupDeviceVariables) | **Get** /v1/config-group/{configGroupId}/device/variables | 
[**UpdateConfigGroupAssociation**](ConfigurationConfigurationGroupApi.md#UpdateConfigGroupAssociation) | **Put** /v1/config-group/{configGroupId}/device/associate | 



## CreateConfigGroup

> string CreateConfigGroup(ctx).Body(body).Execute()





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
    body := "TODO" // string | Config Group (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.CreateConfigGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.CreateConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.CreateConfigGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Config Group | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigGroupAssociation

> CreateConfigGroupAssociation(ctx, configGroupId).Body(body).Execute()





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
    configGroupId := "configGroupId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.CreateConfigGroupAssociation(context.Background(), configGroupId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.CreateConfigGroupAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigGroupAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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


## CreateConfigGroupDeviceVariables

> map[string]interface{} CreateConfigGroupDeviceVariables(ctx, configGroupId).RequestBody(requestBody).Execute()





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
    configGroupId := "configGroupId_example" // string | Config Group Id
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.CreateConfigGroupDeviceVariables(context.Background(), configGroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.CreateConfigGroupDeviceVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigGroupDeviceVariables`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.CreateConfigGroupDeviceVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** | Config Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigGroupDeviceVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**map[string]GetO365PreferredPathFromVAnalyticsRequestValue**](GetO365PreferredPathFromVAnalyticsRequestValue.md) |  | 

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


## CreateConfigGroupDeviceVariables1

> map[string]interface{} CreateConfigGroupDeviceVariables1(ctx, configGroupId).Execute()





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
    configGroupId := "configGroupId_example" // string | Config Group Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.CreateConfigGroupDeviceVariables1(context.Background(), configGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.CreateConfigGroupDeviceVariables1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigGroupDeviceVariables1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.CreateConfigGroupDeviceVariables1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** | Config Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigGroupDeviceVariables1Request struct via the builder pattern


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


## DeleteConfigGroup

> DeleteConfigGroup(ctx, configGroupId).Execute()





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
    configGroupId := "configGroupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.DeleteConfigGroup(context.Background(), configGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.DeleteConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigGroupRequest struct via the builder pattern


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


## DeleteConfigGroupAssociation

> DeleteConfigGroupAssociation(ctx, configGroupId).Body(body).Execute()





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
    configGroupId := "configGroupId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.DeleteConfigGroupAssociation(context.Background(), configGroupId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.DeleteConfigGroupAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigGroupAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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


## DeployConfigGroup

> String DeployConfigGroup(ctx, configGroupId).RequestBody(requestBody).Execute()





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
    configGroupId := "configGroupId_example" // string | Config Group Id
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.DeployConfigGroup(context.Background(), configGroupId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.DeployConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployConfigGroup`: String
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.DeployConfigGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** | Config Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**map[string]GetO365PreferredPathFromVAnalyticsRequestValue**](GetO365PreferredPathFromVAnalyticsRequestValue.md) |  | 

### Return type

[**String**](String.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditConfigGroup

> string EditConfigGroup(ctx, configGroupId).Body(body).Execute()





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
    configGroupId := "configGroupId_example" // string | 
    body := "TODO" // string | Config Group (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.EditConfigGroup(context.Background(), configGroupId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.EditConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditConfigGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.EditConfigGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Config Group | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCedgeConfigGroupSchemaBySchemaType

> string GetCedgeConfigGroupSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





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
    schemaType := "schemaType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.GetCedgeConfigGroupSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.GetCedgeConfigGroupSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCedgeConfigGroupSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.GetCedgeConfigGroupSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCedgeConfigGroupSchemaBySchemaTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaType** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigGroup

> ConfigGroup GetConfigGroup(ctx, configGroupId).Execute()





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
    configGroupId := "configGroupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.GetConfigGroup(context.Background(), configGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.GetConfigGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigGroup`: ConfigGroup
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.GetConfigGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigGroup**](ConfigGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigGroupAssociation

> GetConfigGroupAssociation(ctx, configGroupId).Execute()





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
    configGroupId := "configGroupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.GetConfigGroupAssociation(context.Background(), configGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.GetConfigGroupAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigGroupAssociationRequest struct via the builder pattern


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


## GetConfigGroupBySolution

> string GetConfigGroupBySolution(ctx).Solution(solution).Execute()





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
    solution := "solution_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.GetConfigGroupBySolution(context.Background()).Solution(solution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.GetConfigGroupBySolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigGroupBySolution`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.GetConfigGroupBySolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigGroupBySolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **solution** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigGroupDeviceConfigurationPreview

> map[string]interface{} GetConfigGroupDeviceConfigurationPreview(ctx, configGroupId, deviceId).RequestBody(requestBody).Execute()





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
    configGroupId := "configGroupId_example" // string | Config Group Id
    deviceId := "deviceId_example" // string | Device Id
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.GetConfigGroupDeviceConfigurationPreview(context.Background(), configGroupId, deviceId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.GetConfigGroupDeviceConfigurationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigGroupDeviceConfigurationPreview`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.GetConfigGroupDeviceConfigurationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** | Config Group Id | 
**deviceId** | **string** | Device Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigGroupDeviceConfigurationPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | [**map[string]GetO365PreferredPathFromVAnalyticsRequestValue**](GetO365PreferredPathFromVAnalyticsRequestValue.md) |  | 

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


## GetConfigGroupDeviceVariables

> map[string]interface{} GetConfigGroupDeviceVariables(ctx, configGroupId).DeviceId(deviceId).Suggestions(suggestions).Execute()





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
    configGroupId := "configGroupId_example" // string | Config Group Id
    deviceId := "deviceId_example" // string | Comma separated device id's like d1,d2 (optional)
    suggestions := true // bool | Suggestions for possible values (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.GetConfigGroupDeviceVariables(context.Background(), configGroupId).DeviceId(deviceId).Suggestions(suggestions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.GetConfigGroupDeviceVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigGroupDeviceVariables`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationConfigurationGroupApi.GetConfigGroupDeviceVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** | Config Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigGroupDeviceVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **string** | Comma separated device id&#39;s like d1,d2 | 
 **suggestions** | **bool** | Suggestions for possible values | 

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


## UpdateConfigGroupAssociation

> UpdateConfigGroupAssociation(ctx, configGroupId).Body(body).Execute()





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
    configGroupId := "configGroupId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationConfigurationGroupApi.UpdateConfigGroupAssociation(context.Background(), configGroupId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationConfigurationGroupApi.UpdateConfigGroupAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigGroupAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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

