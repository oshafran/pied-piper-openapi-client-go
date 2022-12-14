/*
Cisco SD-WAN vManage API

Testing MonitoringDeviceStatisticsDetailsApiService

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

func Test_openapi_MonitoringDeviceStatisticsDetailsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MonitoringDeviceStatisticsDetailsApiService GenerateDeviceStatisticsData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var stateDataType string

        resp, httpRes, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GenerateDeviceStatisticsData(context.Background(), stateDataType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceStatisticsDetailsApiService GetActiveAlarms", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetActiveAlarms(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceStatisticsDetailsApiService GetCountWithStateDataType", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var stateDataType string

        resp, httpRes, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetCountWithStateDataType(context.Background(), stateDataType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceStatisticsDetailsApiService GetStatDataFieldsByStateDataType", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var stateDataType string

        resp, httpRes, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetStatDataFieldsByStateDataType(context.Background(), stateDataType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceStatisticsDetailsApiService GetStatisticsType", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceStatisticsDetailsApi.GetStatisticsType(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
