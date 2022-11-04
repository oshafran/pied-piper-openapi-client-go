/*
Cisco SD-WAN vManage API

Testing MonitoringApplicationAwareRoutingApiService

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

func Test_openapi_MonitoringApplicationAwareRoutingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetAggregationDataAppRoute", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetAggregationDataAppRoute(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetAggregationDataByQuery2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetAggregationDataByQuery2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetApprouteGridStat", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetApprouteGridStat(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetCount4", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetCount4(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetCountPost4", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetCountPost4(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetPostAggregationAppDataByQuery2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetPostAggregationAppDataByQuery2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetPostAggregationDataByQuery2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetPostAggregationDataByQuery2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetPostStatBulkRawData2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetPostStatBulkRawData2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetStatBulkRawData2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetStatBulkRawData2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetStatDataFields4", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetStatDataFields4(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetStatDataRawData2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetStatDataRawData2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetStatDataRawDataAsCSV2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetStatDataRawDataAsCSV2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetStatQueryFields4", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetStatQueryFields4(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetStatsRawData2", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetStatsRawData2(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTransportHealth", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTransportHealth(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTransportHealthSummary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTransportHealthSummary(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTunnel", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTunnel(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTunnelChart", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTunnelChart(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTunnels", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTunnels(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTunnelsHealth", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTunnelsHealth(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MonitoringApplicationAwareRoutingApiService GetTunnelsSummary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var type_ string

        resp, httpRes, err := apiClient.MonitoringApplicationAwareRoutingApi.GetTunnelsSummary(context.Background(), type_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
