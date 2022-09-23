# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.5
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://44.196.44.132*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Class1AuthenticationApi* | [**DataserviceClientTokenGet**](docs/Class1AuthenticationApi.md#dataserviceclienttokenget) | **Get** /dataservice/client/token | Token
*Class1AuthenticationApi* | [**JSecurityCheckPost**](docs/Class1AuthenticationApi.md#jsecuritycheckpost) | **Post** /j_security_check | Authentication
*Class2SDWANFabricDevicesApi* | [**DataserviceDeviceCountersGet**](docs/Class2SDWANFabricDevicesApi.md#dataservicedevicecountersget) | **Get** /dataservice/device/counters | Device Counters
*Class2SDWANFabricDevicesApi* | [**DataserviceDeviceGet**](docs/Class2SDWANFabricDevicesApi.md#dataservicedeviceget) | **Get** /dataservice/device | Fabric Devices
*Class2SDWANFabricDevicesApi* | [**DataserviceDeviceMonitorGet**](docs/Class2SDWANFabricDevicesApi.md#dataservicedevicemonitorget) | **Get** /dataservice/device/monitor | Devices Status
*Class2SDWANFabricDevicesApi* | [**DataserviceStatisticsInterfaceGet**](docs/Class2SDWANFabricDevicesApi.md#dataservicestatisticsinterfaceget) | **Get** /dataservice/statistics/interface | Interface statistics
*Class3SDWANDeviceTemplateApi* | [**DataserviceTemplateFeatureGet**](docs/Class3SDWANDeviceTemplateApi.md#dataservicetemplatefeatureget) | **Get** /dataservice/template/feature | Template Feature
*Class3SDWANDeviceTemplateApi* | [**DataserviceTemplateFeatureTypesGet**](docs/Class3SDWANDeviceTemplateApi.md#dataservicetemplatefeaturetypesget) | **Get** /dataservice/template/feature/types | Template Feature Type
*Class4SDWANDevicePolicyApi* | [**DataserviceTemplatePolicyListGet**](docs/Class4SDWANDevicePolicyApi.md#dataservicetemplatepolicylistget) | **Get** /dataservice/template/policy/list | Policy List
*Class4SDWANDevicePolicyApi* | [**DataserviceTemplatePolicyVedgeDevicesGet**](docs/Class4SDWANDevicePolicyApi.md#dataservicetemplatepolicyvedgedevicesget) | **Get** /dataservice/template/policy/vedge/devices | vEdge Template Policy


## Documentation For Models



## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



