package api

import (
	"net/http"
	"testing"

	"github.com/musicmash/api/internal/testutils"
	"github.com/musicmash/api/pkg/api/subscriptions"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Subscriptions_Get(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock subscriptions service api
	subscriptionsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/v1/subscriptions",
		RawResponse: `[{"artist_id": 1}, {"artist_id": 2}, {"artist_id": 3}]`,
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &subscriptionsServiceCalled,
	})

	// action
	subs, err := subscriptions.Get(client, testutils.UserObjque)

	// assert
	assert.NoError(t, err)
	assert.True(t, subscriptionsServiceCalled)
	assert.Len(t, subs, 3)
}

func TestAPI_Subscriptions_Get_Empty(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock subscriptions service api
	subscriptionsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:         env.Mux,
		URL:         "/v1/subscriptions",
		RawResponse: `[]`,
		Method:      http.MethodGet,
		HTTPStatus:  http.StatusOK,
		CallFlag:    &subscriptionsServiceCalled,
	})

	// action
	subs, err := subscriptions.Get(client, testutils.UserObjque)

	// assert
	assert.NoError(t, err)
	assert.True(t, subscriptionsServiceCalled)
	assert.Len(t, subs, 0)
}

func TestAPI_Subscriptions_Create(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock subscriptions service api
	subscriptionsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/v1/subscriptions",
		Method:     http.MethodPost,
		HTTPStatus: http.StatusOK,
		CallFlag:   &subscriptionsServiceCalled,
	})

	// action
	err := subscriptions.Create(client, testutils.UserObjque, []int64{1, 2, 3})

	// assert
	assert.NoError(t, err)
	assert.True(t, subscriptionsServiceCalled)
}

func TestAPI_Subscriptions_Create_WithLimit(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock subscriptions service api
	subscriptionsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/v1/subscriptions",
		Method:     http.MethodPost,
		HTTPStatus: http.StatusOK,
		CallFlag:   &subscriptionsServiceCalled,
	})

	// action
	ids := make([]int64, 666, 666)
	err := subscriptions.Create(client, testutils.UserObjque, ids)

	// assert
	assert.NoError(t, err)
	assert.True(t, subscriptionsServiceCalled)
}

func TestAPI_Subscriptions_Create_EmptySubs(t *testing.T) {
	setup()
	defer teardown()

	// action
	err := subscriptions.Create(client, testutils.UserObjque, []int64{})

	// assert
	assert.Error(t, err)
}

func TestAPI_Subscriptions_Delete(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock subscriptions service api
	subscriptionsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/v1/subscriptions",
		Method:     http.MethodDelete,
		HTTPStatus: http.StatusOK,
		CallFlag:   &subscriptionsServiceCalled,
	})

	// action
	err := subscriptions.Delete(client, testutils.UserObjque, []int64{1, 2, 3})

	// assert
	assert.NoError(t, err)
	assert.True(t, subscriptionsServiceCalled)
}

func TestAPI_Subscriptions_Delete_WithLimit(t *testing.T) {
	setup()
	defer teardown()

	// arrange
	// mock subscriptions service api
	subscriptionsServiceCalled := false
	testutils.HandleReqWithoutBody(t, testutils.HandleReqWithoutBodyOpts{
		Mux:        env.Mux,
		URL:        "/v1/subscriptions",
		Method:     http.MethodDelete,
		HTTPStatus: http.StatusOK,
		CallFlag:   &subscriptionsServiceCalled,
	})

	// action
	ids := make([]int64, 666, 666)
	err := subscriptions.Delete(client, testutils.UserObjque, ids)

	// assert
	assert.NoError(t, err)
	assert.True(t, subscriptionsServiceCalled)
}
