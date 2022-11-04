/*
Cisco SD-WAN vManage API

Testing ConfigurationPolicyQosMapDefinitionBuilderApiService

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

func Test_openapi_ConfigurationPolicyQosMapDefinitionBuilderApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService CreatePolicyDefinition1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.CreatePolicyDefinition1(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService DeletePolicyDefinition1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.DeletePolicyDefinition1(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService EditMultiplePolicyDefinition1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.EditMultiplePolicyDefinition1(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService EditPolicyDefinition1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.EditPolicyDefinition1(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService GetDefinitions1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.GetDefinitions1(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService GetPolicyDefinition1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.GetPolicyDefinition1(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService PreviewPolicyDefinition1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinition1(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService PreviewPolicyDefinitionById1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.PreviewPolicyDefinitionById1(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyQosMapDefinitionBuilderApiService SavePolicyDefinitionInBulk1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyQosMapDefinitionBuilderApi.SavePolicyDefinitionInBulk1(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
