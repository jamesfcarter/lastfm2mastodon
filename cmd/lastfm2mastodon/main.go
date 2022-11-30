package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jamesfcarter/lastfm2mastodon/internal/config"
	"github.com/jamesfcarter/lastfm2mastodon/internal/lastfm"
	"github.com/jamesfcarter/lastfm2mastodon/internal/mastodon"
)

func defaultConfig() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".lastfm2mastodon")
}

type track struct {
	artist string
	title  string
}

type trackMonitor struct {
	playing track
	played  track
}

func (tm *trackMonitor) NewTrack(t *lastfm.Track) bool {
	track := track{
		artist: t.Artist,
		title:  t.Title,
	}
	if t.CurrentlyPlaying {
		if track == tm.playing {
			return false
		}
		if tm.playing != tm.played {
			tm.played = tm.playing
		}
		tm.playing = track
		return true
	}
	if track == tm.played {
		return false
	}
	if track == tm.playing {
		tm.played = track
		return false
	}
	if track != tm.played {
		tm.played = track
		return true
	}
	return false
}

func main() {
	configFile := flag.String("config", defaultConfig(), "the configuration file")
	flag.Parse()
	config, err := config.FromFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	lastFM := lastfm.New(
		config.LastFM.Key,
		config.LastFM.Secret,
		config.LastFM.UserName,
	)
	mastodon := mastodon.New(
		config.Mastodon.ClientID,
		config.Mastodon.ClientSecret,
		config.Mastodon.AccessToken,
		config.Mastodon.URL,
	)

	pollTime := config.LastFM.PollTime()
	monitor := &trackMonitor{}

	log.SetFlags(log.Ldate | log.Ltime)
	for {
		track, err := lastFM.CurrentlyPlaying()
		switch {
		case err != nil:
			log.Println(err)
		case track != nil:
			if !monitor.NewTrack(track) {
				break
			}
			err := mastodon.Toot(fmt.Sprintf("%s - %s\n%s", track.Artist, track.Title, track.URL))
			if err != nil {
				log.Println(err)
				break
			}
			log.Printf("%d: %s - %s\n", track.Count, track.Artist, track.Title)
		}
		time.Sleep(pollTime)
	}
}
