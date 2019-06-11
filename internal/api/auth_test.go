package api

import (
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/testutils"
	"github.com/musicmash/auth/pkg/api"
	"github.com/musicmash/auth/pkg/api/auth"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Auth(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	endpointCalled := false
	authProvider = api.NewProvider(env.Server.URL)
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/auth",
		RawResponse: `{"token": "123"}`,
		Method:      http.MethodPost,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &endpointCalled,
	})
	authProvider = api.NewProvider(env.Server.URL)
	apiProvider := api.NewProvider(server.URL)

	// action
	token, err := auth.Auth(apiProvider, &auth.Payload{Service: "google", Token: "xxxx"})

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "123", token.Token)
}
