/*
Cisco SD-WAN vManage API

Testing MDPOnboardingApiService

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

func Test_openapi_MDPOnboardingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MDPOnboardingApiService GetMDPOnboardingStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MDPOnboardingApi.GetMDPOnboardingStatus(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MDPOnboardingApiService OnboardMDP", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MDPOnboardingApi.OnboardMDP(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MDPOnboardingApiService UpdateOnboardingPayload", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nmsId string

        resp, httpRes, err := apiClient.MDPOnboardingApi.UpdateOnboardingPayload(context.Background(), nmsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
