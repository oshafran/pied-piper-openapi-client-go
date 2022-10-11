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


// ConfigurationNetworkDesignTemplatesApiService ConfigurationNetworkDesignTemplatesApi service
type ConfigurationNetworkDesignTemplatesApiService service

type ApiEditDeviceProfileTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationNetworkDesignTemplatesApiService
	templateId string
	body *map[string]interface{}
}

// Global template
func (r ApiEditDeviceProfileTemplateRequest) Body(body map[string]interface{}) ApiEditDeviceProfileTemplateRequest {
	r.body = &body
	return r
}

func (r ApiEditDeviceProfileTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditDeviceProfileTemplateExecute(r)
}

/*
EditDeviceProfileTemplate Method for EditDeviceProfileTemplate

Edit device profile template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId Template Id
 @return ApiEditDeviceProfileTemplateRequest
*/
func (a *ConfigurationNetworkDesignTemplatesApiService) EditDeviceProfileTemplate(ctx context.Context, templateId string) ApiEditDeviceProfileTemplateRequest {
	return ApiEditDeviceProfileTemplateRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *ConfigurationNetworkDesignTemplatesApiService) EditDeviceProfileTemplateExecute(r ApiEditDeviceProfileTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationNetworkDesignTemplatesApiService.EditDeviceProfileTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networkdesign/profile/template/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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

type ApiEditGlobalTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationNetworkDesignTemplatesApiService
	templateId string
	body *map[string]interface{}
}

// Global template
func (r ApiEditGlobalTemplateRequest) Body(body map[string]interface{}) ApiEditGlobalTemplateRequest {
	r.body = &body
	return r
}

func (r ApiEditGlobalTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditGlobalTemplateExecute(r)
}

/*
EditGlobalTemplate Method for EditGlobalTemplate

Edit global template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId Template Id
 @return ApiEditGlobalTemplateRequest
*/
func (a *ConfigurationNetworkDesignTemplatesApiService) EditGlobalTemplate(ctx context.Context, templateId string) ApiEditGlobalTemplateRequest {
	return ApiEditGlobalTemplateRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *ConfigurationNetworkDesignTemplatesApiService) EditGlobalTemplateExecute(r ApiEditGlobalTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationNetworkDesignTemplatesApiService.EditGlobalTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networkdesign/global/template/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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

type ApiGenerateProfileTemplateListRequest struct {
	ctx context.Context
	ApiService *ConfigurationNetworkDesignTemplatesApiService
}

func (r ApiGenerateProfileTemplateListRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GenerateProfileTemplateListExecute(r)
}

/*
GenerateProfileTemplateList Method for GenerateProfileTemplateList

Generate profile template list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateProfileTemplateListRequest
*/
func (a *ConfigurationNetworkDesignTemplatesApiService) GenerateProfileTemplateList(ctx context.Context) ApiGenerateProfileTemplateListRequest {
	return ApiGenerateProfileTemplateListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationNetworkDesignTemplatesApiService) GenerateProfileTemplateListExecute(r ApiGenerateProfileTemplateListRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationNetworkDesignTemplatesApiService.GenerateProfileTemplateList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networkdesign/profile/template"

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

type ApiGetDeviceProfileFeatureTemplateListRequest struct {
	ctx context.Context
	ApiService *ConfigurationNetworkDesignTemplatesApiService
}

func (r ApiGetDeviceProfileFeatureTemplateListRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceProfileFeatureTemplateListExecute(r)
}

/*
GetDeviceProfileFeatureTemplateList Method for GetDeviceProfileFeatureTemplateList

Generate device profile template list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDeviceProfileFeatureTemplateListRequest
*/
func (a *ConfigurationNetworkDesignTemplatesApiService) GetDeviceProfileFeatureTemplateList(ctx context.Context) ApiGetDeviceProfileFeatureTemplateListRequest {
	return ApiGetDeviceProfileFeatureTemplateListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *ConfigurationNetworkDesignTemplatesApiService) GetDeviceProfileFeatureTemplateListExecute(r ApiGetDeviceProfileFeatureTemplateListRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationNetworkDesignTemplatesApiService.GetDeviceProfileFeatureTemplateList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networkdesign/profile/feature"

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

type ApiGetDeviceProfileTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationNetworkDesignTemplatesApiService
	templateId string
}

func (r ApiGetDeviceProfileTemplateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceProfileTemplateExecute(r)
}

/*
GetDeviceProfileTemplate Method for GetDeviceProfileTemplate

Get device profile template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId Template Id
 @return ApiGetDeviceProfileTemplateRequest
*/
func (a *ConfigurationNetworkDesignTemplatesApiService) GetDeviceProfileTemplate(ctx context.Context, templateId string) ApiGetDeviceProfileTemplateRequest {
	return ApiGetDeviceProfileTemplateRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationNetworkDesignTemplatesApiService) GetDeviceProfileTemplateExecute(r ApiGetDeviceProfileTemplateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationNetworkDesignTemplatesApiService.GetDeviceProfileTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networkdesign/profile/template/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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

type ApiGetGlobalTemplateRequest struct {
	ctx context.Context
	ApiService *ConfigurationNetworkDesignTemplatesApiService
	templateId string
}

func (r ApiGetGlobalTemplateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetGlobalTemplateExecute(r)
}

/*
GetGlobalTemplate Method for GetGlobalTemplate

Get global template

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId Template Id
 @return ApiGetGlobalTemplateRequest
*/
func (a *ConfigurationNetworkDesignTemplatesApiService) GetGlobalTemplate(ctx context.Context, templateId string) ApiGetGlobalTemplateRequest {
	return ApiGetGlobalTemplateRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigurationNetworkDesignTemplatesApiService) GetGlobalTemplateExecute(r ApiGetGlobalTemplateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationNetworkDesignTemplatesApiService.GetGlobalTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networkdesign/global/template/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterToString(r.templateId, "")), -1)

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
