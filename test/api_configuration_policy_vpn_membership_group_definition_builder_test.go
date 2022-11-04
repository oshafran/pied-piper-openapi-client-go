/*
Cisco SD-WAN vManage API

Testing ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService

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

func Test_openapi_ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService CreatePolicyDefinition6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.CreatePolicyDefinition6(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService DeletePolicyDefinition6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.DeletePolicyDefinition6(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService EditMultiplePolicyDefinition6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.EditMultiplePolicyDefinition6(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService EditPolicyDefinition6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.EditPolicyDefinition6(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService GetDefinitions6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.GetDefinitions6(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService GetPolicyDefinition6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.GetPolicyDefinition6(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService PreviewPolicyDefinition6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.PreviewPolicyDefinition6(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService PreviewPolicyDefinitionById6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.PreviewPolicyDefinitionById6(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApiService SavePolicyDefinitionInBulk6", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyVPNMembershipGroupDefinitionBuilderApi.SavePolicyDefinitionInBulk6(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
