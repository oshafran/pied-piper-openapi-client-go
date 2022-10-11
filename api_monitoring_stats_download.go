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


// MonitoringStatsDownloadApiService MonitoringStatsDownloadApi service
type MonitoringStatsDownloadApiService service

type ApiDownload1Request struct {
	ctx context.Context
	ApiService *MonitoringStatsDownloadApiService
	processType string
	fileType string
	queue string
	deviceIp string
	token string
	fileName string
}

func (r ApiDownload1Request) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.Download1Execute(r)
}

/*
Download1 Method for Download1

Downloading stats file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param processType Process type
 @param fileType File type
 @param queue Queue name
 @param deviceIp Device IP
 @param token Token
 @param fileName File name
 @return ApiDownload1Request
*/
func (a *MonitoringStatsDownloadApiService) Download1(ctx context.Context, processType string, fileType string, queue string, deviceIp string, token string, fileName string) ApiDownload1Request {
	return ApiDownload1Request{
		ApiService: a,
		ctx: ctx,
		processType: processType,
		fileType: fileType,
		queue: queue,
		deviceIp: deviceIp,
		token: token,
		fileName: fileName,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MonitoringStatsDownloadApiService) Download1Execute(r ApiDownload1Request) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringStatsDownloadApiService.Download1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/statistics/download/{processType}/file/{fileType}/{queue}/{deviceIp}/{token}/{fileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"processType"+"}", url.PathEscape(parameterToString(r.processType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileType"+"}", url.PathEscape(parameterToString(r.fileType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queue"+"}", url.PathEscape(parameterToString(r.queue, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deviceIp"+"}", url.PathEscape(parameterToString(r.deviceIp, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterToString(r.token, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileName"+"}", url.PathEscape(parameterToString(r.fileName, "")), -1)

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

type ApiDownloadListRequest struct {
	ctx context.Context
	ApiService *MonitoringStatsDownloadApiService
	processType string
	requestBody *map[string]GetO365PreferredPathFromVAnalyticsRequestValue
}

func (r ApiDownloadListRequest) RequestBody(requestBody map[string]GetO365PreferredPathFromVAnalyticsRequestValue) ApiDownloadListRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiDownloadListRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadListExecute(r)
}

/*
DownloadList Method for DownloadList

Downloading list of stats file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param processType Possible types are: remoteprocessing, dr
 @return ApiDownloadListRequest
*/
func (a *MonitoringStatsDownloadApiService) DownloadList(ctx context.Context, processType string) ApiDownloadListRequest {
	return ApiDownloadListRequest{
		ApiService: a,
		ctx: ctx,
		processType: processType,
	}
}

// Execute executes the request
func (a *MonitoringStatsDownloadApiService) DownloadListExecute(r ApiDownloadListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringStatsDownloadApiService.DownloadList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/statistics/download/{processType}/filelist"
	localVarPath = strings.Replace(localVarPath, "{"+"processType"+"}", url.PathEscape(parameterToString(r.processType, "")), -1)

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
	localVarPostBody = r.requestBody
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

type ApiFetchListRequest struct {
	ctx context.Context
	ApiService *MonitoringStatsDownloadApiService
	processType string
}

func (r ApiFetchListRequest) Execute() (*http.Response, error) {
	return r.ApiService.FetchListExecute(r)
}

/*
FetchList Method for FetchList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param processType
 @return ApiFetchListRequest
*/
func (a *MonitoringStatsDownloadApiService) FetchList(ctx context.Context, processType string) ApiFetchListRequest {
	return ApiFetchListRequest{
		ApiService: a,
		ctx: ctx,
		processType: processType,
	}
}

// Execute executes the request
func (a *MonitoringStatsDownloadApiService) FetchListExecute(r ApiFetchListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoringStatsDownloadApiService.FetchList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/statistics/download/{processType}/fetchvManageList"
	localVarPath = strings.Replace(localVarPath, "{"+"processType"+"}", url.PathEscape(parameterToString(r.processType, "")), -1)

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