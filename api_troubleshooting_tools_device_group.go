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


// TroubleshootingToolsDeviceGroupApiService TroubleshootingToolsDeviceGroupApi service
type TroubleshootingToolsDeviceGroupApiService service

type TroubleshootingToolsDeviceGroupApiListDeviceGroupListRequest struct {
	ctx context.Context
	ApiService *TroubleshootingToolsDeviceGroupApiService
}

func (r TroubleshootingToolsDeviceGroupApiListDeviceGroupListRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListDeviceGroupListExecute(r)
}

/*
ListDeviceGroupList Method for ListDeviceGroupList

Retrieve device group list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TroubleshootingToolsDeviceGroupApiListDeviceGroupListRequest
*/
func (a *TroubleshootingToolsDeviceGroupApiService) ListDeviceGroupList(ctx context.Context) TroubleshootingToolsDeviceGroupApiListDeviceGroupListRequest {
	return TroubleshootingToolsDeviceGroupApiListDeviceGroupListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TroubleshootingToolsDeviceGroupApiService) ListDeviceGroupListExecute(r TroubleshootingToolsDeviceGroupApiListDeviceGroupListRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TroubleshootingToolsDeviceGroupApiService.ListDeviceGroupList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type TroubleshootingToolsDeviceGroupApiListDeviceGroupsRequest struct {
	ctx context.Context
	ApiService *TroubleshootingToolsDeviceGroupApiService
}

func (r TroubleshootingToolsDeviceGroupApiListDeviceGroupsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListDeviceGroupsExecute(r)
}

/*
ListDeviceGroups Method for ListDeviceGroups

Retrieve device groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TroubleshootingToolsDeviceGroupApiListDeviceGroupsRequest
*/
func (a *TroubleshootingToolsDeviceGroupApiService) ListDeviceGroups(ctx context.Context) TroubleshootingToolsDeviceGroupApiListDeviceGroupsRequest {
	return TroubleshootingToolsDeviceGroupApiListDeviceGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TroubleshootingToolsDeviceGroupApiService) ListDeviceGroupsExecute(r TroubleshootingToolsDeviceGroupApiListDeviceGroupsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TroubleshootingToolsDeviceGroupApiService.ListDeviceGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group/device"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest struct {
	ctx context.Context
	ApiService *TroubleshootingToolsDeviceGroupApiService
	groupId *string
	ssh *bool
	vpnId *[]VPNID
}

// Group Id
func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest) GroupId(groupId string) TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest {
	r.groupId = &groupId
	return r
}

// SSH
func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest) Ssh(ssh bool) TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest {
	r.ssh = &ssh
	return r
}

// VPN Id
func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest) VpnId(vpnId []VPNID) TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest {
	r.vpnId = &vpnId
	return r
}

func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListGroupDevicesExecute(r)
}

/*
ListGroupDevices Method for ListGroupDevices

Retrieve devices in group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest
*/
func (a *TroubleshootingToolsDeviceGroupApiService) ListGroupDevices(ctx context.Context) TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest {
	return TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TroubleshootingToolsDeviceGroupApiService) ListGroupDevicesExecute(r TroubleshootingToolsDeviceGroupApiListGroupDevicesRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TroubleshootingToolsDeviceGroupApiService.ListGroupDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group/devices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupId == nil {
		return localVarReturnValue, nil, reportError("groupId is required and must be specified")
	}

	localVarQueryParams.Add("groupId", parameterToString(*r.groupId, ""))
	if r.ssh != nil {
		localVarQueryParams.Add("ssh", parameterToString(*r.ssh, ""))
	}
	if r.vpnId != nil {
		t := *r.vpnId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("vpnId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("vpnId", parameterToString(t, "multi"))
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

type TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest struct {
	ctx context.Context
	ApiService *TroubleshootingToolsDeviceGroupApiService
	groupId *string
	vpnId *[]VPNID
}

// Group Id
func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest) GroupId(groupId string) TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest {
	r.groupId = &groupId
	return r
}

// VPN Id
func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest) VpnId(vpnId []VPNID) TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest {
	r.vpnId = &vpnId
	return r
}

func (r TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListGroupDevicesForMapExecute(r)
}

/*
ListGroupDevicesForMap Method for ListGroupDevicesForMap

Retrieve group devices for map

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest
*/
func (a *TroubleshootingToolsDeviceGroupApiService) ListGroupDevicesForMap(ctx context.Context) TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest {
	return TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TroubleshootingToolsDeviceGroupApiService) ListGroupDevicesForMapExecute(r TroubleshootingToolsDeviceGroupApiListGroupDevicesForMapRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TroubleshootingToolsDeviceGroupApiService.ListGroupDevicesForMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group/map/devices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupId == nil {
		return localVarReturnValue, nil, reportError("groupId is required and must be specified")
	}

	localVarQueryParams.Add("groupId", parameterToString(*r.groupId, ""))
	if r.vpnId != nil {
		t := *r.vpnId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("vpnId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("vpnId", parameterToString(t, "multi"))
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

type TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest struct {
	ctx context.Context
	ApiService *TroubleshootingToolsDeviceGroupApiService
	groupId *string
}

// Group Id
func (r TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest) GroupId(groupId string) TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest {
	r.groupId = &groupId
	return r
}

func (r TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListGroupLinksForMapExecute(r)
}

/*
ListGroupLinksForMap Method for ListGroupLinksForMap

Retrieve devices in group for map

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest
*/
func (a *TroubleshootingToolsDeviceGroupApiService) ListGroupLinksForMap(ctx context.Context) TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest {
	return TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TroubleshootingToolsDeviceGroupApiService) ListGroupLinksForMapExecute(r TroubleshootingToolsDeviceGroupApiListGroupLinksForMapRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TroubleshootingToolsDeviceGroupApiService.ListGroupLinksForMap")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/group/map/devices/links"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupId == nil {
		return localVarReturnValue, nil, reportError("groupId is required and must be specified")
	}

	localVarQueryParams.Add("groupId", parameterToString(*r.groupId, ""))
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
