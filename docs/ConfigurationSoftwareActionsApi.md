# \ConfigurationSoftwareActionsApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRemoteServer**](ConfigurationSoftwareActionsApi.md#AddRemoteServer) | **Post** /device/action/remote-server | 
[**CreateImageURL**](ConfigurationSoftwareActionsApi.md#CreateImageURL) | **Post** /device/action/software | 
[**DeleteImageURL**](ConfigurationSoftwareActionsApi.md#DeleteImageURL) | **Delete** /device/action/software/{versionId} | 
[**DeleteRemoteServer**](ConfigurationSoftwareActionsApi.md#DeleteRemoteServer) | **Delete** /device/action/remote-server/{id} | 
[**FindSoftwareImages**](ConfigurationSoftwareActionsApi.md#FindSoftwareImages) | **Get** /device/action/software | 
[**FindSoftwareImagesWithFilters**](ConfigurationSoftwareActionsApi.md#FindSoftwareImagesWithFilters) | **Get** /device/action/software/images | 
[**FindSoftwareVersion**](ConfigurationSoftwareActionsApi.md#FindSoftwareVersion) | **Get** /device/action/software/version | 
[**FindVEdgeSoftwareVersion**](ConfigurationSoftwareActionsApi.md#FindVEdgeSoftwareVersion) | **Get** /device/action/software/vedge/version | 
[**FindZtpSoftwareVersion**](ConfigurationSoftwareActionsApi.md#FindZtpSoftwareVersion) | **Get** /device/action/software/ztp/version | 
[**GetImageProperties**](ConfigurationSoftwareActionsApi.md#GetImageProperties) | **Get** /device/action/software/imageProperties/{versionId} | 
[**GetPnfProperties**](ConfigurationSoftwareActionsApi.md#GetPnfProperties) | **Get** /device/action/software/pnfproperties/{pnfType} | 
[**GetRemoteServerById**](ConfigurationSoftwareActionsApi.md#GetRemoteServerById) | **Get** /device/action/remote-server/{id} | 
[**GetRemoteServerList**](ConfigurationSoftwareActionsApi.md#GetRemoteServerList) | **Get** /device/action/remote-server | 
[**GetVnfProperties**](ConfigurationSoftwareActionsApi.md#GetVnfProperties) | **Get** /device/action/software/vnfproperties/{versionId} | 
[**UpdateImageURL**](ConfigurationSoftwareActionsApi.md#UpdateImageURL) | **Put** /device/action/software/{versionId} | 
[**UpdateRemoteServer**](ConfigurationSoftwareActionsApi.md#UpdateRemoteServer) | **Put** /device/action/remote-server/{id} | 



## AddRemoteServer

> AddRemoteServer(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Request payload for a new remote server entry. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.AddRemoteServer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.AddRemoteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRemoteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Request payload for a new remote server entry. | 

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


## CreateImageURL

> CreateImageURL(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Create software image request payload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.CreateImageURL(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.CreateImageURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Create software image request payload | 

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


## DeleteImageURL

> DeleteImageURL(ctx, versionId).Execute()





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
    versionId := "versionId_example" // string | Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.DeleteImageURL(context.Background(), versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.DeleteImageURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageURLRequest struct via the builder pattern


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


## DeleteRemoteServer

> []map[string]interface{} DeleteRemoteServer(ctx, id).Execute()





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
    id := "id_example" // string | Remote Server ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.DeleteRemoteServer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.DeleteRemoteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRemoteServer`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.DeleteRemoteServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Remote Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteServerRequest struct via the builder pattern


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


## FindSoftwareImages

> []map[string]interface{} FindSoftwareImages(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.FindSoftwareImages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.FindSoftwareImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSoftwareImages`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.FindSoftwareImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindSoftwareImagesRequest struct via the builder pattern


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


## FindSoftwareImagesWithFilters

> []map[string]interface{} FindSoftwareImagesWithFilters(ctx).ImageType(imageType).VnfType(vnfType).Execute()





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
    imageType := "imageType_example" // string | Image type
    vnfType := "vnfType_example" // string | VNF type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.FindSoftwareImagesWithFilters(context.Background()).ImageType(imageType).VnfType(vnfType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.FindSoftwareImagesWithFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSoftwareImagesWithFilters`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.FindSoftwareImagesWithFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindSoftwareImagesWithFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageType** | **string** | Image type | 
 **vnfType** | **string** | VNF type | 

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


## FindSoftwareVersion

> map[string]interface{} FindSoftwareVersion(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.FindSoftwareVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.FindSoftwareVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSoftwareVersion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.FindSoftwareVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindSoftwareVersionRequest struct via the builder pattern


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


## FindVEdgeSoftwareVersion

> map[string]interface{} FindVEdgeSoftwareVersion(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.FindVEdgeSoftwareVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.FindVEdgeSoftwareVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindVEdgeSoftwareVersion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.FindVEdgeSoftwareVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindVEdgeSoftwareVersionRequest struct via the builder pattern


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


## FindZtpSoftwareVersion

> map[string]interface{} FindZtpSoftwareVersion(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.FindZtpSoftwareVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.FindZtpSoftwareVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindZtpSoftwareVersion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.FindZtpSoftwareVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindZtpSoftwareVersionRequest struct via the builder pattern


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


## GetImageProperties

> map[string]interface{} GetImageProperties(ctx, versionId).Execute()





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
    versionId := "versionId_example" // string | Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.GetImageProperties(context.Background(), versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.GetImageProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.GetImageProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagePropertiesRequest struct via the builder pattern


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


## GetPnfProperties

> map[string]interface{} GetPnfProperties(ctx, pnfType).Execute()





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
    pnfType := "pnfType_example" // string | PNF type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.GetPnfProperties(context.Background(), pnfType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.GetPnfProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPnfProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.GetPnfProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pnfType** | **string** | PNF type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPnfPropertiesRequest struct via the builder pattern


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


## GetRemoteServerById

> []map[string]interface{} GetRemoteServerById(ctx, id).Execute()





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
    id := "id_example" // string | Remote Server ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.GetRemoteServerById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.GetRemoteServerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteServerById`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.GetRemoteServerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Remote Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteServerByIdRequest struct via the builder pattern


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


## GetRemoteServerList

> []map[string]interface{} GetRemoteServerList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.GetRemoteServerList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.GetRemoteServerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteServerList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.GetRemoteServerList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteServerListRequest struct via the builder pattern


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


## GetVnfProperties

> map[string]interface{} GetVnfProperties(ctx, versionId).Execute()





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
    versionId := "versionId_example" // string | Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.GetVnfProperties(context.Background(), versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.GetVnfProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVnfProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.GetVnfProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVnfPropertiesRequest struct via the builder pattern


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


## UpdateImageURL

> UpdateImageURL(ctx, versionId).Body(body).Execute()





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
    versionId := "versionId_example" // string | Version
    body := map[string]interface{}{ ... } // map[string]interface{} | Update software image request payload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.UpdateImageURL(context.Background(), versionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.UpdateImageURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Update software image request payload | 

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


## UpdateRemoteServer

> []map[string]interface{} UpdateRemoteServer(ctx, id).Body(body).Execute()





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
    id := "id_example" // string | Remote Server ID
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationSoftwareActionsApi.UpdateRemoteServer(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationSoftwareActionsApi.UpdateRemoteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteServer`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationSoftwareActionsApi.UpdateRemoteServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Remote Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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

