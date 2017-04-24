package auth

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
type TwitterAuth struct {
	consumer     *oauth.Consumer
	AccessToken  *oauth.AccessToken
	RequestToken *oauth.RequestToken
}

func (ad *TwitterAuth) MakeHttpClient() (*http.Client, error) {
	return ad.consumer.MakeHttpClient(ad.AccessToken)
}

func (ad *TwitterAuth) AuthorizeToken(requestToken *oauth.RequestToken, verificationCode string) (*oauth.AccessToken, error) {
	return ad.consumer.AuthorizeToken(requestToken, verificationCode)
}

func (ad *TwitterAuth) GetRequestTokenAndUrl(callback string) (*oauth.RequestToken, string, error) {
	return ad.consumer.GetRequestTokenAndUrl(callback)
}

// Creates a new twitter authentication type given name of an application registered
func NewTwitterAuth(applicationName string) *TwitterAuth {
	var auth TwitterAuth
	var creds Credentials

	// find applicationName within the settings file
	cfg, err := os.Open("config/settings.yaml")
	if err != nil {
		log.Fatal("config file not found")
	}
	viper.SetConfigType("yaml")
	viper.ReadConfig(cfg)

	err = viper.UnmarshalKey(applicationName, &creds)
	if err != nil {
		log.Fatalf("Unable to decode credentials from settings file")
	}

	auth.consumer = oauth.NewConsumer(
		creds.ConsumerAccessKey,
		creds.ConsumerSecretKey,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	return &auth
}
