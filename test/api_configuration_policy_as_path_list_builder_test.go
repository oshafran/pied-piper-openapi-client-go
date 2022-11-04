/*
Cisco SD-WAN vManage API

Testing ConfigurationPolicyASPathListBuilderApiService

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

func Test_openapi_ConfigurationPolicyASPathListBuilderApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService CreatePolicyList5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.CreatePolicyList5(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService DeletePolicyList5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.DeletePolicyList5(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService DeletePolicyListsWithInfoTag5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.DeletePolicyListsWithInfoTag5(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService EditPolicyList5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.EditPolicyList5(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService GetListsById5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.GetListsById5(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService GetPolicyLists5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.GetPolicyLists5(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService GetPolicyListsWithInfoTag5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.GetPolicyListsWithInfoTag5(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService PreviewPolicyList5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.PreviewPolicyList5(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationPolicyASPathListBuilderApiService PreviewPolicyListById5", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id string

        resp, httpRes, err := apiClient.ConfigurationPolicyASPathListBuilderApi.PreviewPolicyListById5(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
