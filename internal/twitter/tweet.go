package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	log "github.com/sirupsen/logrus"
	"github.com/zacharygilliom/MarsWeatherBot/internal/client"
	"github.com/zacharygilliom/MarsWeatherBot/internal/nasaData"
	"strconv"
)

// Send tweet to Twitter API
func NewTweet(client *twitter.Client, body string) {
	tweet, resp, err := client.Statuses.Update(body, nil)
	if err != nil {
		log.Debug("Tweet Unsuccessful: %v", err)
		log.Debug(tweet, resp)
		/*
			fmt.Println("Tweet Unsuccessful:")
			fmt.Println(tweet)
			fmt.Println("Request Response:")
			fmt.Println(resp)
			fmt.Println("Error:")
			fmt.Println(err)
		*/
	} else {
		log.Info("Tweet Successfully Posted")
	}
}

// Will be used in future. Will use this later
func GetTimeline(client *twitter.Client) []twitter.Tweet {
	tweets, resp, err := client.Timelines.MentionTimeline(nil)
	if err != nil {
		log.Debug("Error Connecting to users Timeline: %v", err)
		log.Debug("Error Response: %v", resp)
	}
	return tweets
}

// Create a stream of all messages sent to the specified user.
func Stream(client *twitter.Client) *twitter.Stream {
	params := &twitter.StreamFilterParams{
		Track:         []string{"@MarsWeatherBot"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		log.Debug("Stream Connection Failed: %v", err)
		return nil
	}
	return stream
}

// Initialize our channel in a goroutine and choose what we do with the tweet.  Channel can be stopped with "Ctrl-C".
func GetMessages() (twitter.SwitchDemux, *twitter.Client) {
	demux := twitter.NewSwitchDemux()
	client := client.NewOauth1()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
		fmt.Println(tweet.User.ScreenName)
		message := tweet.Text
		user := tweet.User.ScreenName
		ReplySwitch(message, user, client)
	}
	return demux, client
}

// When we have yesterdays weather available, function will tweet it out.
func OneDayTweet(day int, data nasaData.SolDay) string {
	currentTemp := nasaData.GetDayTemp(day, data)
	strDay := strconv.Itoa(day)
	currentTempStr := fmt.Sprintf("%.2f", currentTemp)
	tweet := "Most up to date Sol Day on Mars is: " +
		strDay + "\n" + "The Average Temperature is " + currentTempStr + " degrees celsius"
	return tweet
}

// If yesterday's weather is not available, we will tweet out the most recent readings.
func MultiDayTweet(days *[]int, data nasaData.SolDay) string {
	var tweet string
	tweet = "No New Data. Most Recent Readings:\n"
	for _, day := range *days {
		fmt.Println(day)
		currentTemp := nasaData.GetDayTemp(day, data)
		fmt.Println(currentTemp)
		currentTempStr := fmt.Sprintf("%.2f", currentTemp)
		strday := strconv.Itoa(day)
		tweet += "Sol Day - " + strday + ": " + currentTempStr + " °C\n"
	}
	return tweet
}

// Decide whether a tweet will be a single day tweet or a multi day tweet
func ConfigureTweet(days *[]int, day int, data nasaData.SolDay) string {
	check := intInSlice(day, days)
	temp := nasaData.GetDayTemp(day, data)
	if check == true && temp != 0 {
		tweet := OneDayTweet(day, data)
		return tweet
	} else {
		tweet := MultiDayTweet(days, data)
		return tweet
	}
}

// Check for value in a list
func intInSlice(a int, list *[]int) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

// Send request to NASA API and tweet out our daily automated tweet
func PostTextTweet() {
	nasadata := nasaData.GetData()
	client := client.NewOauth1()
	if nasadata == nil {
		log.Debug("Error While Parsing JSON Data")
	}
	solDay := nasaData.GetSolDay()
	listday := nasaData.GetListDays(nasadata)
	tweet := ConfigureTweet(listday, solDay, nasadata)
	NewTweet(client, tweet)
}
