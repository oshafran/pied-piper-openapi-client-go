# \ConfigurationTemplateMasterApi

All URIs are relative to *https://44.196.44.132*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTemplateResourceGroup1**](ConfigurationTemplateMasterApi.md#ChangeTemplateResourceGroup1) | **Post** /template/device/resource-group/{resourceGroupName}/{templateId} | 
[**CreateCLITemplate**](ConfigurationTemplateMasterApi.md#CreateCLITemplate) | **Post** /template/device/cli | 
[**CreateMasterTemplate**](ConfigurationTemplateMasterApi.md#CreateMasterTemplate) | **Post** /template/device/feature | 
[**DeleteMasterTemplate**](ConfigurationTemplateMasterApi.md#DeleteMasterTemplate) | **Delete** /template/device/{templateId} | 
[**EditMasterTemplate**](ConfigurationTemplateMasterApi.md#EditMasterTemplate) | **Put** /template/device/{templateId} | 
[**GenerateMasterTemplateList**](ConfigurationTemplateMasterApi.md#GenerateMasterTemplateList) | **Get** /template/device | 
[**GenerateTemplateForMigration**](ConfigurationTemplateMasterApi.md#GenerateTemplateForMigration) | **Get** /template/device/migration | 
[**GetMasterTemplateDefinition**](ConfigurationTemplateMasterApi.md#GetMasterTemplateDefinition) | **Get** /template/device/object/{templateId} | 
[**GetOutOfSyncDevices**](ConfigurationTemplateMasterApi.md#GetOutOfSyncDevices) | **Get** /template/device/syncstatus/{templateId} | 
[**GetOutOfSyncTemplates**](ConfigurationTemplateMasterApi.md#GetOutOfSyncTemplates) | **Get** /template/device/syncstatus | 
[**IsMigrationRequired**](ConfigurationTemplateMasterApi.md#IsMigrationRequired) | **Get** /template/device/is_migration_required | 
[**MigrateTemplates**](ConfigurationTemplateMasterApi.md#MigrateTemplates) | **Post** /template/device/migration | 
[**MigrationInfo**](ConfigurationTemplateMasterApi.md#MigrationInfo) | **Get** /template/device/migration_info | 



## ChangeTemplateResourceGroup1

> ChangeTemplateResourceGroup1(ctx, templateId, resourceGroupName).Execute()





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
    templateId := "templateId_example" // string | Template Id
    resourceGroupName := "resourceGroupName_example" // string | Resource group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.ChangeTemplateResourceGroup1(context.Background(), templateId, resourceGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.ChangeTemplateResourceGroup1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 
**resourceGroupName** | **string** | Resource group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTemplateResourceGroup1Request struct via the builder pattern


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


## CreateCLITemplate

> map[string]interface{} CreateCLITemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Create template request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.CreateCLITemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.CreateCLITemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCLITemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.CreateCLITemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCLITemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Create template request | 

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


## CreateMasterTemplate

> map[string]interface{} CreateMasterTemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Create template request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.CreateMasterTemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.CreateMasterTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMasterTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.CreateMasterTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMasterTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Create template request | 

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


## DeleteMasterTemplate

> DeleteMasterTemplate(ctx, templateId).Execute()





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
    templateId := "templateId_example" // string | Template Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.DeleteMasterTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.DeleteMasterTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMasterTemplateRequest struct via the builder pattern


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


## EditMasterTemplate

> map[string]interface{} EditMasterTemplate(ctx, templateId).Body(body).Execute()





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
    templateId := "templateId_example" // string | Template Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.EditMasterTemplate(context.Background(), templateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.EditMasterTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMasterTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.EditMasterTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMasterTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Template | 

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


## GenerateMasterTemplateList

> []map[string]interface{} GenerateMasterTemplateList(ctx).Feature(feature).Execute()





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
    feature := "feature_example" // string | Feature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.GenerateMasterTemplateList(context.Background()).Feature(feature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.GenerateMasterTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateMasterTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.GenerateMasterTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMasterTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feature** | **string** | Feature | 

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


## GenerateTemplateForMigration

> []map[string]interface{} GenerateTemplateForMigration(ctx).HasAAA(hasAAA).Execute()





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
    hasAAA := true // bool | Return only those uses AAA (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.GenerateTemplateForMigration(context.Background()).HasAAA(hasAAA).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.GenerateTemplateForMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateTemplateForMigration`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.GenerateTemplateForMigration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTemplateForMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hasAAA** | **bool** | Return only those uses AAA | 

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


## GetMasterTemplateDefinition

> map[string]interface{} GetMasterTemplateDefinition(ctx, templateId).Execute()





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
    templateId := "templateId_example" // string | Template Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.GetMasterTemplateDefinition(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.GetMasterTemplateDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMasterTemplateDefinition`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.GetMasterTemplateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterTemplateDefinitionRequest struct via the builder pattern


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


## GetOutOfSyncDevices

> []map[string]interface{} GetOutOfSyncDevices(ctx, templateId).Execute()





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
    templateId := "templateId_example" // string | Template Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.GetOutOfSyncDevices(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.GetOutOfSyncDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutOfSyncDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.GetOutOfSyncDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutOfSyncDevicesRequest struct via the builder pattern


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


## GetOutOfSyncTemplates

> []map[string]interface{} GetOutOfSyncTemplates(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.GetOutOfSyncTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.GetOutOfSyncTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutOfSyncTemplates`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.GetOutOfSyncTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutOfSyncTemplatesRequest struct via the builder pattern


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


## IsMigrationRequired

> map[string]interface{} IsMigrationRequired(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.IsMigrationRequired(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.IsMigrationRequired``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IsMigrationRequired`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.IsMigrationRequired`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIsMigrationRequiredRequest struct via the builder pattern


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


## MigrateTemplates

> map[string]interface{} MigrateTemplates(ctx).Id(id).Prefix(prefix).IncludeAll(includeAll).Execute()





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
    id := []string{"Inner_example"} // []string | Template Id
    prefix := "prefix_example" // string | Prefix (optional) (default to "cisco")
    includeAll := true // bool | Include all flag (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.MigrateTemplates(context.Background()).Id(id).Prefix(prefix).IncludeAll(includeAll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.MigrateTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrateTemplates`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.MigrateTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Template Id | 
 **prefix** | **string** | Prefix | [default to &quot;cisco&quot;]
 **includeAll** | **bool** | Include all flag | [default to true]

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


## MigrationInfo

> map[string]interface{} MigrationInfo(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationTemplateMasterApi.MigrationInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTemplateMasterApi.MigrationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrationInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationTemplateMasterApi.MigrationInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationInfoRequest struct via the builder pattern


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

