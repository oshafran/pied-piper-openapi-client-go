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


// RealTimeMonitoringVMApiService RealTimeMonitoringVMApi service
type RealTimeMonitoringVMApiService service

type ApiGetCloudDockVMLifecycleNicsRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringVMApiService
	userGroup *string
}

// userGroup Name
func (r ApiGetCloudDockVMLifecycleNicsRequest) UserGroup(userGroup string) ApiGetCloudDockVMLifecycleNicsRequest {
	r.userGroup = &userGroup
	return r
}

func (r ApiGetCloudDockVMLifecycleNicsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetCloudDockVMLifecycleNicsExecute(r)
}

/*
GetCloudDockVMLifecycleNics Method for GetCloudDockVMLifecycleNics

Get CloudDock vm lifecycle state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCloudDockVMLifecycleNicsRequest
*/
func (a *RealTimeMonitoringVMApiService) GetCloudDockVMLifecycleNics(ctx context.Context) ApiGetCloudDockVMLifecycleNicsRequest {
	return ApiGetCloudDockVMLifecycleNicsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringVMApiService) GetCloudDockVMLifecycleNicsExecute(r ApiGetCloudDockVMLifecycleNicsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringVMApiService.GetCloudDockVMLifecycleNics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/vm/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userGroup == nil {
		return localVarReturnValue, nil, reportError("userGroup is required and must be specified")
	}

	localVarQueryParams.Add("userGroup", parameterToString(*r.userGroup, ""))
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

type ApiGetVMLifeCycleStateRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringVMApiService
	deviceId *string
}

// Device IP
func (r ApiGetVMLifeCycleStateRequest) DeviceId(deviceId string) ApiGetVMLifeCycleStateRequest {
	r.deviceId = &deviceId
	return r
}

func (r ApiGetVMLifeCycleStateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetVMLifeCycleStateExecute(r)
}

/*
GetVMLifeCycleState Method for GetVMLifeCycleState

Get vm lifecycle state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVMLifeCycleStateRequest
*/
func (a *RealTimeMonitoringVMApiService) GetVMLifeCycleState(ctx context.Context) ApiGetVMLifeCycleStateRequest {
	return ApiGetVMLifeCycleStateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringVMApiService) GetVMLifeCycleStateExecute(r ApiGetVMLifeCycleStateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringVMApiService.GetVMLifeCycleState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/vm/state"

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

type ApiGetVbranchVMLifecycleRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringVMApiService
	deviceId *string
}

// Device IP
func (r ApiGetVbranchVMLifecycleRequest) DeviceId(deviceId string) ApiGetVbranchVMLifecycleRequest {
	r.deviceId = &deviceId
	return r
}

func (r ApiGetVbranchVMLifecycleRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetVbranchVMLifecycleExecute(r)
}

/*
GetVbranchVMLifecycle Method for GetVbranchVMLifecycle

Get vbranch vm lifecycle state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVbranchVMLifecycleRequest
*/
func (a *RealTimeMonitoringVMApiService) GetVbranchVMLifecycle(ctx context.Context) ApiGetVbranchVMLifecycleRequest {
	return ApiGetVbranchVMLifecycleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringVMApiService) GetVbranchVMLifecycleExecute(r ApiGetVbranchVMLifecycleRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringVMApiService.GetVbranchVMLifecycle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/vm/oper/state"

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

type ApiGetVbranchVMLifecycleNicsRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringVMApiService
	deviceId *string
}

// Device IP
func (r ApiGetVbranchVMLifecycleNicsRequest) DeviceId(deviceId string) ApiGetVbranchVMLifecycleNicsRequest {
	r.deviceId = &deviceId
	return r
}

func (r ApiGetVbranchVMLifecycleNicsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetVbranchVMLifecycleNicsExecute(r)
}

/*
GetVbranchVMLifecycleNics Method for GetVbranchVMLifecycleNics

Get vbranch vm lifecycle state (NIC)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVbranchVMLifecycleNicsRequest
*/
func (a *RealTimeMonitoringVMApiService) GetVbranchVMLifecycleNics(ctx context.Context) ApiGetVbranchVMLifecycleNicsRequest {
	return ApiGetVbranchVMLifecycleNicsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringVMApiService) GetVbranchVMLifecycleNicsExecute(r ApiGetVbranchVMLifecycleNicsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringVMApiService.GetVbranchVMLifecycleNics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/vm/nics"

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