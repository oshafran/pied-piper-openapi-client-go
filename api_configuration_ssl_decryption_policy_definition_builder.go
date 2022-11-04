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


// ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService ConfigurationSSLDecryptionPolicyDefinitionBuilderApi service
type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService service

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request) Body(body map[string]interface{}) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request {
	r.body = &body
	return r
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreatePolicyDefinition3Execute(r)
}

/*
CreatePolicyDefinition3 Method for CreatePolicyDefinition3

Create policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) CreatePolicyDefinition3(ctx context.Context) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) CreatePolicyDefinition3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiCreatePolicyDefinition3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.CreatePolicyDefinition3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption"

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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiDeletePolicyDefinition3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	id string
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiDeletePolicyDefinition3Request) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyDefinition3Execute(r)
}

/*
DeletePolicyDefinition3 Method for DeletePolicyDefinition3

Delete policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiDeletePolicyDefinition3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) DeletePolicyDefinition3(ctx context.Context, id string) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiDeletePolicyDefinition3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiDeletePolicyDefinition3Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) DeletePolicyDefinition3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiDeletePolicyDefinition3Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.DeletePolicyDefinition3")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/{id}"
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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request) Body(body map[string]interface{}) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request {
	r.body = &body
	return r
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditMultiplePolicyDefinition3Execute(r)
}

/*
EditMultiplePolicyDefinition3 Method for EditMultiplePolicyDefinition3

Edit multiple policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) EditMultiplePolicyDefinition3(ctx context.Context, id string) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) EditMultiplePolicyDefinition3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditMultiplePolicyDefinition3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.EditMultiplePolicyDefinition3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/multiple/{id}"
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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	id string
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request) Body(body map[string]interface{}) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request {
	r.body = &body
	return r
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditPolicyDefinition3Execute(r)
}

/*
EditPolicyDefinition3 Method for EditPolicyDefinition3

Edit a policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) EditPolicyDefinition3(ctx context.Context, id string) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) EditPolicyDefinition3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiEditPolicyDefinition3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.EditPolicyDefinition3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/{id}"
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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetDefinitions3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetDefinitions3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDefinitions3Execute(r)
}

/*
GetDefinitions3 Method for GetDefinitions3

Get policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetDefinitions3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) GetDefinitions3(ctx context.Context) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetDefinitions3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetDefinitions3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) GetDefinitions3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetDefinitions3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.GetDefinitions3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption"

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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetPolicyDefinition3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	id string
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetPolicyDefinition3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetPolicyDefinition3Execute(r)
}

/*
GetPolicyDefinition3 Method for GetPolicyDefinition3

Get a specific policy definitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetPolicyDefinition3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) GetPolicyDefinition3(ctx context.Context, id string) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetPolicyDefinition3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetPolicyDefinition3Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) GetPolicyDefinition3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiGetPolicyDefinition3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.GetPolicyDefinition3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/{id}"
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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request) Body(body map[string]interface{}) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request {
	r.body = &body
	return r
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyDefinition3Execute(r)
}

/*
PreviewPolicyDefinition3 Method for PreviewPolicyDefinition3

Preview policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) PreviewPolicyDefinition3(ctx context.Context) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) PreviewPolicyDefinition3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinition3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.PreviewPolicyDefinition3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/preview"

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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinitionById3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	id string
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinitionById3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PreviewPolicyDefinitionById3Execute(r)
}

/*
PreviewPolicyDefinitionById3 Method for PreviewPolicyDefinitionById3

Preview policy definition

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Policy Id
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinitionById3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) PreviewPolicyDefinitionById3(ctx context.Context, id string) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinitionById3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinitionById3Request{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) PreviewPolicyDefinitionById3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiPreviewPolicyDefinitionById3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.PreviewPolicyDefinitionById3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/preview/{id}"
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

type ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request struct {
	ctx context.Context
	ApiService *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService
	body *map[string]interface{}
}

// Policy definition
func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request) Body(body map[string]interface{}) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request {
	r.body = &body
	return r
}

func (r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SavePolicyDefinitionInBulk3Execute(r)
}

/*
SavePolicyDefinitionInBulk3 Method for SavePolicyDefinitionInBulk3

Create/Edit policy definitions in bulk

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request
*/
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) SavePolicyDefinitionInBulk3(ctx context.Context) ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request {
	return ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService) SavePolicyDefinitionInBulk3Execute(r ConfigurationSSLDecryptionPolicyDefinitionBuilderApiSavePolicyDefinitionInBulk3Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSSLDecryptionPolicyDefinitionBuilderApiService.SavePolicyDefinitionInBulk3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/definition/ssldecryption/bulk"

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
