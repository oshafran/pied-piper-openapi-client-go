# \RealTimeMonitoringIPsecApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptoIpsecIdentity**](RealTimeMonitoringIPsecApi.md#CreateCryptoIpsecIdentity) | **Get** /device/ipsec/identity | 
[**CreateCryptov1LocalSAList**](RealTimeMonitoringIPsecApi.md#CreateCryptov1LocalSAList) | **Get** /device/ipsec/ikev1 | 
[**CreateCryptov2LocalSAList**](RealTimeMonitoringIPsecApi.md#CreateCryptov2LocalSAList) | **Get** /device/ipsec/ikev2 | 
[**CreateIPsecPWKInboundConnections**](RealTimeMonitoringIPsecApi.md#CreateIPsecPWKInboundConnections) | **Get** /device/ipsec/pwk/inbound | 
[**CreateIPsecPWKLocalSA**](RealTimeMonitoringIPsecApi.md#CreateIPsecPWKLocalSA) | **Get** /device/ipsec/pwk/localsa | 
[**CreateIPsecPWKOutboundConnections**](RealTimeMonitoringIPsecApi.md#CreateIPsecPWKOutboundConnections) | **Get** /device/ipsec/pwk/outbound | 
[**CreateIkeInboundList**](RealTimeMonitoringIPsecApi.md#CreateIkeInboundList) | **Get** /device/ipsec/ike/inbound | 
[**CreateIkeOutboundList**](RealTimeMonitoringIPsecApi.md#CreateIkeOutboundList) | **Get** /device/ipsec/ike/outbound | 
[**CreateIkeSessions**](RealTimeMonitoringIPsecApi.md#CreateIkeSessions) | **Get** /device/ipsec/ike/sessions | 
[**CreateInBoundList**](RealTimeMonitoringIPsecApi.md#CreateInBoundList) | **Get** /device/ipsec/inbound | 
[**CreateLocalSAList**](RealTimeMonitoringIPsecApi.md#CreateLocalSAList) | **Get** /device/ipsec/localsa | 
[**CreateOutBoundList**](RealTimeMonitoringIPsecApi.md#CreateOutBoundList) | **Get** /device/ipsec/outbound | 



## CreateCryptoIpsecIdentity

> map[string]interface{} CreateCryptoIpsecIdentity(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)
    localTlocColor := "localTlocColor_example" // string | Local tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateCryptoIpsecIdentity(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateCryptoIpsecIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCryptoIpsecIdentity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateCryptoIpsecIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptoIpsecIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 
 **localTlocColor** | **string** | Local tloc color | 

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


## CreateCryptov1LocalSAList

> map[string]interface{} CreateCryptov1LocalSAList(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateCryptov1LocalSAList(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateCryptov1LocalSAList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCryptov1LocalSAList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateCryptov1LocalSAList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptov1LocalSAListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 

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


## CreateCryptov2LocalSAList

> map[string]interface{} CreateCryptov2LocalSAList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateCryptov2LocalSAList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateCryptov2LocalSAList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCryptov2LocalSAList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateCryptov2LocalSAList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptov2LocalSAListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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


## CreateIPsecPWKInboundConnections

> map[string]interface{} CreateIPsecPWKInboundConnections(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)
    localTlocColor := "localTlocColor_example" // string | Local tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateIPsecPWKInboundConnections(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateIPsecPWKInboundConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPsecPWKInboundConnections`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateIPsecPWKInboundConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPsecPWKInboundConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 
 **localTlocColor** | **string** | Local tloc color | 

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


## CreateIPsecPWKLocalSA

> map[string]interface{} CreateIPsecPWKLocalSA(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)
    localTlocColor := "localTlocColor_example" // string | Local tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateIPsecPWKLocalSA(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateIPsecPWKLocalSA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPsecPWKLocalSA`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateIPsecPWKLocalSA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPsecPWKLocalSARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 
 **localTlocColor** | **string** | Local tloc color | 

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


## CreateIPsecPWKOutboundConnections

> map[string]interface{} CreateIPsecPWKOutboundConnections(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)
    localTlocColor := "localTlocColor_example" // string | Local tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateIPsecPWKOutboundConnections(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateIPsecPWKOutboundConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIPsecPWKOutboundConnections`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateIPsecPWKOutboundConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPsecPWKOutboundConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 
 **localTlocColor** | **string** | Local tloc color | 

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


## CreateIkeInboundList

> []map[string]interface{} CreateIkeInboundList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateIkeInboundList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateIkeInboundList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIkeInboundList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateIkeInboundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIkeInboundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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


## CreateIkeOutboundList

> []map[string]interface{} CreateIkeOutboundList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateIkeOutboundList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateIkeOutboundList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIkeOutboundList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateIkeOutboundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIkeOutboundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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


## CreateIkeSessions

> []map[string]interface{} CreateIkeSessions(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateIkeSessions(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateIkeSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIkeSessions`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateIkeSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIkeSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 

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


## CreateInBoundList

> []map[string]interface{} CreateInBoundList(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)
    localTlocColor := "localTlocColor_example" // string | Local tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateInBoundList(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).LocalTlocColor(localTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateInBoundList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInBoundList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateInBoundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInBoundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 
 **localTlocColor** | **string** | Local tloc color | 

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


## CreateLocalSAList

> []map[string]interface{} CreateLocalSAList(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device IP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateLocalSAList(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateLocalSAList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocalSAList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateLocalSAList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalSAListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 

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


## CreateOutBoundList

> []map[string]interface{} CreateOutBoundList(ctx).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).Execute()





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
    deviceId := "deviceId_example" // string | Device IP
    remoteTlocAddress := "remoteTlocAddress_example" // string | Remote TLOC address (optional)
    remoteTlocColor := "remoteTlocColor_example" // string | Remote tloc color (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealTimeMonitoringIPsecApi.CreateOutBoundList(context.Background()).DeviceId(deviceId).RemoteTlocAddress(remoteTlocAddress).RemoteTlocColor(remoteTlocColor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimeMonitoringIPsecApi.CreateOutBoundList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOutBoundList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RealTimeMonitoringIPsecApi.CreateOutBoundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOutBoundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device IP | 
 **remoteTlocAddress** | **string** | Remote TLOC address | 
 **remoteTlocColor** | **string** | Remote tloc color | 

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

