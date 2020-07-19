package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func NewTweet(client *twitter.Client, body string) {
	tweet, resp, err := client.Statuses.Update(body, nil)
	if err != nil {
		fmt.Println("Tweet Unsuccessful:")
		fmt.Println(tweet)
		fmt.Println("Request Response:")
		fmt.Println(resp)
		fmt.Println("Error:")
		fmt.Println(err)
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

// Create a stream of all messages sent to the specified user.
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

// Initialize our channel in a goroutine and choose what we do with the tweet.  Channel can be stopped with "Ctrl-C".
func GetMessages() {
	demux := twitter.NewSwitchDemux()
	client := client.NewOauth1()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
		fmt.Println(tweet.User.ScreenName)
		message := tweet.Text
		user := tweet.User.ScreenName
		ReplySwitch(message, user, client)
	}
	stream := Stream(client)
	fmt.Println("Stream Started...")
	go demux.HandleChan(stream.Messages)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	stream.Stop()
	fmt.Println("Stream Stopped...")
}

//When we have yesterdays weather available, function will tweet it out.
func OneDayTweet(day int, data nasaData.SolDay) string {
	currentTemp := nasaData.GetDayTemp(day, data)
	strDay := strconv.Itoa(day)
	currentTempStr := fmt.Sprintf("%.2f", currentTemp)
	tweet := "Most up to date Sol Day on Mars is: " +
		strDay + "\n" + "The Average Temperature is " + currentTempStr + " degrees celsius"
	return tweet
}

//If yesterday's weather is not available, we will tweet out the most recent readings.
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

//Decide whether a tweet will be a single day tweet or a multi day tweet
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

func intInSlice(a int, list *[]int) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

func PostTweet() {
	nasadata := nasaData.GetData()
	client := client.NewOauth1()
	if nasadata == nil {
		fmt.Println("Error While Parsing JSON Data")
	}
	solDay := nasaData.GetSolDay()
	listday := nasaData.GetListDays(nasadata)
	tweet := ConfigureTweet(listday, solDay, nasadata)
	NewTweet(client, tweet)
}
