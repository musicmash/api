package auth

import "net/http"

const (
	HeaderToken    = "x-auth-token"
	HeaderUserName = "user_name"
)

func setUserName(r *http.Request, userName string) {
	r.Header.Set(HeaderUserName, userName)
}

func GetUserName(r *http.Request) string {
	return r.Header.Get(HeaderUserName)
}

func GetToken(r *http.Request) string {
	return r.Header.Get(HeaderToken)
}
