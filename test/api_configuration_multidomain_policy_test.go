/*
Cisco SD-WAN vManage API

Testing ConfigurationMultidomainPolicyApiService

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

func Test_openapi_ConfigurationMultidomainPolicyApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationMultidomainPolicyApiService AddInternalPolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.AddInternalPolicy(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService AttachDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.AttachDevices(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService DetachDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.DetachDevices(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService DisconnectFromMdp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.DisconnectFromMdp(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService EditAttachedDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.EditAttachedDevices(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService GetMDPOnboardingStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.GetMDPOnboardingStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService Offboard", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.Offboard(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService OnboardMDP", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.OnboardMDP(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService RetrieveMDPAttachedDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPAttachedDevices(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService RetrieveMDPConfigObject", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var deviceId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPConfigObject(context.Background(), deviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService RetrieveMDPPolicies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPPolicies(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService RetrieveMDPSupportedDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.RetrieveMDPSupportedDevices(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService UpdateOnboardingPayload", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.UpdateOnboardingPayload(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationMultidomainPolicyApiService UpdatePolicyStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.ConfigurationMultidomainPolicyApi.UpdatePolicyStatus(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
