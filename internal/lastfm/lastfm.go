package lastfm

import (
	"github.com/shkh/lastfm-go/lastfm"
)

type API struct {
	user string
	api  *lastfm.Api
}

type Track struct {
	Artist           string
	Title            string
	URL              string
	Count            int
	CurrentlyPlaying bool
}

func New(key, secret, user string) *API {
	return &API{
		user: user,
		api:  lastfm.New(key, secret),
	}
}

func (a *API) CurrentlyPlaying() (*Track, error) {
	recent, err := a.api.User.GetRecentTracks(lastfm.P{
		"limit": 1,
		"user":  a.user,
	})
	if err != nil {
		return nil, err
	}
	if len(recent.Tracks) == 0 {
		return nil, nil
	}
	return &Track{
		Artist:           recent.Tracks[0].Artist.Name,
		Title:            recent.Tracks[0].Name,
		URL:              recent.Tracks[0].Url,
		Count:            recent.Total,
		CurrentlyPlaying: recent.Tracks[0].NowPlaying == "true",
	}, nil
}
