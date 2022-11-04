/*
Cisco SD-WAN vManage API

Testing RealTimeMonitoringDeviceFeatureListApiService

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

func Test_openapi_RealTimeMonitoringDeviceFeatureListApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RealTimeMonitoringDeviceFeatureListApiService GetFeatureList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RealTimeMonitoringDeviceFeatureListApi.GetFeatureList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RealTimeMonitoringDeviceFeatureListApiService GetSyncedFeatureList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RealTimeMonitoringDeviceFeatureListApi.GetSyncedFeatureList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}