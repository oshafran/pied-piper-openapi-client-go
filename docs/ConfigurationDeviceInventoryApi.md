# \ConfigurationDeviceInventoryApi

All URIs are relative to *http://$VMANAGEHOST*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckSelfSignedCert**](ConfigurationDeviceInventoryApi.md#CheckSelfSignedCert) | **Get** /system/device/selfsignedcert/iscreated | 
[**ClaimDevices**](ConfigurationDeviceInventoryApi.md#ClaimDevices) | **Post** /system/device/claimDevices | 
[**CreateDevice**](ConfigurationDeviceInventoryApi.md#CreateDevice) | **Post** /system/device | 
[**DecommissionVedgeCloud**](ConfigurationDeviceInventoryApi.md#DecommissionVedgeCloud) | **Put** /system/device/decommission/{uuid} | 
[**DeleteDevice**](ConfigurationDeviceInventoryApi.md#DeleteDevice) | **Delete** /system/device/{uuid} | 
[**DevicesWithoutSubjectSudi**](ConfigurationDeviceInventoryApi.md#DevicesWithoutSubjectSudi) | **Get** /system/device/devicesWithoutSubjectSudi | 
[**EditDevice**](ConfigurationDeviceInventoryApi.md#EditDevice) | **Put** /system/device/{uuid} | 
[**FormPost**](ConfigurationDeviceInventoryApi.md#FormPost) | **Post** /system/device/fileupload | 
[**GenerateBootstrapConfigForVedge**](ConfigurationDeviceInventoryApi.md#GenerateBootstrapConfigForVedge) | **Get** /system/device/bootstrap/device/{uuid} | 
[**GenerateBootstrapConfigForVedges**](ConfigurationDeviceInventoryApi.md#GenerateBootstrapConfigForVedges) | **Post** /system/device/bootstrap/devices | 
[**GenerateGenericBootstrapConfigForVedges**](ConfigurationDeviceInventoryApi.md#GenerateGenericBootstrapConfigForVedges) | **Get** /system/device/bootstrap/generic/devices | 
[**GetAllUnclaimedDevices**](ConfigurationDeviceInventoryApi.md#GetAllUnclaimedDevices) | **Get** /system/device/unclaimedDevices | 
[**GetBootstrapConfigZip**](ConfigurationDeviceInventoryApi.md#GetBootstrapConfigZip) | **Get** /system/device/bootstrap/download/{id} | 
[**GetCloudDockDataBasedOnDeviceType**](ConfigurationDeviceInventoryApi.md#GetCloudDockDataBasedOnDeviceType) | **Get** /system/device/type/{deviceCategory} | 
[**GetCloudDockDefaultConfigBasedOnDeviceType**](ConfigurationDeviceInventoryApi.md#GetCloudDockDefaultConfigBasedOnDeviceType) | **Get** /system/device/type/{deviceCategory}/defaultConfig | 
[**GetControllerVEdgeSyncStatus**](ConfigurationDeviceInventoryApi.md#GetControllerVEdgeSyncStatus) | **Get** /system/device/controllers/vedge/status | 
[**GetDevicesDetails**](ConfigurationDeviceInventoryApi.md#GetDevicesDetails) | **Get** /system/device/{deviceCategory} | 
[**GetManagementSystemIPInfo**](ConfigurationDeviceInventoryApi.md#GetManagementSystemIPInfo) | **Get** /system/device/management/systemip | 
[**GetRMACandidates**](ConfigurationDeviceInventoryApi.md#GetRMACandidates) | **Get** /system/device/rma/candidates/{deviceType} | 
[**GetRootCertStatusAll**](ConfigurationDeviceInventoryApi.md#GetRootCertStatusAll) | **Get** /system/device/rootcertchain/status | 
[**GetTenantManagementSystemIPs**](ConfigurationDeviceInventoryApi.md#GetTenantManagementSystemIPs) | **Get** /system/device/tenant/management/systemip | 
[**InvalidateVmanageRootCA**](ConfigurationDeviceInventoryApi.md#InvalidateVmanageRootCA) | **Delete** /system/device/vmanagerootca/{uuid} | 
[**MigrateDevice**](ConfigurationDeviceInventoryApi.md#MigrateDevice) | **Put** /system/device/migrateDevice/{uuid} | 
[**ResetVedgeCloud**](ConfigurationDeviceInventoryApi.md#ResetVedgeCloud) | **Put** /system/device/reset/{uuid} | 
[**SetLifeCycle**](ConfigurationDeviceInventoryApi.md#SetLifeCycle) | **Post** /system/device/lifecycle/management/{uuid} | 
[**SyncDevices**](ConfigurationDeviceInventoryApi.md#SyncDevices) | **Post** /system/device/smartaccount/sync | 
[**SyncRootCertChain**](ConfigurationDeviceInventoryApi.md#SyncRootCertChain) | **Get** /system/device/sync/rootcertchain | 
[**UpdateDeviceSubjectSUDI**](ConfigurationDeviceInventoryApi.md#UpdateDeviceSubjectSUDI) | **Put** /system/device/updateDeviceSubjectSUDI/{uuid} | 
[**ValidateUser**](ConfigurationDeviceInventoryApi.md#ValidateUser) | **Post** /system/device/smartaccount/authenticate | 
[**ValidateUser1**](ConfigurationDeviceInventoryApi.md#ValidateUser1) | **Post** /system/device/generate-payg | 



## CheckSelfSignedCert

> map[string]interface{} CheckSelfSignedCert(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.CheckSelfSignedCert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.CheckSelfSignedCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckSelfSignedCert`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.CheckSelfSignedCert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckSelfSignedCertRequest struct via the builder pattern


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


## ClaimDevices

> []map[string]interface{} ClaimDevices(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Claim device request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.ClaimDevices(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.ClaimDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.ClaimDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClaimDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Claim device request | 

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


## CreateDevice

> CreateDevice(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Create device request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.CreateDevice(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Create device request | 

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


## DecommissionVedgeCloud

> map[string]interface{} DecommissionVedgeCloud(ctx, uuid).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.DecommissionVedgeCloud(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.DecommissionVedgeCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecommissionVedgeCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.DecommissionVedgeCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecommissionVedgeCloudRequest struct via the builder pattern


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


## DeleteDevice

> DeleteDevice(ctx, uuid).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.DeleteDevice(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## DevicesWithoutSubjectSudi

> []map[string]interface{} DevicesWithoutSubjectSudi(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.DevicesWithoutSubjectSudi(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.DevicesWithoutSubjectSudi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesWithoutSubjectSudi`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.DevicesWithoutSubjectSudi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesWithoutSubjectSudiRequest struct via the builder pattern


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


## EditDevice

> EditDevice(ctx, uuid).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device config (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.EditDevice(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.EditDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Device config | 

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


## FormPost

> FormPost(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.FormPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.FormPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFormPostRequest struct via the builder pattern


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


## GenerateBootstrapConfigForVedge

> map[string]interface{} GenerateBootstrapConfigForVedge(ctx, uuid).Configtype(configtype).InclDefRootCert(inclDefRootCert).Version(version).Execute()





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
    configtype := "configtype_example" // string | Device config type
    inclDefRootCert := true // bool | Include default root certs flag (default to true)
    version := "version_example" // string | cloud-init format version (optional) (default to "v1")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GenerateBootstrapConfigForVedge(context.Background(), uuid).Configtype(configtype).InclDefRootCert(inclDefRootCert).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GenerateBootstrapConfigForVedge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateBootstrapConfigForVedge`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GenerateBootstrapConfigForVedge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBootstrapConfigForVedgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configtype** | **string** | Device config type | 
 **inclDefRootCert** | **bool** | Include default root certs flag | [default to true]
 **version** | **string** | cloud-init format version | [default to &quot;v1&quot;]

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


## GenerateBootstrapConfigForVedges

> map[string]interface{} GenerateBootstrapConfigForVedges(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device bootstrap type and id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GenerateBootstrapConfigForVedges(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GenerateBootstrapConfigForVedges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateBootstrapConfigForVedges`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GenerateBootstrapConfigForVedges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBootstrapConfigForVedgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device bootstrap type and id | 

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


## GenerateGenericBootstrapConfigForVedges

> map[string]interface{} GenerateGenericBootstrapConfigForVedges(ctx).Wanif(wanif).Execute()





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
    wanif := "wanif_example" // string | Device WAN interface (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GenerateGenericBootstrapConfigForVedges(context.Background()).Wanif(wanif).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GenerateGenericBootstrapConfigForVedges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateGenericBootstrapConfigForVedges`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GenerateGenericBootstrapConfigForVedges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGenericBootstrapConfigForVedgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wanif** | **string** | Device WAN interface | 

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


## GetAllUnclaimedDevices

> []map[string]interface{} GetAllUnclaimedDevices(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetAllUnclaimedDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetAllUnclaimedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllUnclaimedDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetAllUnclaimedDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUnclaimedDevicesRequest struct via the builder pattern


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


## GetBootstrapConfigZip

> map[string]interface{} GetBootstrapConfigZip(ctx, id).Execute()





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
    id := "id_example" // string | Bootstrap config id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetBootstrapConfigZip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetBootstrapConfigZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBootstrapConfigZip`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetBootstrapConfigZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Bootstrap config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBootstrapConfigZipRequest struct via the builder pattern


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


## GetCloudDockDataBasedOnDeviceType

> []map[string]interface{} GetCloudDockDataBasedOnDeviceType(ctx, deviceCategory).Execute()





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
    deviceCategory := "deviceCategory_example" // string | Device category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetCloudDockDataBasedOnDeviceType(context.Background(), deviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetCloudDockDataBasedOnDeviceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDockDataBasedOnDeviceType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetCloudDockDataBasedOnDeviceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategory** | **string** | Device category | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDockDataBasedOnDeviceTypeRequest struct via the builder pattern


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


## GetCloudDockDefaultConfigBasedOnDeviceType

> []map[string]interface{} GetCloudDockDefaultConfigBasedOnDeviceType(ctx, deviceCategory).Execute()





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
    deviceCategory := "deviceCategory_example" // string | Device category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetCloudDockDefaultConfigBasedOnDeviceType(context.Background(), deviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetCloudDockDefaultConfigBasedOnDeviceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDockDefaultConfigBasedOnDeviceType`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetCloudDockDefaultConfigBasedOnDeviceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategory** | **string** | Device category | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDockDefaultConfigBasedOnDeviceTypeRequest struct via the builder pattern


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


## GetControllerVEdgeSyncStatus

> []map[string]interface{} GetControllerVEdgeSyncStatus(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetControllerVEdgeSyncStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetControllerVEdgeSyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetControllerVEdgeSyncStatus`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetControllerVEdgeSyncStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetControllerVEdgeSyncStatusRequest struct via the builder pattern


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


## GetDevicesDetails

> []map[string]interface{} GetDevicesDetails(ctx, deviceCategory).Model(model).State(state).Uuid(uuid).DeviceIP(deviceIP).Validity(validity).Family(family).Execute()





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
    deviceCategory := "deviceCategory_example" // string | Device category
    model := "model_example" // string | Device model (optional)
    state := []openapiclient.CertificateStates{*openapiclient.NewCertificateStates()} // []CertificateStates | List of states (optional)
    uuid := []openapiclient.DeviceUuid{*openapiclient.NewDeviceUuid()} // []DeviceUuid | List of device uuid (optional)
    deviceIP := []openapiclient.DeviceIP{*openapiclient.NewDeviceIP()} // []DeviceIP | List of device system IP (optional)
    validity := []openapiclient.CertificateValidity{*openapiclient.NewCertificateValidity()} // []CertificateValidity | List of device validity (optional)
    family := "family_example" // string | The platform family to filter for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetDevicesDetails(context.Background(), deviceCategory).Model(model).State(state).Uuid(uuid).DeviceIP(deviceIP).Validity(validity).Family(family).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetDevicesDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicesDetails`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetDevicesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategory** | **string** | Device category | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **string** | Device model | 
 **state** | [**[]CertificateStates**](CertificateStates.md) | List of states | 
 **uuid** | [**[]DeviceUuid**](DeviceUuid.md) | List of device uuid | 
 **deviceIP** | [**[]DeviceIP**](DeviceIP.md) | List of device system IP | 
 **validity** | [**[]CertificateValidity**](CertificateValidity.md) | List of device validity | 
 **family** | **string** | The platform family to filter for | 

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


## GetManagementSystemIPInfo

> map[string]interface{} GetManagementSystemIPInfo(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetManagementSystemIPInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetManagementSystemIPInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManagementSystemIPInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetManagementSystemIPInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagementSystemIPInfoRequest struct via the builder pattern


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


## GetRMACandidates

> []map[string]interface{} GetRMACandidates(ctx, deviceType).Uuid(uuid).Execute()





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
    deviceType := "deviceType_example" // string | Device Type
    uuid := "uuid_example" // string | Excluded currently selected uuid (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetRMACandidates(context.Background(), deviceType).Uuid(uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetRMACandidates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRMACandidates`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetRMACandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceType** | **string** | Device Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRMACandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uuid** | **string** | Excluded currently selected uuid | 

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


## GetRootCertStatusAll

> []map[string]interface{} GetRootCertStatusAll(ctx).State(state).Execute()





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
    state := "state_example" // string | Root certificate state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetRootCertStatusAll(context.Background()).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetRootCertStatusAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRootCertStatusAll`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetRootCertStatusAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRootCertStatusAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Root certificate state | 

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


## GetTenantManagementSystemIPs

> []map[string]interface{} GetTenantManagementSystemIPs(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.GetTenantManagementSystemIPs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.GetTenantManagementSystemIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantManagementSystemIPs`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.GetTenantManagementSystemIPs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantManagementSystemIPsRequest struct via the builder pattern


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


## InvalidateVmanageRootCA

> InvalidateVmanageRootCA(ctx, uuid).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.InvalidateVmanageRootCA(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.InvalidateVmanageRootCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateVmanageRootCARequest struct via the builder pattern


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


## MigrateDevice

> map[string]interface{} MigrateDevice(ctx, uuid).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.MigrateDevice(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.MigrateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrateDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.MigrateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateDeviceRequest struct via the builder pattern


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


## ResetVedgeCloud

> map[string]interface{} ResetVedgeCloud(ctx, uuid).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.ResetVedgeCloud(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.ResetVedgeCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetVedgeCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.ResetVedgeCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetVedgeCloudRequest struct via the builder pattern


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


## SetLifeCycle

> map[string]interface{} SetLifeCycle(ctx, uuid).Enable(enable).Execute()





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
    enable := true // bool | lifecycle needed flag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.SetLifeCycle(context.Background(), uuid).Enable(enable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.SetLifeCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLifeCycle`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.SetLifeCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLifeCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enable** | **bool** | lifecycle needed flag | 

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


## SyncDevices

> map[string]interface{} SyncDevices(ctx).SmartAccountModel(smartAccountModel).Execute()





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
    smartAccountModel := *openapiclient.NewSmartAccountModel() // SmartAccountModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.SyncDevices(context.Background()).SmartAccountModel(smartAccountModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.SyncDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.SyncDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartAccountModel** | [**SmartAccountModel**](SmartAccountModel.md) |  | 

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


## SyncRootCertChain

> SyncRootCertChain(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.SyncRootCertChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.SyncRootCertChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncRootCertChainRequest struct via the builder pattern


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


## UpdateDeviceSubjectSUDI

> UpdateDeviceSubjectSUDI(ctx, uuid).Execute()





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
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.UpdateDeviceSubjectSUDI(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.UpdateDeviceSubjectSUDI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Device uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSubjectSUDIRequest struct via the builder pattern


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


## ValidateUser

> map[string]interface{} ValidateUser(ctx).SmartAccountModel(smartAccountModel).Execute()





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
    smartAccountModel := *openapiclient.NewSmartAccountModel() // SmartAccountModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.ValidateUser(context.Background()).SmartAccountModel(smartAccountModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.ValidateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.ValidateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartAccountModel** | [**SmartAccountModel**](SmartAccountModel.md) |  | 

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


## ValidateUser1

> map[string]interface{} ValidateUser1(ctx).RequestBody(requestBody).Execute()





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
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationDeviceInventoryApi.ValidateUser1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationDeviceInventoryApi.ValidateUser1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateUser1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationDeviceInventoryApi.ValidateUser1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateUser1Request struct via the builder pattern


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

