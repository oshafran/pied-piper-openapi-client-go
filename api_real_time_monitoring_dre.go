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


// RealTimeMonitoringDREApiService RealTimeMonitoringDREApi service
type RealTimeMonitoringDREApiService service

type RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringDREApiService
	deviceId *string
	appqoeDreAutoBypassServerIp *string
	appqoeDreAutoBypassPort *float32
}

// Device IP
func (r RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest) DeviceId(deviceId string) RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest {
	r.deviceId = &deviceId
	return r
}

// Server IP
func (r RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest) AppqoeDreAutoBypassServerIp(appqoeDreAutoBypassServerIp string) RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest {
	r.appqoeDreAutoBypassServerIp = &appqoeDreAutoBypassServerIp
	return r
}

// Port
func (r RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest) AppqoeDreAutoBypassPort(appqoeDreAutoBypassPort float32) RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest {
	r.appqoeDreAutoBypassPort = &appqoeDreAutoBypassPort
	return r
}

func (r RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDreAutoBypassStatsExecute(r)
}

/*
GetDreAutoBypassStats Method for GetDreAutoBypassStats

Get DRE auto-bypass statistics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest
*/
func (a *RealTimeMonitoringDREApiService) GetDreAutoBypassStats(ctx context.Context) RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest {
	return RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringDREApiService) GetDreAutoBypassStatsExecute(r RealTimeMonitoringDREApiGetDreAutoBypassStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringDREApiService.GetDreAutoBypassStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/dre/auto-bypass-stats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
	}

	if r.appqoeDreAutoBypassServerIp != nil {
		localVarQueryParams.Add("appqoe-dre-auto-bypass-server-ip", parameterToString(*r.appqoeDreAutoBypassServerIp, ""))
	}
	if r.appqoeDreAutoBypassPort != nil {
		localVarQueryParams.Add("appqoe-dre-auto-bypass-port", parameterToString(*r.appqoeDreAutoBypassPort, ""))
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

type RealTimeMonitoringDREApiGetDrePeerStatsRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringDREApiService
	deviceId *string
	appqoeDreStatsPeerSystemIp *string
	appqoeDreStatsPeerPeerNo *float32
}

// Device IP
func (r RealTimeMonitoringDREApiGetDrePeerStatsRequest) DeviceId(deviceId string) RealTimeMonitoringDREApiGetDrePeerStatsRequest {
	r.deviceId = &deviceId
	return r
}

// System IP
func (r RealTimeMonitoringDREApiGetDrePeerStatsRequest) AppqoeDreStatsPeerSystemIp(appqoeDreStatsPeerSystemIp string) RealTimeMonitoringDREApiGetDrePeerStatsRequest {
	r.appqoeDreStatsPeerSystemIp = &appqoeDreStatsPeerSystemIp
	return r
}

// Peer Number
func (r RealTimeMonitoringDREApiGetDrePeerStatsRequest) AppqoeDreStatsPeerPeerNo(appqoeDreStatsPeerPeerNo float32) RealTimeMonitoringDREApiGetDrePeerStatsRequest {
	r.appqoeDreStatsPeerPeerNo = &appqoeDreStatsPeerPeerNo
	return r
}

func (r RealTimeMonitoringDREApiGetDrePeerStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDrePeerStatsExecute(r)
}

/*
GetDrePeerStats Method for GetDrePeerStats

Get DRE peer statistics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringDREApiGetDrePeerStatsRequest
*/
func (a *RealTimeMonitoringDREApiService) GetDrePeerStats(ctx context.Context) RealTimeMonitoringDREApiGetDrePeerStatsRequest {
	return RealTimeMonitoringDREApiGetDrePeerStatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringDREApiService) GetDrePeerStatsExecute(r RealTimeMonitoringDREApiGetDrePeerStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringDREApiService.GetDrePeerStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/dre/peer-stats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
	}

	if r.appqoeDreStatsPeerSystemIp != nil {
		localVarQueryParams.Add("appqoe-dre-stats-peer-system-ip", parameterToString(*r.appqoeDreStatsPeerSystemIp, ""))
	}
	if r.appqoeDreStatsPeerPeerNo != nil {
		localVarQueryParams.Add("appqoe-dre-stats-peer-peer-no", parameterToString(*r.appqoeDreStatsPeerPeerNo, ""))
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

type RealTimeMonitoringDREApiGetDreStatsRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringDREApiService
	deviceId *string
}

// Device IP
func (r RealTimeMonitoringDREApiGetDreStatsRequest) DeviceId(deviceId string) RealTimeMonitoringDREApiGetDreStatsRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringDREApiGetDreStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDreStatsExecute(r)
}

/*
GetDreStats Method for GetDreStats

Get DRE statistics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringDREApiGetDreStatsRequest
*/
func (a *RealTimeMonitoringDREApiService) GetDreStats(ctx context.Context) RealTimeMonitoringDREApiGetDreStatsRequest {
	return RealTimeMonitoringDREApiGetDreStatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringDREApiService) GetDreStatsExecute(r RealTimeMonitoringDREApiGetDreStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringDREApiService.GetDreStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/dre/dre-stats"

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

type RealTimeMonitoringDREApiGetDreStatusRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringDREApiService
	deviceId *string
}

// Device IP
func (r RealTimeMonitoringDREApiGetDreStatusRequest) DeviceId(deviceId string) RealTimeMonitoringDREApiGetDreStatusRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringDREApiGetDreStatusRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDreStatusExecute(r)
}

/*
GetDreStatus Method for GetDreStatus

Get DRE status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringDREApiGetDreStatusRequest
*/
func (a *RealTimeMonitoringDREApiService) GetDreStatus(ctx context.Context) RealTimeMonitoringDREApiGetDreStatusRequest {
	return RealTimeMonitoringDREApiGetDreStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringDREApiService) GetDreStatusExecute(r RealTimeMonitoringDREApiGetDreStatusRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringDREApiService.GetDreStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/dre/dre-status"

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
