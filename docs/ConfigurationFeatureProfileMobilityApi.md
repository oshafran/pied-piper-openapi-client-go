# \ConfigurationFeatureProfileMobilityApi

All URIs are relative to *https://2.2.2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBasicProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateBasicProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/basic | 
[**CreateCellularProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateCellularProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/cellular | 
[**CreateEthernetProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateEthernetProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/ethernet | 
[**CreateMobilityConfigProfileParcelForCli**](ConfigurationFeatureProfileMobilityApi.md#CreateMobilityConfigProfileParcelForCli) | **Post** /v1/feature-profile/mobility/cli/{cliId}/config | 
[**CreateNetworkProtocolProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateNetworkProtocolProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/networkProtocol | 
[**CreateSecurityPolicyProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateSecurityPolicyProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/securityPolicy | 
[**CreateVpnProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateVpnProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/vpn | 
[**CreateWifiProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#CreateWifiProfileParcelForMobility) | **Post** /v1/feature-profile/mobility/global/{profileId}/wifi | 
[**DeleteACellularProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#DeleteACellularProfileParcelForMobility) | **Delete** /v1/feature-profile/mobility/global/{profileId}/cellular/{cellularId} | 
[**DeleteAVpnProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#DeleteAVpnProfileParcelForMobility) | **Delete** /v1/feature-profile/mobility/global/{profileId}/vpn/{vpnId} | 
[**DeleteBasicProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#DeleteBasicProfileParcelForMobility) | **Delete** /v1/feature-profile/mobility/global/{profileId}/basic/{parcelId} | 
[**DeleteEthernetProfileParcelForSystem**](ConfigurationFeatureProfileMobilityApi.md#DeleteEthernetProfileParcelForSystem) | **Delete** /v1/feature-profile/mobility/global/{profileId}/ethernet/{ethernetId} | 
[**DeleteMobilityConfigProfileParcelForCLI**](ConfigurationFeatureProfileMobilityApi.md#DeleteMobilityConfigProfileParcelForCLI) | **Delete** /v1/feature-profile/mobility/cli/{cliId}/config/{configId} | 
[**DeleteNetworkProtocolProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#DeleteNetworkProtocolProfileParcelForMobility) | **Delete** /v1/feature-profile/mobility/global/{profileId}/networkProtocol/{networkProtocolId} | 
[**DeleteSecurityPolicyProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#DeleteSecurityPolicyProfileParcelForMobility) | **Delete** /v1/feature-profile/mobility/global/{profileId}/securityPolicy/{securityPolicyId} | 
[**DeleteWifiProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#DeleteWifiProfileParcelForMobility) | **Delete** /v1/feature-profile/mobility/global/{profileId}/wifi/{wifiId} | 
[**EditBasicProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#EditBasicProfileParcelForMobility) | **Put** /v1/feature-profile/mobility/global/{profileId}/basic/{parcelId} | 
[**EditCellularProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#EditCellularProfileParcelForMobility) | **Put** /v1/feature-profile/mobility/global/{profileId}/cellular/{cellularId} | 
[**EditEthernetProfileParcelForSystem**](ConfigurationFeatureProfileMobilityApi.md#EditEthernetProfileParcelForSystem) | **Put** /v1/feature-profile/mobility/global/{profileId}/ethernet/{ethernetId} | 
[**EditMobilityConfigProfileParcelForCLI**](ConfigurationFeatureProfileMobilityApi.md#EditMobilityConfigProfileParcelForCLI) | **Put** /v1/feature-profile/mobility/cli/{cliId}/config/{configId} | 
[**EditNetworkProtocolProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#EditNetworkProtocolProfileParcelForMobility) | **Put** /v1/feature-profile/mobility/global/{profileId}/networkProtocol/{networkProtocolId} | 
[**EditSecurityPolicyProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#EditSecurityPolicyProfileParcelForMobility) | **Put** /v1/feature-profile/mobility/global/{profileId}/securityPolicy/{securityPolicyId} | 
[**EditVpnProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#EditVpnProfileParcelForMobility) | **Put** /v1/feature-profile/mobility/global/{profileId}/vpn/{vpnId} | 
[**EditWifiProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#EditWifiProfileParcelForMobility) | **Put** /v1/feature-profile/mobility/global/{profileId}/wifi/{wifiId} | 
[**GetBasicProfileParcelByParcelIdForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetBasicProfileParcelByParcelIdForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/basic/{parcelId} | 
[**GetBasicProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetBasicProfileParcelForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/basic | 
[**GetCellularProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetCellularProfileParcelForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/cellular/{cellularId} | 
[**GetCellularProfileParcelListForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetCellularProfileParcelListForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/cellular | 
[**GetEthernetProfileParcel**](ConfigurationFeatureProfileMobilityApi.md#GetEthernetProfileParcel) | **Get** /v1/feature-profile/mobility/global/{profileId}/ethernet/{ethernetId} | 
[**GetEthernetProfileParcels**](ConfigurationFeatureProfileMobilityApi.md#GetEthernetProfileParcels) | **Get** /v1/feature-profile/mobility/global/{profileId}/ethernet | 
[**GetMobilityConfigProfileParcelByParcelIdForCLI**](ConfigurationFeatureProfileMobilityApi.md#GetMobilityConfigProfileParcelByParcelIdForCLI) | **Get** /v1/feature-profile/mobility/cli/{cliId}/config/{configId} | 
[**GetMobilityConfigProfileParcelForCLI**](ConfigurationFeatureProfileMobilityApi.md#GetMobilityConfigProfileParcelForCLI) | **Get** /v1/feature-profile/mobility/cli/{cliId}/config | 
[**GetMobilityFeatureProfileByGlobalId**](ConfigurationFeatureProfileMobilityApi.md#GetMobilityFeatureProfileByGlobalId) | **Get** /v1/feature-profile/mobility/global/{profileId} | 
[**GetMobilityGlobalBasicParcelSchemaBySchemaType**](ConfigurationFeatureProfileMobilityApi.md#GetMobilityGlobalBasicParcelSchemaBySchemaType) | **Get** /v1/feature-profile/mobility/global/basic/schema | 
[**GetNetworkProtocolProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetNetworkProtocolProfileParcelForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/networkProtocol/{networkProtocolId} | 
[**GetNetworkProtocolProfileParcelListForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetNetworkProtocolProfileParcelListForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/networkProtocol | 
[**GetSecurityPolicyProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetSecurityPolicyProfileParcelForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/securityPolicy/{securityPolicyId} | 
[**GetSecurityPolicyProfileParcelListForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetSecurityPolicyProfileParcelListForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/securityPolicy | 
[**GetVpnProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetVpnProfileParcelForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/vpn/{vpnId} | 
[**GetVpnProfileParcelListForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetVpnProfileParcelListForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/vpn | 
[**GetWifiProfileParcelForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetWifiProfileParcelForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/wifi/{wifiId} | 
[**GetWifiProfileParcelListForMobility**](ConfigurationFeatureProfileMobilityApi.md#GetWifiProfileParcelListForMobility) | **Get** /v1/feature-profile/mobility/global/{profileId}/wifi | 



## CreateBasicProfileParcelForMobility

> string CreateBasicProfileParcelForMobility(ctx, profileId).Body(body).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    body := "TODO" // string | Basic Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateBasicProfileParcelForMobility(context.Background(), profileId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateBasicProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBasicProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateBasicProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBasicProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Basic Profile Parcel | 

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


## CreateCellularProfileParcelForMobility

> string CreateCellularProfileParcelForMobility(ctx, profileId).CellularProfile(cellularProfile).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    cellularProfile := *openapiclient.NewCellularProfile() // CellularProfile | Cellular Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateCellularProfileParcelForMobility(context.Background(), profileId).CellularProfile(cellularProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateCellularProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCellularProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateCellularProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCellularProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cellularProfile** | [**CellularProfile**](CellularProfile.md) | Cellular Profile Parcel | 

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


## CreateEthernetProfileParcelForMobility

> string CreateEthernetProfileParcelForMobility(ctx, profileId).Ethernet(ethernet).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    ethernet := *openapiclient.NewEthernet("Wifi profile Parcel", "wifi") // Ethernet | Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateEthernetProfileParcelForMobility(context.Background(), profileId).Ethernet(ethernet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateEthernetProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEthernetProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateEthernetProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEthernetProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ethernet** | [**Ethernet**](Ethernet.md) | Ethernet Profile Parcel | 

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


## CreateMobilityConfigProfileParcelForCli

> string CreateMobilityConfigProfileParcelForCli(ctx, cliId).Body(body).Execute()





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
    cliId := "cliId_example" // string | Feature Profile ID
    body := "refer to schema for documentation and example" // string | cli config Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateMobilityConfigProfileParcelForCli(context.Background(), cliId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateMobilityConfigProfileParcelForCli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobilityConfigProfileParcelForCli`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateMobilityConfigProfileParcelForCli`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobilityConfigProfileParcelForCliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | cli config Profile Parcel | 

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


## CreateNetworkProtocolProfileParcelForMobility

> string CreateNetworkProtocolProfileParcelForMobility(ctx, profileId).NetworkProtocol(networkProtocol).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    networkProtocol := *openapiclient.NewNetworkProtocol("Wifi profile Parcel", "wifi") // NetworkProtocol | NetworkProtocol Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateNetworkProtocolProfileParcelForMobility(context.Background(), profileId).NetworkProtocol(networkProtocol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateNetworkProtocolProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkProtocolProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateNetworkProtocolProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProtocolProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkProtocol** | [**NetworkProtocol**](NetworkProtocol.md) | NetworkProtocol Profile Parcel | 

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


## CreateSecurityPolicyProfileParcelForMobility

> string CreateSecurityPolicyProfileParcelForMobility(ctx, profileId).SecurityPolicy(securityPolicy).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    securityPolicy := *openapiclient.NewSecurityPolicy("Wifi profile Parcel", "wifi") // SecurityPolicy | SecurityPolicy Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateSecurityPolicyProfileParcelForMobility(context.Background(), profileId).SecurityPolicy(securityPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateSecurityPolicyProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityPolicyProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateSecurityPolicyProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityPolicyProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityPolicy** | [**SecurityPolicy**](SecurityPolicy.md) | SecurityPolicy Profile Parcel | 

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


## CreateVpnProfileParcelForMobility

> string CreateVpnProfileParcelForMobility(ctx, profileId).Vpn(vpn).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    vpn := *openapiclient.NewVpn(*openapiclient.NewSiteToSiteVpn("LocalInterface_example", "LocalPrivateSubnet_example", "PreSharedSecret_example", "RemotePrivateSubnets_example", "RemotePublicIp_example"), "Wifi profile Parcel", "wifi") // Vpn | Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateVpnProfileParcelForMobility(context.Background(), profileId).Vpn(vpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateVpnProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVpnProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateVpnProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpn** | [**Vpn**](Vpn.md) | Vpn Profile Parcel | 

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


## CreateWifiProfileParcelForMobility

> string CreateWifiProfileParcelForMobility(ctx, profileId).Wifi(wifi).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    wifi := *openapiclient.NewWifi("Wifi profile Parcel", "wifi") // Wifi | Wifi Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.CreateWifiProfileParcelForMobility(context.Background(), profileId).Wifi(wifi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.CreateWifiProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWifiProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.CreateWifiProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWifiProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wifi** | [**Wifi**](Wifi.md) | Wifi Profile Parcel | 

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


## DeleteACellularProfileParcelForMobility

> DeleteACellularProfileParcelForMobility(ctx, profileId, cellularId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    cellularId := "cellularId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteACellularProfileParcelForMobility(context.Background(), profileId, cellularId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteACellularProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**cellularId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACellularProfileParcelForMobilityRequest struct via the builder pattern


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


## DeleteAVpnProfileParcelForMobility

> DeleteAVpnProfileParcelForMobility(ctx, profileId, vpnId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteAVpnProfileParcelForMobility(context.Background(), profileId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteAVpnProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAVpnProfileParcelForMobilityRequest struct via the builder pattern


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


## DeleteBasicProfileParcelForMobility

> DeleteBasicProfileParcelForMobility(ctx, profileId, parcelId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    parcelId := "parcelId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteBasicProfileParcelForMobility(context.Background(), profileId, parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteBasicProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**parcelId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBasicProfileParcelForMobilityRequest struct via the builder pattern


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


## DeleteEthernetProfileParcelForSystem

> DeleteEthernetProfileParcelForSystem(ctx, profileId, ethernetId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    ethernetId := "ethernetId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteEthernetProfileParcelForSystem(context.Background(), profileId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteEthernetProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**ethernetId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEthernetProfileParcelForSystemRequest struct via the builder pattern


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


## DeleteMobilityConfigProfileParcelForCLI

> DeleteMobilityConfigProfileParcelForCLI(ctx, cliId, configId).Execute()





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
    cliId := "cliId_example" // string | Feature Profile ID
    configId := "configId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteMobilityConfigProfileParcelForCLI(context.Background(), cliId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteMobilityConfigProfileParcelForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 
**configId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobilityConfigProfileParcelForCLIRequest struct via the builder pattern


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


## DeleteNetworkProtocolProfileParcelForMobility

> DeleteNetworkProtocolProfileParcelForMobility(ctx, profileId, networkProtocolId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    networkProtocolId := "networkProtocolId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteNetworkProtocolProfileParcelForMobility(context.Background(), profileId, networkProtocolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteNetworkProtocolProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**networkProtocolId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProtocolProfileParcelForMobilityRequest struct via the builder pattern


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


## DeleteSecurityPolicyProfileParcelForMobility

> DeleteSecurityPolicyProfileParcelForMobility(ctx, profileId, securityPolicyId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    securityPolicyId := "securityPolicyId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteSecurityPolicyProfileParcelForMobility(context.Background(), profileId, securityPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteSecurityPolicyProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**securityPolicyId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityPolicyProfileParcelForMobilityRequest struct via the builder pattern


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


## DeleteWifiProfileParcelForMobility

> DeleteWifiProfileParcelForMobility(ctx, profileId, wifiId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    wifiId := "wifiId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.DeleteWifiProfileParcelForMobility(context.Background(), profileId, wifiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.DeleteWifiProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**wifiId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWifiProfileParcelForMobilityRequest struct via the builder pattern


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


## EditBasicProfileParcelForMobility

> string EditBasicProfileParcelForMobility(ctx, profileId, parcelId).Body(body).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    parcelId := "parcelId_example" // string | Profile Parcel ID
    body := "TODO" // string | Basic Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditBasicProfileParcelForMobility(context.Background(), profileId, parcelId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditBasicProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBasicProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.EditBasicProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**parcelId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBasicProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Basic Profile Parcel | 

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


## EditCellularProfileParcelForMobility

> EditCellularProfileParcelForMobility(ctx, profileId, cellularId).Cellular(cellular).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    cellularId := "cellularId_example" // string | Profile Parcel ID
    cellular := *openapiclient.NewCellular(*openapiclient.NewSimSlotConfig(), "Wifi profile Parcel", "wifi") // Cellular | Cellular Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditCellularProfileParcelForMobility(context.Background(), profileId, cellularId).Cellular(cellular).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditCellularProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**cellularId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCellularProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cellular** | [**Cellular**](Cellular.md) | Cellular Profile Parcel | 

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


## EditEthernetProfileParcelForSystem

> EditEthernetProfileParcelForSystem(ctx, profileId, ethernetId).Body(body).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    ethernetId := "ethernetId_example" // string | Profile Parcel ID
    body := "{"type":"ethernet","ethernetInterfaceList":[{"interfaceName":"GigabitEthernet0/0","wanConfiguration":"Active","portType":"WAN","ipAssignment":"static","staticIpAddress":"1.1.1.2","staticIpAddressSubnetMask":"255.255.0.0","staticRouteIp":"3.3.3.3"},{"interfaceName":"GigabitEthernet0/1","adminState":"enabled"},{"interfaceName":"GigabitEthernet0/2","adminState":"disabled"}]}" // string | Ethernet Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditEthernetProfileParcelForSystem(context.Background(), profileId, ethernetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditEthernetProfileParcelForSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**ethernetId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEthernetProfileParcelForSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | Ethernet Profile Parcel | 

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


## EditMobilityConfigProfileParcelForCLI

> string EditMobilityConfigProfileParcelForCLI(ctx, cliId, configId).Body(body).Execute()





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
    cliId := "cliId_example" // string | Feature Profile ID
    configId := "configId_example" // string | Profile Parcel ID
    body := "refer to schema for documentation and example" // string | cli config Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditMobilityConfigProfileParcelForCLI(context.Background(), cliId, configId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditMobilityConfigProfileParcelForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMobilityConfigProfileParcelForCLI`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.EditMobilityConfigProfileParcelForCLI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 
**configId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMobilityConfigProfileParcelForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | cli config Profile Parcel | 

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


## EditNetworkProtocolProfileParcelForMobility

> EditNetworkProtocolProfileParcelForMobility(ctx, profileId, networkProtocolId).NetworkProtocol(networkProtocol).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    networkProtocolId := "networkProtocolId_example" // string | Profile Parcel ID
    networkProtocol := *openapiclient.NewNetworkProtocol("Wifi profile Parcel", "wifi") // NetworkProtocol | Network Protocol Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditNetworkProtocolProfileParcelForMobility(context.Background(), profileId, networkProtocolId).NetworkProtocol(networkProtocol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditNetworkProtocolProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**networkProtocolId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditNetworkProtocolProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkProtocol** | [**NetworkProtocol**](NetworkProtocol.md) | Network Protocol Profile Parcel | 

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


## EditSecurityPolicyProfileParcelForMobility

> EditSecurityPolicyProfileParcelForMobility(ctx, profileId, securityPolicyId).SecurityPolicy(securityPolicy).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    securityPolicyId := "securityPolicyId_example" // string | Profile Parcel ID
    securityPolicy := *openapiclient.NewSecurityPolicy("Wifi profile Parcel", "wifi") // SecurityPolicy | Security Policy Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditSecurityPolicyProfileParcelForMobility(context.Background(), profileId, securityPolicyId).SecurityPolicy(securityPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditSecurityPolicyProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**securityPolicyId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSecurityPolicyProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **securityPolicy** | [**SecurityPolicy**](SecurityPolicy.md) | Security Policy Profile Parcel | 

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


## EditVpnProfileParcelForMobility

> EditVpnProfileParcelForMobility(ctx, profileId, vpnId).Vpn(vpn).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID
    vpn := *openapiclient.NewVpn(*openapiclient.NewSiteToSiteVpn("LocalInterface_example", "LocalPrivateSubnet_example", "PreSharedSecret_example", "RemotePrivateSubnets_example", "RemotePublicIp_example"), "Wifi profile Parcel", "wifi") // Vpn | Vpn Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditVpnProfileParcelForMobility(context.Background(), profileId, vpnId).Vpn(vpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditVpnProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVpnProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpn** | [**Vpn**](Vpn.md) | Vpn Profile Parcel | 

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


## EditWifiProfileParcelForMobility

> EditWifiProfileParcelForMobility(ctx, profileId, wifiId).Wifi(wifi).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    wifiId := "wifiId_example" // string | Profile Parcel ID
    wifi := *openapiclient.NewWifi("Wifi profile Parcel", "wifi") // Wifi | Wifi Profile Parcel (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.EditWifiProfileParcelForMobility(context.Background(), profileId, wifiId).Wifi(wifi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.EditWifiProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**wifiId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWifiProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wifi** | [**Wifi**](Wifi.md) | Wifi Profile Parcel | 

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


## GetBasicProfileParcelByParcelIdForMobility

> string GetBasicProfileParcelByParcelIdForMobility(ctx, profileId, parcelId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    parcelId := "parcelId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetBasicProfileParcelByParcelIdForMobility(context.Background(), profileId, parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetBasicProfileParcelByParcelIdForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicProfileParcelByParcelIdForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetBasicProfileParcelByParcelIdForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**parcelId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicProfileParcelByParcelIdForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetBasicProfileParcelForMobility

> string GetBasicProfileParcelForMobility(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetBasicProfileParcelForMobility(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetBasicProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetBasicProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCellularProfileParcelForMobility

> string GetCellularProfileParcelForMobility(ctx, profileId, cellularId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    cellularId := "cellularId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetCellularProfileParcelForMobility(context.Background(), profileId, cellularId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetCellularProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetCellularProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**cellularId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetCellularProfileParcelListForMobility

> string GetCellularProfileParcelListForMobility(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetCellularProfileParcelListForMobility(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetCellularProfileParcelListForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellularProfileParcelListForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetCellularProfileParcelListForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCellularProfileParcelListForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetEthernetProfileParcel

> string GetEthernetProfileParcel(ctx, profileId, ethernetId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    ethernetId := "ethernetId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetEthernetProfileParcel(context.Background(), profileId, ethernetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetEthernetProfileParcel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEthernetProfileParcel`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetEthernetProfileParcel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**ethernetId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEthernetProfileParcelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetEthernetProfileParcels

> string GetEthernetProfileParcels(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetEthernetProfileParcels(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetEthernetProfileParcels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEthernetProfileParcels`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetEthernetProfileParcels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEthernetProfileParcelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetMobilityConfigProfileParcelByParcelIdForCLI

> string GetMobilityConfigProfileParcelByParcelIdForCLI(ctx, cliId, configId).Execute()





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
    cliId := "cliId_example" // string | Feature Profile ID
    configId := "configId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetMobilityConfigProfileParcelByParcelIdForCLI(context.Background(), cliId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetMobilityConfigProfileParcelByParcelIdForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobilityConfigProfileParcelByParcelIdForCLI`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetMobilityConfigProfileParcelByParcelIdForCLI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 
**configId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobilityConfigProfileParcelByParcelIdForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetMobilityConfigProfileParcelForCLI

> string GetMobilityConfigProfileParcelForCLI(ctx, cliId).Execute()





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
    cliId := "cliId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetMobilityConfigProfileParcelForCLI(context.Background(), cliId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetMobilityConfigProfileParcelForCLI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobilityConfigProfileParcelForCLI`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetMobilityConfigProfileParcelForCLI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobilityConfigProfileParcelForCLIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetMobilityFeatureProfileByGlobalId

> map[string]interface{} GetMobilityFeatureProfileByGlobalId(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetMobilityFeatureProfileByGlobalId(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetMobilityFeatureProfileByGlobalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobilityFeatureProfileByGlobalId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetMobilityFeatureProfileByGlobalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobilityFeatureProfileByGlobalIdRequest struct via the builder pattern


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


## GetMobilityGlobalBasicParcelSchemaBySchemaType

> string GetMobilityGlobalBasicParcelSchemaBySchemaType(ctx).SchemaType(schemaType).Execute()





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
    schemaType := "schemaType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetMobilityGlobalBasicParcelSchemaBySchemaType(context.Background()).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetMobilityGlobalBasicParcelSchemaBySchemaType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobilityGlobalBasicParcelSchemaBySchemaType`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetMobilityGlobalBasicParcelSchemaBySchemaType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMobilityGlobalBasicParcelSchemaBySchemaTypeRequest struct via the builder pattern


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


## GetNetworkProtocolProfileParcelForMobility

> string GetNetworkProtocolProfileParcelForMobility(ctx, profileId, networkProtocolId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    networkProtocolId := "networkProtocolId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetNetworkProtocolProfileParcelForMobility(context.Background(), profileId, networkProtocolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetNetworkProtocolProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProtocolProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetNetworkProtocolProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**networkProtocolId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProtocolProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetNetworkProtocolProfileParcelListForMobility

> string GetNetworkProtocolProfileParcelListForMobility(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetNetworkProtocolProfileParcelListForMobility(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetNetworkProtocolProfileParcelListForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProtocolProfileParcelListForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetNetworkProtocolProfileParcelListForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProtocolProfileParcelListForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSecurityPolicyProfileParcelForMobility

> string GetSecurityPolicyProfileParcelForMobility(ctx, profileId, securityPolicyId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    securityPolicyId := "securityPolicyId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetSecurityPolicyProfileParcelForMobility(context.Background(), profileId, securityPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetSecurityPolicyProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityPolicyProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetSecurityPolicyProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**securityPolicyId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPolicyProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSecurityPolicyProfileParcelListForMobility

> string GetSecurityPolicyProfileParcelListForMobility(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetSecurityPolicyProfileParcelListForMobility(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetSecurityPolicyProfileParcelListForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityPolicyProfileParcelListForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetSecurityPolicyProfileParcelListForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPolicyProfileParcelListForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetVpnProfileParcelForMobility

> string GetVpnProfileParcelForMobility(ctx, profileId, vpnId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    vpnId := "vpnId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetVpnProfileParcelForMobility(context.Background(), profileId, vpnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetVpnProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVpnProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetVpnProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**vpnId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetVpnProfileParcelListForMobility

> string GetVpnProfileParcelListForMobility(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetVpnProfileParcelListForMobility(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetVpnProfileParcelListForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVpnProfileParcelListForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetVpnProfileParcelListForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnProfileParcelListForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetWifiProfileParcelForMobility

> string GetWifiProfileParcelForMobility(ctx, profileId, wifiId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID
    wifiId := "wifiId_example" // string | Profile Parcel ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetWifiProfileParcelForMobility(context.Background(), profileId, wifiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetWifiProfileParcelForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWifiProfileParcelForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetWifiProfileParcelForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 
**wifiId** | **string** | Profile Parcel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWifiProfileParcelForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetWifiProfileParcelListForMobility

> string GetWifiProfileParcelListForMobility(ctx, profileId).Execute()





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
    profileId := "profileId_example" // string | Feature Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationFeatureProfileMobilityApi.GetWifiProfileParcelListForMobility(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationFeatureProfileMobilityApi.GetWifiProfileParcelListForMobility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWifiProfileParcelListForMobility`: string
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationFeatureProfileMobilityApi.GetWifiProfileParcelListForMobility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Feature Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWifiProfileParcelListForMobilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

