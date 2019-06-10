package api

import (
	"net/http/httptest"
)

var (
	server *httptest.Server
)

func setup() {
	server = httptest.NewServer(getMux())
}

func teardown() {
	server.Close()
}
