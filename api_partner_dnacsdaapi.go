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
	"strings"
)


// PartnerDNACSDAAPIApiService PartnerDNACSDAAPIApi service
type PartnerDNACSDAAPIApiService service

type PartnerDNACSDAAPIApiCreateSDAConfigRequest struct {
	ctx context.Context
	ApiService *PartnerDNACSDAAPIApiService
	partnerId string
	body *map[string]interface{}
}

// Device SDA configuration
func (r PartnerDNACSDAAPIApiCreateSDAConfigRequest) Body(body map[string]interface{}) PartnerDNACSDAAPIApiCreateSDAConfigRequest {
	r.body = &body
	return r
}

func (r PartnerDNACSDAAPIApiCreateSDAConfigRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateSDAConfigExecute(r)
}

/*
CreateSDAConfig Method for CreateSDAConfig

Create SDA enabled device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return PartnerDNACSDAAPIApiCreateSDAConfigRequest
*/
func (a *PartnerDNACSDAAPIApiService) CreateSDAConfig(ctx context.Context, partnerId string) PartnerDNACSDAAPIApiCreateSDAConfigRequest {
	return PartnerDNACSDAAPIApiCreateSDAConfigRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerDNACSDAAPIApiService) CreateSDAConfigExecute(r PartnerDNACSDAAPIApiCreateSDAConfigRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerDNACSDAAPIApiService.CreateSDAConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/dnac/sda/config/{partnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest struct {
	ctx context.Context
	ApiService *PartnerDNACSDAAPIApiService
	partnerId string
	body *map[string]interface{}
}

// Device SDA configuration
func (r PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest) Body(body map[string]interface{}) PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest {
	r.body = &body
	return r
}

func (r PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateSDAConfigFromNetconfExecute(r)
}

/*
CreateSDAConfigFromNetconf Method for CreateSDAConfigFromNetconf

Create SDA enabled device from Netconf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest
*/
func (a *PartnerDNACSDAAPIApiService) CreateSDAConfigFromNetconf(ctx context.Context, partnerId string) PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest {
	return PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerDNACSDAAPIApiService) CreateSDAConfigFromNetconfExecute(r PartnerDNACSDAAPIApiCreateSDAConfigFromNetconfRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerDNACSDAAPIApiService.CreateSDAConfigFromNetconf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/dnac/sda/netconfconfig/{partnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type PartnerDNACSDAAPIApiGetDeviceDetailsRequest struct {
	ctx context.Context
	ApiService *PartnerDNACSDAAPIApiService
	partnerId string
	uuid string
}

func (r PartnerDNACSDAAPIApiGetDeviceDetailsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceDetailsExecute(r)
}

/*
GetDeviceDetails Method for GetDeviceDetails

Get SDA enabled devices detail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @param uuid Device uuid
 @return PartnerDNACSDAAPIApiGetDeviceDetailsRequest
*/
func (a *PartnerDNACSDAAPIApiService) GetDeviceDetails(ctx context.Context, partnerId string, uuid string) PartnerDNACSDAAPIApiGetDeviceDetailsRequest {
	return PartnerDNACSDAAPIApiGetDeviceDetailsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PartnerDNACSDAAPIApiService) GetDeviceDetailsExecute(r PartnerDNACSDAAPIApiGetDeviceDetailsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerDNACSDAAPIApiService.GetDeviceDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/dnac/sda/device/{partnerId}/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterToString(r.uuid, "")), -1)

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

type PartnerDNACSDAAPIApiGetOverlayVPNListRequest struct {
	ctx context.Context
	ApiService *PartnerDNACSDAAPIApiService
}

func (r PartnerDNACSDAAPIApiGetOverlayVPNListRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetOverlayVPNListExecute(r)
}

/*
GetOverlayVPNList Method for GetOverlayVPNList

Get Overlay VPN list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PartnerDNACSDAAPIApiGetOverlayVPNListRequest
*/
func (a *PartnerDNACSDAAPIApiService) GetOverlayVPNList(ctx context.Context) PartnerDNACSDAAPIApiGetOverlayVPNListRequest {
	return PartnerDNACSDAAPIApiGetOverlayVPNListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PartnerDNACSDAAPIApiService) GetOverlayVPNListExecute(r PartnerDNACSDAAPIApiGetOverlayVPNListRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerDNACSDAAPIApiService.GetOverlayVPNList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/dnac/sda/vpn"

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

type PartnerDNACSDAAPIApiGetSDAEnabledDevicesRequest struct {
	ctx context.Context
	ApiService *PartnerDNACSDAAPIApiService
	partnerId string
}

func (r PartnerDNACSDAAPIApiGetSDAEnabledDevicesRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetSDAEnabledDevicesExecute(r)
}

/*
GetSDAEnabledDevices Method for GetSDAEnabledDevices

Get SDA enabled devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return PartnerDNACSDAAPIApiGetSDAEnabledDevicesRequest
*/
func (a *PartnerDNACSDAAPIApiService) GetSDAEnabledDevices(ctx context.Context, partnerId string) PartnerDNACSDAAPIApiGetSDAEnabledDevicesRequest {
	return PartnerDNACSDAAPIApiGetSDAEnabledDevicesRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PartnerDNACSDAAPIApiService) GetSDAEnabledDevicesExecute(r PartnerDNACSDAAPIApiGetSDAEnabledDevicesRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerDNACSDAAPIApiService.GetSDAEnabledDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/dnac/sda/device/{partnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

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

type PartnerDNACSDAAPIApiGetSitesForPartnerRequest struct {
	ctx context.Context
	ApiService *PartnerDNACSDAAPIApiService
	partnerId string
}

func (r PartnerDNACSDAAPIApiGetSitesForPartnerRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetSitesForPartnerExecute(r)
}

/*
GetSitesForPartner Method for GetSitesForPartner

Get SDA enabled devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return PartnerDNACSDAAPIApiGetSitesForPartnerRequest
*/
func (a *PartnerDNACSDAAPIApiService) GetSitesForPartner(ctx context.Context, partnerId string) PartnerDNACSDAAPIApiGetSitesForPartnerRequest {
	return PartnerDNACSDAAPIApiGetSitesForPartnerRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PartnerDNACSDAAPIApiService) GetSitesForPartnerExecute(r PartnerDNACSDAAPIApiGetSitesForPartnerRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerDNACSDAAPIApiService.GetSitesForPartner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/dnac/sda/site/{partnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

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
