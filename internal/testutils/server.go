package testutils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Env represents a testing environment for all resources.
type Env struct {
	Mux    *http.ServeMux
	Server *httptest.Server
}

// Setup prepares the new testing environment.
func Setup() *Env {
	mux := http.NewServeMux()
	return &Env{
		Mux:    mux,
		Server: httptest.NewServer(mux),
	}
}

// TearDown releases the testing environment.
func (t *Env) TearDown() {
	t.Server.Close()
}

// HandleReqWithoutBodyOpts represents options for test handler without request
// body.
type HandleReqWithoutBodyOpts struct {
	Mux         *http.ServeMux
	URL         string
	RawResponse string
	Method      string
	HTTPStatus  int
	CallFlag    *bool
}

// HandleReqWithoutBody provides the HTTP endpoint to test requests without
// body.
func HandleReqWithoutBody(t *testing.T, opts HandleReqWithoutBodyOpts) {
	opts.Mux.HandleFunc(opts.URL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(opts.HTTPStatus)
		fmt.Fprintf(w, opts.RawResponse)

		if r.Method != opts.Method {
			t.Fatalf("expected %s method but got %s", opts.Method, r.Method)
		}

		*opts.CallFlag = true
	})
}
