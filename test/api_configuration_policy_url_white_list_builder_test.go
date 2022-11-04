/*
Cisco SD-WAN vManage API

Testing ConfigurationPolicyURLWhiteListBuilderApiService

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

func Test_openapi_ConfigurationPolicyURLWhiteListBuilderApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService CreatePolicyList38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.CreatePolicyList38(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService DeletePolicyList38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.DeletePolicyList38(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService DeletePolicyListsWithInfoTag38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.DeletePolicyListsWithInfoTag38(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService EditPolicyList38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.EditPolicyList38(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService GetListsById38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.GetListsById38(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService GetPolicyLists35", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.GetPolicyLists35(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService GetPolicyListsWithInfoTag38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.GetPolicyListsWithInfoTag38(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService PreviewPolicyList38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.PreviewPolicyList38(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyURLWhiteListBuilderApiService PreviewPolicyListById38", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyURLWhiteListBuilderApi.PreviewPolicyListById38(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}