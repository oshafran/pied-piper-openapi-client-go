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


// ConfigurationDeviceFirmwareUpdateApiService ConfigurationDeviceFirmwareUpdateApi service
type ConfigurationDeviceFirmwareUpdateApiService service

type ApiActivateFirmwareImageRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
	body *string
}

func (r ApiActivateFirmwareImageRequest) Body(body string) ApiActivateFirmwareImageRequest {
	r.body = &body
	return r
}

func (r ApiActivateFirmwareImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.ActivateFirmwareImageExecute(r)
}

/*
ActivateFirmwareImage Method for ActivateFirmwareImage

Activate firmware on device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiActivateFirmwareImageRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) ActivateFirmwareImage(ctx context.Context) ApiActivateFirmwareImageRequest {
	return ApiActivateFirmwareImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) ActivateFirmwareImageExecute(r ApiActivateFirmwareImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.ActivateFirmwareImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware/activate"

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

type ApiDeleteFirmwareImageRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
	versionId string
}

func (r ApiDeleteFirmwareImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFirmwareImageExecute(r)
}

/*
DeleteFirmwareImage Method for DeleteFirmwareImage

Delete firmware image package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param versionId Firmware image version
 @return ApiDeleteFirmwareImageRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) DeleteFirmwareImage(ctx context.Context, versionId string) ApiDeleteFirmwareImageRequest {
	return ApiDeleteFirmwareImageRequest{
		ApiService: a,
		ctx: ctx,
		versionId: versionId,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) DeleteFirmwareImageExecute(r ApiDeleteFirmwareImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.DeleteFirmwareImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware/{versionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterToString(r.versionId, "")), -1)

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

type ApiGetDevicesFWUpgradeRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
}

func (r ApiGetDevicesFWUpgradeRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetDevicesFWUpgradeExecute(r)
}

/*
GetDevicesFWUpgrade Method for GetDevicesFWUpgrade

Get list of devices that support firmware upgrade

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDevicesFWUpgradeRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) GetDevicesFWUpgrade(ctx context.Context) ApiGetDevicesFWUpgradeRequest {
	return ApiGetDevicesFWUpgradeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) GetDevicesFWUpgradeExecute(r ApiGetDevicesFWUpgradeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.GetDevicesFWUpgrade")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware/devices"

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

type ApiGetFirmwareImageDetailsRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
	versionId string
}

func (r ApiGetFirmwareImageDetailsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetFirmwareImageDetailsExecute(r)
}

/*
GetFirmwareImageDetails Method for GetFirmwareImageDetails

Get firmware image details for a given version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param versionId Firmware image version
 @return ApiGetFirmwareImageDetailsRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) GetFirmwareImageDetails(ctx context.Context, versionId string) ApiGetFirmwareImageDetailsRequest {
	return ApiGetFirmwareImageDetailsRequest{
		ApiService: a,
		ctx: ctx,
		versionId: versionId,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) GetFirmwareImageDetailsExecute(r ApiGetFirmwareImageDetailsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.GetFirmwareImageDetails")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware/{versionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterToString(r.versionId, "")), -1)

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

type ApiGetFirmwareImagesRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
}

func (r ApiGetFirmwareImagesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetFirmwareImagesExecute(r)
}

/*
GetFirmwareImages Method for GetFirmwareImages

Get list of firmware images in the repository

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFirmwareImagesRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) GetFirmwareImages(ctx context.Context) ApiGetFirmwareImagesRequest {
	return ApiGetFirmwareImagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) GetFirmwareImagesExecute(r ApiGetFirmwareImagesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.GetFirmwareImages")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware"

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

type ApiInstallFirmwareImageRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
	body *string
}

func (r ApiInstallFirmwareImageRequest) Body(body string) ApiInstallFirmwareImageRequest {
	r.body = &body
	return r
}

func (r ApiInstallFirmwareImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.InstallFirmwareImageExecute(r)
}

/*
InstallFirmwareImage Method for InstallFirmwareImage

Install firmware on device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInstallFirmwareImageRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) InstallFirmwareImage(ctx context.Context) ApiInstallFirmwareImageRequest {
	return ApiInstallFirmwareImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) InstallFirmwareImageExecute(r ApiInstallFirmwareImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.InstallFirmwareImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware/install"

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

type ApiProcessFirmwareImageRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
}

func (r ApiProcessFirmwareImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProcessFirmwareImageExecute(r)
}

/*
ProcessFirmwareImage Method for ProcessFirmwareImage

Upload firmware image package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProcessFirmwareImageRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) ProcessFirmwareImage(ctx context.Context) ApiProcessFirmwareImageRequest {
	return ApiProcessFirmwareImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) ProcessFirmwareImageExecute(r ApiProcessFirmwareImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.ProcessFirmwareImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware"

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

type ApiRemoveFirmwareImageRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceFirmwareUpdateApiService
	body *string
}

func (r ApiRemoveFirmwareImageRequest) Body(body string) ApiRemoveFirmwareImageRequest {
	r.body = &body
	return r
}

func (r ApiRemoveFirmwareImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveFirmwareImageExecute(r)
}

/*
RemoveFirmwareImage Method for RemoveFirmwareImage

Remove firmware on device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveFirmwareImageRequest
*/
func (a *ConfigurationDeviceFirmwareUpdateApiService) RemoveFirmwareImage(ctx context.Context) ApiRemoveFirmwareImageRequest {
	return ApiRemoveFirmwareImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceFirmwareUpdateApiService) RemoveFirmwareImageExecute(r ApiRemoveFirmwareImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceFirmwareUpdateApiService.RemoveFirmwareImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/firmware/remove"

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
