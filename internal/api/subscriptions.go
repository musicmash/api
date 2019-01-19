package api

import (
	"net/http"
)

func createSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return
}

func deleteSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return
}

func getUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return
}
