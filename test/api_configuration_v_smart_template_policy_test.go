/*
Cisco SD-WAN vManage API

Testing ConfigurationVSmartTemplatePolicyApiService

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

func Test_openapi_ConfigurationVSmartTemplatePolicyApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService ActivatePolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.ActivatePolicy(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService ActivatePolicyForCloudServices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.ActivatePolicyForCloudServices(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService CheckVSmartConnectivityStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.CheckVSmartConnectivityStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService CreateVSmartTemplate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.CreateVSmartTemplate(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService DeActivatePolicy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.DeActivatePolicy(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService DeleteVSmartTemplate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.DeleteVSmartTemplate(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService EditTemplateWithoutLockChecks", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.EditTemplateWithoutLockChecks(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService EditVSmartTemplate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.EditVSmartTemplate(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService GenerateVSmartPolicyTemplateList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.GenerateVSmartPolicyTemplateList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService GetTemplateByPolicyId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var policyId string

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.GetTemplateByPolicyId(context.Background(), policyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationVSmartTemplatePolicyApiService QosmosNbarMigrationWarning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationVSmartTemplatePolicyApi.QosmosNbarMigrationWarning(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
