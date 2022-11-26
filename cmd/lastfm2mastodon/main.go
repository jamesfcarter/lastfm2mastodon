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
	var count int
	var artist string
	var title string

	log.SetFlags(log.Ldate | log.Ltime)
	for {
		track, err := lastFM.CurrentlyPlaying()
		switch {
		case err != nil:
			log.Println(err)
		case track != nil:
			if track.Count == count &&
				track.Artist == artist &&
				track.Title == title {
				break
			}
			count = track.Count
			artist = track.Artist
			title = track.Title
			err := mastodon.Toot(fmt.Sprintf("%s - %s\n%s", artist, title, track.URL))
			if err != nil {
				log.Println(err)
				break
			}
			log.Printf("%s - %s\n", artist, title)
		}
		time.Sleep(pollTime)
	}
}
