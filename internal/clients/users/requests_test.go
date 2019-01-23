package users

import (
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestClients_Users_Get(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/users/objque@me",
		Method:     http.MethodGet,
		HTTPStatus: http.StatusOK,
		CallFlag:   &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	err := Get(client, "objque@me")

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
}

func TestClients_Users_Get_NotFound(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/users/objque@me",
		Method:     http.MethodGet,
		HTTPStatus: http.StatusNotFound,
		CallFlag:   &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	err := Get(client, "objque@me")

	// assert
	assert.Error(t, err)
	assert.True(t, endpointCalled)
}
