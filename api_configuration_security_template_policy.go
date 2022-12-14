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


// ConfigurationSecurityTemplatePolicyApiService ConfigurationSecurityTemplatePolicyApi service
type ConfigurationSecurityTemplatePolicyApiService service

type ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	body *map[string]interface{}
}

// Policy template
func (r ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest) Body(body map[string]interface{}) ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest {
	r.body = &body
	return r
}

func (r ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateSecurityTemplateExecute(r)
}

/*
CreateSecurityTemplate Method for CreateSecurityTemplate

Create Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) CreateSecurityTemplate(ctx context.Context) ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest {
	return ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationSecurityTemplatePolicyApiService) CreateSecurityTemplateExecute(r ConfigurationSecurityTemplatePolicyApiCreateSecurityTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.CreateSecurityTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
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

type ConfigurationSecurityTemplatePolicyApiDeleteSecurityTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	policyId string
}

func (r ConfigurationSecurityTemplatePolicyApiDeleteSecurityTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSecurityTemplateExecute(r)
}

/*
DeleteSecurityTemplate Method for DeleteSecurityTemplate

Delete Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId Policy Id
 @return ConfigurationSecurityTemplatePolicyApiDeleteSecurityTemplateRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) DeleteSecurityTemplate(ctx context.Context, policyId string) ConfigurationSecurityTemplatePolicyApiDeleteSecurityTemplateRequest {
	return ConfigurationSecurityTemplatePolicyApiDeleteSecurityTemplateRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
func (a *ConfigurationSecurityTemplatePolicyApiService) DeleteSecurityTemplateExecute(r ConfigurationSecurityTemplatePolicyApiDeleteSecurityTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.DeleteSecurityTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

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

type ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	policyId string
	body *map[string]interface{}
}

// Policy template
func (r ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest) Body(body map[string]interface{}) ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest {
	r.body = &body
	return r
}

func (r ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditSecurityTemplateExecute(r)
}

/*
EditSecurityTemplate Method for EditSecurityTemplate

Edit Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId Policy Id
 @return ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) EditSecurityTemplate(ctx context.Context, policyId string) ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest {
	return ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) EditSecurityTemplateExecute(r ConfigurationSecurityTemplatePolicyApiEditSecurityTemplateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.EditSecurityTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

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

type ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	policyId string
	body *map[string]interface{}
}

// Policy template
func (r ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest) Body(body map[string]interface{}) ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest {
	r.body = &body
	return r
}

func (r ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EditTemplateWithLenientLockExecute(r)
}

/*
EditTemplateWithLenientLock Method for EditTemplateWithLenientLock

Edit Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId Policy Id
 @return ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) EditTemplateWithLenientLock(ctx context.Context, policyId string) ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest {
	return ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) EditTemplateWithLenientLockExecute(r ConfigurationSecurityTemplatePolicyApiEditTemplateWithLenientLockRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.EditTemplateWithLenientLock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/staging/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

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

type ConfigurationSecurityTemplatePolicyApiGenerateSecurityPolicySummaryRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
}

func (r ConfigurationSecurityTemplatePolicyApiGenerateSecurityPolicySummaryRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GenerateSecurityPolicySummaryExecute(r)
}

/*
GenerateSecurityPolicySummary Method for GenerateSecurityPolicySummary

Generate security policy summary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSecurityTemplatePolicyApiGenerateSecurityPolicySummaryRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) GenerateSecurityPolicySummary(ctx context.Context) ConfigurationSecurityTemplatePolicyApiGenerateSecurityPolicySummaryRequest {
	return ConfigurationSecurityTemplatePolicyApiGenerateSecurityPolicySummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) GenerateSecurityPolicySummaryExecute(r ConfigurationSecurityTemplatePolicyApiGenerateSecurityPolicySummaryRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.GenerateSecurityPolicySummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/summary"

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

type ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	mode *string
}

// Mode
func (r ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest) Mode(mode string) ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest {
	r.mode = &mode
	return r
}

func (r ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GenerateSecurityTemplateListExecute(r)
}

/*
GenerateSecurityTemplateList Method for GenerateSecurityTemplateList

Generate template list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) GenerateSecurityTemplateList(ctx context.Context) ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest {
	return ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) GenerateSecurityTemplateListExecute(r ConfigurationSecurityTemplatePolicyApiGenerateSecurityTemplateListRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.GenerateSecurityTemplateList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mode != nil {
		localVarQueryParams.Add("mode", parameterToString(*r.mode, ""))
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

type ConfigurationSecurityTemplatePolicyApiGetDeviceListByIdRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	policyId string
}

func (r ConfigurationSecurityTemplatePolicyApiGetDeviceListByIdRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceListByIdExecute(r)
}

/*
GetDeviceListById Method for GetDeviceListById

Get device list by Id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId Policy Id
 @return ConfigurationSecurityTemplatePolicyApiGetDeviceListByIdRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) GetDeviceListById(ctx context.Context, policyId string) ConfigurationSecurityTemplatePolicyApiGetDeviceListByIdRequest {
	return ConfigurationSecurityTemplatePolicyApiGetDeviceListByIdRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) GetDeviceListByIdExecute(r ConfigurationSecurityTemplatePolicyApiGetDeviceListByIdRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.GetDeviceListById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/devices/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

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

type ConfigurationSecurityTemplatePolicyApiGetSecurityPolicyDeviceListRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
}

func (r ConfigurationSecurityTemplatePolicyApiGetSecurityPolicyDeviceListRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetSecurityPolicyDeviceListExecute(r)
}

/*
GetSecurityPolicyDeviceList Method for GetSecurityPolicyDeviceList

Get device list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationSecurityTemplatePolicyApiGetSecurityPolicyDeviceListRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) GetSecurityPolicyDeviceList(ctx context.Context) ConfigurationSecurityTemplatePolicyApiGetSecurityPolicyDeviceListRequest {
	return ConfigurationSecurityTemplatePolicyApiGetSecurityPolicyDeviceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) GetSecurityPolicyDeviceListExecute(r ConfigurationSecurityTemplatePolicyApiGetSecurityPolicyDeviceListRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.GetSecurityPolicyDeviceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/devices"

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

type ConfigurationSecurityTemplatePolicyApiGetSecurityTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	policyId string
}

func (r ConfigurationSecurityTemplatePolicyApiGetSecurityTemplateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetSecurityTemplateExecute(r)
}

/*
GetSecurityTemplate Method for GetSecurityTemplate

Get Template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId Policy Id
 @return ConfigurationSecurityTemplatePolicyApiGetSecurityTemplateRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) GetSecurityTemplate(ctx context.Context, policyId string) ConfigurationSecurityTemplatePolicyApiGetSecurityTemplateRequest {
	return ConfigurationSecurityTemplatePolicyApiGetSecurityTemplateRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) GetSecurityTemplateExecute(r ConfigurationSecurityTemplatePolicyApiGetSecurityTemplateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.GetSecurityTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/definition/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

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

type ConfigurationSecurityTemplatePolicyApiGetSecurityTemplatesForDeviceRequest struct {
	ctx context.Context
	ApiService *ConfigurationSecurityTemplatePolicyApiService
	deviceModel DeviceModel
}

func (r ConfigurationSecurityTemplatePolicyApiGetSecurityTemplatesForDeviceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetSecurityTemplatesForDeviceExecute(r)
}

/*
GetSecurityTemplatesForDevice Method for GetSecurityTemplatesForDevice

Get templates that map a device model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceModel Device model
 @return ConfigurationSecurityTemplatePolicyApiGetSecurityTemplatesForDeviceRequest
*/
func (a *ConfigurationSecurityTemplatePolicyApiService) GetSecurityTemplatesForDevice(ctx context.Context, deviceModel DeviceModel) ConfigurationSecurityTemplatePolicyApiGetSecurityTemplatesForDeviceRequest {
	return ConfigurationSecurityTemplatePolicyApiGetSecurityTemplatesForDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceModel: deviceModel,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationSecurityTemplatePolicyApiService) GetSecurityTemplatesForDeviceExecute(r ConfigurationSecurityTemplatePolicyApiGetSecurityTemplatesForDeviceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationSecurityTemplatePolicyApiService.GetSecurityTemplatesForDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/template/policy/security/{deviceModel}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceModel"+"}", url.PathEscape(parameterToString(r.deviceModel, "")), -1)

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
