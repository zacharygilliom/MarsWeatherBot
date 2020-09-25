package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/zacharygilliom/MarsWeatherBot/internal/nasaData"
	"strings"
)

// Read in the tweet by the user and decide how to respond.
func ReplySwitch(message, user string, client *twitter.Client) {
	if strings.Contains(message, "--Weather") {
		ReplyDayAndWeather(user, client)
	} else if strings.Contains(message, "--About") {
		ReplyAbout(user, client)
	} else if strings.Contains(message, "--Help") {
		ReplyHelp(user, client)
	} else {
		ReplyNull(user, client)
	}
}

func ReplyAbout(user string, client *twitter.Client) {
	body := "@" + user + " This is a twitter bot created with the weather data gathered by the Mars Insight Rover.  This is not affiliated with NASA, but uses the free open source API provided by NASA."
	NewTweet(client, body)
}

// Check if data queried is not empty and tweet out the weather.
func ReplyDayAndWeather(user string, client *twitter.Client) {
	nasadata := nasaData.GetData()
	if len(nasadata) == 0 {
		tweetBody := "@" + user + " I'm sorry, but the weather data has not been updated in the past seven days.  Please check back tomorrow."
		NewTweet(client, tweetBody)
	} else {
		currentDay := nasaData.GetSolDay()
		listday := nasaData.GetListDays(nasadata)
		body := ConfigureTweet(listday, currentDay, nasadata)
		tweetBody := "@" + user + " " + body
		NewTweet(client, tweetBody)
	}
}

func ReplyHelp(user string, client *twitter.Client) {
	body := "@" + user + " Check out the current flag options below \n --Weather\n --About\n --Help"
	NewTweet(client, body)
}

func ReplyNull(user string, client *twitter.Client) {
	body := "@" + user + " Not a valid tweet flag.  Type '--Help' for more information."
	NewTweet(client, body)
}
