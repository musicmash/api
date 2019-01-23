package artists

import (
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestClients_Artists_Search(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	answer := `[{"name": "Slipknot","poster": "https://cdn/pic"}]`
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/objque@me/artists",
		RawResponse: answer,
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	artists, err := Search(client, "objque@me", "slipknot")

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
	assert.Len(t, artists, 1)
	assert.Equal(t, "Slipknot", artists[0].Name)
	assert.Equal(t, "https://cdn/pic", artists[0].Poster)
}

func TestClients_Artists_GetDetails(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	answer := `{
  "name": "Slipknot",
  "poster": "https://cdn/pic",
  "releases": {
    "announced": [],
    "released": [
      {
        "artist_name": "Slipknot",
        "poster": "https://cdn/pic",
        "released": "2018-10-31T00:00:00Z",
        "stores": [
          {
            "id": "77287002",
            "name": "deezer",
            "url": "https://deezer.com/en/album/77287002"
          },
          {
            "id": "1440513655",
            "name": "itunes",
            "url": "https://itunes.apple.com/us/album/1440513655"
          }
        ],
        "title": "All Out Life"
      }
    ]
  },
  "stores": [
    {
      "id": "6907568",
      "name": "itunes",
      "url": "https://itunes.apple.com/us/artist/6907568"
    },
    {
      "id": "117",
      "name": "deezer",
      "url": "https://deezer.com/artist/117"
    }
  ]
}`
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/objque@me/artists/slipknot",
		RawResponse: answer,
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	details, err := GetDetails(client, "objque@me", "slipknot")

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
	assert.Equal(t, "Slipknot", details.Name)
	assert.Equal(t, "https://cdn/pic", details.Poster)
	assert.Empty(t, details.Releases.Announced)
	assert.Len(t, details.Releases.Recent, 1)
	assert.Equal(t, "Slipknot", details.Releases.Recent[0].ArtistName)
	assert.Equal(t, "https://cdn/pic", details.Releases.Recent[0].Poster)
	assert.Len(t, details.Releases.Recent[0].Stores, 2)
	assert.Equal(t, "All Out Life", details.Releases.Recent[0].Title)
	assert.Len(t, details.Stores, 2)
}
