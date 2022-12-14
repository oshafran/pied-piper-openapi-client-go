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
)


// RealTimeMonitoringStaticRouteTrackerServiceApiService RealTimeMonitoringStaticRouteTrackerServiceApi service
type RealTimeMonitoringStaticRouteTrackerServiceApiService service

type RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest struct {
	ctx context.Context
	ApiService *RealTimeMonitoringStaticRouteTrackerServiceApiService
	deviceId *string
}

// Device Id
func (r RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest) DeviceId(deviceId string) RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest {
	r.deviceId = &deviceId
	return r
}

func (r RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetStaticRouteTrackerInfoExecute(r)
}

/*
GetStaticRouteTrackerInfo Method for GetStaticRouteTrackerInfo

Get single static route tracker info from device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest
*/
func (a *RealTimeMonitoringStaticRouteTrackerServiceApiService) GetStaticRouteTrackerInfo(ctx context.Context) RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest {
	return RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RealTimeMonitoringStaticRouteTrackerServiceApiService) GetStaticRouteTrackerInfoExecute(r RealTimeMonitoringStaticRouteTrackerServiceApiGetStaticRouteTrackerInfoRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RealTimeMonitoringStaticRouteTrackerServiceApiService.GetStaticRouteTrackerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device/staticRouteTracker"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return localVarReturnValue, nil, reportError("deviceId is required and must be specified")
	}

	localVarQueryParams.Add("deviceId", parameterToString(*r.deviceId, ""))
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
