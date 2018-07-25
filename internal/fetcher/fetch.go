package fetcher

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/objque/musicmash/internal/config"
	"github.com/objque/musicmash/internal/db"
	"github.com/objque/musicmash/internal/itunes"
	"github.com/objque/musicmash/internal/log"
	"github.com/pkg/errors"
)

func saveIfNewestRelease(artist string, release *itunes.Release) bool {
	if !release.IsLatest() {
		return false
	}

	log.Infof("Found a new release from '%s' ('%s'): '%s'", release.ArtistName, artist, release.CollectionName)
	db.DbMgr.EnsureReleaseExists(&db.Release{
		// NOTE (m.kalinin): we provide artist because if artist releases feat with someone, then
		// release will contain incorrect name.
		ArtistName: artist,
		StoreID:    release.TrackID,
	})
	return true
}

func fetch() error {
	// load all artists from the db
	artists, err := db.DbMgr.GetAllArtists()
	if err != nil {
		return errors.Wrap(err, "can't load artists from the db")
	}

	// load releases from the store
	for _, artist := range artists {
		lastRelease, err := itunes.GetLatestAlbumRelease(artist.SearchName)
		if err != nil {
			log.Error(errors.Wrapf(err, "can't load artist/album '%s' from the iTunes", artist.SearchName))
			continue
		}
		if saveIfNewestRelease(artist.Name, lastRelease) {
			// NOTE (m.kalinin): continue because album may be full ep or single and we
			// do not need to search single track release
			continue
		}

		lastRelease, err = itunes.GetLatestTrackRelease(artist.SearchName)
		if err != nil {
			log.Error(errors.Wrapf(err, "can't load artist/track '%s' from the iTunes", artist.SearchName))
			continue
		}
		saveIfNewestRelease(artist.Name, lastRelease)

		// NOTE (m.kalinin): iTunes has rate-limit 20 requests per minute.
		// Until we won't use proxies, we should sleep.
		time.Sleep(time.Second * 3)
	}
	return nil
}

func isMustFetch() bool {
	last, err := db.DbMgr.GetLastFetch()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}

		log.Error(err)
		return false
	}
	return calcDiffHours(last.Date) > config.Config.Fetching.CountOfSkippedHoursToFetch
}

func Run() {
	for {
		if isMustFetch() {
			now := time.Now().UTC()
			log.Infof("Start fetching stage for '%s'...", now.String())
			if err := fetch(); err != nil {
				log.Error(err)
			} else {
				log.Infof("Finish fetching stage '%s'...", time.Now().UTC().String())
				db.DbMgr.SetLastFetch(now)
			}
		}

		time.Sleep(time.Hour)
	}
}
