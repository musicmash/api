package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDB_Users_EnsureExists(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := DbMgr.EnsureUserExists("objque@me")

	// assert
	assert.NoError(t, err)
	user, err := DbMgr.FindUserByName("objque@me")
	assert.NoError(t, err)
	assert.Equal(t, "objque@me", user.Name)
}

func TestDB_Users_List(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	assert.NoError(t, DbMgr.EnsureUserExists("objque@me"))
	assert.NoError(t, DbMgr.EnsureUserExists("jade@abuse"))

	// action
	users, err := DbMgr.GetAllUsers()

	// assert
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "objque@me", users[0].Name)
	assert.Equal(t, "jade@abuse", users[1].Name)
}

func TestDB_Users_GetUsersWithReleases(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	const artist = "Architects"
	now := time.Now().UTC()
	assert.NoError(t, DbMgr.EnsureUserExists("objque@me"))
	assert.NoError(t, DbMgr.EnsureUserExists("jade@abuse"))
	assert.NoError(t, DbMgr.EnsureUserExists("jake@worrow"))
	assert.NoError(t, DbMgr.EnsureArtistExists("architects"))
	assert.NoError(t, DbMgr.EnsureSubscriptionExists("objque@me", artist))
	assert.NoError(t, DbMgr.EnsureSubscriptionExists("jake@worrow", artist))
	assert.NoError(t, DbMgr.EnsureReleaseExists(&Release{ArtistName: artist, CreatedAt: now}))

	// action
	users, err := DbMgr.GetUsersWithReleases(now)

	// assert
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.EqualValues(t, []string{"jake@worrow", "objque@me"}, users)
}

func TestDB_Users_GetUsersWithReleases_NoReleases_ForProvidedHour(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	const artist = "Architects"
	now := time.Now().UTC()
	assert.NoError(t, DbMgr.EnsureUserExists("objque@me"))
	assert.NoError(t, DbMgr.EnsureUserExists("jade@abuse"))
	assert.NoError(t, DbMgr.EnsureUserExists("jake@worrow"))
	assert.NoError(t, DbMgr.EnsureArtistExists("architects"))
	assert.NoError(t, DbMgr.EnsureSubscriptionExists("objque@me", artist))
	assert.NoError(t, DbMgr.EnsureSubscriptionExists("jake@worrow", artist))
	assert.NoError(t, DbMgr.EnsureReleaseExists(&Release{ArtistName: artist, CreatedAt: now.Add(-time.Hour)}))

	// action
	users, err := DbMgr.GetUsersWithReleases(now)

	// assert
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}
