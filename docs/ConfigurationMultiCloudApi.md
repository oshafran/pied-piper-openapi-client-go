# \ConfigurationMultiCloudApi

All URIs are relative to */dataservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEdgeGlobalSettings**](ConfigurationMultiCloudApi.md#AddEdgeGlobalSettings) | **Post** /multicloud/settings/edge/global | 
[**AddGlobalSettings**](ConfigurationMultiCloudApi.md#AddGlobalSettings) | **Post** /multicloud/settings/global | 
[**AttachSites**](ConfigurationMultiCloudApi.md#AttachSites) | **Post** /multicloud/cloudgateway/{cloudGatewayName}/site | 
[**Audit**](ConfigurationMultiCloudApi.md#Audit) | **Post** /multicloud/audit | 
[**AuditDryRun**](ConfigurationMultiCloudApi.md#AuditDryRun) | **Get** /multicloud/audit | 
[**CleanUpAllConnectivityGatewaysInLocalDB**](ConfigurationMultiCloudApi.md#CleanUpAllConnectivityGatewaysInLocalDB) | **Delete** /multicloud/connectivitygateway | 
[**CreateCgw**](ConfigurationMultiCloudApi.md#CreateCgw) | **Post** /multicloud/cloudgateway | 
[**CreateConnectivityGateway**](ConfigurationMultiCloudApi.md#CreateConnectivityGateway) | **Post** /multicloud/connectivitygateway | 
[**CreateDeviceLink**](ConfigurationMultiCloudApi.md#CreateDeviceLink) | **Post** /multicloud/devicelink/edge | 
[**CreateEdgeConnectivity**](ConfigurationMultiCloudApi.md#CreateEdgeConnectivity) | **Post** /multicloud/connectivity/edge | 
[**CreateIcgw**](ConfigurationMultiCloudApi.md#CreateIcgw) | **Post** /multicloud/gateway/edge | 
[**CreateVirtualWan**](ConfigurationMultiCloudApi.md#CreateVirtualWan) | **Post** /multicloud/vwan | 
[**DeleteAccount**](ConfigurationMultiCloudApi.md#DeleteAccount) | **Delete** /multicloud/accounts/{accountId} | 
[**DeleteCgw**](ConfigurationMultiCloudApi.md#DeleteCgw) | **Delete** /multicloud/cloudgateway/{cloudGatewayName} | 
[**DeleteConnectivityGateway**](ConfigurationMultiCloudApi.md#DeleteConnectivityGateway) | **Delete** /multicloud/connectivitygateway/{cloudProvider}/{connectivityGatewayName} | 
[**DeleteDeviceLink**](ConfigurationMultiCloudApi.md#DeleteDeviceLink) | **Delete** /multicloud/devicelink/edge/{deviceLinkName} | 
[**DeleteEdgeAccount**](ConfigurationMultiCloudApi.md#DeleteEdgeAccount) | **Delete** /multicloud/accounts/edge/{accountId} | 
[**DeleteEdgeAccount1**](ConfigurationMultiCloudApi.md#DeleteEdgeAccount1) | **Delete** /multicloud/locations/edge/{edgeType} | 
[**DeleteEdgeConnectivity**](ConfigurationMultiCloudApi.md#DeleteEdgeConnectivity) | **Delete** /multicloud/connectivity/edge/{connectionName} | 
[**DeleteIcgw**](ConfigurationMultiCloudApi.md#DeleteIcgw) | **Delete** /multicloud/gateway/edge/{edgeGatewayName} | 
[**DeleteVirtualWan**](ConfigurationMultiCloudApi.md#DeleteVirtualWan) | **Delete** /multicloud/vwan/{cloudProvider}/{vWanName} | 
[**DetachSites1**](ConfigurationMultiCloudApi.md#DetachSites1) | **Delete** /multicloud/cloudgateway/{cloudGatewayName}/site | 
[**EdgeAudit**](ConfigurationMultiCloudApi.md#EdgeAudit) | **Post** /multicloud/audit/edge | 
[**EdgeAuditDryRun**](ConfigurationMultiCloudApi.md#EdgeAuditDryRun) | **Get** /multicloud/audit/edge | 
[**EditTag**](ConfigurationMultiCloudApi.md#EditTag) | **Put** /multicloud/hostvpc/tags | 
[**GetAllCloudAccounts**](ConfigurationMultiCloudApi.md#GetAllCloudAccounts) | **Get** /multicloud/accounts | 
[**GetAzureNetworkVirtualAppliances**](ConfigurationMultiCloudApi.md#GetAzureNetworkVirtualAppliances) | **Get** /multicloud/cloudgateway/nvas | 
[**GetAzureNvaSkuList**](ConfigurationMultiCloudApi.md#GetAzureNvaSkuList) | **Get** /multicloud/cloudgateway/nvasku | 
[**GetAzureResourceGroups**](ConfigurationMultiCloudApi.md#GetAzureResourceGroups) | **Get** /multicloud/cloudgateway/resourceGroups | 
[**GetAzureVirtualHubs**](ConfigurationMultiCloudApi.md#GetAzureVirtualHubs) | **Get** /multicloud/cloudgateway/vhubs | 
[**GetAzureVirtualWans**](ConfigurationMultiCloudApi.md#GetAzureVirtualWans) | **Get** /multicloud/cloudgateway/vwans | 
[**GetCgwAssociatedMappings**](ConfigurationMultiCloudApi.md#GetCgwAssociatedMappings) | **Get** /multicloud/mapping/{cloudType} | 
[**GetCgwCustomSettingDetails**](ConfigurationMultiCloudApi.md#GetCgwCustomSettingDetails) | **Get** /multicloud/cloudgatewaysetting/{cloudGatewayName} | 
[**GetCgwDetails**](ConfigurationMultiCloudApi.md#GetCgwDetails) | **Get** /multicloud/cloudgateway/{cloudGatewayName} | 
[**GetCgwOrgResources**](ConfigurationMultiCloudApi.md#GetCgwOrgResources) | **Get** /multicloud/cloudgateway/resource | 
[**GetCgwTypes**](ConfigurationMultiCloudApi.md#GetCgwTypes) | **Get** /multicloud/cloudgatewaytype | 
[**GetCgws**](ConfigurationMultiCloudApi.md#GetCgws) | **Get** /multicloud/cloudgateway | 
[**GetCloudAccountDetails**](ConfigurationMultiCloudApi.md#GetCloudAccountDetails) | **Get** /multicloud/accounts/{accountId} | 
[**GetCloudConnectedSites**](ConfigurationMultiCloudApi.md#GetCloudConnectedSites) | **Get** /multicloud/connected-sites/{cloudType} | 
[**GetCloudConnectedSites1**](ConfigurationMultiCloudApi.md#GetCloudConnectedSites1) | **Get** /multicloud/connected-sites/edge/{edgeType} | 
[**GetCloudDevices**](ConfigurationMultiCloudApi.md#GetCloudDevices) | **Get** /multicloud/devices/{cloudType} | 
[**GetCloudDevices1**](ConfigurationMultiCloudApi.md#GetCloudDevices1) | **Get** /multicloud/devices/edge/{edgeType} | 
[**GetCloudGateways**](ConfigurationMultiCloudApi.md#GetCloudGateways) | **Get** /multicloud/cloudgateways/{cloudType} | 
[**GetCloudRegions**](ConfigurationMultiCloudApi.md#GetCloudRegions) | **Get** /multicloud/regions | 
[**GetCloudTypes**](ConfigurationMultiCloudApi.md#GetCloudTypes) | **Get** /multicloud/types | 
[**GetCloudWidget**](ConfigurationMultiCloudApi.md#GetCloudWidget) | **Get** /multicloud/widget/{cloudType} | 
[**GetConnectivityGatewayCreationOptions**](ConfigurationMultiCloudApi.md#GetConnectivityGatewayCreationOptions) | **Get** /multicloud/connectivitygatewaycreationoptions | 
[**GetConnectivityGateways**](ConfigurationMultiCloudApi.md#GetConnectivityGateways) | **Get** /multicloud/connectivitygateway | 
[**GetDashboardEdgeInfo**](ConfigurationMultiCloudApi.md#GetDashboardEdgeInfo) | **Get** /multicloud/dashboard/edge | 
[**GetDefaultMappingValues**](ConfigurationMultiCloudApi.md#GetDefaultMappingValues) | **Get** /multicloud/map/defaults | 
[**GetDeviceLinks**](ConfigurationMultiCloudApi.md#GetDeviceLinks) | **Get** /multicloud/devicelink/edge | 
[**GetDlPortSpeed**](ConfigurationMultiCloudApi.md#GetDlPortSpeed) | **Get** /multicloud/devicelink/edge/portspeed/{edgeType} | 
[**GetEdgeAccountDetails**](ConfigurationMultiCloudApi.md#GetEdgeAccountDetails) | **Get** /multicloud/accounts/edge/{accountId} | 
[**GetEdgeAccounts**](ConfigurationMultiCloudApi.md#GetEdgeAccounts) | **Get** /multicloud/accounts/edge | 
[**GetEdgeBillingAccounts**](ConfigurationMultiCloudApi.md#GetEdgeBillingAccounts) | **Get** /multicloud/billingaccounts/edge/{edgeType}/{edgeAccountId} | 
[**GetEdgeConnectivityDetailByName**](ConfigurationMultiCloudApi.md#GetEdgeConnectivityDetailByName) | **Get** /multicloud/connectivity/edge/{connectivityName} | 
[**GetEdgeConnectivityDetails**](ConfigurationMultiCloudApi.md#GetEdgeConnectivityDetails) | **Get** /multicloud/connectivity/edge | 
[**GetEdgeGateways**](ConfigurationMultiCloudApi.md#GetEdgeGateways) | **Get** /multicloud/gateways/edge/{edgeType} | 
[**GetEdgeGlobalSettings**](ConfigurationMultiCloudApi.md#GetEdgeGlobalSettings) | **Get** /multicloud/settings/edge/global | 
[**GetEdgeLocationsInfo**](ConfigurationMultiCloudApi.md#GetEdgeLocationsInfo) | **Get** /multicloud/locations/edge/{edgeType} | 
[**GetEdgeMappingTags**](ConfigurationMultiCloudApi.md#GetEdgeMappingTags) | **Get** /multicloud/map/tags/edge | 
[**GetEdgeTypes**](ConfigurationMultiCloudApi.md#GetEdgeTypes) | **Get** /multicloud/types/edge | 
[**GetEdgeWanDevices**](ConfigurationMultiCloudApi.md#GetEdgeWanDevices) | **Get** /multicloud/edge/{edgeType}/device | 
[**GetEdgeWidget**](ConfigurationMultiCloudApi.md#GetEdgeWidget) | **Get** /multicloud/widget/edge/{edgeType} | 
[**GetGlobalSettings**](ConfigurationMultiCloudApi.md#GetGlobalSettings) | **Get** /multicloud/settings/global | 
[**GetHostVpcs**](ConfigurationMultiCloudApi.md#GetHostVpcs) | **Get** /multicloud/hostvpc | 
[**GetIcgwCustomSettingDetails**](ConfigurationMultiCloudApi.md#GetIcgwCustomSettingDetails) | **Get** /multicloud/gateway/edge/setting/{edgeGatewayName} | 
[**GetIcgwDetails**](ConfigurationMultiCloudApi.md#GetIcgwDetails) | **Get** /multicloud/gateway/edge/{edgeGatewayName} | 
[**GetIcgwTypes**](ConfigurationMultiCloudApi.md#GetIcgwTypes) | **Get** /multicloud/gateway/edge/types | 
[**GetIcgws**](ConfigurationMultiCloudApi.md#GetIcgws) | **Get** /multicloud/gateway/edge | 
[**GetMappingMatrix**](ConfigurationMultiCloudApi.md#GetMappingMatrix) | **Get** /multicloud/map | 
[**GetMappingStatus**](ConfigurationMultiCloudApi.md#GetMappingStatus) | **Get** /multicloud/map/status | 
[**GetMappingSummary**](ConfigurationMultiCloudApi.md#GetMappingSummary) | **Get** /multicloud/map/summary | 
[**GetMappingTags**](ConfigurationMultiCloudApi.md#GetMappingTags) | **Get** /multicloud/map/tags | 
[**GetMappingVpns**](ConfigurationMultiCloudApi.md#GetMappingVpns) | **Get** /multicloud/map/vpns | 
[**GetNvaSecurityRules**](ConfigurationMultiCloudApi.md#GetNvaSecurityRules) | **Get** /multicloud/cloudgateway/nvaSecurityRules/{cloudGatewayName} | 
[**GetPartnerPorts**](ConfigurationMultiCloudApi.md#GetPartnerPorts) | **Get** /multicloud/partnerports/edge | 
[**GetPortSpeed**](ConfigurationMultiCloudApi.md#GetPortSpeed) | **Get** /multicloud/portSpeed/edge/{edgeType}/{edgeAccountId}/{connectivityType} | 
[**GetPostAggregationDataByQuery25**](ConfigurationMultiCloudApi.md#GetPostAggregationDataByQuery25) | **Post** /multicloud/statistics/interface/aggregation | 
[**GetSites**](ConfigurationMultiCloudApi.md#GetSites) | **Get** /multicloud/site | 
[**GetSites1**](ConfigurationMultiCloudApi.md#GetSites1) | **Get** /multicloud/cloudgateway/{cloudGatewayName}/site | 
[**GetSshKeyList**](ConfigurationMultiCloudApi.md#GetSshKeyList) | **Get** /multicloud/sshkeys | 
[**GetSupportedEdgeImageNames**](ConfigurationMultiCloudApi.md#GetSupportedEdgeImageNames) | **Get** /multicloud/imagename/edge | 
[**GetSupportedEdgeInstanceSize**](ConfigurationMultiCloudApi.md#GetSupportedEdgeInstanceSize) | **Get** /multicloud/instancesize/edge | 
[**GetSupportedInstanceSize**](ConfigurationMultiCloudApi.md#GetSupportedInstanceSize) | **Get** /multicloud/instancesize | 
[**GetSupportedLoopbackCgwColors**](ConfigurationMultiCloudApi.md#GetSupportedLoopbackCgwColors) | **Get** /multicloud/loopbackCGWColor/edge | 
[**GetSupportedSoftwareImageList**](ConfigurationMultiCloudApi.md#GetSupportedSoftwareImageList) | **Get** /multicloud/swimages | 
[**GetTunnelNames**](ConfigurationMultiCloudApi.md#GetTunnelNames) | **Get** /multicloud/tunnels/{cloudType} | 
[**GetVHubs**](ConfigurationMultiCloudApi.md#GetVHubs) | **Get** /multicloud/vhubs | 
[**GetVWans**](ConfigurationMultiCloudApi.md#GetVWans) | **Get** /multicloud/vwans | 
[**GetVpcTags**](ConfigurationMultiCloudApi.md#GetVpcTags) | **Get** /multicloud/hostvpc/tags | 
[**GetWanDevices**](ConfigurationMultiCloudApi.md#GetWanDevices) | **Get** /multicloud/device | 
[**GetWanInterfaceColors**](ConfigurationMultiCloudApi.md#GetWanInterfaceColors) | **Get** /multicloud/interfacecolor | 
[**HostvpcTagging**](ConfigurationMultiCloudApi.md#HostvpcTagging) | **Post** /multicloud/hostvpc/tags | 
[**ProcessMapping**](ConfigurationMultiCloudApi.md#ProcessMapping) | **Post** /multicloud/map | 
[**Telemetry**](ConfigurationMultiCloudApi.md#Telemetry) | **Post** /multicloud/telemetry | 
[**TunnelScaling**](ConfigurationMultiCloudApi.md#TunnelScaling) | **Put** /multicloud/cloudgateway/{cloudGatewayName}/site | 
[**UnTag**](ConfigurationMultiCloudApi.md#UnTag) | **Delete** /multicloud/hostvpc/tags/{tagName} | 
[**UpdateAccount**](ConfigurationMultiCloudApi.md#UpdateAccount) | **Put** /multicloud/accounts/{accountId} | 
[**UpdateCgw**](ConfigurationMultiCloudApi.md#UpdateCgw) | **Put** /multicloud/cloudgateway/{cloudGatewayName} | 
[**UpdateDeviceLink**](ConfigurationMultiCloudApi.md#UpdateDeviceLink) | **Put** /multicloud/devicelink/edge | 
[**UpdateEdgeAccount**](ConfigurationMultiCloudApi.md#UpdateEdgeAccount) | **Put** /multicloud/accounts/edge/{accountId} | 
[**UpdateEdgeConnectivity**](ConfigurationMultiCloudApi.md#UpdateEdgeConnectivity) | **Put** /multicloud/connectivity/edge | 
[**UpdateEdgeGlobalSettings**](ConfigurationMultiCloudApi.md#UpdateEdgeGlobalSettings) | **Put** /multicloud/settings/edge/global | 
[**UpdateEdgeLocationsInfo**](ConfigurationMultiCloudApi.md#UpdateEdgeLocationsInfo) | **Put** /multicloud/locations/edge/{edgeType}/accountId/{accountId} | 
[**UpdateGlobalSettings**](ConfigurationMultiCloudApi.md#UpdateGlobalSettings) | **Put** /multicloud/settings/global | 
[**UpdateIcgw**](ConfigurationMultiCloudApi.md#UpdateIcgw) | **Put** /multicloud/gateway/edge/{edgeGatewayName} | 
[**UpdateNvaSecurityRules**](ConfigurationMultiCloudApi.md#UpdateNvaSecurityRules) | **Put** /multicloud/cloudgateway/nvaSecurityRules/{cloudGatewayName} | 
[**ValidateAccountAdd**](ConfigurationMultiCloudApi.md#ValidateAccountAdd) | **Post** /multicloud/accounts | 
[**ValidateAccountUpdateCredentials**](ConfigurationMultiCloudApi.md#ValidateAccountUpdateCredentials) | **Put** /multicloud/accounts/{accountId}/credentials | 
[**ValidateEdgeAccountAdd**](ConfigurationMultiCloudApi.md#ValidateEdgeAccountAdd) | **Post** /multicloud/accounts/edge | 
[**ValidateEdgeAccountUpdateCredentials**](ConfigurationMultiCloudApi.md#ValidateEdgeAccountUpdateCredentials) | **Put** /multicloud/accounts/edge/{accountId}/credentials | 



## AddEdgeGlobalSettings

> AddEdgeGlobalSettings(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Global setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.AddEdgeGlobalSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.AddEdgeGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEdgeGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Global setting | 

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


## AddGlobalSettings

> AddGlobalSettings(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Global setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.AddGlobalSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.AddGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Global setting | 

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


## AttachSites

> map[string]interface{} AttachSites(ctx, cloudGatewayName).Body(body).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    body := map[string]interface{}{ ... } // map[string]interface{} | Site information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.AttachSites(context.Background(), cloudGatewayName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.AttachSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachSites`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.AttachSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Site information | 

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


## Audit

> map[string]interface{} Audit(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Audit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.Audit(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.Audit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Audit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.Audit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Audit | 

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


## AuditDryRun

> AuditDryRun(ctx).CloudType(cloudType).CloudRegion(cloudRegion).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type (optional)
    cloudRegion := "cloudRegion_example" // string | Region (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.AuditDryRun(context.Background()).CloudType(cloudType).CloudRegion(cloudRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.AuditDryRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **cloudRegion** | **string** | Region | 

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


## CleanUpAllConnectivityGatewaysInLocalDB

> map[string]interface{} CleanUpAllConnectivityGatewaysInLocalDB(ctx).DeletionType(deletionType).Execute()





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
    deletionType := "deletionType_example" // string | Deletion Type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CleanUpAllConnectivityGatewaysInLocalDB(context.Background()).DeletionType(deletionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CleanUpAllConnectivityGatewaysInLocalDB``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CleanUpAllConnectivityGatewaysInLocalDB`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CleanUpAllConnectivityGatewaysInLocalDB`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCleanUpAllConnectivityGatewaysInLocalDBRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletionType** | **string** | Deletion Type | 

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


## CreateCgw

> map[string]interface{} CreateCgw(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Cloud gateway (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CreateCgw(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CreateCgw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCgw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CreateCgw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCgwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Cloud gateway | 

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


## CreateConnectivityGateway

> map[string]interface{} CreateConnectivityGateway(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Connectivity gateway (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CreateConnectivityGateway(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CreateConnectivityGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectivityGateway`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CreateConnectivityGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectivityGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Connectivity gateway | 

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


## CreateDeviceLink

> map[string]interface{} CreateDeviceLink(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device Link (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CreateDeviceLink(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CreateDeviceLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLink`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CreateDeviceLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device Link | 

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


## CreateEdgeConnectivity

> map[string]interface{} CreateEdgeConnectivity(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Edge connectivity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CreateEdgeConnectivity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CreateEdgeConnectivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEdgeConnectivity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CreateEdgeConnectivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Edge connectivity | 

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


## CreateIcgw

> map[string]interface{} CreateIcgw(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Interconnect Gateway (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CreateIcgw(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CreateIcgw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIcgw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CreateIcgw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIcgwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Interconnect Gateway | 

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


## CreateVirtualWan

> map[string]interface{} CreateVirtualWan(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Virtual WAN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.CreateVirtualWan(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.CreateVirtualWan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualWan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.CreateVirtualWan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualWanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Virtual WAN | 

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


## DeleteAccount

> DeleteAccount(ctx, accountId).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## DeleteCgw

> map[string]interface{} DeleteCgw(ctx, cloudGatewayName).DeleteAllResources(deleteAllResources).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    deleteAllResources := true // bool | Optional Flag for deletion of Azure Resource Group, Default: True (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteCgw(context.Background(), cloudGatewayName).DeleteAllResources(deleteAllResources).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteCgw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCgw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DeleteCgw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCgwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAllResources** | **bool** | Optional Flag for deletion of Azure Resource Group, Default: True | 

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


## DeleteConnectivityGateway

> map[string]interface{} DeleteConnectivityGateway(ctx, cloudProvider, connectivityGatewayName).ConnectivityType(connectivityType).Execute()





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
    cloudProvider := "cloudProvider_example" // string | Cloud Provider
    connectivityGatewayName := "connectivityGatewayName_example" // string | Connectivity gateway name
    connectivityType := "connectivityType_example" // string | Cloud Connectivity Type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteConnectivityGateway(context.Background(), cloudProvider, connectivityGatewayName).ConnectivityType(connectivityType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteConnectivityGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnectivityGateway`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DeleteConnectivityGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** | Cloud Provider | 
**connectivityGatewayName** | **string** | Connectivity gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectivityGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectivityType** | **string** | Cloud Connectivity Type | 

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


## DeleteDeviceLink

> map[string]interface{} DeleteDeviceLink(ctx, deviceLinkName).Execute()





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
    deviceLinkName := "deviceLinkName_example" // string | Device Link Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteDeviceLink(context.Background(), deviceLinkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteDeviceLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeviceLink`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DeleteDeviceLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceLinkName** | **string** | Device Link Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceLinkRequest struct via the builder pattern


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


## DeleteEdgeAccount

> DeleteEdgeAccount(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | Edge Account Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteEdgeAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteEdgeAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Edge Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEdgeAccountRequest struct via the builder pattern


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


## DeleteEdgeAccount1

> DeleteEdgeAccount1(ctx, edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge Type (default to "MEGAPORT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteEdgeAccount1(context.Background(), edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteEdgeAccount1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge Type | [default to &quot;MEGAPORT&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEdgeAccount1Request struct via the builder pattern


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


## DeleteEdgeConnectivity

> map[string]interface{} DeleteEdgeConnectivity(ctx, connectionName).Execute()





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
    connectionName := "connectionName_example" // string | Edge connectivity name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteEdgeConnectivity(context.Background(), connectionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteEdgeConnectivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEdgeConnectivity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DeleteEdgeConnectivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string** | Edge connectivity name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEdgeConnectivityRequest struct via the builder pattern


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


## DeleteIcgw

> map[string]interface{} DeleteIcgw(ctx, edgeGatewayName).Execute()





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
    edgeGatewayName := "edgeGatewayName_example" // string | Edge gateway name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteIcgw(context.Background(), edgeGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteIcgw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIcgw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DeleteIcgw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeGatewayName** | **string** | Edge gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIcgwRequest struct via the builder pattern


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


## DeleteVirtualWan

> map[string]interface{} DeleteVirtualWan(ctx, cloudProvider, vWanName).AccountId(accountId).ResourceGroup(resourceGroup).Execute()





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
    cloudProvider := "cloudProvider_example" // string | Cloud Provider
    vWanName := "vWanName_example" // string | Virtual Wan name
    accountId := "accountId_example" // string | Account Id (optional)
    resourceGroup := "resourceGroup_example" // string | Resource Group (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DeleteVirtualWan(context.Background(), cloudProvider, vWanName).AccountId(accountId).ResourceGroup(resourceGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DeleteVirtualWan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualWan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DeleteVirtualWan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** | Cloud Provider | 
**vWanName** | **string** | Virtual Wan name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualWanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **string** | Account Id | 
 **resourceGroup** | **string** | Resource Group | 

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


## DetachSites1

> map[string]interface{} DetachSites1(ctx, cloudGatewayName).Body(body).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    body := map[string]interface{}{ ... } // map[string]interface{} | Site information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.DetachSites1(context.Background(), cloudGatewayName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.DetachSites1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachSites1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.DetachSites1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachSites1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Site information | 

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


## EdgeAudit

> EdgeAudit(ctx).EdgeType(edgeType).CloudType(cloudType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)
    cloudType := "cloudType_example" // string | Cloud type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.EdgeAudit(context.Background()).EdgeType(edgeType).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.EdgeAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 
 **cloudType** | **string** | Cloud type | 

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


## EdgeAuditDryRun

> EdgeAuditDryRun(ctx).EdgeType(edgeType).CloudType(cloudType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)
    cloudType := "cloudType_example" // string | Cloud type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.EdgeAuditDryRun(context.Background()).EdgeType(edgeType).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.EdgeAuditDryRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeAuditDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 
 **cloudType** | **string** | Cloud type | 

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


## EditTag

> map[string]interface{} EditTag(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | VPC tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.EditTag(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.EditTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTag`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.EditTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | VPC tag | 

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


## GetAllCloudAccounts

> map[string]interface{} GetAllCloudAccounts(ctx).CloudType(cloudType).CloudGatewayEnabled(cloudGatewayEnabled).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type (optional)
    cloudGatewayEnabled := true // bool | Cloud gateway enabled flag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetAllCloudAccounts(context.Background()).CloudType(cloudType).CloudGatewayEnabled(cloudGatewayEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetAllCloudAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCloudAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetAllCloudAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCloudAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **cloudGatewayEnabled** | **bool** | Cloud gateway enabled flag | 

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


## GetAzureNetworkVirtualAppliances

> map[string]interface{} GetAzureNetworkVirtualAppliances(ctx).CloudType(cloudType).AccoundId(accoundId).Region(region).ResourceGroupName(resourceGroupName).ResourceGroupSource(resourceGroupSource).VhubName(vhubName).VhubSource(vhubSource).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accoundId := "accoundId_example" // string | Account ID
    region := "region_example" // string | Region
    resourceGroupName := "resourceGroupName_example" // string | Resource Group Name
    resourceGroupSource := "resourceGroupSource_example" // string | Resource Group Source
    vhubName := "vhubName_example" // string | VHUB name
    vhubSource := "vhubSource_example" // string | VHUB source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetAzureNetworkVirtualAppliances(context.Background()).CloudType(cloudType).AccoundId(accoundId).Region(region).ResourceGroupName(resourceGroupName).ResourceGroupSource(resourceGroupSource).VhubName(vhubName).VhubSource(vhubSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetAzureNetworkVirtualAppliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureNetworkVirtualAppliances`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetAzureNetworkVirtualAppliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureNetworkVirtualAppliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accoundId** | **string** | Account ID | 
 **region** | **string** | Region | 
 **resourceGroupName** | **string** | Resource Group Name | 
 **resourceGroupSource** | **string** | Resource Group Source | 
 **vhubName** | **string** | VHUB name | 
 **vhubSource** | **string** | VHUB source | 

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


## GetAzureNvaSkuList

> map[string]interface{} GetAzureNvaSkuList(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetAzureNvaSkuList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetAzureNvaSkuList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureNvaSkuList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetAzureNvaSkuList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureNvaSkuListRequest struct via the builder pattern


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


## GetAzureResourceGroups

> map[string]interface{} GetAzureResourceGroups(ctx).CloudType(cloudType).AccountId(accountId).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accountId := "accountId_example" // string | Account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetAzureResourceGroups(context.Background()).CloudType(cloudType).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetAzureResourceGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureResourceGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetAzureResourceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureResourceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accountId** | **string** | Account ID | 

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


## GetAzureVirtualHubs

> map[string]interface{} GetAzureVirtualHubs(ctx).CloudType(cloudType).AccoundId(accoundId).Region(region).ResourceGroupName(resourceGroupName).ResourceGroupSource(resourceGroupSource).VwanName(vwanName).VwanSource(vwanSource).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accoundId := "accoundId_example" // string | Account ID
    region := "region_example" // string | Region
    resourceGroupName := "resourceGroupName_example" // string | Resource Group Name
    resourceGroupSource := "resourceGroupSource_example" // string | Resource Group Source
    vwanName := "vwanName_example" // string | VWAN name
    vwanSource := "vwanSource_example" // string | VWAN source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetAzureVirtualHubs(context.Background()).CloudType(cloudType).AccoundId(accoundId).Region(region).ResourceGroupName(resourceGroupName).ResourceGroupSource(resourceGroupSource).VwanName(vwanName).VwanSource(vwanSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetAzureVirtualHubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureVirtualHubs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetAzureVirtualHubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureVirtualHubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accoundId** | **string** | Account ID | 
 **region** | **string** | Region | 
 **resourceGroupName** | **string** | Resource Group Name | 
 **resourceGroupSource** | **string** | Resource Group Source | 
 **vwanName** | **string** | VWAN name | 
 **vwanSource** | **string** | VWAN source | 

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


## GetAzureVirtualWans

> map[string]interface{} GetAzureVirtualWans(ctx).CloudType(cloudType).AccoundId(accoundId).ResourceGroupName(resourceGroupName).ResourceGroupSource(resourceGroupSource).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accoundId := "accoundId_example" // string | Account ID
    resourceGroupName := "resourceGroupName_example" // string | Resource Group Name
    resourceGroupSource := "resourceGroupSource_example" // string | Resource Group Source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetAzureVirtualWans(context.Background()).CloudType(cloudType).AccoundId(accoundId).ResourceGroupName(resourceGroupName).ResourceGroupSource(resourceGroupSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetAzureVirtualWans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureVirtualWans`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetAzureVirtualWans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureVirtualWansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accoundId** | **string** | Account ID | 
 **resourceGroupName** | **string** | Resource Group Name | 
 **resourceGroupSource** | **string** | Resource Group Source | 

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


## GetCgwAssociatedMappings

> map[string]interface{} GetCgwAssociatedMappings(ctx, cloudType).CloudGatewayName(cloudGatewayName).SiteUuid(siteUuid).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud Gateway Name
    siteUuid := "siteUuid_example" // string | Site Device UUID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCgwAssociatedMappings(context.Background(), cloudType).CloudGatewayName(cloudGatewayName).SiteUuid(siteUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCgwAssociatedMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCgwAssociatedMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCgwAssociatedMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudType** | **string** | Cloud type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCgwAssociatedMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGatewayName** | **string** | Cloud Gateway Name | 
 **siteUuid** | **string** | Site Device UUID | 

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


## GetCgwCustomSettingDetails

> map[string]interface{} GetCgwCustomSettingDetails(ctx, cloudGatewayName).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCgwCustomSettingDetails(context.Background(), cloudGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCgwCustomSettingDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCgwCustomSettingDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCgwCustomSettingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCgwCustomSettingDetailsRequest struct via the builder pattern


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


## GetCgwDetails

> map[string]interface{} GetCgwDetails(ctx, cloudGatewayName).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCgwDetails(context.Background(), cloudGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCgwDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCgwDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCgwDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCgwDetailsRequest struct via the builder pattern


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


## GetCgwOrgResources

> map[string]interface{} GetCgwOrgResources(ctx).CloudGatewayName(cloudGatewayName).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCgwOrgResources(context.Background()).CloudGatewayName(cloudGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCgwOrgResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCgwOrgResources`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCgwOrgResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCgwOrgResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudGatewayName** | **string** | Cloud gateway name | 

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


## GetCgwTypes

> map[string]interface{} GetCgwTypes(ctx).CloudType(cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCgwTypes(context.Background()).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCgwTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCgwTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCgwTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCgwTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 

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


## GetCgws

> map[string]interface{} GetCgws(ctx).CloudType(cloudType).AccountId(accountId).Region(region).CloudGatewayName(cloudGatewayName).ConnectivityState(connectivityState).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type (optional)
    accountId := "accountId_example" // string | Account Id (optional)
    region := "region_example" // string | Region (optional)
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name (optional)
    connectivityState := "connectivityState_example" // string | Connectivity State (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCgws(context.Background()).CloudType(cloudType).AccountId(accountId).Region(region).CloudGatewayName(cloudGatewayName).ConnectivityState(connectivityState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCgws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCgws`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCgws`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCgwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accountId** | **string** | Account Id | 
 **region** | **string** | Region | 
 **cloudGatewayName** | **string** | Cloud gateway name | 
 **connectivityState** | **string** | Connectivity State | 

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


## GetCloudAccountDetails

> map[string]interface{} GetCloudAccountDetails(ctx, accountId).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudAccountDetails(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudAccountDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudAccountDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudAccountDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudAccountDetailsRequest struct via the builder pattern


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


## GetCloudConnectedSites

> map[string]interface{} GetCloudConnectedSites(ctx, cloudType).CloudGatewayName(cloudGatewayName).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud Gateway Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudConnectedSites(context.Background(), cloudType).CloudGatewayName(cloudGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudConnectedSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudConnectedSites`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudConnectedSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudType** | **string** | Cloud type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudConnectedSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGatewayName** | **string** | Cloud Gateway Name | 

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


## GetCloudConnectedSites1

> map[string]interface{} GetCloudConnectedSites1(ctx, edgeType).EdgeGatewayName(edgeGatewayName).Execute()





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
    edgeType := "edgeType_example" // string | Edge type
    edgeGatewayName := "edgeGatewayName_example" // string | Interconnect Gateway Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudConnectedSites1(context.Background(), edgeType).EdgeGatewayName(edgeGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudConnectedSites1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudConnectedSites1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudConnectedSites1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudConnectedSites1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeGatewayName** | **string** | Interconnect Gateway Name | 

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


## GetCloudDevices

> map[string]interface{} GetCloudDevices(ctx, cloudType).CloudGatewayName(cloudGatewayName).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud Gateway Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudDevices(context.Background(), cloudType).CloudGatewayName(cloudGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudType** | **string** | Cloud type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGatewayName** | **string** | Cloud Gateway Name | 

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


## GetCloudDevices1

> map[string]interface{} GetCloudDevices1(ctx, edgeType).EdgeGatewayName(edgeGatewayName).Execute()





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
    edgeType := "edgeType_example" // string | Edge type
    edgeGatewayName := "edgeGatewayName_example" // string | Edge Gateway Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudDevices1(context.Background(), edgeType).EdgeGatewayName(edgeGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudDevices1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudDevices1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudDevices1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDevices1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeGatewayName** | **string** | Edge Gateway Name | 

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


## GetCloudGateways

> map[string]interface{} GetCloudGateways(ctx, cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudGateways(context.Background(), cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudGateways`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudType** | **string** | Cloud type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudGatewaysRequest struct via the builder pattern


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


## GetCloudRegions

> map[string]interface{} GetCloudRegions(ctx).CloudType(cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudRegions(context.Background()).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudRegions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 

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


## GetCloudTypes

> map[string]interface{} GetCloudTypes(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudTypesRequest struct via the builder pattern


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


## GetCloudWidget

> map[string]interface{} GetCloudWidget(ctx, cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetCloudWidget(context.Background(), cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetCloudWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudWidget`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetCloudWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudType** | **string** | Cloud type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudWidgetRequest struct via the builder pattern


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


## GetConnectivityGatewayCreationOptions

> map[string]interface{} GetConnectivityGatewayCreationOptions(ctx).AccountId(accountId).CloudType(cloudType).ConnectivityType(connectivityType).Refresh(refresh).Execute()





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
    accountId := "accountId_example" // string | Account Id (optional)
    cloudType := "cloudType_example" // string | Cloud Type (optional)
    connectivityType := "connectivityType_example" // string | Cloud Connectivity Type (optional)
    refresh := "refresh_example" // string | Refresh (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetConnectivityGatewayCreationOptions(context.Background()).AccountId(accountId).CloudType(cloudType).ConnectivityType(connectivityType).Refresh(refresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetConnectivityGatewayCreationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectivityGatewayCreationOptions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetConnectivityGatewayCreationOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectivityGatewayCreationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account Id | 
 **cloudType** | **string** | Cloud Type | 
 **connectivityType** | **string** | Cloud Connectivity Type | 
 **refresh** | **string** | Refresh | 

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


## GetConnectivityGateways

> map[string]interface{} GetConnectivityGateways(ctx).AccountId(accountId).CloudType(cloudType).ConnectivityType(connectivityType).ConnectivityGatewayName(connectivityGatewayName).Region(region).Network(network).State(state).Refresh(refresh).EdgeType(edgeType).Execute()





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
    accountId := "accountId_example" // string | Account Id (optional)
    cloudType := "cloudType_example" // string | Cloud Type (optional)
    connectivityType := "connectivityType_example" // string | Cloud Connectivity Type (optional)
    connectivityGatewayName := "connectivityGatewayName_example" // string | Connectivity Gateway Name (optional)
    region := "region_example" // string | Region (optional)
    network := "network_example" // string | Network (optional)
    state := "state_example" // string | State (optional)
    refresh := "refresh_example" // string | Refresh (optional)
    edgeType := "edgeType_example" // string | Edge type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetConnectivityGateways(context.Background()).AccountId(accountId).CloudType(cloudType).ConnectivityType(connectivityType).ConnectivityGatewayName(connectivityGatewayName).Region(region).Network(network).State(state).Refresh(refresh).EdgeType(edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetConnectivityGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectivityGateways`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetConnectivityGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectivityGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account Id | 
 **cloudType** | **string** | Cloud Type | 
 **connectivityType** | **string** | Cloud Connectivity Type | 
 **connectivityGatewayName** | **string** | Connectivity Gateway Name | 
 **region** | **string** | Region | 
 **network** | **string** | Network | 
 **state** | **string** | State | 
 **refresh** | **string** | Refresh | 
 **edgeType** | **string** | Edge type | 

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


## GetDashboardEdgeInfo

> map[string]interface{} GetDashboardEdgeInfo(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetDashboardEdgeInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetDashboardEdgeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardEdgeInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetDashboardEdgeInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardEdgeInfoRequest struct via the builder pattern


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


## GetDefaultMappingValues

> map[string]interface{} GetDefaultMappingValues(ctx).CloudType(cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetDefaultMappingValues(context.Background()).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetDefaultMappingValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultMappingValues`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetDefaultMappingValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultMappingValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 

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


## GetDeviceLinks

> map[string]interface{} GetDeviceLinks(ctx).EdgeType(edgeType).DeviceLinkName(deviceLinkName).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)
    deviceLinkName := "deviceLinkName_example" // string | Device Link Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetDeviceLinks(context.Background()).EdgeType(edgeType).DeviceLinkName(deviceLinkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetDeviceLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLinks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetDeviceLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 
 **deviceLinkName** | **string** | Device Link Name | 

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


## GetDlPortSpeed

> map[string]interface{} GetDlPortSpeed(ctx, edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Interconnect Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetDlPortSpeed(context.Background(), edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetDlPortSpeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDlPortSpeed`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetDlPortSpeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Interconnect Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDlPortSpeedRequest struct via the builder pattern


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


## GetEdgeAccountDetails

> map[string]interface{} GetEdgeAccountDetails(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | Edge Account Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeAccountDetails(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeAccountDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeAccountDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeAccountDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Edge Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeAccountDetailsRequest struct via the builder pattern


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


## GetEdgeAccounts

> map[string]interface{} GetEdgeAccounts(ctx).EdgeType(edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeAccounts(context.Background()).EdgeType(edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 

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


## GetEdgeBillingAccounts

> map[string]interface{} GetEdgeBillingAccounts(ctx, edgeType, edgeAccountId).Region(region).Execute()





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
    edgeType := "edgeType_example" // string | Interconnect Provider
    edgeAccountId := "edgeAccountId_example" // string | Interconnect Provider Account ID
    region := "region_example" // string | Region (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeBillingAccounts(context.Background(), edgeType, edgeAccountId).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeBillingAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeBillingAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeBillingAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Interconnect Provider | 
**edgeAccountId** | **string** | Interconnect Provider Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeBillingAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **region** | **string** | Region | 

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


## GetEdgeConnectivityDetailByName

> map[string]interface{} GetEdgeConnectivityDetailByName(ctx, connectivityName).Execute()





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
    connectivityName := "connectivityName_example" // string | IC-GW connectivity name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeConnectivityDetailByName(context.Background(), connectivityName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeConnectivityDetailByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeConnectivityDetailByName`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeConnectivityDetailByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectivityName** | **string** | IC-GW connectivity name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeConnectivityDetailByNameRequest struct via the builder pattern


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


## GetEdgeConnectivityDetails

> map[string]interface{} GetEdgeConnectivityDetails(ctx).EdgeType(edgeType).ConnectivityName(connectivityName).ConnectivityType(connectivityType).EdgeGatewayName(edgeGatewayName).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)
    connectivityName := "connectivityName_example" // string | Connectivity Name (optional)
    connectivityType := "connectivityType_example" // string | Connectivity Type (optional)
    edgeGatewayName := "edgeGatewayName_example" // string | Interconnect Gateway name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeConnectivityDetails(context.Background()).EdgeType(edgeType).ConnectivityName(connectivityName).ConnectivityType(connectivityType).EdgeGatewayName(edgeGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeConnectivityDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeConnectivityDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeConnectivityDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeConnectivityDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 
 **connectivityName** | **string** | Connectivity Name | 
 **connectivityType** | **string** | Connectivity Type | 
 **edgeGatewayName** | **string** | Interconnect Gateway name | 

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


## GetEdgeGateways

> map[string]interface{} GetEdgeGateways(ctx, edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeGateways(context.Background(), edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeGateways`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeGatewaysRequest struct via the builder pattern


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


## GetEdgeGlobalSettings

> map[string]interface{} GetEdgeGlobalSettings(ctx).EdgeType(edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeGlobalSettings(context.Background()).EdgeType(edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeGlobalSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeGlobalSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 

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


## GetEdgeLocationsInfo

> map[string]interface{} GetEdgeLocationsInfo(ctx, edgeType).AccountId(accountId).Region(region).Execute()





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
    edgeType := "edgeType_example" // string | Edge Type (default to "MEGAPORT")
    accountId := "accountId_example" // string | Edge Account Id (optional)
    region := "region_example" // string | Region (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeLocationsInfo(context.Background(), edgeType).AccountId(accountId).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeLocationsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeLocationsInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeLocationsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge Type | [default to &quot;MEGAPORT&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeLocationsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | Edge Account Id | 
 **region** | **string** | Region | 

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


## GetEdgeMappingTags

> map[string]interface{} GetEdgeMappingTags(ctx).CloudType(cloudType).AccountId(accountId).ResourceGroup(resourceGroup).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accountId := "accountId_example" // string | Cloud Account Id (optional)
    resourceGroup := "resourceGroup_example" // string | Resource Group (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeMappingTags(context.Background()).CloudType(cloudType).AccountId(accountId).ResourceGroup(resourceGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeMappingTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeMappingTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeMappingTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeMappingTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accountId** | **string** | Cloud Account Id | 
 **resourceGroup** | **string** | Resource Group | 

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


## GetEdgeTypes

> map[string]interface{} GetEdgeTypes(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeTypesRequest struct via the builder pattern


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


## GetEdgeWanDevices

> map[string]interface{} GetEdgeWanDevices(ctx, edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge Type (default to "MEGAPORT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeWanDevices(context.Background(), edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeWanDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeWanDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeWanDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge Type | [default to &quot;MEGAPORT&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeWanDevicesRequest struct via the builder pattern


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


## GetEdgeWidget

> map[string]interface{} GetEdgeWidget(ctx, edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetEdgeWidget(context.Background(), edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetEdgeWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdgeWidget`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetEdgeWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgeWidgetRequest struct via the builder pattern


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


## GetGlobalSettings

> map[string]interface{} GetGlobalSettings(ctx).CloudType(cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetGlobalSettings(context.Background()).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetGlobalSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 

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


## GetHostVpcs

> map[string]interface{} GetHostVpcs(ctx).CloudType(cloudType).AccountIds(accountIds).Region(region).Untagged(untagged).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accountIds := "accountIds_example" // string | Account Id (optional)
    region := "region_example" // string | Region (optional)
    untagged := "untagged_example" // string | Untagged flag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetHostVpcs(context.Background()).CloudType(cloudType).AccountIds(accountIds).Region(region).Untagged(untagged).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetHostVpcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostVpcs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetHostVpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accountIds** | **string** | Account Id | 
 **region** | **string** | Region | 
 **untagged** | **string** | Untagged flag | 

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


## GetIcgwCustomSettingDetails

> map[string]interface{} GetIcgwCustomSettingDetails(ctx, edgeGatewayName).Execute()





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
    edgeGatewayName := "edgeGatewayName_example" // string | Edge gateway name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetIcgwCustomSettingDetails(context.Background(), edgeGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetIcgwCustomSettingDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIcgwCustomSettingDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetIcgwCustomSettingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeGatewayName** | **string** | Edge gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIcgwCustomSettingDetailsRequest struct via the builder pattern


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


## GetIcgwDetails

> map[string]interface{} GetIcgwDetails(ctx, edgeGatewayName).Execute()





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
    edgeGatewayName := "edgeGatewayName_example" // string | Edge gateway name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetIcgwDetails(context.Background(), edgeGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetIcgwDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIcgwDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetIcgwDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeGatewayName** | **string** | Edge gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIcgwDetailsRequest struct via the builder pattern


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


## GetIcgwTypes

> map[string]interface{} GetIcgwTypes(ctx).EdgeType(edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetIcgwTypes(context.Background()).EdgeType(edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetIcgwTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIcgwTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetIcgwTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIcgwTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 

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


## GetIcgws

> map[string]interface{} GetIcgws(ctx).EdgeType(edgeType).AccountId(accountId).Region(region).RegionId(regionId).ResourceState(resourceState).EdgeGatewayName(edgeGatewayName).BillingAccountId(billingAccountId).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)
    accountId := "accountId_example" // string | Account Id (optional)
    region := "region_example" // string | Region (optional)
    regionId := "regionId_example" // string | Region Id (optional)
    resourceState := "resourceState_example" // string | Resource State (optional)
    edgeGatewayName := "edgeGatewayName_example" // string | Edge gateway name (optional)
    billingAccountId := "billingAccountId_example" // string | billing Account Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetIcgws(context.Background()).EdgeType(edgeType).AccountId(accountId).Region(region).RegionId(regionId).ResourceState(resourceState).EdgeGatewayName(edgeGatewayName).BillingAccountId(billingAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetIcgws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIcgws`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetIcgws`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIcgwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 
 **accountId** | **string** | Account Id | 
 **region** | **string** | Region | 
 **regionId** | **string** | Region Id | 
 **resourceState** | **string** | Resource State | 
 **edgeGatewayName** | **string** | Edge gateway name | 
 **billingAccountId** | **string** | billing Account Id | 

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


## GetMappingMatrix

> map[string]interface{} GetMappingMatrix(ctx).CloudType(cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetMappingMatrix(context.Background()).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetMappingMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappingMatrix`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetMappingMatrix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 

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


## GetMappingStatus

> map[string]interface{} GetMappingStatus(ctx).CloudType(cloudType).Region(region).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type (optional)
    region := "region_example" // string | Region (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetMappingStatus(context.Background()).CloudType(cloudType).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetMappingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappingStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetMappingStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **region** | **string** | Region | 

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


## GetMappingSummary

> map[string]interface{} GetMappingSummary(ctx).VpnTunnelStatus(vpnTunnelStatus).CloudType(cloudType).Execute()





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
    vpnTunnelStatus := true // bool | VPN tunnel status (optional)
    cloudType := "cloudType_example" // string | Cloud type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetMappingSummary(context.Background()).VpnTunnelStatus(vpnTunnelStatus).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetMappingSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappingSummary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetMappingSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnTunnelStatus** | **bool** | VPN tunnel status | 
 **cloudType** | **string** | Cloud type | 

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


## GetMappingTags

> map[string]interface{} GetMappingTags(ctx).CloudType(cloudType).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetMappingTags(context.Background()).CloudType(cloudType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetMappingTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappingTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetMappingTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 

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


## GetMappingVpns

> map[string]interface{} GetMappingVpns(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetMappingVpns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetMappingVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappingVpns`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetMappingVpns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingVpnsRequest struct via the builder pattern


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


## GetNvaSecurityRules

> map[string]interface{} GetNvaSecurityRules(ctx, cloudGatewayName).Body(body).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    body := map[string]interface{}{ ... } // map[string]interface{} | Get NVA security Rules (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetNvaSecurityRules(context.Background(), cloudGatewayName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetNvaSecurityRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNvaSecurityRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetNvaSecurityRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNvaSecurityRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Get NVA security Rules | 

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


## GetPartnerPorts

> map[string]interface{} GetPartnerPorts(ctx).EdgeType(edgeType).AccountId(accountId).CloudType(cloudType).ConnectType(connectType).VxcPermitted(vxcPermitted).AuthorizationKey(authorizationKey).Refresh(refresh).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional)
    accountId := "accountId_example" // string | Edge Account Id (optional)
    cloudType := "cloudType_example" // string | Cloud Type (optional)
    connectType := "connectType_example" // string | Connect Type filter (optional)
    vxcPermitted := "vxcPermitted_example" // string | VXC Permitted on the port (optional)
    authorizationKey := "authorizationKey_example" // string | authorization Key (optional)
    refresh := "refresh_example" // string | Refresh (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetPartnerPorts(context.Background()).EdgeType(edgeType).AccountId(accountId).CloudType(cloudType).ConnectType(connectType).VxcPermitted(vxcPermitted).AuthorizationKey(authorizationKey).Refresh(refresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetPartnerPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartnerPorts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetPartnerPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | 
 **accountId** | **string** | Edge Account Id | 
 **cloudType** | **string** | Cloud Type | 
 **connectType** | **string** | Connect Type filter | 
 **vxcPermitted** | **string** | VXC Permitted on the port | 
 **authorizationKey** | **string** | authorization Key | 
 **refresh** | **string** | Refresh | 

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


## GetPortSpeed

> map[string]interface{} GetPortSpeed(ctx, edgeType, edgeAccountId, connectivityType).CloudType(cloudType).CloudAccountId(cloudAccountId).ConnectType(connectType).ConnectSubType(connectSubType).ConnectivityGateway(connectivityGateway).PartnerPort(partnerPort).Execute()





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
    edgeType := "edgeType_example" // string | Interconnect Provider
    edgeAccountId := "edgeAccountId_example" // string | Interconnect Provider Account ID
    connectivityType := "connectivityType_example" // string | Interconnect Connectivity Type
    cloudType := "cloudType_example" // string | Cloud Service Provider (optional)
    cloudAccountId := "cloudAccountId_example" // string | Cloud Service Provider Account ID (optional)
    connectType := "connectType_example" // string | Connection Type filter (optional)
    connectSubType := "connectSubType_example" // string | Connection Sub-Type filter (optional)
    connectivityGateway := "connectivityGateway_example" // string | Connectivity Gateway (optional)
    partnerPort := "partnerPort_example" // string | partnerPort (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetPortSpeed(context.Background(), edgeType, edgeAccountId, connectivityType).CloudType(cloudType).CloudAccountId(cloudAccountId).ConnectType(connectType).ConnectSubType(connectSubType).ConnectivityGateway(connectivityGateway).PartnerPort(partnerPort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetPortSpeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPortSpeed`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetPortSpeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Interconnect Provider | 
**edgeAccountId** | **string** | Interconnect Provider Account ID | 
**connectivityType** | **string** | Interconnect Connectivity Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortSpeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cloudType** | **string** | Cloud Service Provider | 
 **cloudAccountId** | **string** | Cloud Service Provider Account ID | 
 **connectType** | **string** | Connection Type filter | 
 **connectSubType** | **string** | Connection Sub-Type filter | 
 **connectivityGateway** | **string** | Connectivity Gateway | 
 **partnerPort** | **string** | partnerPort | 

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


## GetPostAggregationDataByQuery25

> map[string]interface{} GetPostAggregationDataByQuery25(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Stats query string (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetPostAggregationDataByQuery25(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetPostAggregationDataByQuery25``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostAggregationDataByQuery25`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetPostAggregationDataByQuery25`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPostAggregationDataByQuery25Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Stats query string | 

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


## GetSites

> map[string]interface{} GetSites(ctx).Color(color).Attached(attached).Execute()





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
    color := "color_example" // string | Color (optional)
    attached := true // bool | Is endpoint attached (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSites(context.Background()).Color(color).Attached(attached).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSites`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **color** | **string** | Color | 
 **attached** | **bool** | Is endpoint attached | 

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


## GetSites1

> map[string]interface{} GetSites1(ctx, cloudGatewayName).SystemIp(systemIp).SiteId(siteId).Color(color).VpnTunnelStatus(vpnTunnelStatus).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    systemIp := "systemIp_example" // string | System IP (optional)
    siteId := "siteId_example" // string | Site Id (optional)
    color := "color_example" // string | Color (optional)
    vpnTunnelStatus := true // bool | Fetch vpnTunnelStatus (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSites1(context.Background(), cloudGatewayName).SystemIp(systemIp).SiteId(siteId).Color(color).VpnTunnelStatus(vpnTunnelStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSites1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSites1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSites1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSites1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemIp** | **string** | System IP | 
 **siteId** | **string** | Site Id | 
 **color** | **string** | Color | 
 **vpnTunnelStatus** | **bool** | Fetch vpnTunnelStatus | 

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


## GetSshKeyList

> map[string]interface{} GetSshKeyList(ctx).CloudType(cloudType).AccountId(accountId).CloudRegion(cloudRegion).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accountId := "accountId_example" // string | Account Id
    cloudRegion := "cloudRegion_example" // string | Region

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSshKeyList(context.Background()).CloudType(cloudType).AccountId(accountId).CloudRegion(cloudRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSshKeyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSshKeyList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSshKeyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accountId** | **string** | Account Id | 
 **cloudRegion** | **string** | Region | 

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


## GetSupportedEdgeImageNames

> map[string]interface{} GetSupportedEdgeImageNames(ctx).EdgeType(edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional) (default to "MEGAPORT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSupportedEdgeImageNames(context.Background()).EdgeType(edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSupportedEdgeImageNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedEdgeImageNames`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSupportedEdgeImageNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedEdgeImageNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | [default to &quot;MEGAPORT&quot;]

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


## GetSupportedEdgeInstanceSize

> map[string]interface{} GetSupportedEdgeInstanceSize(ctx).EdgeType(edgeType).Execute()





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
    edgeType := "edgeType_example" // string | Edge type (optional) (default to "MEGAPORT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSupportedEdgeInstanceSize(context.Background()).EdgeType(edgeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSupportedEdgeInstanceSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedEdgeInstanceSize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSupportedEdgeInstanceSize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedEdgeInstanceSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeType** | **string** | Edge type | [default to &quot;MEGAPORT&quot;]

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


## GetSupportedInstanceSize

> map[string]interface{} GetSupportedInstanceSize(ctx).CloudType(cloudType).AccountId(accountId).CloudRegion(cloudRegion).Execute()





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
    cloudType := "cloudType_example" // string | 
    accountId := "accountId_example" // string |  (optional)
    cloudRegion := "cloudRegion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSupportedInstanceSize(context.Background()).CloudType(cloudType).AccountId(accountId).CloudRegion(cloudRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSupportedInstanceSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedInstanceSize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSupportedInstanceSize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedInstanceSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** |  | 
 **accountId** | **string** |  | 
 **cloudRegion** | **string** |  | 

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


## GetSupportedLoopbackCgwColors

> map[string]interface{} GetSupportedLoopbackCgwColors(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSupportedLoopbackCgwColors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSupportedLoopbackCgwColors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedLoopbackCgwColors`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSupportedLoopbackCgwColors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedLoopbackCgwColorsRequest struct via the builder pattern


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


## GetSupportedSoftwareImageList

> map[string]interface{} GetSupportedSoftwareImageList(ctx).CloudType(cloudType).AccountId(accountId).CloudRegion(cloudRegion).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    accountId := "accountId_example" // string | Account Id (optional)
    cloudRegion := "cloudRegion_example" // string | Region (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetSupportedSoftwareImageList(context.Background()).CloudType(cloudType).AccountId(accountId).CloudRegion(cloudRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetSupportedSoftwareImageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedSoftwareImageList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetSupportedSoftwareImageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedSoftwareImageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **accountId** | **string** | Account Id | 
 **cloudRegion** | **string** | Region | 

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


## GetTunnelNames

> map[string]interface{} GetTunnelNames(ctx, cloudType).CloudGatewayName(cloudGatewayName).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetTunnelNames(context.Background(), cloudType).CloudGatewayName(cloudGatewayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetTunnelNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTunnelNames`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetTunnelNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudType** | **string** | Cloud type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTunnelNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudGatewayName** | **string** | Cloud gateway name | 

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


## GetVHubs

> map[string]interface{} GetVHubs(ctx).CloudType(cloudType).AccountId(accountId).ResourceGroup(resourceGroup).VWanName(vWanName).VNetTags(vNetTags).Execute()





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
    cloudType := "cloudType_example" // string | Cloud Type (optional)
    accountId := "accountId_example" // string | Account Id (optional)
    resourceGroup := "resourceGroup_example" // string | Resource Group (optional)
    vWanName := "vWanName_example" // string | VWan Name (optional)
    vNetTags := "vNetTags_example" // string | VNet Tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetVHubs(context.Background()).CloudType(cloudType).AccountId(accountId).ResourceGroup(resourceGroup).VWanName(vWanName).VNetTags(vNetTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetVHubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVHubs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetVHubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVHubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud Type | 
 **accountId** | **string** | Account Id | 
 **resourceGroup** | **string** | Resource Group | 
 **vWanName** | **string** | VWan Name | 
 **vNetTags** | **string** | VNet Tags | 

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


## GetVWans

> map[string]interface{} GetVWans(ctx).AccountId(accountId).CloudType(cloudType).ResourceGroup(resourceGroup).Refresh(refresh).Execute()





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
    accountId := "accountId_example" // string | Account Id (optional)
    cloudType := "cloudType_example" // string | Cloud Type (optional)
    resourceGroup := "resourceGroup_example" // string | Resource Group (optional)
    refresh := "refresh_example" // string | Refresh (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetVWans(context.Background()).AccountId(accountId).CloudType(cloudType).ResourceGroup(resourceGroup).Refresh(refresh).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetVWans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVWans`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetVWans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVWansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account Id | 
 **cloudType** | **string** | Cloud Type | 
 **resourceGroup** | **string** | Resource Group | 
 **refresh** | **string** | Refresh | 

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


## GetVpcTags

> map[string]interface{} GetVpcTags(ctx).CloudType(cloudType).Region(region).TagName(tagName).Execute()





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
    cloudType := "cloudType_example" // string | Cloud type
    region := "region_example" // string | Region (optional)
    tagName := "tagName_example" // string | Tag name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetVpcTags(context.Background()).CloudType(cloudType).Region(region).TagName(tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetVpcTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVpcTags`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetVpcTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudType** | **string** | Cloud type | 
 **region** | **string** | Region | 
 **tagName** | **string** | Tag name | 

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


## GetWanDevices

> map[string]interface{} GetWanDevices(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetWanDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetWanDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetWanDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanDevicesRequest struct via the builder pattern


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


## GetWanInterfaceColors

> map[string]interface{} GetWanInterfaceColors(ctx).Execute()





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
    resp, r, err := apiClient.ConfigurationMultiCloudApi.GetWanInterfaceColors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.GetWanInterfaceColors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWanInterfaceColors`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.GetWanInterfaceColors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWanInterfaceColorsRequest struct via the builder pattern


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


## HostvpcTagging

> map[string]interface{} HostvpcTagging(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | VPC tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.HostvpcTagging(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.HostvpcTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HostvpcTagging`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.HostvpcTagging`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostvpcTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | VPC tag | 

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


## ProcessMapping

> map[string]interface{} ProcessMapping(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | VPC mapping (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.ProcessMapping(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.ProcessMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessMapping`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.ProcessMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | VPC mapping | 

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


## Telemetry

> map[string]interface{} Telemetry(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | telemetry (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.Telemetry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.Telemetry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Telemetry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.Telemetry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTelemetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | telemetry | 

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


## TunnelScaling

> map[string]interface{} TunnelScaling(ctx, cloudGatewayName).Body(body).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    body := map[string]interface{}{ ... } // map[string]interface{} | Site information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.TunnelScaling(context.Background(), cloudGatewayName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.TunnelScaling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TunnelScaling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.TunnelScaling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTunnelScalingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Site information | 

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


## UnTag

> map[string]interface{} UnTag(ctx, tagName).Execute()





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
    tagName := "tagName_example" // string | Tag name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UnTag(context.Background(), tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UnTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnTag`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UnTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | Tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnTagRequest struct via the builder pattern


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


## UpdateAccount

> UpdateAccount(ctx, accountId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Multicloud account info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateAccount(context.Background(), accountId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Multicloud account info | 

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


## UpdateCgw

> map[string]interface{} UpdateCgw(ctx, cloudGatewayName).Body(body).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | Cloud gateway (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateCgw(context.Background(), cloudGatewayName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateCgw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCgw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UpdateCgw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCgwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Cloud gateway | 

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


## UpdateDeviceLink

> map[string]interface{} UpdateDeviceLink(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Device Link (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateDeviceLink(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateDeviceLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceLink`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UpdateDeviceLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Device Link | 

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


## UpdateEdgeAccount

> UpdateEdgeAccount(ctx, accountId).Body(body).Execute()





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
    accountId := "accountId_example" // string | Multicloud Edge Account Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Multicloud edge account info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateEdgeAccount(context.Background(), accountId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateEdgeAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Multicloud Edge Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Multicloud edge account info | 

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


## UpdateEdgeConnectivity

> map[string]interface{} UpdateEdgeConnectivity(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Edge connectivity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateEdgeConnectivity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateEdgeConnectivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEdgeConnectivity`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UpdateEdgeConnectivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Edge connectivity | 

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


## UpdateEdgeGlobalSettings

> UpdateEdgeGlobalSettings(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Global setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateEdgeGlobalSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateEdgeGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Global setting | 

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


## UpdateEdgeLocationsInfo

> map[string]interface{} UpdateEdgeLocationsInfo(ctx, edgeType, accountId).Execute()





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
    edgeType := "edgeType_example" // string | Edge Type (default to "MEGAPORT")
    accountId := "accountId_example" // string | Edge Account Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateEdgeLocationsInfo(context.Background(), edgeType, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateEdgeLocationsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEdgeLocationsInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UpdateEdgeLocationsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeType** | **string** | Edge Type | [default to &quot;MEGAPORT&quot;]
**accountId** | **string** | Edge Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeLocationsInfoRequest struct via the builder pattern


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


## UpdateGlobalSettings

> UpdateGlobalSettings(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Global setting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateGlobalSettings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Global setting | 

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


## UpdateIcgw

> map[string]interface{} UpdateIcgw(ctx, edgeGatewayName).RequestBody(requestBody).Execute()





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
    edgeGatewayName := "edgeGatewayName_example" // string | Edge gateway name
    requestBody := map[string]GetO365PreferredPathFromVAnalyticsRequestValue{"key": *openapiclient.NewGetO365PreferredPathFromVAnalyticsRequestValue()} // map[string]GetO365PreferredPathFromVAnalyticsRequestValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateIcgw(context.Background(), edgeGatewayName).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateIcgw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIcgw`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UpdateIcgw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeGatewayName** | **string** | Edge gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIcgwRequest struct via the builder pattern


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


## UpdateNvaSecurityRules

> map[string]interface{} UpdateNvaSecurityRules(ctx, cloudGatewayName).Body(body).Execute()





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
    cloudGatewayName := "cloudGatewayName_example" // string | Cloud gateway name
    body := map[string]interface{}{ ... } // map[string]interface{} | Update NVA security Rules (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.UpdateNvaSecurityRules(context.Background(), cloudGatewayName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.UpdateNvaSecurityRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNvaSecurityRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationMultiCloudApi.UpdateNvaSecurityRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudGatewayName** | **string** | Cloud gateway name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNvaSecurityRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Update NVA security Rules | 

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


## ValidateAccountAdd

> ValidateAccountAdd(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Multicloud account info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.ValidateAccountAdd(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.ValidateAccountAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAccountAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Multicloud account info | 

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


## ValidateAccountUpdateCredentials

> ValidateAccountUpdateCredentials(ctx, accountId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Multicloud account info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.ValidateAccountUpdateCredentials(context.Background(), accountId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.ValidateAccountUpdateCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAccountUpdateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Multicloud account info | 

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


## ValidateEdgeAccountAdd

> ValidateEdgeAccountAdd(ctx).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} | Multicloud edge account info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.ValidateEdgeAccountAdd(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.ValidateEdgeAccountAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateEdgeAccountAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Multicloud edge account info | 

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


## ValidateEdgeAccountUpdateCredentials

> ValidateEdgeAccountUpdateCredentials(ctx, accountId).Body(body).Execute()





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
    accountId := "accountId_example" // string | Multicloud Edge Account Id
    body := map[string]interface{}{ ... } // map[string]interface{} | Multicloud edge account info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationMultiCloudApi.ValidateEdgeAccountUpdateCredentials(context.Background(), accountId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationMultiCloudApi.ValidateEdgeAccountUpdateCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Multicloud Edge Account Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateEdgeAccountUpdateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Multicloud edge account info | 

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

