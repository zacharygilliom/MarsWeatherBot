package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
)

func ReplySwitch(message string) {

}

func ReplyAbout(user string, client *twitter.Client) {
	body := "@" + user + " This is a twitter bot created to help familiarize people with the weather data gathered by the Mars Insight Rover.  This is not affiliated with NASA, but uses the free open source API provided by NASA."
	NewTweet(client, body)

}

func ReplyDayAndWeather(user string, client *twitter.Client) {
	data := nasaData.GetData()
	currentDay := nasaData.GetSolDay()
	currentDayTemp := nasaData.GetDayTemp(currentDay, data)

	body := "@" + user
}
