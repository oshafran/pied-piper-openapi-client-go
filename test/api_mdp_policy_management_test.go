/*
Cisco SD-WAN vManage API

Testing MDPPolicyManagementApiService

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

func Test_openapi_MDPPolicyManagementApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MDPPolicyManagementApiService RetrieveMDPPolicies", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.MDPPolicyManagementApi.RetrieveMDPPolicies(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MDPPolicyManagementApiService UpdatePolicyStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.MDPPolicyManagementApi.UpdatePolicyStatus(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
