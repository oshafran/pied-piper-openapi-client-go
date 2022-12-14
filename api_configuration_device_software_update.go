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
	"reflect"
)


// ConfigurationDeviceSoftwareUpdateApiService ConfigurationDeviceSoftwareUpdateApi service
type ConfigurationDeviceSoftwareUpdateApiService service

type ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceSoftwareUpdateApiService
	fileName string
	imageType *string
}

// Image type
func (r ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest) ImageType(imageType string) ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest {
	r.imageType = &imageType
	return r
}

func (r ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadPackageFileExecute(r)
}

/*
DownloadPackageFile Method for DownloadPackageFile

Download software package file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileName Pakcage file name
 @return ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest
*/
func (a *ConfigurationDeviceSoftwareUpdateApiService) DownloadPackageFile(ctx context.Context, fileName string) ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest {
	return ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest{
		ApiService: a,
		ctx: ctx,
		fileName: fileName,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceSoftwareUpdateApiService) DownloadPackageFileExecute(r ConfigurationDeviceSoftwareUpdateApiDownloadPackageFileRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceSoftwareUpdateApiService.DownloadPackageFile")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/software/package/{fileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"fileName"+"}", url.PathEscape(parameterToString(r.fileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.imageType == nil {
		return nil, reportError("imageType is required and must be specified")
	}

	localVarQueryParams.Add("imageType", parameterToString(*r.imageType, ""))
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

type ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceSoftwareUpdateApiService
	versionId string
	body *string
}

func (r ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest) Body(body string) ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest {
	r.body = &body
	return r
}

func (r ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditImageMetadataExecute(r)
}

/*
EditImageMetadata Method for EditImageMetadata

Update Package Metadata

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param versionId Image ID
 @return ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest
*/
func (a *ConfigurationDeviceSoftwareUpdateApiService) EditImageMetadata(ctx context.Context, versionId string) ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest {
	return ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest{
		ApiService: a,
		ctx: ctx,
		versionId: versionId,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceSoftwareUpdateApiService) EditImageMetadataExecute(r ConfigurationDeviceSoftwareUpdateApiEditImageMetadataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceSoftwareUpdateApiService.EditImageMetadata")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/software/package/{versionId}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"versionId"+"}", url.PathEscape(parameterToString(r.versionId, "")), -1)

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

type ConfigurationDeviceSoftwareUpdateApiGetImageMetadataRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceSoftwareUpdateApiService
	versionId string
}

func (r ConfigurationDeviceSoftwareUpdateApiGetImageMetadataRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetImageMetadataExecute(r)
}

/*
GetImageMetadata Method for GetImageMetadata

Update Package Metadata

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param versionId Image ID
 @return ConfigurationDeviceSoftwareUpdateApiGetImageMetadataRequest
*/
func (a *ConfigurationDeviceSoftwareUpdateApiService) GetImageMetadata(ctx context.Context, versionId string) ConfigurationDeviceSoftwareUpdateApiGetImageMetadataRequest {
	return ConfigurationDeviceSoftwareUpdateApiGetImageMetadataRequest{
		ApiService: a,
		ctx: ctx,
		versionId: versionId,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceSoftwareUpdateApiService) GetImageMetadataExecute(r ConfigurationDeviceSoftwareUpdateApiGetImageMetadataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceSoftwareUpdateApiService.GetImageMetadata")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/software/package/{versionId}/metadata"
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

type ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceSoftwareUpdateApiService
	imageType *[]string
}

// Image type
func (r ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest) ImageType(imageType []string) ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest {
	r.imageType = &imageType
	return r
}

func (r ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetUploadImagesCountExecute(r)
}

/*
GetUploadImagesCount Method for GetUploadImagesCount

Number of software image presented in vManage repository

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest
*/
func (a *ConfigurationDeviceSoftwareUpdateApiService) GetUploadImagesCount(ctx context.Context) ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest {
	return ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceSoftwareUpdateApiService) GetUploadImagesCountExecute(r ConfigurationDeviceSoftwareUpdateApiGetUploadImagesCountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceSoftwareUpdateApiService.GetUploadImagesCount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/software/package/imageCount"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.imageType == nil {
		return nil, reportError("imageType is required and must be specified")
	}

	{
		t := *r.imageType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("imageType", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("imageType", parameterToString(t, "multi"))
		}
	}
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

type ConfigurationDeviceSoftwareUpdateApiInstallPkgRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceSoftwareUpdateApiService
}

func (r ConfigurationDeviceSoftwareUpdateApiInstallPkgRequest) Execute() (*http.Response, error) {
	return r.ApiService.InstallPkgExecute(r)
}

/*
InstallPkg Method for InstallPkg

Install software package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ConfigurationDeviceSoftwareUpdateApiInstallPkgRequest
*/
func (a *ConfigurationDeviceSoftwareUpdateApiService) InstallPkg(ctx context.Context) ConfigurationDeviceSoftwareUpdateApiInstallPkgRequest {
	return ConfigurationDeviceSoftwareUpdateApiInstallPkgRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceSoftwareUpdateApiService) InstallPkgExecute(r ConfigurationDeviceSoftwareUpdateApiInstallPkgRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceSoftwareUpdateApiService.InstallPkg")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/software/package"

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

type ConfigurationDeviceSoftwareUpdateApiProcessSoftwareImageRequest struct {
	ctx context.Context
	ApiService *ConfigurationDeviceSoftwareUpdateApiService
	imageType string
}

func (r ConfigurationDeviceSoftwareUpdateApiProcessSoftwareImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProcessSoftwareImageExecute(r)
}

/*
ProcessSoftwareImage Method for ProcessSoftwareImage

Install software image package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageType Image type
 @return ConfigurationDeviceSoftwareUpdateApiProcessSoftwareImageRequest
*/
func (a *ConfigurationDeviceSoftwareUpdateApiService) ProcessSoftwareImage(ctx context.Context, imageType string) ConfigurationDeviceSoftwareUpdateApiProcessSoftwareImageRequest {
	return ConfigurationDeviceSoftwareUpdateApiProcessSoftwareImageRequest{
		ApiService: a,
		ctx: ctx,
		imageType: imageType,
	}
}

// Execute executes the request
func (a *ConfigurationDeviceSoftwareUpdateApiService) ProcessSoftwareImageExecute(r ConfigurationDeviceSoftwareUpdateApiProcessSoftwareImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationDeviceSoftwareUpdateApiService.ProcessSoftwareImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/action/software/package/{imageType}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageType"+"}", url.PathEscape(parameterToString(r.imageType, "")), -1)

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
