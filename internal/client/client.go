package client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var APIKEY, APISECRETKEY, ACCESSTOKEN, ACCESSTOKENSECRET string = configs.GetTwitterCredentials()

func NewOauth1() *twitter.Client {
	config := oauth1.NewConfig(APIKEY, APISECRETKEY)
	token := oauth1.NewToken(ACCESSTOKEN, ACCESSTOKENSECRET)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}

func NewOauth2() *twitter.Client {
	config := &clientcredentials.Config{
		ClientID:     APIKEY,
		ClientSecret: APISECRETKEY,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(oauth2.NoContext)
	client := twitter.NewClient(httpClient)
	return client

}
