package db

import (
	"time"
)

type Artist struct {
	ID         int64 `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"-"`
	CreatedAt  time.Time
	Name       string
	SearchName string
}

type ArtistMgr interface {
	CreateArtist(artist *Artist) error
	FindArtistByName(id string) (*Artist, error)
	GetAllArtists() ([]*Artist, error)
	EnsureArtistExists(artist *Artist) error
}

func (mgr *AppDatabaseMgr) FindArtistByName(name string) (*Artist, error) {
	artist := Artist{}
	if err := mgr.db.Where("name = ?", name).First(&artist).Error; err != nil {
		return nil, err
	}

	return &artist, mgr.db.Find(&artist).Error
}

func (mgr *AppDatabaseMgr) GetAllArtists() ([]*Artist, error) {
	var artists = make([]*Artist, 0)
	return artists, mgr.db.Find(&artists).Error
}

func (mgr *AppDatabaseMgr) CreateArtist(artist *Artist) error {
	return mgr.db.Create(artist).Error
}

func (mgr *AppDatabaseMgr) EnsureArtistExists(artist *Artist) error {
	_, err := mgr.FindArtistByName(artist.Name)
	if err != nil {
		return mgr.CreateArtist(artist)
	}
	return nil
}
