# lastfm2mastodon

lastfm2mastodon is a simple tool that posts songs scrobbled to last.fm. It is configured via a file in TOML format. By default this is `~/.lastfm2mastodon` but may be specified via the `-config` command line flag.

The configuration file looks like this:
```
[lastfm]
key = "last_fm_key"
secret = "last_fm_secret"
user_name = "last_fm_username"
poll_time_seconds = 60

[mastodon]
access_token = "mastodon_access_token"
client_id = "mastodon_client_id"
client_secret = "mastodon_client_secret"
url = "https://mastodon.social/"
```

An API key for last.fm is required. You can create one here (you only need to specify email and application name): https://www.last.fm/api/account/create

You also need an application key for your Masotodon instance. Log in and go to your preferences. Then go to the _Development_ section and click on _New Application_. You need to specify an application name and make sure the scopes include `write`.


Once started the application will run forever logging each song that it tooted as well as any errors.

