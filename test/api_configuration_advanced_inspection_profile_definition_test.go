/*
Cisco SD-WAN vManage API

Testing ConfigurationAdvancedInspectionProfileDefinitionApiService

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

func Test_openapi_ConfigurationAdvancedInspectionProfileDefinitionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService CreatePolicyDefinition10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.CreatePolicyDefinition10(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService DeletePolicyDefinition10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.DeletePolicyDefinition10(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService EditMultiplePolicyDefinition10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.EditMultiplePolicyDefinition10(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService EditPolicyDefinition10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.EditPolicyDefinition10(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService GetDefinitions10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.GetDefinitions10(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService GetPolicyDefinition10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.GetPolicyDefinition10(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService PreviewPolicyDefinition10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.PreviewPolicyDefinition10(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService PreviewPolicyDefinitionById10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.PreviewPolicyDefinitionById10(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationAdvancedInspectionProfileDefinitionApiService SavePolicyDefinitionInBulk10", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationAdvancedInspectionProfileDefinitionApi.SavePolicyDefinitionInBulk10(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
