/*
Cisco SD-WAN vManage API

Testing ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService

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

func Test_openapi_ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService CreatePolicyDefinition28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.CreatePolicyDefinition28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService DeletePolicyDefinition28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.DeletePolicyDefinition28(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService EditMultiplePolicyDefinition28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.EditMultiplePolicyDefinition28(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService EditPolicyDefinition28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.EditPolicyDefinition28(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService GetDefinitions28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.GetDefinitions28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService GetPolicyDefinition28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.GetPolicyDefinition28(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService PreviewPolicyDefinition28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.PreviewPolicyDefinition28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService PreviewPolicyDefinitionById28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.PreviewPolicyDefinitionById28(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyFXSDIDPortDefinitionBuilderApiService SavePolicyDefinitionInBulk28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyFXSDIDPortDefinitionBuilderApi.SavePolicyDefinitionInBulk28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
