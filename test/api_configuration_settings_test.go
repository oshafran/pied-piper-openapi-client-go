/*
Cisco SD-WAN vManage API

Testing ConfigurationSettingsApiService

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

func Test_openapi_ConfigurationSettingsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationSettingsApiService CreateAnalyticsDataFile", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.CreateAnalyticsDataFile(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService EditCertConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var settingType string

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.EditCertConfiguration(context.Background(), settingType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService EditConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var settingType string

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.EditConfiguration(context.Background(), settingType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService GetBanner", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.GetBanner(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService GetCertConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var settingType string

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.GetCertConfiguration(context.Background(), settingType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService GetConfigurationBySettingType", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var settingType string

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.GetConfigurationBySettingType(context.Background(), settingType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService GetGoogleMapKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.GetGoogleMapKey(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService GetMaintenanceWindow", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.GetMaintenanceWindow(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService GetSessionTimout", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.GetSessionTimout(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService NewCertConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var settingType string

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.NewCertConfiguration(context.Background(), settingType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationSettingsApiService NewConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var settingType string

        resp, httpRes, err := apiClient.ConfigurationSettingsApi.NewConfiguration(context.Background(), settingType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
