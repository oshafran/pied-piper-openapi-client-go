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
	"reflect"
)


// RealTimeMonitoringSystemApiService RealTimeMonitoringSystemApi service
type RealTimeMonitoringSystemApiService service

type RealTimeMonitoringSystemApiCreateDeviceInfoListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *[]string
}

// Device Id
func (r RealTimeMonitoringSystemApiCreateDeviceInfoListRequest) DeviceId(deviceId []string) RealTimeMonitoringSystemApiCreateDeviceInfoListRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceInfoListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceInfoListExecute(r)
}

/*
CreateDeviceInfoList Method for CreateDeviceInfoList

Get device information list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceInfoListRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceInfoList(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceInfoListRequest {
	return RealTimeMonitoringSystemApiCreateDeviceInfoListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceInfoListExecute(r RealTimeMonitoringSystemApiCreateDeviceInfoListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceInfoList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/system/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
	}

	{
		t := *r.deviceId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("deviceId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("deviceId", parameterToString(t, "multi"))
		}
	}
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

type RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device Id
func (r RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceSystemProcessListExecute(r)
}

/*
CreateDeviceSystemProcessList Method for CreateDeviceSystemProcessList

Get device system process list from device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemProcessList(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest {
	return RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemProcessListExecute(r RealTimeMonitoringSystemApiCreateDeviceSystemProcessListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceSystemProcessList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/csp/system/processlist"

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

type RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device Id
func (r RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceSystemSettingExecute(r)
}

/*
CreateDeviceSystemSetting Method for CreateDeviceSystemSetting

Get device system settings from device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemSetting(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest {
	return RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemSettingExecute(r RealTimeMonitoringSystemApiCreateDeviceSystemSettingRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceSystemSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/csp/system/settings"

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

type RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device Id
func (r RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceSystemSettingNativeInfoExecute(r)
}

/*
CreateDeviceSystemSettingNativeInfo Method for CreateDeviceSystemSettingNativeInfo

Get device system native settings from device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemSettingNativeInfo(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest {
	return RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemSettingNativeInfoExecute(r RealTimeMonitoringSystemApiCreateDeviceSystemSettingNativeInfoRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceSystemSettingNativeInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/csp/system/native"

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

type RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device IP
func (r RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceSystemStatsListExecute(r)
}

/*
CreateDeviceSystemStatsList Method for CreateDeviceSystemStatsList

Get device system stats list (Real Time)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemStatsList(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest {
	return RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemStatsListExecute(r RealTimeMonitoringSystemApiCreateDeviceSystemStatsListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceSystemStatsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/system/statistics"

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

type RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device Id
func (r RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceSystemStatusExecute(r)
}

/*
CreateDeviceSystemStatus Method for CreateDeviceSystemStatus

Get device system status from device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemStatus(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest {
	return RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemStatusExecute(r RealTimeMonitoringSystemApiCreateDeviceSystemStatusRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceSystemStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/csp/system/status"

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

type RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device IP
func (r RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDeviceSystemStatusListExecute(r)
}

/*
CreateDeviceSystemStatusList Method for CreateDeviceSystemStatusList

Get device system status list (Real Time)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemStatusList(ctx context.Context) RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest {
	return RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateDeviceSystemStatusListExecute(r RealTimeMonitoringSystemApiCreateDeviceSystemStatusListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateDeviceSystemStatusList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/system/status"

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

type RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringSystemApiService
	deviceId *string
}

// Device IP
func (r RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest) DeviceId(deviceId string) RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateSyncedDeviceSystemStatusListExecute(r)
}

/*
CreateSyncedDeviceSystemStatusList Method for CreateSyncedDeviceSystemStatusList

Get device system stats list synchronously

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest
*/
func (a *RealTimeMonitoringSystemApiService) CreateSyncedDeviceSystemStatusList(ctx context.Context) RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest {
	return RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringSystemApiService) CreateSyncedDeviceSystemStatusListExecute(r RealTimeMonitoringSystemApiCreateSyncedDeviceSystemStatusListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringSystemApiService.CreateSyncedDeviceSystemStatusList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/system/synced/status"

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
