package mastodon

import (
	"context"

	"github.com/mattn/go-mastodon"
)

type API struct {
	client *mastodon.Client
}

func New(clientID, clientSecret, accessToken, url string) *API {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       url,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	})
	return &API{
		client: c,
	}
}

func (a *API) Toot(msg string) error {
	toot := &mastodon.Toot{
		Status: msg,
	}
	_, err := a.client.PostStatus(context.Background(), toot)
	return err
}
