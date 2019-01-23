package feed

import (
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/clients"
	"github.com/musicmash/api/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestClients_Feed_Get(t *testing.T) {
	env := testutils.Setup()
	defer env.TearDown()

	// arrange
	answer :=
		`{
  "announced": [
    {
      "artist_name": "Juice WRLD",
      "poster": "http://pic.cdn",
      "released": "2019-02-01T00:00:00Z",
      "stores": [
        {
          "id": "1449799389",
          "name": "itunes",
          "url": "https://itunes.apple.com/us/album/1449799389"
        }
      ],
      "title": "Juice WRLD"
    }
  ],
  "date": "2019-01-16T18:31:17.610881069Z",
  "released": [
    {
      "artist_name": "Lil Keed",
      "poster": "http://pic.cdn",
      "released": "2019-01-21T00:00:00Z",
      "stores": [
        {
          "id": "1450198294",
          "name": "itunes",
          "url": "https://itunes.apple.com/us/album/1450198294"
        }
      ],
      "title": "No One Standing (feat. Lil Keed)"
    }
  ]
}`
	endpointCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/objque@me/feed",
		RawResponse: answer,
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &endpointCalled,
	})
	client := clients.NewProvider(env.Server.URL)

	// action
	feed, err := Get(client, "objque@me", nil)

	// assert
	assert.NoError(t, err)
	assert.True(t, endpointCalled)
	assert.Len(t, feed.Announced, 1)
	assert.Len(t, feed.Released, 1)
}
