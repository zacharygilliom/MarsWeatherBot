package twitterPost

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
)

func CreateClient() *twitter.Client {
	apikey, apisecretkey, accesstoken, accesstokensecret := configs.GetTwitterCredentials()
	config := oauth1.NewConfig(apikey, apisecretkey)
	token := oauth1.NewToken(accesstoken, accesstokensecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	return client

}

func PostTweet(client *twitter.Client) {
	tweet, resp, err := client.Statuses.Update("Test Tweet", nil)
	if err != nil {
		fmt.Println("Tweet Unsuccessful: %d, response: %d, error: %d", tweet, resp, err)
	} else {
		fmt.Println("Tweet Successfully Posted")
	}
}
