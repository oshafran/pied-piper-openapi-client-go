/*
Cisco SD-WAN vManage API

Testing SystemReverseProxyApiService

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

func Test_openapi_SystemReverseProxyApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SystemReverseProxyApiService CreateReverseProxyMappings", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uuid string

        resp, httpRes, err := apiClient.SystemReverseProxyApi.CreateReverseProxyMappings(context.Background(), uuid).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SystemReverseProxyApiService GetReverseProxyMappings", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uuid string

        resp, httpRes, err := apiClient.SystemReverseProxyApi.GetReverseProxyMappings(context.Background(), uuid).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
