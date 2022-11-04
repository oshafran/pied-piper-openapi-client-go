/*
Cisco SD-WAN vManage API

Testing TenantManagementApiService

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

func Test_openapi_TenantManagementApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TenantManagementApiService CreateTenant", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.CreateTenant(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService CreateTenantAsync", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.CreateTenantAsync(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService CreateTenantAsyncBulk", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.CreateTenantAsyncBulk(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService DeleteTenant", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.TenantManagementApi.DeleteTenant(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService DeleteTenantAsyncBulk", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.DeleteTenantAsyncBulk(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService ForceStatusCollection", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.ForceStatusCollection(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService GetAllTenantStatuses", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.GetAllTenantStatuses(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService GetAllTenants", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.GetAllTenants(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService GetTenant", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.TenantManagementApi.GetTenant(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService GetTenantHostingCapacityOnvSmarts", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.GetTenantHostingCapacityOnvSmarts(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService GetTenantvSmartMapping", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.GetTenantvSmartMapping(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService SwitchTenant", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.TenantManagementApi.SwitchTenant(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService TenantvSmartMtMigrate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TenantManagementApi.TenantvSmartMtMigrate(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService UpdateTenant", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.TenantManagementApi.UpdateTenant(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TenantManagementApiService VSessionId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var tenantId string

        resp, httpRes, err := apiClient.TenantManagementApi.VSessionId(context.Background(), tenantId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
