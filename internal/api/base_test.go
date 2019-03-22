package api

import (
	"net/http/httptest"

	"github.com/musicmash/api/internal/api/middleware/auth"
)

var (
	server *httptest.Server
)

type MockAuthorizer struct{}

func (m *MockAuthorizer) Authorize(token string) (username string, err error) {
	return "test.user1@gmail.com", nil
}

func setup() {
	server = httptest.NewServer(getMux(auth.NewMiddleware(&MockAuthorizer{})))
}

func teardown() {
	server.Close()
}
