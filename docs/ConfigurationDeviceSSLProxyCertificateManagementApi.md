# \ConfigurationDeviceSSLProxyCertificateManagementApi

All URIs are relative to *http://11.222.33.444*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWANEdge**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#AddWANEdge) | **Post** /sslproxy/certificate/wanedge/{deviceId} | 
[**GenerateSSLProxyCSR**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GenerateSSLProxyCSR) | **Post** /sslproxy/generate/vmanage/csr | 
[**GenerateSslProxyCSR**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GenerateSslProxyCSR) | **Post** /sslproxy/generate/csr/sslproxy | 
[**GetAllDeviceCSR**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetAllDeviceCSR) | **Post** /sslproxy/devicecsr | 
[**GetAllDeviceCertificates**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetAllDeviceCertificates) | **Post** /sslproxy/devicecertificates | 
[**GetCertificateState**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetCertificateState) | **Get** /sslproxy/settings/certificate | 
[**GetEnterpriseCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetEnterpriseCertificate) | **Get** /sslproxy/settings/enterprise/certificate | 
[**GetProxyCertOfEdge**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetProxyCertOfEdge) | **Get** /sslproxy/certificate | 
[**GetSelfSignedCert**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetSelfSignedCert) | **Get** /certificate/vmanage/selfsignedcert | 
[**GetSslProxyCSR**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetSslProxyCSR) | **Get** /sslproxy/csr | 
[**GetSslProxyList**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetSslProxyList) | **Get** /sslproxy/list | 
[**GetVManageEnterpriseRootCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetVManageEnterpriseRootCertificate) | **Get** /sslproxy/settings/enterprise/rootca | 
[**GetvManageCSR**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetvManageCSR) | **Get** /sslproxy/settings/vmanage/csr | 
[**GetvManageCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetvManageCertificate) | **Get** /sslproxy/settings/vmanage/certificate | 
[**GetvManageRootCA**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#GetvManageRootCA) | **Get** /sslproxy/settings/vmanage/rootca | 
[**RenewCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#RenewCertificate) | **Post** /sslproxy/renew | 
[**RevokeCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#RevokeCertificate) | **Post** /sslproxy/revoke | 
[**RevokeRenewCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#RevokeRenewCertificate) | **Post** /sslproxy/revokerenew | 
[**SetEnterpriseCert**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#SetEnterpriseCert) | **Post** /sslproxy/settings/enterprise/certificate | 
[**SetEnterpriseRootCaCert**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#SetEnterpriseRootCaCert) | **Post** /sslproxy/settings/enterprise/rootca | 
[**SetvManageRootCA**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#SetvManageRootCA) | **Post** /sslproxy/settings/vmanage/rootca | 
[**SetvManageintermediateCert**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#SetvManageintermediateCert) | **Post** /sslproxy/settings/vmanage/certificate | 
[**UpdateCertificate**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#UpdateCertificate) | **Put** /sslproxy/certificate | 
[**UploadCertificiates**](ConfigurationDeviceSSLProxyCertificateManagementApi.md#UploadCertificiates) | **Post** /sslproxy/certificates | 



## AddWANEdge

> AddWANEdge(ctx, deviceId).Body(body).Execute()





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
    deviceId := "deviceId_example" // string | Device Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Cert state (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.AddWANEdge(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.AddWANEdge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWANEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Cert state | 

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


## GenerateSSLProxyCSR

> map[string]interface{} GenerateSSLProxyCSR(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | CSR request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GenerateSSLProxyCSR(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GenerateSSLProxyCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSSLProxyCSR`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GenerateSSLProxyCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSSLProxyCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | CSR request | 

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


## GenerateSslProxyCSR

> GenerateSslProxyCSR(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | CSR request for edge (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GenerateSslProxyCSR(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GenerateSslProxyCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSslProxyCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | CSR request for edge | 

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


## GetAllDeviceCSR

> map[string]interface{} GetAllDeviceCSR(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetAllDeviceCSR(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetAllDeviceCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDeviceCSR`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetAllDeviceCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeviceCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device list | 

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


## GetAllDeviceCertificates

> map[string]interface{} GetAllDeviceCertificates(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetAllDeviceCertificates(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetAllDeviceCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDeviceCertificates`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetAllDeviceCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeviceCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device list | 

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


## GetCertificateState

> map[string]interface{} GetCertificateState(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetCertificateState(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetCertificateState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateState`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetCertificateState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateStateRequest struct via the builder pattern


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


## GetEnterpriseCertificate

> map[string]interface{} GetEnterpriseCertificate(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetEnterpriseCertificate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetEnterpriseCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnterpriseCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetEnterpriseCertificate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnterpriseCertificateRequest struct via the builder pattern


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


## GetProxyCertOfEdge

> map[string]interface{} GetProxyCertOfEdge(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetProxyCertOfEdge(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetProxyCertOfEdge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProxyCertOfEdge`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetProxyCertOfEdge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProxyCertOfEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device Id | 

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


## GetSelfSignedCert

> map[string]interface{} GetSelfSignedCert(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetSelfSignedCert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetSelfSignedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfSignedCert`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetSelfSignedCert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfSignedCertRequest struct via the builder pattern


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


## GetSslProxyCSR

> map[string]interface{} GetSslProxyCSR(ctx).DeviceId(deviceId).Execute()





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
    deviceId := "8d86d8b2-2239-402e-9fef-467f7bad3f2f" // string | device UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetSslProxyCSR(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetSslProxyCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslProxyCSR`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetSslProxyCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSslProxyCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | device UUID | 

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


## GetSslProxyList

> []map[string]interface{} GetSslProxyList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetSslProxyList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetSslProxyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslProxyList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetSslProxyList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslProxyListRequest struct via the builder pattern


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


## GetVManageEnterpriseRootCertificate

> map[string]interface{} GetVManageEnterpriseRootCertificate(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetVManageEnterpriseRootCertificate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetVManageEnterpriseRootCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVManageEnterpriseRootCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetVManageEnterpriseRootCertificate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVManageEnterpriseRootCertificateRequest struct via the builder pattern


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


## GetvManageCSR

> map[string]interface{} GetvManageCSR(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageCSR(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetvManageCSR`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageCSR`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetvManageCSRRequest struct via the builder pattern


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


## GetvManageCertificate

> map[string]interface{} GetvManageCertificate(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageCertificate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetvManageCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageCertificate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetvManageCertificateRequest struct via the builder pattern


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


## GetvManageRootCA

> map[string]interface{} GetvManageRootCA(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageRootCA(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageRootCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetvManageRootCA`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.GetvManageRootCA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetvManageRootCARequest struct via the builder pattern


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


## RenewCertificate

> map[string]interface{} RenewCertificate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Renew device certificate request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.RenewCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.RenewCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.RenewCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenewCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Renew device certificate request | 

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


## RevokeCertificate

> map[string]interface{} RevokeCertificate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Revoke device certificate request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.RevokeCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.RevokeCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.RevokeCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Revoke device certificate request | 

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


## RevokeRenewCertificate

> map[string]interface{} RevokeRenewCertificate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Revoke device certificate request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.RevokeRenewCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.RevokeRenewCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeRenewCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.RevokeRenewCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRenewCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Revoke device certificate request | 

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


## SetEnterpriseCert

> map[string]interface{} SetEnterpriseCert(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Config enterprise certificate request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.SetEnterpriseCert(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.SetEnterpriseCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEnterpriseCert`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.SetEnterpriseCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEnterpriseCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Config enterprise certificate request | 

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


## SetEnterpriseRootCaCert

> map[string]interface{} SetEnterpriseRootCaCert(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Set enterprise root CA request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.SetEnterpriseRootCaCert(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.SetEnterpriseRootCaCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEnterpriseRootCaCert`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.SetEnterpriseRootCaCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEnterpriseRootCaCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Set enterprise root CA request | 

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


## SetvManageRootCA

> map[string]interface{} SetvManageRootCA(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Set vManage root CA request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.SetvManageRootCA(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.SetvManageRootCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetvManageRootCA`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.SetvManageRootCA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetvManageRootCARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Set vManage root CA request | 

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


## SetvManageintermediateCert

> map[string]interface{} SetvManageintermediateCert(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Set vManage intermediate CA request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.SetvManageintermediateCert(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.SetvManageintermediateCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetvManageintermediateCert`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.SetvManageintermediateCert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetvManageintermediateCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Set vManage intermediate CA request | 

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


## UpdateCertificate

> map[string]interface{} UpdateCertificate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Upload device certificate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.UpdateCertificate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.UpdateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceSSLProxyCertificateManagementApi.UpdateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Upload device certificate | 

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


## UploadCertificiates

> UploadCertificiates(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceSSLProxyCertificateManagementApi.UploadCertificiates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceSSLProxyCertificateManagementApi.UploadCertificiates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCertificiatesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

