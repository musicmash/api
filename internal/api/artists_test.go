package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/testutils"
	"github.com/musicmash/api/pkg/api/artists/search"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Artists_Search(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock artists service api
	artistsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/v1/search",
		RawResponse: fmt.Sprintf(`[{"id":69, "name":"%s"}]`, testutils.ArtistSkrillex),
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &artistsServiceCalled,
	})

	// action
	artists, err := search.Do(client, testutils.ArtistSkrillex)

	// assert
	assert.NoError(t, err)
	assert.True(t, artistsServiceCalled)
	assert.Len(t, artists, 1)
	assert.Equal(t, int64(69), artists[0].ID)
	assert.Equal(t, testutils.ArtistSkrillex, artists[0].Name)
}

func TestAPI_Artists_Search_NameWithSpaces(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock artists service api
	artistsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/v1/search",
		RawResponse: fmt.Sprintf(`[{"id":69, "name":"%s"}]`, testutils.ArtistWolvesAtTheGate),
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &artistsServiceCalled,
	})

	// action
	artists, err := search.Do(client, testutils.ArtistWolvesAtTheGate)

	// assert
	assert.NoError(t, err)
	assert.True(t, artistsServiceCalled)
	assert.Len(t, artists, 1)
	assert.Equal(t, int64(69), artists[0].ID)
	assert.Equal(t, testutils.ArtistWolvesAtTheGate, artists[0].Name)
}

func TestAPI_Artists_Search_Empty(t *testing.T) {
	setup()
	defer teardown()

	// action
	artists, err := search.Do(client, "")

	// assert
	assert.Error(t, err)
	assert.Nil(t, artists)
}
