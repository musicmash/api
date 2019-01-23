package releases

import "time"

type Release struct {
	ArtistName string          `json:"artist_name"`
	Title      string          `json:"title" gorm:"size:1000"`
	Poster     string          `json:"poster"`
	Released   time.Time       `json:"released"`
	Stores     []*ReleaseStore `json:"stores"`
}

type ReleaseStore struct {
	StoreURL  string `json:"url"`
	StoreName string `json:"name"`
	StoreID   string `json:"id"`
}
