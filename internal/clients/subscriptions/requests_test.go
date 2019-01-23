package subscriptions

import (
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestClients_Subscriptions_Get(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	answer := `[{"artist_name": "The Pierces"}, {"artist_name": "The Prodigy"}]`
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/objque@me/subscriptions",
		RawResponse: answer,
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	subs, err := Get(client, "objque@me")

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
	assert.Len(t, subs, 2)
	want := []*Subscription{{ArtistName: "The Pierces"}, {ArtistName: "The Prodigy"}}
	assert.Equal(t, want, subs)
}

func TestClients_Subscriptions_Subscribe(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/objque@me/subscriptions",
		Method:     http.MethodPost,
		HTTPStatus: http.StatusAccepted,
		CallFlag:   &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	err := Subscribe(client, "objque@me", []string{"the prodigy"})

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
}

func TestClients_Subscriptions_UnSubscribe(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/objque@me/subscriptions",
		Method:     http.MethodDelete,
		HTTPStatus: http.StatusOK,
		CallFlag:   &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	err := UnSubscribe(client, "objque@me", []string{"the prodigy"})

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
}
