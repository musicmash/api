package validators

import (
	"fmt"
	"net/http"
)

func IsUserExits(w http.ResponseWriter, name string) error {
	w.WriteHeader(http.StatusHTTPVersionNotSupported)
	return fmt.Errorf("user '%s' not found", name)
}
