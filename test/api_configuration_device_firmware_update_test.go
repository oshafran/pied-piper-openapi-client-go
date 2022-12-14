/*
Cisco SD-WAN vManage API

Testing ConfigurationDeviceFirmwareUpdateApiService

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

func Test_openapi_ConfigurationDeviceFirmwareUpdateApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService ActivateFirmwareImage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.ActivateFirmwareImage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService DeleteFirmwareImage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var versionId string

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.DeleteFirmwareImage(context.Background(), versionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService GetDevicesFWUpgrade", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.GetDevicesFWUpgrade(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService GetFirmwareImageDetails", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var versionId string

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.GetFirmwareImageDetails(context.Background(), versionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService GetFirmwareImages", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.GetFirmwareImages(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService InstallFirmwareImage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.InstallFirmwareImage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService ProcessFirmwareImage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.ProcessFirmwareImage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceFirmwareUpdateApiService RemoveFirmwareImage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceFirmwareUpdateApi.RemoveFirmwareImage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
