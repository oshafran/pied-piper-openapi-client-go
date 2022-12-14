/*
Cisco SD-WAN vManage API

Testing TroubleshootingToolsDiagnosticsApiService

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

func Test_openapi_TroubleshootingToolsDiagnosticsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TroubleshootingToolsDiagnosticsApiService ClearSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.ClearSession(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService DisableDeviceLog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.DisableDeviceLog(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService DisablePacketCaptureSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.DisablePacketCaptureSession(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService DisableSpeedTestSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.DisableSpeedTestSession(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService DownloadDebugLog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.DownloadDebugLog(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService DownloadFile", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.DownloadFile(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService ForceStopPcapSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.ForceStopPcapSession(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService FormPostPacketCapture", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var deviceUUID string
        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.FormPostPacketCapture(context.Background(), deviceUUID, sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetAggregationDataByQuery27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetAggregationDataByQuery28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetAggregationDataByQuery28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetConcurrentData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetConcurrentData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetConcurrentDomainData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetConcurrentDomainData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetCount29", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCount29(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetCount30", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCount30(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetCountPost29", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCountPost29(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetCountPost30", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetCountPost30(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetDBSchema", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetDBSchema(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetDeviceLog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetDeviceLog(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetDomainMetric", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetDomainMetric(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetFileDownloadStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFileDownloadStatus(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetFinalizedData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFinalizedData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetFinalizedDomainData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFinalizedDomainData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetFlowDetail", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFlowDetail(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetFlowMetric", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetFlowMetric(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetInterfaceBandwidth", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetInterfaceBandwidth(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetLogType", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetLogType(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetMonitorState", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetMonitorState(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetNwpiDscp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetNwpiDscp(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetNwpiNbarAppGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetNwpiNbarAppGroup(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetNwpiProtocol", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetNwpiProtocol(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPacketFeatures", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPacketFeatures(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPostAggregationAppDataByQuery26", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery26(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPostAggregationAppDataByQuery27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationAppDataByQuery27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPostAggregationDataByQuery27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPostAggregationDataByQuery28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostAggregationDataByQuery28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPostStatBulkRawData27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetPostStatBulkRawData28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetPostStatBulkRawData28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSession(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetSessionInfoCapture", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSessionInfoCapture(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetSessionInfoLog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSessionInfoLog(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetSessions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSessions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetSpeedTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSpeedTest(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetSpeedTestStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetSpeedTestStatus(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatBulkRawData27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatBulkRawData28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatBulkRawData28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatDataFields29", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataFields29(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatDataFields30", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataFields30(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatDataRawData26", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawData26(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatDataRawData27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawData27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatDataRawDataAsCSV27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatDataRawDataAsCSV28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatDataRawDataAsCSV28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatQueryFields29", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatQueryFields29(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatQueryFields30", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatQueryFields30(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatsRawData27", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatsRawData27(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetStatsRawData28", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetStatsRawData28(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetThreadPools", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetThreadPools(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetTraceFlow", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetTraceFlow(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetTraceHistory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetTraceHistory(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService GetVnicInfoByVnfId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var vnfId string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.GetVnicInfoByVnfId(context.Background(), vnfId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService MonitorStart", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.MonitorStart(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService MonitorStop", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.MonitorStop(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService ProcessDeviceStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var deviceUUID string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.ProcessDeviceStatus(context.Background(), deviceUUID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService RenewSessionInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.RenewSessionInfo(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService SaveSpeedTestResults", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var deviceUUID string
        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.SaveSpeedTestResults(context.Background(), deviceUUID, sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService SearchDeviceLog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.SearchDeviceLog(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService StartPcapSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.StartPcapSession(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService StartSpeedTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.StartSpeedTest(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService StopPcapSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.StopPcapSession(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService StopSpeedTest", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId Uuid

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.StopSpeedTest(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService StreamLog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var logType string
        var deviceUUID string
        var sessionId string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.StreamLog(context.Background(), logType, deviceUUID, sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService TraceDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceDelete(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService TraceFinFlowWithQuery", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceFinFlowWithQuery(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService TraceStart", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceStart(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TroubleshootingToolsDiagnosticsApiService TraceStop", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var traceId string

        resp, httpRes, err := apiClient.TroubleshootingToolsDiagnosticsApi.TraceStop(context.Background(), traceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
