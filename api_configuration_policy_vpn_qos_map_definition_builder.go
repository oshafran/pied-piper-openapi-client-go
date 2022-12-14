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


// ConfigurationPolicyVPNQosMapDefinitionBuilderApiService ConfigurationPolicyVPNQosMapDefinitionBuilderApi service
type ConfigurationPolicyVPNQosMapDefinitionBuilderApiService service

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request) Body(body map[string]interface{}) ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreatePolicyDefinition2Execute(r)
}

/*
CreatePolicyDefinition2 Method for CreatePolicyDefinition2

Create policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) CreatePolicyDefinition2(ctx context.Context) ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) CreatePolicyDefinition2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiCreatePolicyDefinition2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.CreatePolicyDefinition2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap"

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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiDeletePolicyDefinition2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	id string
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiDeletePolicyDefinition2Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyDefinition2Execute(r)
}

/*
DeletePolicyDefinition2 Method for DeletePolicyDefinition2

Delete policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiDeletePolicyDefinition2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) DeletePolicyDefinition2(ctx context.Context, id string) ConfigurationPolicyVPNQosMapDefinitionBuilderApiDeletePolicyDefinition2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiDeletePolicyDefinition2Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) DeletePolicyDefinition2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiDeletePolicyDefinition2Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.DeletePolicyDefinition2")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/{id}"
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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request) Body(body map[string]interface{}) ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditMultiplePolicyDefinition2Execute(r)
}

/*
EditMultiplePolicyDefinition2 Method for EditMultiplePolicyDefinition2

Edit multiple policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) EditMultiplePolicyDefinition2(ctx context.Context, id string) ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) EditMultiplePolicyDefinition2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditMultiplePolicyDefinition2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.EditMultiplePolicyDefinition2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/multiple/{id}"
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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request) Body(body map[string]interface{}) ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditPolicyDefinition2Execute(r)
}

/*
EditPolicyDefinition2 Method for EditPolicyDefinition2

Edit a policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) EditPolicyDefinition2(ctx context.Context, id string) ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) EditPolicyDefinition2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiEditPolicyDefinition2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.EditPolicyDefinition2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/{id}"
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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetDefinitions2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetDefinitions2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDefinitions2Execute(r)
}

/*
GetDefinitions2 Method for GetDefinitions2

Get policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetDefinitions2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) GetDefinitions2(ctx context.Context) ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetDefinitions2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetDefinitions2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) GetDefinitions2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetDefinitions2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.GetDefinitions2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap"

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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetPolicyDefinition2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	id string
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetPolicyDefinition2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyDefinition2Execute(r)
}

/*
GetPolicyDefinition2 Method for GetPolicyDefinition2

Get a specific policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetPolicyDefinition2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) GetPolicyDefinition2(ctx context.Context, id string) ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetPolicyDefinition2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetPolicyDefinition2Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) GetPolicyDefinition2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiGetPolicyDefinition2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.GetPolicyDefinition2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/{id}"
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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request) Body(body map[string]interface{}) ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyDefinition2Execute(r)
}

/*
PreviewPolicyDefinition2 Method for PreviewPolicyDefinition2

Preview policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) PreviewPolicyDefinition2(ctx context.Context) ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) PreviewPolicyDefinition2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinition2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.PreviewPolicyDefinition2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/preview"

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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinitionById2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	id string
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinitionById2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyDefinitionById2Execute(r)
}

/*
PreviewPolicyDefinitionById2 Method for PreviewPolicyDefinitionById2

Preview policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinitionById2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) PreviewPolicyDefinitionById2(ctx context.Context, id string) ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinitionById2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinitionById2Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) PreviewPolicyDefinitionById2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiPreviewPolicyDefinitionById2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.PreviewPolicyDefinitionById2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/preview/{id}"
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

type ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request struct {
	ctx context.Context
	ApiService *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request) Body(body map[string]interface{}) ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request {
	r.body = &body
	return r
}

func (r ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SavePolicyDefinitionInBulk2Execute(r)
}

/*
SavePolicyDefinitionInBulk2 Method for SavePolicyDefinitionInBulk2

Create/Edit policy definitions in bulk

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request
*/
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) SavePolicyDefinitionInBulk2(ctx context.Context) ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request {
	return ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationPolicyVPNQosMapDefinitionBuilderApiService) SavePolicyDefinitionInBulk2Execute(r ConfigurationPolicyVPNQosMapDefinitionBuilderApiSavePolicyDefinitionInBulk2Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationPolicyVPNQosMapDefinitionBuilderApiService.SavePolicyDefinitionInBulk2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/vpnqosmap/bulk"

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
