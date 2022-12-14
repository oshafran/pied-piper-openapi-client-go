/*
Cisco SD-WAN vManage API

Testing ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService

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

func Test_openapi_ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService CreatePolicyList9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.CreatePolicyList9(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService DeletePolicyList9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.DeletePolicyList9(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService DeletePolicyListsWithInfoTag9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.DeletePolicyListsWithInfoTag9(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService EditPolicyList9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.EditPolicyList9(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService GetListsById9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.GetListsById9(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService GetListsForAllDataPrefixes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.GetListsForAllDataPrefixes(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService GetPolicyListsWithInfoTag9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.GetPolicyListsWithInfoTag9(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService PreviewPolicyList9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.PreviewPolicyList9(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApiService PreviewPolicyListById9", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyDataPrefixAllTypesOfListsBuilderApi.PreviewPolicyListById9(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
