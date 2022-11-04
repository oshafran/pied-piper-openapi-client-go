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


// ConfigurationPolicyExtendedCommunityListBuilderApiService ConfigurationPolicyExtendedCommunityListBuilderApi service
type ConfigurationPolicyExtendedCommunityListBuilderApiService service

type ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request) Body(body map[string]interface{}) ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreatePolicyList12Execute(r)
}

/*
CreatePolicyList12 Method for CreatePolicyList12

Create policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) CreatePolicyList12(ctx context.Context) ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) CreatePolicyList12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiCreatePolicyList12Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.CreatePolicyList12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity"

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

type ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyList12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	id string
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyList12Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyList12Execute(r)
}

/*
DeletePolicyList12 Method for DeletePolicyList12

Delete policy list entry for a specific type of policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyList12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) DeletePolicyList12(ctx context.Context, id string) ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyList12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyList12Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) DeletePolicyList12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyList12Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.DeletePolicyList12")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	infoTag *string
}

// InfoTag
func (r ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request) InfoTag(infoTag string) ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request {
	r.infoTag = &infoTag
	return r
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeletePolicyListsWithInfoTag12Execute(r)
}

/*
DeletePolicyListsWithInfoTag12 Method for DeletePolicyListsWithInfoTag12

Delete policy lists with specific info tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) DeletePolicyListsWithInfoTag12(ctx context.Context) ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) DeletePolicyListsWithInfoTag12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiDeletePolicyListsWithInfoTag12Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.DeletePolicyListsWithInfoTag12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.infoTag != nil {
		localVarQueryParams.Add("infoTag", parameterToString(*r.infoTag, ""))
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

type ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request) Body(body map[string]interface{}) ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditPolicyList12Execute(r)
}

/*
EditPolicyList12 Method for EditPolicyList12

Edit policy list entries for a specific type of policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) EditPolicyList12(ctx context.Context, id string) ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) EditPolicyList12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiEditPolicyList12Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.EditPolicyList12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ConfigurationPolicyExtendedCommunityListBuilderApiGetListsById12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	id string
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiGetListsById12Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetListsById12Execute(r)
}

/*
GetListsById12 Method for GetListsById12

Get a specific policy list based on the id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyExtendedCommunityListBuilderApiGetListsById12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) GetListsById12(ctx context.Context, id string) ConfigurationPolicyExtendedCommunityListBuilderApiGetListsById12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiGetListsById12Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) GetListsById12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiGetListsById12Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.GetListsById12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyLists11Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyLists11Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyLists11Execute(r)
}

/*
GetPolicyLists11 Method for GetPolicyLists11

Get policy lists

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyLists11Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) GetPolicyLists11(ctx context.Context) ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyLists11Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyLists11Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) GetPolicyLists11Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyLists11Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.GetPolicyLists11")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity"

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

type ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	infoTag *string
}

// InfoTag
func (r ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request) InfoTag(infoTag string) ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request {
	r.infoTag = &infoTag
	return r
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyListsWithInfoTag12Execute(r)
}

/*
GetPolicyListsWithInfoTag12 Method for GetPolicyListsWithInfoTag12

Get policy lists with specific info tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) GetPolicyListsWithInfoTag12(ctx context.Context) ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) GetPolicyListsWithInfoTag12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiGetPolicyListsWithInfoTag12Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.GetPolicyListsWithInfoTag12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity/filtered"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.infoTag != nil {
		localVarQueryParams.Add("infoTag", parameterToString(*r.infoTag, ""))
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

type ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request) Body(body map[string]interface{}) ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyList12Execute(r)
}

/*
PreviewPolicyList12 Method for PreviewPolicyList12

Preview a policy list based on the policy list type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) PreviewPolicyList12(ctx context.Context) ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) PreviewPolicyList12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyList12Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.PreviewPolicyList12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity/preview"

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

type ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyListById12Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyExtendedCommunityListBuilderApiService
	id string
}

func (r ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyListById12Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyListById12Execute(r)
}

/*
PreviewPolicyListById12 Method for PreviewPolicyListById12

Preview a specific policy list entry based on id provided

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyListById12Request
*/
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) PreviewPolicyListById12(ctx context.Context, id string) ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyListById12Request {
	return ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyListById12Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyExtendedCommunityListBuilderApiService) PreviewPolicyListById12Execute(r ConfigurationPolicyExtendedCommunityListBuilderApiPreviewPolicyListById12Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyExtendedCommunityListBuilderApiService.PreviewPolicyListById12")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/extcommunity/preview/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
