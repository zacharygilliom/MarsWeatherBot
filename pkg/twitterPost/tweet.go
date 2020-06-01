package twitterPost

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
)

var APIKEY, APISECRETKEY, ACCESSTOKEN, ACCESSTOKENSECRET string = configs.GetTwitterCredentials()

func CreateConfig() *oauth1.Config {
	config := oauth1.NewConfig(APIKEY, APISECRETKEY)
	return config
}

func CreateToken() *oauth1.Token {
	token := oauth1.NewToken(ACCESSTOKEN, ACCESSTOKENSECRET)
	return token
}

func CreateClient() *twitter.Client {
	config := CreateConfig()
	token := CreateToken()
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client
}

func NewTweet(client *twitter.Client, day *string, temp float64) {
	if temp != 0 {
		currentTemp := fmt.Sprintf("%.2f", temp)
		tweet, resp, err := client.Statuses.Update("Most up to date Sol Day on Mars is: "+
			*day+"\n"+"The Average Temperature is "+currentTemp+" degrees celsius", nil)
		if err != nil {
			fmt.Println("Tweet Unsuccessful: %s, response: %s, error: %s", tweet, resp, err)
		} else {
			fmt.Println("Tweet Successfully Posted")
		}
	} else {
		tweet, resp, err := client.Statuses.Update("No temperature data found for Sol Day "+
			*day+"\n"+"Check Back Tomorrow", nil)
		if err != nil {
			fmt.Println("Tweet Unsuccessful: %s, response: %s, error: %s", tweet, resp, err)
		} else {
			fmt.Println("Tweet Successfully Posted")
		}
	}

}
