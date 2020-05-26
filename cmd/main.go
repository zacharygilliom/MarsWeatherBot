package main

import (
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitterPost"
)

func main() {
	// Query our data and configure into our struct
	nasadata := nasaData.GetData()
	fmt.Println("GetData configured correctly")
	if nasadata == nil {
		fmt.Println(nasadata)
	}

	// Pull the Average Temperature data out and format
	dailyTemp := nasadata["530"].AT.Av
	tweet := fmt.Sprintf("%.2f", dailyTemp)

	// Create our Twitter Client through the go-twitter API.
	// And Post the tweet
	client := twitterPost.CreateClient()
	twitterPost.PostTweet(client, tweet)
}
