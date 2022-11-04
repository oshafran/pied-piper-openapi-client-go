/*
Cisco SD-WAN vManage API

Testing ConfigurationDeviceSoftwarePackageApiService

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

func Test_openapi_ConfigurationDeviceSoftwarePackageApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ConfigurationDeviceSoftwarePackageApiService CreateVnfPackage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ConfigurationDeviceSoftwarePackageApi.CreateVnfPackage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceSoftwarePackageApiService EditConfigFile", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uuid string

        resp, httpRes, err := apiClient.ConfigurationDeviceSoftwarePackageApi.EditConfigFile(context.Background(), uuid).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceSoftwarePackageApiService GetFileContents", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uuid string

        resp, httpRes, err := apiClient.ConfigurationDeviceSoftwarePackageApi.GetFileContents(context.Background(), uuid).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ConfigurationDeviceSoftwarePackageApiService UploadImageFile", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.ConfigurationDeviceSoftwarePackageApi.UploadImageFile(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}