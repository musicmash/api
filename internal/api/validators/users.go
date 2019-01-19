package validators

import (
	"net/http"
)

func IsUserExits(w http.ResponseWriter, name string) error {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
}
