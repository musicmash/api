package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/musicmash/api/internal/clients/feed"
	"github.com/musicmash/api/internal/log"
)

func getUserFeed(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	since := r.URL.Query().Get("since")
	weekAgo := time.Now().UTC().Add(-time.Hour * 24 * 7)
	if since != "" {
		var err error
		weekAgo, err = time.Parse("2006-01-02", since)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if weekAgo.After(time.Now().UTC().Truncate(time.Hour)) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	userFeed, err := feed.Get(feedProvider, userName, &feed.Options{Since: &weekAgo})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	buffer, err := json.Marshal(&userFeed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(buffer)
}
