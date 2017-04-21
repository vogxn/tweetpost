package config

import (
	_ "fmt"
	"log"
	"net/http"
	"os"

	"github.com/mrjones/oauth"
	"github.com/spf13/viper"
)

type Credentials struct {
	ConsumerAccessKey string `mapstructure:"consumer_access_key"`
	ConsumerSecretKey string `mapstructure:"consumer_secret_key"`
}

// Encapsulates all that we need to pass around controller methods
type AuthDetail struct {
	Consumer     *oauth.Consumer
	AccessToken  *oauth.AccessToken
	RequestToken *oauth.RequestToken
}

var Creds Credentials
var Auth AuthDetail

func (ad *AuthDetail) MakeHttpClient() (*http.Client, error) {
	return ad.Consumer.MakeHttpClient(ad.AccessToken)
}

func init() {
	cfg, err := os.Open("config/settings.yaml")
	if err != nil {
		log.Fatal("config file not found")
	}
	viper.SetConfigType("yaml")
	viper.ReadConfig(cfg)

	err = viper.UnmarshalKey("tweetpost", &Creds)
	if err != nil {
		log.Fatalf("Unable to decode credentials from settings file")
	}

	Auth.Consumer = oauth.NewConsumer(
		Creds.ConsumerAccessKey,
		Creds.ConsumerSecretKey,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
}
