# \TroubleshootingToolsDeviceConnectivityApi

All URIs are relative to *http://$VMANAGE_EXTERNAL_IP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyAdminTechOnDevice**](TroubleshootingToolsDeviceConnectivityApi.md#CopyAdminTechOnDevice) | **Post** /device/tools/admintech/copy | 
[**CreateAdminTech**](TroubleshootingToolsDeviceConnectivityApi.md#CreateAdminTech) | **Post** /device/tools/admintech | 
[**DeleteAdminTechFile**](TroubleshootingToolsDeviceConnectivityApi.md#DeleteAdminTechFile) | **Delete** /device/tools/admintech/{requestID} | 
[**DeleteAdminTechOnDevice**](TroubleshootingToolsDeviceConnectivityApi.md#DeleteAdminTechOnDevice) | **Delete** /device/tools/admintech/delete | 
[**DownloadAdminTechFile**](TroubleshootingToolsDeviceConnectivityApi.md#DownloadAdminTechFile) | **Get** /device/tools/admintech/download/{filename} | 
[**FactoryReset**](TroubleshootingToolsDeviceConnectivityApi.md#FactoryReset) | **Post** /device/tools/factoryreset | 
[**GetControlConnections**](TroubleshootingToolsDeviceConnectivityApi.md#GetControlConnections) | **Get** /troubleshooting/control/{uuid} | 
[**GetDeviceConfiguration**](TroubleshootingToolsDeviceConnectivityApi.md#GetDeviceConfiguration) | **Get** /troubleshooting/devicebringup | 
[**GetInProgressCount**](TroubleshootingToolsDeviceConnectivityApi.md#GetInProgressCount) | **Get** /device/tools/admintechs/inprogress | 
[**InvalidateDevice**](TroubleshootingToolsDeviceConnectivityApi.md#InvalidateDevice) | **Post** /certificate/device/invalidate | 
[**ListAdminTechs**](TroubleshootingToolsDeviceConnectivityApi.md#ListAdminTechs) | **Get** /device/tools/admintechs | 
[**ListAdminTechsOnDevice**](TroubleshootingToolsDeviceConnectivityApi.md#ListAdminTechsOnDevice) | **Post** /device/tools/admintechlist | 
[**NpingDevice**](TroubleshootingToolsDeviceConnectivityApi.md#NpingDevice) | **Post** /device/tools/nping/{deviceIP} | 
[**PingDevice**](TroubleshootingToolsDeviceConnectivityApi.md#PingDevice) | **Post** /device/tools/ping/{deviceIP} | 
[**ProcessInterfaceReset**](TroubleshootingToolsDeviceConnectivityApi.md#ProcessInterfaceReset) | **Post** /device/tools/reset/interface/{deviceIP} | 
[**ProcessPortHopColor**](TroubleshootingToolsDeviceConnectivityApi.md#ProcessPortHopColor) | **Post** /device/tools/porthopcolor/{deviceIP} | 
[**ProcessResetUser**](TroubleshootingToolsDeviceConnectivityApi.md#ProcessResetUser) | **Post** /device/tools/resetuser/{deviceIP} | 
[**ServicePath**](TroubleshootingToolsDeviceConnectivityApi.md#ServicePath) | **Post** /device/tools/servicepath/{deviceIP} | 
[**StageDevice**](TroubleshootingToolsDeviceConnectivityApi.md#StageDevice) | **Post** /certificate/device/stage | 
[**TracerouteDevice**](TroubleshootingToolsDeviceConnectivityApi.md#TracerouteDevice) | **Post** /device/tools/traceroute/{deviceIP} | 
[**TunnelPath**](TroubleshootingToolsDeviceConnectivityApi.md#TunnelPath) | **Post** /device/tools/tunnelpath/{deviceIP} | 
[**UploadAdminTech**](TroubleshootingToolsDeviceConnectivityApi.md#UploadAdminTech) | **Post** /device/tools/admintechs/upload | 



## CopyAdminTechOnDevice

> CopyAdminTechOnDevice(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Admin tech copy request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.CopyAdminTechOnDevice(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.CopyAdminTechOnDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopyAdminTechOnDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Admin tech copy request | 

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


## CreateAdminTech

> CreateAdminTech(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Admin tech generation request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.CreateAdminTech(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.CreateAdminTech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminTechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Admin tech generation request | 

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


## DeleteAdminTechFile

> DeleteAdminTechFile(ctx, requestID).Execute()





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
    requestID := "requestID_example" // string | Request Id of admin tech generation request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.DeleteAdminTechFile(context.Background(), requestID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.DeleteAdminTechFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** | Request Id of admin tech generation request | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminTechFileRequest struct via the builder pattern


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


## DeleteAdminTechOnDevice

> DeleteAdminTechOnDevice(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Admin tech copy request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.DeleteAdminTechOnDevice(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.DeleteAdminTechOnDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdminTechOnDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Admin tech copy request | 

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


## DownloadAdminTechFile

> map[string]interface{} DownloadAdminTechFile(ctx, filename).Execute()





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
    filename := "filename_example" // string | Admin tech file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.DownloadAdminTechFile(context.Background(), filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.DownloadAdminTechFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAdminTechFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.DownloadAdminTechFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Admin tech file | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAdminTechFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FactoryReset

> FactoryReset(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device factory reset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.FactoryReset(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.FactoryReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFactoryResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device factory reset | 

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


## GetControlConnections

> map[string]interface{} GetControlConnections(ctx, uuid).Execute()





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
    uuid := "uuid_example" // string | Device uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.GetControlConnections(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.GetControlConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControlConnections`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.GetControlConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetControlConnectionsRequest struct via the builder pattern


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


## GetDeviceConfiguration

> map[string]interface{} GetDeviceConfiguration(ctx).Uuid(uuid).Execute()





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
    uuid := "uuid_example" // string | Device uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.GetDeviceConfiguration(context.Background()).Uuid(uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.GetDeviceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.GetDeviceConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** | Device uuid | 

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


## GetInProgressCount

> map[string]interface{} GetInProgressCount(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.GetInProgressCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.GetInProgressCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInProgressCount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.GetInProgressCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInProgressCountRequest struct via the builder pattern


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


## InvalidateDevice

> map[string]interface{} InvalidateDevice(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | vEdge device info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.InvalidateDevice(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.InvalidateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvalidateDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.InvalidateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | vEdge device info | 

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


## ListAdminTechs

> map[string]interface{} ListAdminTechs(ctx).Execute()





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
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.ListAdminTechs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.ListAdminTechs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAdminTechs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.ListAdminTechs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdminTechsRequest struct via the builder pattern


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


## ListAdminTechsOnDevice

> map[string]interface{} ListAdminTechsOnDevice(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Admin tech listing request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.ListAdminTechsOnDevice(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.ListAdminTechsOnDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAdminTechsOnDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.ListAdminTechsOnDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdminTechsOnDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Admin tech listing request | 

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


## NpingDevice

> NpingDevice(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | NPing parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.NpingDevice(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.NpingDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiNpingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | NPing parameter | 

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


## PingDevice

> PingDevice(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | Ping parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.PingDevice(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.PingDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Ping parameter | 

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


## ProcessInterfaceReset

> ProcessInterfaceReset(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | Device interface (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.ProcessInterfaceReset(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.ProcessInterfaceReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessInterfaceResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device interface | 

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


## ProcessPortHopColor

> ProcessPortHopColor(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | Device port hop color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.ProcessPortHopColor(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.ProcessPortHopColor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessPortHopColorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device port hop color | 

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


## ProcessResetUser

> ProcessResetUser(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | Device user reset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.ProcessResetUser(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.ProcessResetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessResetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device user reset | 

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


## ServicePath

> ServicePath(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | Service path parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.ServicePath(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.ServicePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Service path parameter | 

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


## StageDevice

> map[string]interface{} StageDevice(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | vEdge device info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.StageDevice(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.StageDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StageDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingToolsDeviceConnectivityApi.StageDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStageDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | vEdge device info | 

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


## TracerouteDevice

> TracerouteDevice(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | Traceroute parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.TracerouteDevice(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.TracerouteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiTracerouteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Traceroute parameter | 

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


## TunnelPath

> TunnelPath(ctx, deviceIP).Body(body).Execute()





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
    deviceIP := "deviceIP_example" // string | Device IP
    body := map[string]interface{}{ ... } // map[string]interface{} | TunnelPath parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.TunnelPath(context.Background(), deviceIP).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.TunnelPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIP** | **string** | Device IP | 

### Other Parameters

Other parameters are passed through a pointer to a apiTunnelPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | TunnelPath parameter | 

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


## UploadAdminTech

> UploadAdminTech(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Admin tech upload request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootingToolsDeviceConnectivityApi.UploadAdminTech(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingToolsDeviceConnectivityApi.UploadAdminTech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAdminTechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Admin tech upload request | 

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

