package api

import (
	"net/http"
)

func createSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
}

func deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
}

func getUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
}
