# \ConfigurationNetworkDesignApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquireAttachLock**](ConfigurationNetworkDesignApi.md#AcquireAttachLock) | **Post** /networkdesign/profile/lock/{profileId} | 
[**AcquireEditLock**](ConfigurationNetworkDesignApi.md#AcquireEditLock) | **Post** /networkdesign/lock/edit | 
[**CreateNetworkDesign**](ConfigurationNetworkDesignApi.md#CreateNetworkDesign) | **Post** /networkdesign | 
[**EditNetworkDesign**](ConfigurationNetworkDesignApi.md#EditNetworkDesign) | **Put** /networkdesign | 
[**GetDeviceProfileConfigStatus**](ConfigurationNetworkDesignApi.md#GetDeviceProfileConfigStatus) | **Get** /networkdesign/profile/status | 
[**GetDeviceProfileConfigStatusByProfileId**](ConfigurationNetworkDesignApi.md#GetDeviceProfileConfigStatusByProfileId) | **Get** /networkdesign/profile/status/{profileId} | 
[**GetDeviceProfileTaskCount**](ConfigurationNetworkDesignApi.md#GetDeviceProfileTaskCount) | **Get** /networkdesign/profile/task/count | 
[**GetDeviceProfileTaskStatus**](ConfigurationNetworkDesignApi.md#GetDeviceProfileTaskStatus) | **Get** /networkdesign/profile/task/status | 
[**GetDeviceProfileTaskStatusByProfileId**](ConfigurationNetworkDesignApi.md#GetDeviceProfileTaskStatusByProfileId) | **Get** /networkdesign/profile/task/status/{profileId} | 
[**GetGlobalParameters**](ConfigurationNetworkDesignApi.md#GetGlobalParameters) | **Get** /networkdesign/global/parameters | 
[**GetNetworkDesign**](ConfigurationNetworkDesignApi.md#GetNetworkDesign) | **Get** /networkdesign | 
[**GetServiceProfileConfig**](ConfigurationNetworkDesignApi.md#GetServiceProfileConfig) | **Get** /networkdesign/serviceProfileConfig/{profileId} | 
[**PushDeviceProfileTemplate**](ConfigurationNetworkDesignApi.md#PushDeviceProfileTemplate) | **Post** /networkdesign/profile/attachment/{profileId} | 
[**PushNetworkDesign**](ConfigurationNetworkDesignApi.md#PushNetworkDesign) | **Post** /networkdesign/attachment | 
[**RunMyTest**](ConfigurationNetworkDesignApi.md#RunMyTest) | **Get** /networkdesign/mytest/{name} | 



## AcquireAttachLock

> map[string]interface{} AcquireAttachLock(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Device profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.AcquireAttachLock(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.AcquireAttachLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireAttachLock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.AcquireAttachLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Device profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireAttachLockRequest struct via the builder pattern


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


## AcquireEditLock

> map[string]interface{} AcquireEditLock(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.AcquireEditLock(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.AcquireEditLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcquireEditLock`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.AcquireEditLock`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcquireEditLockRequest struct via the builder pattern


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


## CreateNetworkDesign

> map[string]interface{} CreateNetworkDesign(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Network design payload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.CreateNetworkDesign(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.CreateNetworkDesign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkDesign`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.CreateNetworkDesign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDesignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Network design payload | 

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


## EditNetworkDesign

> map[string]interface{} EditNetworkDesign(ctx).Id(id).Body(body).Execute()





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
    id := "id_example" // string | Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Network design payload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.EditNetworkDesign(context.Background()).Id(id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.EditNetworkDesign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditNetworkDesign`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.EditNetworkDesign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditNetworkDesignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Id | 
 **body** | **map[string]interface{}** | Network design payload | 

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


## GetDeviceProfileConfigStatus

> map[string]interface{} GetDeviceProfileConfigStatus(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileConfigStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileConfigStatusRequest struct via the builder pattern


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


## GetDeviceProfileConfigStatusByProfileId

> map[string]interface{} GetDeviceProfileConfigStatusByProfileId(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Device profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatusByProfileId(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatusByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileConfigStatusByProfileId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatusByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Device profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileConfigStatusByProfileIdRequest struct via the builder pattern


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


## GetDeviceProfileTaskCount

> map[string]interface{} GetDeviceProfileTaskCount(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileTaskCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetDeviceProfileTaskCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileTaskCount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetDeviceProfileTaskCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileTaskCountRequest struct via the builder pattern


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


## GetDeviceProfileTaskStatus

> map[string]interface{} GetDeviceProfileTaskStatus(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileTaskStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileTaskStatusRequest struct via the builder pattern


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


## GetDeviceProfileTaskStatusByProfileId

> map[string]interface{} GetDeviceProfileTaskStatusByProfileId(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Device profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatusByProfileId(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatusByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfileTaskStatusByProfileId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatusByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Device profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfileTaskStatusByProfileIdRequest struct via the builder pattern


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


## GetGlobalParameters

> map[string]interface{} GetGlobalParameters(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetGlobalParameters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetGlobalParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalParameters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetGlobalParameters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalParametersRequest struct via the builder pattern


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


## GetNetworkDesign

> map[string]interface{} GetNetworkDesign(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetNetworkDesign(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetNetworkDesign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDesign`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetNetworkDesign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDesignRequest struct via the builder pattern


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


## GetServiceProfileConfig

> map[string]interface{} GetServiceProfileConfig(ctx, profileId).DeviceModel(deviceModel).Execute()





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
    profileId := "profileId_example" // string | Device profile Id
    deviceModel := "deviceModel_example" // string | Device model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.GetServiceProfileConfig(context.Background(), profileId).DeviceModel(deviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.GetServiceProfileConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceProfileConfig`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.GetServiceProfileConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Device profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceProfileConfigRequest struct via the builder pattern


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


## PushDeviceProfileTemplate

> map[string]interface{} PushDeviceProfileTemplate(ctx, profileId).Body(body).Execute()





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
    profileId := "profileId_example" // string | Device profile Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Device template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.PushDeviceProfileTemplate(context.Background(), profileId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.PushDeviceProfileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushDeviceProfileTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.PushDeviceProfileTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Device profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushDeviceProfileTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device template | 

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


## PushNetworkDesign

> map[string]interface{} PushNetworkDesign(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.PushNetworkDesign(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.PushNetworkDesign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushNetworkDesign`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.PushNetworkDesign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushNetworkDesignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device template | 

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


## RunMyTest

> map[string]interface{} RunMyTest(ctx, name).Execute()





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
    name := "name_example" // string | Test bane

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationNetworkDesignApi.RunMyTest(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationNetworkDesignApi.RunMyTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunMyTest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationNetworkDesignApi.RunMyTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Test bane | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunMyTestRequest struct via the builder pattern


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

