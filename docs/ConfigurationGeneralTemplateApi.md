# \ConfigurationGeneralTemplateApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTemplateResourceGroup**](ConfigurationGeneralTemplateApi.md#ChangeTemplateResourceGroup) | **Post** /template/feature/resource-group/{resourceGroupName}/{templateId} | 
[**CloneTemplate**](ConfigurationGeneralTemplateApi.md#CloneTemplate) | **Post** /template/feature/clone | 
[**CreateFeatureTemplate**](ConfigurationGeneralTemplateApi.md#CreateFeatureTemplate) | **Post** /template/feature | 
[**CreateLITemplate**](ConfigurationGeneralTemplateApi.md#CreateLITemplate) | **Post** /template/feature/li | 
[**DeleteGeneralTemplate**](ConfigurationGeneralTemplateApi.md#DeleteGeneralTemplate) | **Delete** /template/feature/{templateId} | 
[**EditFeatureTemplate**](ConfigurationGeneralTemplateApi.md#EditFeatureTemplate) | **Put** /template/feature/{templateId} | 
[**EditLITemplate**](ConfigurationGeneralTemplateApi.md#EditLITemplate) | **Put** /template/feature/li/{templateId} | 
[**GenerateFeatureTemplateList**](ConfigurationGeneralTemplateApi.md#GenerateFeatureTemplateList) | **Get** /template/feature | 
[**GenerateMasterTemplateDefinition**](ConfigurationGeneralTemplateApi.md#GenerateMasterTemplateDefinition) | **Get** /template/feature/master/{type_name} | 
[**GenerateTemplateByDeviceType**](ConfigurationGeneralTemplateApi.md#GenerateTemplateByDeviceType) | **Get** /template/feature/{deviceType} | 
[**GenerateTemplateTypeDefinition**](ConfigurationGeneralTemplateApi.md#GenerateTemplateTypeDefinition) | **Get** /template/feature/types/definition/{type_name}/{version} | 
[**GenerateTemplateTypes**](ConfigurationGeneralTemplateApi.md#GenerateTemplateTypes) | **Get** /template/feature/types | 
[**GetDefaultNetworks**](ConfigurationGeneralTemplateApi.md#GetDefaultNetworks) | **Get** /template/feature/default/networks | 
[**GetDeviceTemplatesAttachedToFeature**](ConfigurationGeneralTemplateApi.md#GetDeviceTemplatesAttachedToFeature) | **Get** /template/feature/devicetemplates/{templateId} | 
[**GetEncryptedString**](ConfigurationGeneralTemplateApi.md#GetEncryptedString) | **Post** /template/security/encryptText/encrypt | 
[**GetGeneralTemplate**](ConfigurationGeneralTemplateApi.md#GetGeneralTemplate) | **Get** /template/feature/object/{templateId} | 
[**GetNetworkInterface**](ConfigurationGeneralTemplateApi.md#GetNetworkInterface) | **Get** /template/feature/default/networkinterface | 
[**GetTemplateDefinition**](ConfigurationGeneralTemplateApi.md#GetTemplateDefinition) | **Get** /template/feature/definition/{templateId} | 
[**GetTemplateForMigration**](ConfigurationGeneralTemplateApi.md#GetTemplateForMigration) | **Get** /template/feature/migration | 
[**ListLITemplate**](ConfigurationGeneralTemplateApi.md#ListLITemplate) | **Get** /template/feature/li | 



## ChangeTemplateResourceGroup

> ChangeTemplateResourceGroup(ctx, templateId, resourceGroupName).Execute()





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
    resourceGroupName := "resourceGroupName_example" // string | Resrouce group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.ChangeTemplateResourceGroup(context.Background(), templateId, resourceGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.ChangeTemplateResourceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 
**resourceGroupName** | **string** | Resrouce group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeTemplateResourceGroupRequest struct via the builder pattern


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


## CloneTemplate

> map[string]interface{} CloneTemplate(ctx).Id(id).Name(name).Desc(desc).Execute()





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
    id := "220a6bd0-28ef-4c88-92e6-ee7539396fd7" // string | Template Id to clone from
    name := "BR2-VPN10-Feature" // string | Name for the cloned template
    desc := "desc_example" // string | Description for the cloned template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.CloneTemplate(context.Background()).Id(id).Name(name).Desc(desc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.CloneTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.CloneTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloneTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Template Id to clone from | 
 **name** | **string** | Name for the cloned template | 
 **desc** | **string** | Description for the cloned template | 

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


## CreateFeatureTemplate

> map[string]interface{} CreateFeatureTemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Feature template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.CreateFeatureTemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.CreateFeatureTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeatureTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.CreateFeatureTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Feature template | 

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


## CreateLITemplate

> map[string]interface{} CreateLITemplate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | LI template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.CreateLITemplate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.CreateLITemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLITemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.CreateLITemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLITemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | LI template | 

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


## DeleteGeneralTemplate

> DeleteGeneralTemplate(ctx, templateId).Execute()





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
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.DeleteGeneralTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.DeleteGeneralTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGeneralTemplateRequest struct via the builder pattern


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


## EditFeatureTemplate

> map[string]interface{} EditFeatureTemplate(ctx, templateId).Body(body).Execute()





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
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.EditFeatureTemplate(context.Background(), templateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.EditFeatureTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditFeatureTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.EditFeatureTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditFeatureTemplateRequest struct via the builder pattern


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


## EditLITemplate

> map[string]interface{} EditLITemplate(ctx, templateId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | LI template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.EditLITemplate(context.Background(), templateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.EditLITemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditLITemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.EditLITemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLITemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | LI template | 

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


## GenerateFeatureTemplateList

> []map[string]interface{} GenerateFeatureTemplateList(ctx).Summary(summary).Offset(offset).Limit(limit).Execute()





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
    summary := true // bool | Flag to include template definition (optional) (default to false)
    offset := int32(56) // int32 | Pagination offset (optional)
    limit := int32(56) // int32 | Pagination limit on templateId (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GenerateFeatureTemplateList(context.Background()).Summary(summary).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GenerateFeatureTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateFeatureTemplateList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GenerateFeatureTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateFeatureTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **summary** | **bool** | Flag to include template definition | [default to false]
 **offset** | **int32** | Pagination offset | 
 **limit** | **int32** | Pagination limit on templateId | [default to 0]

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


## GenerateMasterTemplateDefinition

> map[string]interface{} GenerateMasterTemplateDefinition(ctx, typeName).Execute()





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
    typeName := "typeName_example" // string | Device type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GenerateMasterTemplateDefinition(context.Background(), typeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GenerateMasterTemplateDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateMasterTemplateDefinition`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GenerateMasterTemplateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | Device type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMasterTemplateDefinitionRequest struct via the builder pattern


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


## GenerateTemplateByDeviceType

> []map[string]interface{} GenerateTemplateByDeviceType(ctx, deviceType).Execute()





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
    deviceType := "deviceType_example" // string | Device type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GenerateTemplateByDeviceType(context.Background(), deviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GenerateTemplateByDeviceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateTemplateByDeviceType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GenerateTemplateByDeviceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceType** | **string** | Device type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTemplateByDeviceTypeRequest struct via the builder pattern


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


## GenerateTemplateTypeDefinition

> []map[string]interface{} GenerateTemplateTypeDefinition(ctx, typeName, version).Execute()





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
    typeName := "typeName_example" // string | Feature template type
    version := "version_example" // string | Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GenerateTemplateTypeDefinition(context.Background(), typeName, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GenerateTemplateTypeDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateTemplateTypeDefinition`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GenerateTemplateTypeDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | Feature template type | 
**version** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTemplateTypeDefinitionRequest struct via the builder pattern


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


## GenerateTemplateTypes

> []map[string]interface{} GenerateTemplateTypes(ctx).Type_(type_).Execute()





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
    type_ := "type__example" // string | Device type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GenerateTemplateTypes(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GenerateTemplateTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateTemplateTypes`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GenerateTemplateTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTemplateTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Device type | 

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


## GetDefaultNetworks

> map[string]interface{} GetDefaultNetworks(ctx).DeviceModel(deviceModel).Execute()





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
    deviceModel := "deviceModel_example" // string | Device model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetDefaultNetworks(context.Background()).DeviceModel(deviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetDefaultNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultNetworks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetDefaultNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceModel** | **string** | Device model | 

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


## GetDeviceTemplatesAttachedToFeature

> map[string]interface{} GetDeviceTemplatesAttachedToFeature(ctx, templateId).Execute()





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
    templateId := "templateId_example" // string | Feature template Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetDeviceTemplatesAttachedToFeature(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetDeviceTemplatesAttachedToFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTemplatesAttachedToFeature`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetDeviceTemplatesAttachedToFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Feature template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTemplatesAttachedToFeatureRequest struct via the builder pattern


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


## GetEncryptedString

> map[string]interface{} GetEncryptedString(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Type6 Encryption (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetEncryptedString(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetEncryptedString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEncryptedString`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetEncryptedString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptedStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Type6 Encryption | 

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


## GetGeneralTemplate

> map[string]interface{} GetGeneralTemplate(ctx, templateId).Execute()





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
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetGeneralTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetGeneralTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeneralTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetGeneralTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralTemplateRequest struct via the builder pattern


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


## GetNetworkInterface

> map[string]interface{} GetNetworkInterface(ctx).DeviceModel(deviceModel).Execute()





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
    deviceModel := "deviceModel_example" // string | Device model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetNetworkInterface(context.Background()).DeviceModel(deviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceModel** | **string** | Device model | 

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


## GetTemplateDefinition

> map[string]interface{} GetTemplateDefinition(ctx, templateId).Execute()





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
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetTemplateDefinition(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetTemplateDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateDefinition`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetTemplateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Template Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateDefinitionRequest struct via the builder pattern


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


## GetTemplateForMigration

> []map[string]interface{} GetTemplateForMigration(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.GetTemplateForMigration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.GetTemplateForMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateForMigration`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.GetTemplateForMigration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateForMigrationRequest struct via the builder pattern


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


## ListLITemplate

> []map[string]interface{} ListLITemplate(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationGeneralTemplateApi.ListLITemplate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationGeneralTemplateApi.ListLITemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLITemplate`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationGeneralTemplateApi.ListLITemplate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLITemplateRequest struct via the builder pattern


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

