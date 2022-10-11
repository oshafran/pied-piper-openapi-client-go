# \ConfigurationCloudOnRampApi

All URIs are relative to *http://VMANAGEIP*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcquireResourcePool**](ConfigurationCloudOnRampApi.md#AcquireResourcePool) | **Post** /template/cor/acquireResourcePool | 
[**AddDevicePair**](ConfigurationCloudOnRampApi.md#AddDevicePair) | **Post** /template/cor/devicepair | 
[**AddTransitVPC**](ConfigurationCloudOnRampApi.md#AddTransitVPC) | **Post** /template/cor/transitvpc | 
[**AuthenticateCloudOnRampCredAndAdd**](ConfigurationCloudOnRampApi.md#AuthenticateCloudOnRampCredAndAdd) | **Post** /template/cor/cloud/authenticate | 
[**AuthenticateCredAndUpdate**](ConfigurationCloudOnRampApi.md#AuthenticateCredAndUpdate) | **Put** /template/cor/cloud/authenticate | 
[**CreateAndMap**](ConfigurationCloudOnRampApi.md#CreateAndMap) | **Post** /template/cor | 
[**CreateResourcePool**](ConfigurationCloudOnRampApi.md#CreateResourcePool) | **Post** /template/cor/createResourcePool | 
[**GetAmiList**](ConfigurationCloudOnRampApi.md#GetAmiList) | **Get** /template/cor/ami | 
[**GetCORStatus**](ConfigurationCloudOnRampApi.md#GetCORStatus) | **Get** /template/cor | 
[**GetCloudAccounts**](ConfigurationCloudOnRampApi.md#GetCloudAccounts) | **Get** /template/cor/cloud/account | 
[**GetCloudHostVPCs**](ConfigurationCloudOnRampApi.md#GetCloudHostVPCs) | **Get** /template/cor/hostvpc | 
[**GetCloudHostVpcAccountDetails**](ConfigurationCloudOnRampApi.md#GetCloudHostVpcAccountDetails) | **Get** /template/cor/cloud/host/accountdetails | 
[**GetCloudList**](ConfigurationCloudOnRampApi.md#GetCloudList) | **Get** /template/cor/cloud | 
[**GetCloudMappedHostAccounts**](ConfigurationCloudOnRampApi.md#GetCloudMappedHostAccounts) | **Get** /template/cor/cloud/mappedhostaccounts | 
[**GetCloudOnRampDevices**](ConfigurationCloudOnRampApi.md#GetCloudOnRampDevices) | **Get** /template/cor/device | 
[**GetExternalId**](ConfigurationCloudOnRampApi.md#GetExternalId) | **Get** /template/cor/externalId | 
[**GetHostVPCs**](ConfigurationCloudOnRampApi.md#GetHostVPCs) | **Get** /template/cor/devicepair/hostvpc | 
[**GetMappedVPCs**](ConfigurationCloudOnRampApi.md#GetMappedVPCs) | **Get** /template/cor/map | 
[**GetPemKeyList**](ConfigurationCloudOnRampApi.md#GetPemKeyList) | **Get** /template/cor/pem | 
[**GetTenantAndHostVpcList**](ConfigurationCloudOnRampApi.md#GetTenantAndHostVpcList) | **Get** /template/cor/hostvpclist | 
[**GetTransitDevicePairAndHostList**](ConfigurationCloudOnRampApi.md#GetTransitDevicePairAndHostList) | **Get** /template/cor/getTransitDevicePairAndHostList | 
[**GetTransitVPCSupportedSize**](ConfigurationCloudOnRampApi.md#GetTransitVPCSupportedSize) | **Get** /template/cor/transitvpc/size | 
[**GetTransitVPCs**](ConfigurationCloudOnRampApi.md#GetTransitVPCs) | **Get** /template/cor/transitvpc | 
[**GetTransitVpcVpnList**](ConfigurationCloudOnRampApi.md#GetTransitVpcVpnList) | **Get** /template/cor/getTransitVpnList | 
[**MapVPCs**](ConfigurationCloudOnRampApi.md#MapVPCs) | **Post** /template/cor/map | 
[**RaiseAlarmForAccount**](ConfigurationCloudOnRampApi.md#RaiseAlarmForAccount) | **Post** /template/cor/account/alarm | 
[**RemoveDeviceId**](ConfigurationCloudOnRampApi.md#RemoveDeviceId) | **Delete** /template/cor/deleteDevicepair | 
[**RemoveTransitVPC**](ConfigurationCloudOnRampApi.md#RemoveTransitVPC) | **Delete** /template/cor/accountid/{accountid} | 
[**ScaleDown**](ConfigurationCloudOnRampApi.md#ScaleDown) | **Post** /template/cor/scale/down | 
[**ScaleUp**](ConfigurationCloudOnRampApi.md#ScaleUp) | **Post** /template/cor/scale/up | 
[**UnmapVPCs**](ConfigurationCloudOnRampApi.md#UnmapVPCs) | **Delete** /template/cor/map | 
[**UpdateHostVpcReachability**](ConfigurationCloudOnRampApi.md#UpdateHostVpcReachability) | **Put** /template/cor/hostvpclist | 
[**UpdateTransitVPC**](ConfigurationCloudOnRampApi.md#UpdateTransitVPC) | **Put** /template/cor/transitvpc | 
[**UpdateTransitVpcAutoscaleProperties**](ConfigurationCloudOnRampApi.md#UpdateTransitVpcAutoscaleProperties) | **Put** /template/cor/transitvpc/autoscale-properties | 
[**Updatestatus**](ConfigurationCloudOnRampApi.md#Updatestatus) | **Post** /template/cor/status | 



## AcquireResourcePool

> AcquireResourcePool(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Add IP from resource pool request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.AcquireResourcePool(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.AcquireResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcquireResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Add IP from resource pool request | 

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


## AddDevicePair

> map[string]interface{} AddDevicePair(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Add device pair request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.AddDevicePair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.AddDevicePair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDevicePair`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.AddDevicePair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicePairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Add device pair request | 

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


## AddTransitVPC

> map[string]interface{} AddTransitVPC(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | VPC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.AddTransitVPC(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.AddTransitVPC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTransitVPC`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.AddTransitVPC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTransitVPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | VPC | 

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


## AuthenticateCloudOnRampCredAndAdd

> map[string]interface{} AuthenticateCloudOnRampCredAndAdd(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Cloud account credential (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.AuthenticateCloudOnRampCredAndAdd(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.AuthenticateCloudOnRampCredAndAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateCloudOnRampCredAndAdd`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.AuthenticateCloudOnRampCredAndAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateCloudOnRampCredAndAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Cloud account credential | 

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


## AuthenticateCredAndUpdate

> map[string]interface{} AuthenticateCredAndUpdate(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Cloud account credential (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.AuthenticateCredAndUpdate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.AuthenticateCredAndUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateCredAndUpdate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.AuthenticateCredAndUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateCredAndUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Cloud account credential | 

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


## CreateAndMap

> map[string]interface{} CreateAndMap(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Map host to transit VPC request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.CreateAndMap(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.CreateAndMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAndMap`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.CreateAndMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Map host to transit VPC request | 

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


## CreateResourcePool

> CreateResourcePool(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Add resource pool request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.CreateResourcePool(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.CreateResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Add resource pool request | 

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


## GetAmiList

> []map[string]interface{} GetAmiList(ctx).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    cloudregion := "cloudregion_example" // string | Cloud region
    cloudtype := "cloudtype_example" // string | Cloud type (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetAmiList(context.Background()).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetAmiList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmiList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetAmiList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAmiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **cloudregion** | **string** | Cloud region | 
 **cloudtype** | **string** | Cloud type | [default to &quot;AWS&quot;]

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


## GetCORStatus

> []map[string]interface{} GetCORStatus(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCORStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCORStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCORStatus`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCORStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCORStatusRequest struct via the builder pattern


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


## GetCloudAccounts

> map[string]interface{} GetCloudAccounts(ctx).Cloudtype(cloudtype).CloudEnvironment(cloudEnvironment).Execute()





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
    cloudtype := "cloudtype_example" // string | Cloud type
    cloudEnvironment := "cloudEnvironment_example" // string | Cloud environment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCloudAccounts(context.Background()).Cloudtype(cloudtype).CloudEnvironment(cloudEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCloudAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCloudAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudtype** | **string** | Cloud type | 
 **cloudEnvironment** | **string** | Cloud environment | 

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


## GetCloudHostVPCs

> []map[string]interface{} GetCloudHostVPCs(ctx).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    cloudregion := "cloudregion_example" // string | Cloud region
    cloudtype := "cloudtype_example" // string | Cloud type (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCloudHostVPCs(context.Background()).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCloudHostVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudHostVPCs`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCloudHostVPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudHostVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **cloudregion** | **string** | Cloud region | 
 **cloudtype** | **string** | Cloud type | [default to &quot;AWS&quot;]

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


## GetCloudHostVpcAccountDetails

> map[string]interface{} GetCloudHostVpcAccountDetails(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCloudHostVpcAccountDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCloudHostVpcAccountDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudHostVpcAccountDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCloudHostVpcAccountDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudHostVpcAccountDetailsRequest struct via the builder pattern


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


## GetCloudList

> []map[string]interface{} GetCloudList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCloudList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCloudList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCloudList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudListRequest struct via the builder pattern


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


## GetCloudMappedHostAccounts

> map[string]interface{} GetCloudMappedHostAccounts(ctx).Accountid(accountid).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    cloudtype := "cloudtype_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCloudMappedHostAccounts(context.Background()).Accountid(accountid).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCloudMappedHostAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudMappedHostAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCloudMappedHostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudMappedHostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **cloudtype** | **string** | Cloud type | 

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


## GetCloudOnRampDevices

> []map[string]interface{} GetCloudOnRampDevices(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetCloudOnRampDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetCloudOnRampDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudOnRampDevices`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetCloudOnRampDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudOnRampDevicesRequest struct via the builder pattern


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


## GetExternalId

> []map[string]interface{} GetExternalId(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetExternalId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetExternalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalId`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetExternalId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalIdRequest struct via the builder pattern


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


## GetHostVPCs

> map[string]interface{} GetHostVPCs(ctx).TransitVpcId(transitVpcId).DevicePairId(devicePairId).Execute()





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
    transitVpcId := "transitVpcId_example" // string | Transit VPC Id
    devicePairId := "devicePairId_example" // string | Device pair Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetHostVPCs(context.Background()).TransitVpcId(transitVpcId).DevicePairId(devicePairId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetHostVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostVPCs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetHostVPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitVpcId** | **string** | Transit VPC Id | 
 **devicePairId** | **string** | Device pair Id | 

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


## GetMappedVPCs

> map[string]interface{} GetMappedVPCs(ctx).Accountid(accountid).Cloudregion(cloudregion).Execute()





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
    accountid := "accountid_example" // string | Account Id
    cloudregion := "cloudregion_example" // string | Cloud region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetMappedVPCs(context.Background()).Accountid(accountid).Cloudregion(cloudregion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetMappedVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappedVPCs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetMappedVPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMappedVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **cloudregion** | **string** | Cloud region | 

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


## GetPemKeyList

> []map[string]interface{} GetPemKeyList(ctx).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    cloudregion := "cloudregion_example" // string | Cloud region
    cloudtype := "cloudtype_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetPemKeyList(context.Background()).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetPemKeyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPemKeyList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetPemKeyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPemKeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **cloudregion** | **string** | Cloud region | 
 **cloudtype** | **string** | Cloud type | 

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


## GetTenantAndHostVpcList

> []map[string]interface{} GetTenantAndHostVpcList(ctx).Intent(intent).Execute()





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
    intent := "intent_example" // string | Intent (optional) (default to "undefined")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetTenantAndHostVpcList(context.Background()).Intent(intent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetTenantAndHostVpcList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantAndHostVpcList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetTenantAndHostVpcList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantAndHostVpcListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **intent** | **string** | Intent | [default to &quot;undefined&quot;]

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


## GetTransitDevicePairAndHostList

> map[string]interface{} GetTransitDevicePairAndHostList(ctx).AccountId(accountId).CloudRegion(cloudRegion).Execute()





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
    accountId := "accountId_example" // string | Account Id
    cloudRegion := "cloudRegion_example" // string | Cloud region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetTransitDevicePairAndHostList(context.Background()).AccountId(accountId).CloudRegion(cloudRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetTransitDevicePairAndHostList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitDevicePairAndHostList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetTransitDevicePairAndHostList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitDevicePairAndHostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account Id | 
 **cloudRegion** | **string** | Cloud region | 

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


## GetTransitVPCSupportedSize

> []map[string]interface{} GetTransitVPCSupportedSize(ctx).CloudEnvironment(cloudEnvironment).Cloudtype(cloudtype).Execute()





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
    cloudEnvironment := "cloudEnvironment_example" // string | Cloud environment
    cloudtype := "cloudtype_example" // string | Cloud type (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetTransitVPCSupportedSize(context.Background()).CloudEnvironment(cloudEnvironment).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetTransitVPCSupportedSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitVPCSupportedSize`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetTransitVPCSupportedSize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitVPCSupportedSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEnvironment** | **string** | Cloud environment | 
 **cloudtype** | **string** | Cloud type | [default to &quot;AWS&quot;]

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


## GetTransitVPCs

> []map[string]interface{} GetTransitVPCs(ctx).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    cloudregion := "cloudregion_example" // string | Cloud region
    cloudtype := "cloudtype_example" // string | Cloud type (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetTransitVPCs(context.Background()).Accountid(accountid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetTransitVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitVPCs`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetTransitVPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **cloudregion** | **string** | Cloud region | 
 **cloudtype** | **string** | Cloud type | [default to &quot;AWS&quot;]

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


## GetTransitVpcVpnList

> []map[string]interface{} GetTransitVpcVpnList(ctx).AccountId(accountId).Execute()





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
    accountId := "accountId_example" // string | Account Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.GetTransitVpcVpnList(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.GetTransitVpcVpnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitVpcVpnList`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.GetTransitVpcVpnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitVpcVpnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account Id | 

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


## MapVPCs

> map[string]interface{} MapVPCs(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Map host to VPC/VNet (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.MapVPCs(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.MapVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MapVPCs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.MapVPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMapVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Map host to VPC/VNet | 

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


## RaiseAlarmForAccount

> RaiseAlarmForAccount(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Account object (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.RaiseAlarmForAccount(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.RaiseAlarmForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRaiseAlarmForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Account object | 

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


## RemoveDeviceId

> map[string]interface{} RemoveDeviceId(ctx).Accountid(accountid).Transitvpcid(transitvpcid).Transitvpcname(transitvpcname).Cloudregion(cloudregion).DevicePairId(devicePairId).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    transitvpcid := "transitvpcid_example" // string | VPC Id
    transitvpcname := "transitvpcname_example" // string | VPC Name
    cloudregion := "cloudregion_example" // string | Cloud region
    devicePairId := "devicePairId_example" // string | Device pair Id
    cloudtype := "cloudtype_example" // string | Cloud type (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.RemoveDeviceId(context.Background()).Accountid(accountid).Transitvpcid(transitvpcid).Transitvpcname(transitvpcname).Cloudregion(cloudregion).DevicePairId(devicePairId).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.RemoveDeviceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveDeviceId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.RemoveDeviceId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDeviceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountid** | **string** | Account Id | 
 **transitvpcid** | **string** | VPC Id | 
 **transitvpcname** | **string** | VPC Name | 
 **cloudregion** | **string** | Cloud region | 
 **devicePairId** | **string** | Device pair Id | 
 **cloudtype** | **string** | Cloud type | [default to &quot;AWS&quot;]

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


## RemoveTransitVPC

> map[string]interface{} RemoveTransitVPC(ctx, accountid).Transitvpcid(transitvpcid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()





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
    accountid := "accountid_example" // string | Account Id
    transitvpcid := "transitvpcid_example" // string | Cloud VPC Id
    cloudregion := "cloudregion_example" // string | Cloud region
    cloudtype := "cloudtype_example" // string | Cloud type (optional) (default to "AWS")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.RemoveTransitVPC(context.Background(), accountid).Transitvpcid(transitvpcid).Cloudregion(cloudregion).Cloudtype(cloudtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.RemoveTransitVPC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTransitVPC`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.RemoveTransitVPC`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountid** | **string** | Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTransitVPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitvpcid** | **string** | Cloud VPC Id | 
 **cloudregion** | **string** | Cloud region | 
 **cloudtype** | **string** | Cloud type | [default to &quot;AWS&quot;]

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


## ScaleDown

> ScaleDown(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Update VPC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.ScaleDown(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.ScaleDown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScaleDownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Update VPC | 

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


## ScaleUp

> ScaleUp(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Update VPC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.ScaleUp(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.ScaleUp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScaleUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Update VPC | 

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


## UnmapVPCs

> map[string]interface{} UnmapVPCs(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Unmap host to VPC/VNet (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.UnmapVPCs(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.UnmapVPCs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmapVPCs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.UnmapVPCs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmapVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Unmap host to VPC/VNet | 

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


## UpdateHostVpcReachability

> UpdateHostVpcReachability(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Update VPC status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.UpdateHostVpcReachability(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.UpdateHostVpcReachability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostVpcReachabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Update VPC status | 

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


## UpdateTransitVPC

> map[string]interface{} UpdateTransitVPC(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | VPC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.UpdateTransitVPC(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.UpdateTransitVPC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransitVPC`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.UpdateTransitVPC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransitVPCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | VPC | 

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


## UpdateTransitVpcAutoscaleProperties

> map[string]interface{} UpdateTransitVpcAutoscaleProperties(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | VPC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.UpdateTransitVpcAutoscaleProperties(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.UpdateTransitVpcAutoscaleProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransitVpcAutoscaleProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationCloudOnRampApi.UpdateTransitVpcAutoscaleProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransitVpcAutoscalePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | VPC | 

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


## Updatestatus

> Updatestatus(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Status object (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationCloudOnRampApi.Updatestatus(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationCloudOnRampApi.Updatestatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatestatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Status object | 

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

