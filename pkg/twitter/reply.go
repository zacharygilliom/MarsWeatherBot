package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	//"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	//"strconv"
	"strings"
)

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
	/*
		fmt.Println(nasadata)
		currentDay := nasaData.GetSolDay()
		fmt.Println(currentDay)
		//strDay := strconv.Itoa(currentDay)
		//currentDayTemp := nasaData.GetDayTemp(currentDay, data)
		//currentTempStr := fmt.Sprintf("%.2f", currentDayTemp)
		listday := nasaData.GetListDays(nasadata)
		fmt.Println(listday)
		body := ConfigureTweet(listday, currentDay, nasadata)
		fmt.Println(body)

		tweetBody := "@" + user + " " + body
		fmt.Println(tweetBody)
		//body := "@" + user + "\n Current Mars Sol Day: " + strDay + "\n Current Temperature: " + currentTempStr
		NewTweet(client, tweetBody)
	*/
}

func ReplyHelp(user string, client *twitter.Client) {
	body := "@" + user + " Check out the current flag options below \n --Weather\n --About\n --Help"
	NewTweet(client, body)
}

func ReplyNull(user string, client *twitter.Client) {
	body := "@" + user + " Not a valid tweet flag.  Type '--Help' for more information."
	NewTweet(client, body)
}
