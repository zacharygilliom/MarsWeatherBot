package client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
)

var APIKEY, APISECRETKEY, ACCESSTOKEN, ACCESSTOKENSECRET string = configs.GetTwitterCredentials()

func New() *twitter.Client {
	config := oauth1.NewConfig(APIKEY, APISECRETKEY)
	token := oauth1.NewToken(ACCESSTOKEN, ACCESSTOKENSECRET)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}
