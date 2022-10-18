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


// PartnerWCMConfigsApiService PartnerWCMConfigsApi service
type PartnerWCMConfigsApiService service

type PartnerWCMConfigsApiPushNetconfConfigsRequest struct {
	ctx context.Context
	ApiService *PartnerWCMConfigsApiService
	nmsId string
	requestBody *[]map[string]interface{}
}

// Netconf configuration
func (r PartnerWCMConfigsApiPushNetconfConfigsRequest) RequestBody(requestBody []map[string]interface{}) PartnerWCMConfigsApiPushNetconfConfigsRequest {
	r.requestBody = &requestBody
	return r
}

func (r PartnerWCMConfigsApiPushNetconfConfigsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PushNetconfConfigsExecute(r)
}

/*
PushNetconfConfigs Method for PushNetconfConfigs

Push device configs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nmsId NMS Id
 @return PartnerWCMConfigsApiPushNetconfConfigsRequest
*/
func (a *PartnerWCMConfigsApiService) PushNetconfConfigs(ctx context.Context, nmsId string) PartnerWCMConfigsApiPushNetconfConfigsRequest {
	return PartnerWCMConfigsApiPushNetconfConfigsRequest{
		ApiService: a,
		ctx: ctx,
		nmsId: nmsId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PartnerWCMConfigsApiService) PushNetconfConfigsExecute(r PartnerWCMConfigsApiPushNetconfConfigsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerWCMConfigsApiService.PushNetconfConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/partner/wcm/netconf/{nmsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"nmsId"+"}", url.PathEscape(parameterToString(r.nmsId, "")), -1)

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
	localVarPostBody = r.requestBody
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
