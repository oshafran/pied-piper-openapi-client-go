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


// ConfigurationPolicyTranslationProfileListBuilderApiService ConfigurationPolicyTranslationProfileListBuilderApi service
type ConfigurationPolicyTranslationProfileListBuilderApiService service

type ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request) Body(body map[string]interface{}) ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreatePolicyList1Execute(r)
}

/*
CreatePolicyList1 Method for CreatePolicyList1

Create policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) CreatePolicyList1(ctx context.Context) ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) CreatePolicyList1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiCreatePolicyList1Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.CreatePolicyList1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile"

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

type ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyList1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	id string
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyList1Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyList1Execute(r)
}

/*
DeletePolicyList1 Method for DeletePolicyList1

Delete policy list entry for a specific type of policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyList1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) DeletePolicyList1(ctx context.Context, id string) ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyList1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyList1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) DeletePolicyList1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyList1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.DeletePolicyList1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile/{id}"
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

type ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	infoTag *string
}

// InfoTag
func (r ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request) InfoTag(infoTag string) ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request {
	r.infoTag = &infoTag
	return r
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeletePolicyListsWithInfoTag1Execute(r)
}

/*
DeletePolicyListsWithInfoTag1 Method for DeletePolicyListsWithInfoTag1

Delete policy lists with specific info tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) DeletePolicyListsWithInfoTag1(ctx context.Context) ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) DeletePolicyListsWithInfoTag1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiDeletePolicyListsWithInfoTag1Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.DeletePolicyListsWithInfoTag1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile"

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

type ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request) Body(body map[string]interface{}) ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditPolicyList1Execute(r)
}

/*
EditPolicyList1 Method for EditPolicyList1

Edit policy list entries for a specific type of policy list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) EditPolicyList1(ctx context.Context, id string) ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) EditPolicyList1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiEditPolicyList1Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.EditPolicyList1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile/{id}"
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

type ConfigurationPolicyTranslationProfileListBuilderApiGetListsById1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	id string
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiGetListsById1Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetListsById1Execute(r)
}

/*
GetListsById1 Method for GetListsById1

Get a specific policy list based on the id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyTranslationProfileListBuilderApiGetListsById1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) GetListsById1(ctx context.Context, id string) ConfigurationPolicyTranslationProfileListBuilderApiGetListsById1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiGetListsById1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) GetListsById1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiGetListsById1Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.GetListsById1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile/{id}"
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

type ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyLists1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyLists1Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyLists1Execute(r)
}

/*
GetPolicyLists1 Method for GetPolicyLists1

Get policy lists

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyLists1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) GetPolicyLists1(ctx context.Context) ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyLists1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyLists1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) GetPolicyLists1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyLists1Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.GetPolicyLists1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile"

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

type ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	infoTag *string
}

// InfoTag
func (r ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request) InfoTag(infoTag string) ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request {
	r.infoTag = &infoTag
	return r
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyListsWithInfoTag1Execute(r)
}

/*
GetPolicyListsWithInfoTag1 Method for GetPolicyListsWithInfoTag1

Get policy lists with specific info tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) GetPolicyListsWithInfoTag1(ctx context.Context) ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) GetPolicyListsWithInfoTag1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiGetPolicyListsWithInfoTag1Request) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.GetPolicyListsWithInfoTag1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile/filtered"

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

type ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	body *map[string]interface{}
}

// Policy list
func (r ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request) Body(body map[string]interface{}) ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyList1Execute(r)
}

/*
PreviewPolicyList1 Method for PreviewPolicyList1

Preview a policy list based on the policy list type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) PreviewPolicyList1(ctx context.Context) ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) PreviewPolicyList1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyList1Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.PreviewPolicyList1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile/preview"

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

type ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyListById1Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyTranslationProfileListBuilderApiService
	id string
}

func (r ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyListById1Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyListById1Execute(r)
}

/*
PreviewPolicyListById1 Method for PreviewPolicyListById1

Preview a specific policy list entry based on id provided

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyListById1Request
*/
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) PreviewPolicyListById1(ctx context.Context, id string) ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyListById1Request {
	return ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyListById1Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyTranslationProfileListBuilderApiService) PreviewPolicyListById1Execute(r ConfigurationPolicyTranslationProfileListBuilderApiPreviewPolicyListById1Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyTranslationProfileListBuilderApiService.PreviewPolicyListById1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/list/translationprofile/preview/{id}"
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
