package config

import (
	"io"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	LastFM   LastFM   `toml:"lastfm"`
	Mastodon Mastodon `toml:"mastodon"`
}

type LastFM struct {
	Key             string `toml:"key"`
	Secret          string `toml:"secret"`
	UserName        string `toml:"user_name"`
	PollTimeSeconds int    `toml:"poll_time_seconds"`
	QuitOnError     bool   `toml:"quit_on_error"`
}

type Mastodon struct {
	AccessToken  string `toml:"access_token"`
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	URL          string `toml:"url"`
}

func (l LastFM) PollTime() time.Duration {
	pollTime := l.PollTimeSeconds
	if pollTime < 1 {
		pollTime = 30
	}
	return time.Duration(pollTime) * time.Second
}

func Load(from io.Reader) (*Config, error) {
	var config Config
	_, err := toml.NewDecoder(from).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func FromFile(fname string) (*Config, error) {
	cf, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	return Load(cf)
}
