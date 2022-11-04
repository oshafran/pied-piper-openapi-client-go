/*
Cisco SD-WAN vManage API

Testing MonitoringDeviceDetailsApiService

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

func Test_openapi_MonitoringDeviceDetailsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MonitoringDeviceDetailsApiService EnableSDAVCOnDevice", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var deviceIP string
        var enable bool

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.EnableSDAVCOnDevice(context.Background(), deviceIP, enable).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GenerateDeviceStateData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var stateDataType string

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GenerateDeviceStateData(context.Background(), stateDataType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GenerateDeviceStateDataFields", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var stateDataType string

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GenerateDeviceStateDataFields(context.Background(), stateDataType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GenerateDeviceStateDataWithQueryString", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var stateDataType string

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GenerateDeviceStateDataWithQueryString(context.Background(), stateDataType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetAllDeviceStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetAllDeviceStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceCounters", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceCounters(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceListAsKeyValue", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceListAsKeyValue(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceModels", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uuid string

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceModels(context.Background(), uuid).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceOnlyStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceOnlyStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceRunningConfig", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceRunningConfig(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceRunningConfigHTML", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceRunningConfigHTML(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceTlocStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceTlocStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceTlocUtil", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceTlocUtil(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetDeviceTlocUtilDetails", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetDeviceTlocUtilDetails(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetHardwareHealthDetails", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetHardwareHealthDetails(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetHardwareHealthSummary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetHardwareHealthSummary(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetStatsQueues", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetStatsQueues(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetSyncQueues", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetSyncQueues(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetUnconfigured", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetUnconfigured(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetVManageSystemIP", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetVManageSystemIP(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetVedgeInventory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetVedgeInventory(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService GetVedgeInventorySummary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.GetVedgeInventorySummary(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService ListAllDeviceModels", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.ListAllDeviceModels(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService ListAllDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.ListAllDevices(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService ListAllMonitorDetailsDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.ListAllMonitorDetailsDevices(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService ListCurrentlySyncingDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.ListCurrentlySyncingDevices(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService ListReachableDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.ListReachableDevices(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService ListUnreachableDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.ListUnreachableDevices(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService RemoveUnreachableDevice", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var deviceIP string

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.RemoveUnreachableDevice(context.Background(), deviceIP).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService SetBlockSync", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.SetBlockSync(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringDeviceDetailsApiService SyncAllDevicesMemDB", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringDeviceDetailsApi.SyncAllDevicesMemDB(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}