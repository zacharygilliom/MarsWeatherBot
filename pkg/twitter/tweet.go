package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	//"github.com/dghubble/oauth1"
	//"github.com/zacharygilliom/MarsWeatherBot/configs"
	//"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	"os"
	"os/signal"
	"syscall"
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

func Stream(client *twitter.Client) *twitter.Stream {
	params := &twitter.StreamFilterParams{
		Track:         []string{"@MarsWeatherBot"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		fmt.Println("Stream Connection Failed")
	}
	return stream
}

func GetMessages(client *twitter.Client) {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
		fmt.Println(tweet.User)
	}
	stream := Stream(client)
	fmt.Println("Stream Started...")
	demux.HandleChan(stream.Messages)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	stream.Stop()
	fmt.Println("Stream Stopped...")
}
