package yandex

import (
	"net/http"
	"testing"

	"github.com/objque/musicmash/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestYandexHandler_Fetch(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	releases := []*db.Release{
		{
			ID:         1,
			ArtistName: "skrillex",
			StoreID:    1433791393,
		},
	}
	// mock yandex auth
	mux.HandleFunc("/api/v2.1/handlers/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"yandexuid": "1234276871451297001"}`))
	})
	// mock itunes lookup api
	mux.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
          "resultCount": 1,
          "results": [
            {
              "artistName": "skrillex",
              "collectionName": "MDR (Remixes) - Single"
            }
          ]
        }`))
	})
	// mock yandex search artist
	mux.HandleFunc("/handlers/music-search.jsx", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
            "text": "gorgon city",
            "artists": {
                "items": [{
                    "id": 817678,
                    "name": "skrillex"
                }]
            }
        }`))
	})
	// mock yandex get artist albums
	mux.HandleFunc("/handlers/artist.jsx", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
                "artist": {
                    "id": 817678,
                    "name": "skrillex"
                },
                "albums": [{
                    "id": 5647716,
                    "title": "MDR (Remixes)",
                    "year": 2018,
                    "releaseDate": "2018-08-10T00:00:00+03:00"
                },{
                    "id": 6564,
                    "title": "The system",
                    "year": 2017,
                    "releaseDate": "2017-01-10T00:00:00+03:00"
                }]
            }`))
	})
	yandex := New(server.URL)

	// action
	yandex.Fetch(releases)

	// assert
	assert.True(t, db.DbMgr.IsReleaseExistsInStore(yandex.GetStoreName(), "5647716"))
}