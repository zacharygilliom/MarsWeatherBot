package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	//"github.com/dghubble/oauth1"
	//"github.com/zacharygilliom/MarsWeatherBot/configs"
	//"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
)

func NewTweet(client *twitter.Client, body string) {
	tweet, resp, err := client.Statuses.Update(body, nil)
	if err != nil {
		fmt.Println("Tweet Unsuccessful: %s, response: %s, error: %v", tweet, resp, err)
	} else {
		fmt.Println("Tweet Successfully Posted")
	}
}

func GetTimeline(client *twitter.Client) []twitter.Tweet {
	tweets, resp, err := client.Timelines.MentionTimeline(nil)
	if err != nil {
		fmt.Println(resp, err)
	}
	return tweets
}
