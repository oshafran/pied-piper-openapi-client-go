/*
Cisco SD-WAN vManage API

The vManage API exposes the functionality of operations maintaining devices and the overlay network

API version: 2.0.0
Contact: vmanage@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// RealTimeMonitoringIPv6NeighboursApiService RealTimeMonitoringIPv6NeighboursApi service
type RealTimeMonitoringIPv6NeighboursApiService service

type ApiGetIpv6InterfaceRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringIPv6NeighboursApiService
	deviceId *string
	vpnId *string
	ifName *string
	mac *string
}

// Device IP
func (r ApiGetIpv6InterfaceRequest) DeviceId(deviceId string) ApiGetIpv6InterfaceRequest {
	r.deviceId = &deviceId
	return r
}

// VPN Id
func (r ApiGetIpv6InterfaceRequest) VpnId(vpnId string) ApiGetIpv6InterfaceRequest {
	r.vpnId = &vpnId
	return r
}

// Interface name
func (r ApiGetIpv6InterfaceRequest) IfName(ifName string) ApiGetIpv6InterfaceRequest {
	r.ifName = &ifName
	return r
}

// Mac address
func (r ApiGetIpv6InterfaceRequest) Mac(mac string) ApiGetIpv6InterfaceRequest {
	r.mac = &mac
	return r
}

func (r ApiGetIpv6InterfaceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetIpv6InterfaceExecute(r)
}

/*
GetIpv6Interface Method for GetIpv6Interface

Get IPv6 Neighbors from device (Real Time)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIpv6InterfaceRequest
*/
func (a *RealTimeMonitoringIPv6NeighboursApiService) GetIpv6Interface(ctx context.Context) ApiGetIpv6InterfaceRequest {
	return ApiGetIpv6InterfaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringIPv6NeighboursApiService) GetIpv6InterfaceExecute(r ApiGetIpv6InterfaceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringIPv6NeighboursApiService.GetIpv6Interface")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/ndv6"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
	}

	if r.vpnId != nil {
		localVarQueryParams.Add("vpn-id", parameterToString(*r.vpnId, ""))
	}
	if r.ifName != nil {
		localVarQueryParams.Add("if-name", parameterToString(*r.ifName, ""))
	}
	if r.mac != nil {
		localVarQueryParams.Add("mac", parameterToString(*r.mac, ""))
	}
	localVarQueryParams.Add("deviceId", parameterToString(*r.deviceId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}