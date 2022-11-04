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


// SystemReverseProxyApiService SystemReverseProxyApi service
type SystemReverseProxyApiService service

type SystemReverseProxyApiCreateReverseProxyMappingsRequest struct {
	ctx context.Context
	ApiService *SystemReverseProxyApiService
	uuid string
	body *map[string]interface{}
}

// Device reverse proxy mappings
func (r SystemReverseProxyApiCreateReverseProxyMappingsRequest) Body(body map[string]interface{}) SystemReverseProxyApiCreateReverseProxyMappingsRequest {
	r.body = &body
	return r
}

func (r SystemReverseProxyApiCreateReverseProxyMappingsRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateReverseProxyMappingsExecute(r)
}

/*
CreateReverseProxyMappings Method for CreateReverseProxyMappings

Create reverse proxy IP/Port mappings for controller

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid Device uuid
 @return SystemReverseProxyApiCreateReverseProxyMappingsRequest
*/
func (a *SystemReverseProxyApiService) CreateReverseProxyMappings(ctx context.Context, uuid string) SystemReverseProxyApiCreateReverseProxyMappingsRequest {
	return SystemReverseProxyApiCreateReverseProxyMappingsRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
func (a *SystemReverseProxyApiService) CreateReverseProxyMappingsExecute(r SystemReverseProxyApiCreateReverseProxyMappingsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemReverseProxyApiService.CreateReverseProxyMappings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/reverseproxy/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterToString(r.uuid, "")), -1)

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

type SystemReverseProxyApiGetReverseProxyMappingsRequest struct {
	ctx context.Context
	ApiService *SystemReverseProxyApiService
	uuid string
}

func (r SystemReverseProxyApiGetReverseProxyMappingsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetReverseProxyMappingsExecute(r)
}

/*
GetReverseProxyMappings Method for GetReverseProxyMappings

Get reverse proxy IP/Port mappings for controller

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid Device uuid
 @return SystemReverseProxyApiGetReverseProxyMappingsRequest
*/
func (a *SystemReverseProxyApiService) GetReverseProxyMappings(ctx context.Context, uuid string) SystemReverseProxyApiGetReverseProxyMappingsRequest {
	return SystemReverseProxyApiGetReverseProxyMappingsRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SystemReverseProxyApiService) GetReverseProxyMappingsExecute(r SystemReverseProxyApiGetReverseProxyMappingsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemReverseProxyApiService.GetReverseProxyMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/reverseproxy/{uuid}"
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
