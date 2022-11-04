/*
Cisco SD-WAN vManage API

Testing ConfigurationNetworkDesignApiService

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

func Test_openapi_ConfigurationNetworkDesignApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationNetworkDesignApiService AcquireAttachLock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var profileId string

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.AcquireAttachLock(context.Background(), profileId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService AcquireEditLock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.AcquireEditLock(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService CreateNetworkDesign", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.CreateNetworkDesign(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService EditNetworkDesign", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.EditNetworkDesign(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetDeviceProfileConfigStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetDeviceProfileConfigStatusByProfileId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var profileId string

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileConfigStatusByProfileId(context.Background(), profileId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetDeviceProfileTaskCount", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileTaskCount(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetDeviceProfileTaskStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetDeviceProfileTaskStatusByProfileId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var profileId string

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetDeviceProfileTaskStatusByProfileId(context.Background(), profileId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetGlobalParameters", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetGlobalParameters(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetNetworkDesign", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetNetworkDesign(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService GetServiceProfileConfig", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var profileId string

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.GetServiceProfileConfig(context.Background(), profileId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService PushDeviceProfileTemplate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var profileId string

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.PushDeviceProfileTemplate(context.Background(), profileId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService PushNetworkDesign", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.PushNetworkDesign(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationNetworkDesignApiService RunMyTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.ConfigurationNetworkDesignApi.RunMyTest(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}