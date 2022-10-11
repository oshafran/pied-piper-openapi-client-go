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


// PartnerACIPolicyBuilderApiService PartnerACIPolicyBuilderApi service
type PartnerACIPolicyBuilderApiService service

type ApiCreateDscpMappingsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
	partnerId string
	body *map[string]interface{}
}

// ACI definition
func (r ApiCreateDscpMappingsRequest) Body(body map[string]interface{}) ApiCreateDscpMappingsRequest {
	r.body = &body
	return r
}

func (r ApiCreateDscpMappingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateDscpMappingsExecute(r)
}

/*
CreateDscpMappings Method for CreateDscpMappings

Create an ACI definition entry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return ApiCreateDscpMappingsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) CreateDscpMappings(ctx context.Context, partnerId string) ApiCreateDscpMappingsRequest {
	return ApiCreateDscpMappingsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) CreateDscpMappingsExecute(r ApiCreateDscpMappingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.CreateDscpMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/dscpmapping/{partnerId}"
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

type ApiDeleteDscpMappingsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
	partnerId string
}

func (r ApiDeleteDscpMappingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteDscpMappingsExecute(r)
}

/*
DeleteDscpMappings Method for DeleteDscpMappings

Delete DSCP mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return ApiDeleteDscpMappingsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) DeleteDscpMappings(ctx context.Context, partnerId string) ApiDeleteDscpMappingsRequest {
	return ApiDeleteDscpMappingsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) DeleteDscpMappingsExecute(r ApiDeleteDscpMappingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.DeleteDscpMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/dscpmapping/{partnerId}"
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

type ApiGetACIDefinitionsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
}

func (r ApiGetACIDefinitionsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetACIDefinitionsExecute(r)
}

/*
GetACIDefinitions Method for GetACIDefinitions

Get ACI definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetACIDefinitionsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) GetACIDefinitions(ctx context.Context) ApiGetACIDefinitionsRequest {
	return ApiGetACIDefinitionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) GetACIDefinitionsExecute(r ApiGetACIDefinitionsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.GetACIDefinitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy"

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

type ApiGetDataPrefixMappingsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
	partnerId string
}

func (r ApiGetDataPrefixMappingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDataPrefixMappingsExecute(r)
}

/*
GetDataPrefixMappings Method for GetDataPrefixMappings

Get prefix mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return ApiGetDataPrefixMappingsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) GetDataPrefixMappings(ctx context.Context, partnerId string) ApiGetDataPrefixMappingsRequest {
	return ApiGetDataPrefixMappingsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) GetDataPrefixMappingsExecute(r ApiGetDataPrefixMappingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.GetDataPrefixMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/prefixmapping/{partnerId}"
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

type ApiGetDataPrefixSequencesRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
}

func (r ApiGetDataPrefixSequencesRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDataPrefixSequencesExecute(r)
}

/*
GetDataPrefixSequences Method for GetDataPrefixSequences

Get data prefix sequence

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDataPrefixSequencesRequest
*/
func (a *PartnerACIPolicyBuilderApiService) GetDataPrefixSequences(ctx context.Context) ApiGetDataPrefixSequencesRequest {
	return ApiGetDataPrefixSequencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) GetDataPrefixSequencesExecute(r ApiGetDataPrefixSequencesRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.GetDataPrefixSequences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/sequences"

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

type ApiGetDscpMappingsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
	partnerId string
}

func (r ApiGetDscpMappingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDscpMappingsExecute(r)
}

/*
GetDscpMappings Method for GetDscpMappings

Get DSCP policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return ApiGetDscpMappingsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) GetDscpMappings(ctx context.Context, partnerId string) ApiGetDscpMappingsRequest {
	return ApiGetDscpMappingsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) GetDscpMappingsExecute(r ApiGetDscpMappingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.GetDscpMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/dscpmapping/{partnerId}"
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

type ApiGetEventsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
	partnerId string
	starttime *int64
	endtime *int64
}

// Start time
func (r ApiGetEventsRequest) Starttime(starttime int64) ApiGetEventsRequest {
	r.starttime = &starttime
	return r
}

// End time
func (r ApiGetEventsRequest) Endtime(endtime int64) ApiGetEventsRequest {
	r.endtime = &endtime
	return r
}

func (r ApiGetEventsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetEventsExecute(r)
}

/*
GetEvents Method for GetEvents

Get ACI events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return ApiGetEventsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) GetEvents(ctx context.Context, partnerId string) ApiGetEventsRequest {
	return ApiGetEventsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) GetEventsExecute(r ApiGetEventsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.GetEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/events/{partnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.starttime != nil {
		localVarQueryParams.Add("starttime", parameterToString(*r.starttime, ""))
	}
	if r.endtime != nil {
		localVarQueryParams.Add("endtime", parameterToString(*r.endtime, ""))
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

type ApiSetDataPrefixMappingsRequest struct {
	ctx context.Context
	ApiService *PartnerACIPolicyBuilderApiService
	partnerId string
	body *map[string]interface{}
}

// Prefix definition
func (r ApiSetDataPrefixMappingsRequest) Body(body map[string]interface{}) ApiSetDataPrefixMappingsRequest {
	r.body = &body
	return r
}

func (r ApiSetDataPrefixMappingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SetDataPrefixMappingsExecute(r)
}

/*
SetDataPrefixMappings Method for SetDataPrefixMappings

Create data prefix mapping

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId Partner Id
 @return ApiSetDataPrefixMappingsRequest
*/
func (a *PartnerACIPolicyBuilderApiService) SetDataPrefixMappings(ctx context.Context, partnerId string) ApiSetDataPrefixMappingsRequest {
	return ApiSetDataPrefixMappingsRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerACIPolicyBuilderApiService) SetDataPrefixMappingsExecute(r ApiSetDataPrefixMappingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerACIPolicyBuilderApiService.SetDataPrefixMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/aci/policy/prefixmapping/{partnerId}"
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
