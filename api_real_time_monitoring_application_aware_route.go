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


// RealTimeMonitoringApplicationAwareRouteApiService RealTimeMonitoringApplicationAwareRouteApi service
type RealTimeMonitoringApplicationAwareRouteApiService service

type RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringApplicationAwareRouteApiService
	deviceId *string
}

// Device IP
func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest) DeviceId(deviceId string) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateAppRouteSlaClassListExecute(r)
}

/*
CreateAppRouteSlaClassList Method for CreateAppRouteSlaClassList

Get SLA class list from device (Real Time)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest
*/
func (a *RealTimeMonitoringApplicationAwareRouteApiService) CreateAppRouteSlaClassList(ctx context.Context) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest {
	return RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringApplicationAwareRouteApiService) CreateAppRouteSlaClassListExecute(r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteSlaClassListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringApplicationAwareRouteApiService.CreateAppRouteSlaClassList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/app-route/sla-class"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
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

type RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringApplicationAwareRouteApiService
	deviceId *string
	remoteSystemIp *string
	localColor *string
	remoteColor *string
}

// Device IP
func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest) DeviceId(deviceId string) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest {
	r.deviceId = &deviceId
	return r
}

// Remote system IP
func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest) RemoteSystemIp(remoteSystemIp string) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest {
	r.remoteSystemIp = &remoteSystemIp
	return r
}

// Local color
func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest) LocalColor(localColor string) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest {
	r.localColor = &localColor
	return r
}

// Remote color
func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest) RemoteColor(remoteColor string) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest {
	r.remoteColor = &remoteColor
	return r
}

func (r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateAppRouteStatisticsListExecute(r)
}

/*
CreateAppRouteStatisticsList Method for CreateAppRouteStatisticsList

Get application-aware routing statistics from device (Real Time)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest
*/
func (a *RealTimeMonitoringApplicationAwareRouteApiService) CreateAppRouteStatisticsList(ctx context.Context) RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest {
	return RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringApplicationAwareRouteApiService) CreateAppRouteStatisticsListExecute(r RealTimeMonitoringApplicationAwareRouteApiCreateAppRouteStatisticsListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringApplicationAwareRouteApiService.CreateAppRouteStatisticsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/app-route/statistics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
	}

	if r.remoteSystemIp != nil {
		localVarQueryParams.Add("remote-system-ip", parameterToString(*r.remoteSystemIp, ""))
	}
	if r.localColor != nil {
		localVarQueryParams.Add("local-color", parameterToString(*r.localColor, ""))
	}
	if r.remoteColor != nil {
		localVarQueryParams.Add("remote-color", parameterToString(*r.remoteColor, ""))
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
