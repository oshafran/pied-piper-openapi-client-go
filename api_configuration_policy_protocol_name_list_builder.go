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


// ConfigurationPolicyProtocolNameListBuilderApiService ConfigurationPolicyProtocolNameListBuilderApi service
type ConfigurationPolicyProtocolNameListBuilderApiService service

type ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request) Body(body map[string]interface{}) ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreatePolicyList28Execute(r)
}

/*
CreatePolicyList28 Method for CreatePolicyList28

Create policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) CreatePolicyList28(ctx context.Context) ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) CreatePolicyList28Execute(r ConfigurationPolicyProtocolNameListBuilderApiCreatePolicyList28Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.CreatePolicyList28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname"

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

type ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyList28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	id string
}

func (r ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyList28Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyList28Execute(r)
}

/*
DeletePolicyList28 Method for DeletePolicyList28

Delete policy list entry for a specific type of policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyList28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) DeletePolicyList28(ctx context.Context, id string) ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyList28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyList28Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) DeletePolicyList28Execute(r ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyList28Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.DeletePolicyList28")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname/{id}"
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

type ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	infoTag *string
}

// InfoTag
func (r ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request) InfoTag(infoTag string) ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request {
	r.infoTag = &infoTag
	return r
}

func (r ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeletePolicyListsWithInfoTag28Execute(r)
}

/*
DeletePolicyListsWithInfoTag28 Method for DeletePolicyListsWithInfoTag28

Delete policy lists with specific info tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) DeletePolicyListsWithInfoTag28(ctx context.Context) ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) DeletePolicyListsWithInfoTag28Execute(r ConfigurationPolicyProtocolNameListBuilderApiDeletePolicyListsWithInfoTag28Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.DeletePolicyListsWithInfoTag28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname"

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

type ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request) Body(body map[string]interface{}) ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditPolicyList28Execute(r)
}

/*
EditPolicyList28 Method for EditPolicyList28

Edit policy list entries for a specific type of policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) EditPolicyList28(ctx context.Context, id string) ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) EditPolicyList28Execute(r ConfigurationPolicyProtocolNameListBuilderApiEditPolicyList28Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.EditPolicyList28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname/{id}"
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

type ConfigurationPolicyProtocolNameListBuilderApiGetListsById28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	id string
}

func (r ConfigurationPolicyProtocolNameListBuilderApiGetListsById28Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetListsById28Execute(r)
}

/*
GetListsById28 Method for GetListsById28

Get a specific policy list based on the id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyProtocolNameListBuilderApiGetListsById28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) GetListsById28(ctx context.Context, id string) ConfigurationPolicyProtocolNameListBuilderApiGetListsById28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiGetListsById28Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) GetListsById28Execute(r ConfigurationPolicyProtocolNameListBuilderApiGetListsById28Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.GetListsById28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname/{id}"
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

type ConfigurationPolicyProtocolNameListBuilderApiGetPolicyLists25Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
}

func (r ConfigurationPolicyProtocolNameListBuilderApiGetPolicyLists25Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyLists25Execute(r)
}

/*
GetPolicyLists25 Method for GetPolicyLists25

Get policy lists

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyProtocolNameListBuilderApiGetPolicyLists25Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) GetPolicyLists25(ctx context.Context) ConfigurationPolicyProtocolNameListBuilderApiGetPolicyLists25Request {
	return ConfigurationPolicyProtocolNameListBuilderApiGetPolicyLists25Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) GetPolicyLists25Execute(r ConfigurationPolicyProtocolNameListBuilderApiGetPolicyLists25Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.GetPolicyLists25")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname"

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

type ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	infoTag *string
}

// InfoTag
func (r ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request) InfoTag(infoTag string) ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request {
	r.infoTag = &infoTag
	return r
}

func (r ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyListsWithInfoTag28Execute(r)
}

/*
GetPolicyListsWithInfoTag28 Method for GetPolicyListsWithInfoTag28

Get policy lists with specific info tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) GetPolicyListsWithInfoTag28(ctx context.Context) ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) GetPolicyListsWithInfoTag28Execute(r ConfigurationPolicyProtocolNameListBuilderApiGetPolicyListsWithInfoTag28Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.GetPolicyListsWithInfoTag28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname/filtered"

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

type ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request) Body(body map[string]interface{}) ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyList28Execute(r)
}

/*
PreviewPolicyList28 Method for PreviewPolicyList28

Preview a policy list based on the policy list type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) PreviewPolicyList28(ctx context.Context) ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) PreviewPolicyList28Execute(r ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyList28Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.PreviewPolicyList28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname/preview"

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

type ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyListById28Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyProtocolNameListBuilderApiService
	id string
}

func (r ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyListById28Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyListById28Execute(r)
}

/*
PreviewPolicyListById28 Method for PreviewPolicyListById28

Preview a specific policy list entry based on id provided

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyListById28Request
*/
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) PreviewPolicyListById28(ctx context.Context, id string) ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyListById28Request {
	return ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyListById28Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyProtocolNameListBuilderApiService) PreviewPolicyListById28Execute(r ConfigurationPolicyProtocolNameListBuilderApiPreviewPolicyListById28Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyProtocolNameListBuilderApiService.PreviewPolicyListById28")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/protocolname/preview/{id}"
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
