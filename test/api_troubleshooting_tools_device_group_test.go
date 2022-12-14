/*
Cisco SD-WAN vManage API

Testing TroubleshootingToolsDeviceGroupApiService

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

func Test_openapi_TroubleshootingToolsDeviceGroupApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TroubleshootingToolsDeviceGroupApiService ListDeviceGroupList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListDeviceGroupList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDeviceGroupApiService ListDeviceGroups", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListDeviceGroups(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDeviceGroupApiService ListGroupDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListGroupDevices(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDeviceGroupApiService ListGroupDevicesForMap", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListGroupDevicesForMap(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDeviceGroupApiService ListGroupLinksForMap", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDeviceGroupApi.ListGroupLinksForMap(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
