/*
SpaceTraders API

Testing SystemsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package spacetradersapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_spacetradersapi_SystemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemsAPIService GetConstruction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetConstruction(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetJumpGate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetJumpGate(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetMarket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetMarket(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetShipyard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetShipyard(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetSystem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetSystem(context.Background(), systemSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetSystemWaypoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetSystemWaypoints(context.Background(), systemSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetSystems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemsAPI.GetSystems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService GetWaypoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.GetWaypoint(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsAPIService SupplyConstruction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemSymbol string
		var waypointSymbol string

		resp, httpRes, err := apiClient.SystemsAPI.SupplyConstruction(context.Background(), systemSymbol, waypointSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
