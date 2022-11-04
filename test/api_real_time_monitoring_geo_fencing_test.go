/*
Cisco SD-WAN vManage API

Testing RealTimeMonitoringGeoFencingApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_RealTimeMonitoringGeoFencingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RealTimeMonitoringGeoFencingApiService GetGeofenceStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RealTimeMonitoringGeoFencingApi.GetGeofenceStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
