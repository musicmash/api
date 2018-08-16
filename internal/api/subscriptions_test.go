package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/objque/musicmash/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestAPI_SubscribeUser(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureUserExists("objque"))
	assert.NoError(t, db.DbMgr.EnsureArtistExists(&db.Artist{Name: "Moderat", StoreID: 00001}))

	// action
	artists := []string{"modeRat"}
	buffer, _ := json.Marshal(&artists)
	resp, err := http.Post(fmt.Sprintf("%s/objque/subscriptions", server.URL), "application/json", bytes.NewReader(buffer))

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusAccepted)
	body := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.Len(t, body, 1)
}

func TestAPI_SubscribeUser_UserNotFound(t *testing.T) {
	setup()
	defer teardown()

	// action
	resp, err := http.Post(fmt.Sprintf("%s/objque/subscriptions", server.URL), "application/json", nil)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestAPI_UnsubscribeUser(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, db.DbMgr.EnsureUserExists("objque"))
	assert.NoError(t, db.DbMgr.EnsureSubscriptionExists(&db.Subscription{UserID: "objque", ArtistName: "Skrillex"}))
	assert.NoError(t, db.DbMgr.EnsureSubscriptionExists(&db.Subscription{UserID: "objque", ArtistName: "Calvin Risk"}))
	assert.NoError(t, db.DbMgr.EnsureSubscriptionExists(&db.Subscription{UserID: "objque", ArtistName: "AC/DC"}))
	assert.NoError(t, db.DbMgr.EnsureSubscriptionExists(&db.Subscription{UserID: "mike", ArtistName: "AC/DC"}))

	// action
	artists := []string{"Calvin Risk"}
	buffer, _ := json.Marshal(&artists)
	resp, err := httpDelete(fmt.Sprintf("%s/objque/subscriptions", server.URL), bytes.NewReader(buffer))

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	subs, err := db.DbMgr.FindAllUserSubscriptions("mike")
	assert.NoError(t, err)
	assert.Len(t, subs, 1)
	assert.Equal(t, "AC/DC", subs[0].ArtistName)

	subs, err = db.DbMgr.FindAllUserSubscriptions("objque")
	assert.NoError(t, err)
	assert.Len(t, subs, 2)
	assert.Equal(t, "AC/DC", subs[0].ArtistName)
	assert.Equal(t, "Skrillex", subs[1].ArtistName)
}

func TestAPI_UnsubscribeUser_UserNotFound(t *testing.T) {
	setup()
	defer teardown()

	// action
	resp, err := httpDelete(fmt.Sprintf("%s/objque/subscriptions", server.URL), nil)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}
