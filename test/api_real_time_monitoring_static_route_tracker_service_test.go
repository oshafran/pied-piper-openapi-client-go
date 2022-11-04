/*
Cisco SD-WAN vManage API

Testing RealTimeMonitoringStaticRouteTrackerServiceApiService

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

func Test_openapi_RealTimeMonitoringStaticRouteTrackerServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RealTimeMonitoringStaticRouteTrackerServiceApiService GetStaticRouteTrackerInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RealTimeMonitoringStaticRouteTrackerServiceApi.GetStaticRouteTrackerInfo(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
